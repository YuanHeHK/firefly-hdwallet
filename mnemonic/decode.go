package mnemonic

import (
	"golang.org/x/crypto/scrypt"
)


func decryptData(cipherText, auth, salt, iv []byte) ([]byte, error) {

	derivedKey, err := scrypt.Key(auth, salt, scryptN, scryptR, scryptP, scryptDKLen)
	if err != nil {
		return nil, err
	}

	plainText, err := aesCTRXOR(derivedKey[:16], cipherText, iv)
	if err != nil {
		return nil, err
	}
	return plainText, err
}
