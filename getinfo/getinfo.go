package getinfo

import (
	"database/sql"  //SQL操作
	"encoding/json" //json相互変換
	"fmt"           //基本的な入出力処理
	"io"            //入出力
	"net/http"      //http通信
	"os"

	"github.com/joho/godotenv" //.envの使用
	_ "github.com/lib/pq"      //ポスグレ使用
)

func GetInfo(w http.ResponseWriter, r *http.Request) {

	err := godotenv.Load(".vscode/.env")

	// もし err がnilではないなら、"読み込み出来ませんでした"が出力されます。
	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	}
	db, err := sql.Open("postgres", "user="+os.Getenv("ENVUSER")+" password="+os.Getenv("ENVPASS")+" dbname=UniGoDB sslmode=disable") //DBに接続
	checkErr(err)
	//データの検索
	type idata struct { //レコード用の構造体
		ID    int
		TITLE string
		ABOUT string
		DATE  string
	}
	var id int
	var title string
	var about string
	var date string
	var datarows []idata                               //データを格納する配列
	rows, err := db.Query("SELECT * FROM public.info") //全取得
	for rows.Next() {                                  //1つずつ処理
		switch err := rows.Scan(&id, &title, &about, &date); err { //エラーの有無でスイッチ
		case sql.ErrNoRows:
			fmt.Println("No rows were returned")
		case nil:
			// 一行毎に配列を追加
			datarows = append(datarows, idata{
				ID:    id,
				TITLE: title,
				ABOUT: about,
				DATE:  date,
			})
		default:
			checkErr(err)
		}
	}
	checkErr(err)
	jsoninfo, _ := json.Marshal(datarows) //jsonに変換
	io.WriteString(w, string(jsoninfo))   //string化したものを送信
}
func checkErr(err error) {
	if err != nil {
		//panic(err)
	}
}
