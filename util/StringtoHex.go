package util

import (
	"crypto/sha256" //ハッシュ関数
	"encoding/hex"  //16進数
)

func StringtoHex(s string) string { //文字列にSHA-256して返す
	byt := sha256.Sum256([]byte(s)) //ハッシュ化
	bina := byt[:]                  //おまじない
	hash := hex.EncodeToString(bina)
	return hash
}
