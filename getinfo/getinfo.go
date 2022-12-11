package getinfo

import (
	"database/sql"  //SQL操作
	"encoding/json" //json相互変換
	"fmt"           //基本的な入出力処理
	"io"            //入出力
	"net/http"      //http通信

	"github.com/Fuses-Garage/UnityGo/util" //自作パッケージ
	_ "github.com/lib/pq"                  //ポスグレ使用
)

func GetInfo(w http.ResponseWriter, r *http.Request) {
	db := util.LoginToDB()
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
	var datarows []idata                        //データを格納する配列
	rows, err := db.Query("SELECT * FROM info") //全取得
	util.CheckErr(err)
	for rows.Next() { //1つずつ処理
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
			util.CheckErr(err)
		}
	}
	jsoninfo, _ := json.Marshal(datarows) //jsonに変換
	io.WriteString(w, string(jsoninfo))   //string化したものを送信
}
