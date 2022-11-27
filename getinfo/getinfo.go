package getinfo

import (
	"database/sql"
	"encoding/json"
	"fmt" //基本的な入出力処理
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func GetInfo(w http.ResponseWriter, r *http.Request) {

	err := godotenv.Load(".vscode/.env")

	// もし err がnilではないなら、"読み込み出来ませんでした"が出力されます。
	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	}
	db, err := sql.Open("postgres", "user="+os.Getenv("ENVUSER")+" password="+os.Getenv("ENVPASS")+" dbname=UniGoDB sslmode=disable")
	checkErr(err)
	//データの検索
	type idata struct {
		ID    int
		TITLE string
		ABOUT string
		DATE  string
	}
	var id int
	var title string
	var about string
	var date string
	var datarows []idata
	rows, err := db.Query("SELECT * FROM public.info")
	for rows.Next() {
		switch err := rows.Scan(&id, &title, &about, &date); err {
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
		//fmt.Fprintf(w, "%d,%s,%s\n", id, title, about)
	}
	checkErr(err)
	jsoninfo, _ := json.Marshal(datarows)
	checkErr(err)

	io.WriteString(w, string(jsoninfo))
}
func checkErr(err error) {
	if err != nil {
		//panic(err)
	}
}
