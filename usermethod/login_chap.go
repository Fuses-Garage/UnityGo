package usermethod

import (
	"fmt"      //基本的な入出力処理
	"io"       //入出力
	"net/http" //http通信

	"github.com/Fuses-Garage/UnityGo/util"
	_ "github.com/lib/pq" //ポスグレ使用
)

func LoginChap(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" { //POSTじゃなければ
		io.WriteString(w, "危険なのでPOSTメソッドでアクセスしてください。") //エラーメッセージを返す
		return
	}
	db := util.LoginToDB() //DBにログイン
	err := r.ParseForm()   //postボディを詠む
	if err != nil {
		fmt.Printf("読み込み失敗: %v", err)
	}
	resp := r.FormValue("pass")
	row := db.QueryRow("SELECT chl FROM chltable WHERE code=$1", r.FormValue("code")) //一時保存していたチャレンジを取得
	db.Exec("DELETE FROM chltable WHERE code=$1", r.FormValue("code"))                //保存していたチャレンジを削除
	var chl string
	err = row.Scan(&chl)
	util.CheckErr(err)
	if r.FormValue("pass") == "" || r.FormValue("loginname") == "" || r.FormValue("code") == "" { //一つでも空欄があれば
		io.WriteString(w, "全てのテキストボックスに値を入力してください。") //エラーメッセージを返す
		return
	}
	var passhash string                                                                                   //受け取ったパスワードをハッシュ化
	row = db.QueryRow("SELECT passhash FROM userdata_login WHERE loginname=$1", r.FormValue("loginname")) //ログイン名が一致するパスワード（のハッシュ値）を探す
	if row.Scan(&passhash) != nil {                                                                       //対応するユーザーがいなかったら
		io.WriteString(w, "ユーザ名かログインIDが間違っています") //エラーメッセージの情報は最低限
		return
	}
	answer := util.StringtoHex(passhash + chl) //正しいレスポンスを計算
	util.CheckErr((err))
	if answer == resp { //2つのレスポンスが合致すれば

		io.WriteString(w, "success "+LoginSuccess(r.FormValue("loginname"), db)) //ログイン成功
		return
	} else { //ヒットしないもしくは複数ヒットしたら
		io.WriteString(w, "ユーザ名かログインIDが間違っています") //エラーメッセージの情報は最低限
		return
	}
}
