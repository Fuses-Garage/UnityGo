package keymethod

import (
	"crypto/rand" //複合のための乱数生成に必要
	"crypto/rsa"  //複合に必要

	"github.com/Fuses-Garage/UnityGo/util"
)

func Decrypto(s string) string { //文字列sを複合する
	result, err := rsa.DecryptPKCS1v15(rand.Reader, getPrvKey(), []byte(s)) //秘密鍵で複合
	util.CheckErr(err)
	return string(result) //Stringにして返す
}
