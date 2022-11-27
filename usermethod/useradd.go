package useradd

import (
	"database/sql" //SQL操作
	"fmt"          //基本的な入出力処理
	"io"           //入出力
	"net/http"     //http通信
	"os"

	"github.com/joho/godotenv" //.envの使用
	_ "github.com/lib/pq"      //ポスグレ使用
)

func GetInfo(w http.ResponseWriter, r *http.Request) {

	err := godotenv.Load(".vscode/.env")
	checkErr(err)
	err = r.ParseForm()
	checkErr(err)
	// もし err がnilではないなら、"読み込み出来ませんでした"が出力されます。
	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	}
	db, err := sql.Open("postgres", "user="+os.Getenv("ENVUSER")+" password="+os.Getenv("ENVPASS")+" dbname=UniGoDB sslmode=disable") //DBに接続
	checkErr(err)

	row := db.QueryRow("SELECT COUNT(*) AS count FROM userdata_login WHERE loginname=$2", r.Form.Get("loginname"))
	var count int
	err = row.Scan(&count)
	checkErr((err))
	if count > 0 {
		io.WriteString(w, "そのユーザーネームは既に登録されています")
		return
	}
	_, err = db.Exec("INSERT INTO userdata_login VALUES(DEFAULT,$1,$2,$3)", r.Form.Get("name"), r.Form.Get("loginname"), r.Form.Get("passhash"))
	checkErr(err)
	io.WriteString(w, "success") //string化したものを送信
}
func checkErr(err error) {
	if err != nil {
		//panic(err)
	}
}
