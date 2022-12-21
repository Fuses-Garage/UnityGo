package usermethod

import (
	//基本的な入出力処理
	"database/sql"
	"fmt"
	"io"       //入出力
	"net/http" //http通信
	"time"

	"github.com/Fuses-Garage/UnityGo/util"
	_ "github.com/lib/pq" //ポスグレ使用
)

func LoginSuccess(loginname string, db *sql.DB) string { //ログイン成功時の処理
	ans := ""
	row := db.QueryRow("SELECT sessioncode FROM userdata_login WHERE loginname=$1", loginname) //ログインしたユーザのセッションコードを検索
	row.Scan(&ans)
	if ans == "" { //まだセッションがないなら
		randstr, err := MakeRandomStr(256)
		util.CheckErr(err)
		ans = util.StringtoHex(loginname + randstr + time.Now().String())                      //セッションIDを生成
		db.Exec("UPDATE userdata_login SET sessioncode=$1 WHERE loginname=$2", ans, loginname) //セッションIDを書き込み
	}
	db.Exec("UPDATE userdata_login SET lastlogin=$1 WHERE loginname=$2", time.Now(), loginname) //最終ログイン日時を更新
	return ans
}
func CheckSession(SID string, db *sql.DB) string {
	row := db.QueryRow("SELECT COUNT(*) AS count FROM userdata_login WHERE sessioncode=$1", SID) //セッションIDが一致するユーザーを探す
	var count int
	err := row.Scan(&count) //ヒット数をintで受ける
	util.CheckErr((err))
	if count == 1 { //1件のみヒットしたら
		namerow := db.QueryRow("SELECT loginname FROM userdata_login WHERE sessioncode=$1", SID) //セッションIDが一致するユーザーを探す
		var loginname string
		err = namerow.Scan(&loginname) //ログイン名を取り出す
		util.CheckErr(err)
		LoginSuccess(loginname, db) //ログイン成功時の処理
		return "success"            //成功！
	} else { //ヒットしないもしくは複数ヒットしたら
		return "login please" //ログインを要求
	}
}
func GetUInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" { //POSTじゃなければ
		io.WriteString(w, "危険なのでPOSTメソッドでアクセスしてください。") //エラーメッセージを返す
		return
	}
	err := r.ParseForm() //postボディを詠む
	util.CheckErr(err)
	if r.FormValue("SID") == "" { //セッションIDが空白なら
		io.WriteString(w, "login please") //ログインを要求する
		return
	} else {
		db := util.LoginToDB()                         //DBにログイン
		result := CheckSession(r.FormValue("SID"), db) //ログイン状態を確認
		if result == "success" {                       //ログイン中なら
			namerow := db.QueryRow("SELECT name FROM userdata_login WHERE sessioncode=$1", r.FormValue("SID")) //セッションIDが一致するユーザーを探す
			name := ""
			namerow.Scan(&name)                        //ユーザ情報を取得
			json := fmt.Sprintf(`{"name":"%s"}`, name) //JSON文字列にする
			io.WriteString(w, result+" "+json)         //ログインを要求する
		} else {
			io.WriteString(w, result) //ログインを要求する
		}

		return
	}
}
