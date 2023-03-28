package main

import (
	"flag"
	"fmt"
	"os"
	"sync"
	"walletmanager/mnemonic"

	_ "github.com/go-sql-driver/mysql"
)

var Wg sync.WaitGroup

var _version_ = ""

func main() {

	var myHelp bool
	flag.BoolVar(&myHelp, "h", false, "-h")
	flag.Parse()
	if myHelp {
		fmt.Println("Version : ",_version_)
		fmt.Println("help : 本例用来使用开源bip39协议将助记词通过客户的密码一起加密写入一个加密文件。")
		fmt.Println("输入参数助记词文件路径。例如：./ff-hdtool usdt.txt")
		return
	}

	err := preConditioning()
	if err != nil {
		return
	}
}

func preConditioning() error {
	args := os.Args
	if len(args) == 2 {
		usdtPath := args[1]
		fmt.Println("input encode password:")
		err := mnemonic.Encrypt(usdtPath)
		if err != nil {
			fmt.Println("encode file error:", err)
			return err
		}
		fmt.Println("encode file success!")
	} else {
		fmt.Println("need more params")
		return nil
	}
	return nil
}
