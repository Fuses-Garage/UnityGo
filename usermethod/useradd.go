package useradd

import (
	"crypto/sha256" //ハッシュ関数
	"database/sql"  //SQL操作
	"encoding/hex"  //16進数
	"fmt"           //基本的な入出力処理
	"io"            //入出力
	"net/http"      //http通信
	"os"

	"github.com/joho/godotenv" //.envの使用
	_ "github.com/lib/pq"      //ポスグレ使用
)

func GetInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" { //一つでも空欄があれば
		io.WriteString(w, "危険なのでPOSTメソッドでアクセスしてください。") //エラーメッセージを返す
		return
	}
	err := r.ParseForm() //postボディを詠む
	checkErr(err)
	if r.Form.Get("name") == "" || r.Form.Get("pass") == "" || r.Form.Get("loginname") == "" { //一つでも空欄があれば
		io.WriteString(w, "全てのテキストボックスに値を入力してください。") //エラーメッセージを返す
		return
	}
	err = godotenv.Load(".vscode/.env")
	checkErr(err)
	db, err := sql.Open("postgres", "user="+os.Getenv("ENVUSER")+" password="+os.Getenv("ENVPASS")+" dbname=UniGoDB sslmode=disable") //DBに接続
	checkErr(err)
	if err != nil {
		fmt.Printf("読み込み失敗: %v", err)
	}
	row := db.QueryRow("SELECT COUNT(*) AS count FROM userdata_login WHERE loginname=$1 OR name=$2", r.Form.Get("loginname"), r.Form.Get("name")) //同じログインネームのユーザーを探す
	var count int
	err = row.Scan(&count) //ヒット数をintで受ける
	checkErr((err))
	if count > 0 { //誰かいたら
		io.WriteString(w, "そのユーザーネームは既に登録されています") //エラーメッセージを返す
		return
	}
	byt := sha256.Sum256([]byte(r.Form.Get("pass")))                                                                           //ハッシュ化
	bina := byt[:]                                                                                                             //おまじない
	hash := hex.EncodeToString(bina)                                                                                           //16進数の文字列に変化
	_, err = db.Exec("INSERT INTO userdata_login VALUES(DEFAULT,$1,$2,$3)", r.Form.Get("name"), r.Form.Get("loginname"), hash) //新しいレコードを追加
	checkErr(err)
	io.WriteString(w, "success") //string化したものを送信
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
