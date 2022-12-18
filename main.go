package main

import (
	"net/http" //HTTPを使った通信に必要

	"github.com/Fuses-Garage/UnityGo/getinfo"
	"github.com/Fuses-Garage/UnityGo/usermethod"
)

func main() {
	http.HandleFunc("/getinfo", getinfo.GetInfo)               //ルートのアクセスにHelloWorldをハンドリング
	http.HandleFunc("/useradd", usermethod.UserAdd)            //ルートのアクセスにHelloWorldをハンドリング
	http.HandleFunc("/login_basic", usermethod.Login_Basic)    //ルートのアクセスにHelloWorldをハンドリング
	http.HandleFunc("/getchallenge", usermethod.MakeChallenge) //ルートのアクセスにHelloWorldをハンドリング
	http.HandleFunc("/login_chap", usermethod.LoginChap)       //ルートのアクセスにHelloWorldをハンドリング
	err := http.ListenAndServe(":80", nil)                     //サーバ起動
	if err != nil {
		panic(err)
	}
	for {
	} //ずっと起動
}
