package main

import (
	"fmt" //基本的な入出力処理
	"goland/getinfo"
	"io"       //入出力に使います
	"log"      //エラーを表示するときに使います
	"net/http" //HTTPを使った通信に必要
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {

	io.WriteString(w, "Hello World!")
}

func main() {
	fmt.Println("Hello, World!")
	http.HandleFunc("/", HelloWorld)             //ルートのアクセスにHelloWorldをハンドリング
	http.HandleFunc("/getinfo", getinfo.GetInfo) //ルートのアクセスにHelloWorldをハンドリング
	err := http.ListenAndServe(":80", nil)       //サーバ起動
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
	for {
	} //ずっと起動
}
