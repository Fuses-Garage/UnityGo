package usermethod

import (
	"fmt"      //基本的な入出力処理
	"io"       //入出力
	"net/http" //http通信

	"github.com/Fuses-Garage/UnityGo/util"
	_ "github.com/lib/pq" //ポスグレ使用
)

func Login_Basic(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" { //一つでも空欄があれば
		io.WriteString(w, "危険なのでPOSTメソッドでアクセスしてください。") //エラーメッセージを返す
		return
	}
	err := r.ParseForm() //postボディを詠む
	util.CheckErr(err)
	if r.FormValue("pass") == "" || r.FormValue("loginname") == "" { //一つでも空欄があれば
		io.WriteString(w, "全てのテキストボックスに値を入力してください。") //エラーメッセージを返す
		return
	}
	db := util.LoginToDB()
	if err != nil {
		fmt.Printf("読み込み失敗: %v", err)
	}
	hash := util.StringtoHex(r.FormValue("pass")) //受け取ったパスワードをハッシュ化
	row := db.QueryRow("SELECT COUNT(*) AS count FROM userdata_login WHERE loginname=$1 AND passhash=$2",
		r.FormValue("loginname"), hash) //ログイン名とパスワード（のハッシュ値）が一致するユーザーを探す
	var count int
	err = row.Scan(&count) //ヒット数をintで受ける
	util.CheckErr((err))
	if count == 1 { //1件のみヒットしたら
		io.WriteString(w, "success") //ログイン成功
		return
	} else { //ヒットしないもしくは複数ヒットしたら
		io.WriteString(w, "ユーザ名かログインIDが間違っています") //エラーメッセージの情報は最低限
		return
	}
}
