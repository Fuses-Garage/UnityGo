package util

import (
	"database/sql" //SQL操作
	"os"

	"github.com/joho/godotenv" //.envの使用
	_ "github.com/lib/pq"      //ポスグレ使用
)

func LoginToDB() *sql.DB {
	err := godotenv.Load(".vscode/.env")
	CheckErr(err)
	db, err := sql.Open("postgres", "user="+os.Getenv("ENVUSER")+" password="+os.Getenv("ENVPASS")+" dbname=UniGoDB sslmode=disable") //DBに接続
	CheckErr(err)
	return db
}
