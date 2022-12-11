package usermethod

import (
	"fmt"      //基本的な入出力処理
	"io"       //入出力
	"net/http" //http通信

	"github.com/Fuses-Garage/UnityGo/util"
	_ "github.com/lib/pq" //ポスグレ使用
)

func UserAdd(w http.ResponseWriter, r *http.Request) {
	fmt.Println("useradd")
	if r.Method != "POST" { //一つでも空欄があれば
		io.WriteString(w, "危険なのでPOSTメソッドでアクセスしてください。") //エラーメッセージを返す
		return
	}
	err := r.ParseForm() //postボディを詠む
	util.CheckErr(err)
	if r.FormValue("name") == "" || r.FormValue("pass") == "" || r.FormValue("loginname") == "" { //一つでも空欄があれば
		io.WriteString(w, "全てのテキストボックスに値を入力してください。") //エラーメッセージを返す
		return
	}
	util.CheckErr(err)
	db := util.LoginToDB()
	util.CheckErr(err)
	if err != nil {
		fmt.Printf("読み込み失敗: %v", err)
	}
	row := db.QueryRow("SELECT COUNT(*) AS count FROM userdata_login WHERE loginname=$1 OR name=$2", r.FormValue("loginname"), r.FormValue("name")) //同じログインネームのユーザーを探す
	var count int
	err = row.Scan(&count) //ヒット数をintで受ける
	util.CheckErr((err))
	if count > 0 { //誰かいたら
		io.WriteString(w, "そのユーザーネームは既に登録されています") //エラーメッセージを返す
		return
	}
	hash := util.StringtoHex(r.FormValue("pass"))
	_, err = db.Exec("INSERT INTO userdata_login VALUES(DEFAULT,$1,$2,$3)", r.FormValue("name"), r.FormValue("loginname"), hash) //新しいレコードを追加
	util.CheckErr(err)
	io.WriteString(w, "success") //string化したものを送信
}
