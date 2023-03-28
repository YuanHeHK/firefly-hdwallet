package mnemonic

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/howeyc/gopass"
	logging "github.com/ipfs/go-log/v2"
	"github.com/tyler-smith/go-bip39"
	"golang.org/x/crypto/scrypt"
)

var log = logging.Logger("encode")

var (
	scryptN     = 1 << 18
	scryptP     = 1
	scryptR     = 8
	scryptDKLen = 32
)

type hiddenInfo struct {
	Hidden []byte `json:"hidden"`
	Iv     []byte `json:"iv"`
	Salt   []byte `json:"salt"`
}

func Encrypt(usdtPath string) error {
	passwd, err := gopass.GetPasswdMasked()
	if err != nil {
		return err
	}

	err = initEncryptMnemonic(passwd, usdtPath)
	if err != nil {
		return err
	}

	return nil
}

func initEncryptMnemonic(passwd []byte, usdtPath string) error {
	data, err := ioutil.ReadFile(usdtPath)
	if err != nil {
		return err
	}

	dir := filepath.Dir(usdtPath)
	data = data[:len(data)-1]
	mnemonic := string(data)
	modelsMnemonic := mnemonic
	ok := bip39.IsMnemonicValid(modelsMnemonic)
	if !ok {
		return errors.New("mnemonic is not valid")
	}

	if err = encryptData(data, passwd, filepath.Join(dir, "./hidde.txt")); err != nil {
		return err
	}
	if err = os.RemoveAll(usdtPath); err != nil {
		return err
	}

	return nil
}

func encryptData(data, auth []byte, path string) error {

	salt := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		panic("reading from crypto/rand failed: " + err.Error())
	}
	derivedKey, err := scrypt.Key(auth, salt, scryptN, scryptR, scryptP, scryptDKLen)
	if err != nil {
		return err
	}
	encryptKey := derivedKey[:16]

	iv := make([]byte, aes.BlockSize) // 16
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic("reading from crypto/rand failed: " + err.Error())
	}
	cipherText, err := aesCTRXOR(encryptKey, data, iv)
	if err != nil {
		return err
	}
	hid := hiddenInfo{
		Hidden: cipherText,
		Iv:     iv,
		Salt:   salt,
	}
	hidData, err := json.Marshal(hid)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path, hidData, 0640)
	if err != nil {
		return err
	}
	return nil
}

func aesCTRXOR(key, inText, iv []byte) ([]byte, error) {
	// AES-128 is selected due to size of encryptKey.
	aesBlock, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	stream := cipher.NewCTR(aesBlock, iv)
	outText := make([]byte, len(inText))
	stream.XORKeyStream(outText, inText)
	return outText, err
}
