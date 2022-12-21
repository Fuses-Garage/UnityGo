package usermethod

import (
	"crypto/rand"
	"errors"
	"fmt"
	"io"       //入出力
	"net/http" //http通信

	"github.com/Fuses-Garage/UnityGo/util"
	_ "github.com/lib/pq" //ポスグレ使用
)

func MakeRandomStr(digit uint32) (string, error) {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// 乱数を生成
	b := make([]byte, digit)
	if _, err := rand.Read(b); err != nil {
		return "", errors.New("unexpected error")
	}

	// letters からランダムに取り出して文字列を生成
	var result string
	for _, v := range b {
		// index が letters の長さに収まるように調整
		result += string(letters[int(v)%len(letters)])
	}
	return result, nil
}
func MakeChallenge(w http.ResponseWriter, r *http.Request) {
	chl, err := MakeRandomStr(64) //チャレンジを生成
	util.CheckErr(err)
	code, err := MakeRandomStr(64) //チャレンジを識別する文字列を生成
	util.CheckErr(err)
	db := util.LoginToDB()                                           //DBにログイン
	db.Exec("INSERT INTO chltable VALUES(DEFAULT,$1,$2)", code, chl) //DBにチャレンジを追加
	json := fmt.Sprintf(`{"code":"%s","chl":"%s"}`, code, chl)       //JSON文字列にする
	io.WriteString(w, json)                                          //生成したJSONを返す
}
