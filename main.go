package main

import (
	"github.com/Fuses-Garage/UnityGo/getinfo"
	useradd "github.com/Fuses-Garage/UnityGo/usermethod"

	//"log"      //エラーを表示するときに使います
	"net/http" //HTTPを使った通信に必要
)

func main() {
	http.HandleFunc("/", getinfo.GetInfo)        //ルートのアクセスにHelloWorldをハンドリング
	http.HandleFunc("/adduser", useradd.UserAdd) //ルートのアクセスにHelloWorldをハンドリング
	err := http.ListenAndServe(":80", nil)       //サーバ起動
	if err != nil {
		//log.Fatal("ListenAndServe:", err)
	}
	for {
	} //ずっと起動
}
