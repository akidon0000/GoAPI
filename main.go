
package main

import (
	"fmt"
	// "net/http"
	// "gen"
	// "github.com/labstack/echo"

	// _ "github.com/go-sql-driver/mysql"
	"os"
	"github.com/joho/godotenv"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


func main() {
	// DBに接続
	db := gormConnect()
	// main関数が終わる際にDBの接続を切る
	defer db.Close()
	// テーブル名の複数形化を無効化します。trueにすると`User`のテーブル名は`user`になります
	db.SingularTable(true)

}


// SQLConnect DB接続
func gormConnect() (database *gorm.DB) {
	err := godotenv.Load(fmt.Sprintf("../%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println(".env読み取り成功")
	}

	DBMS := "mysql"
	PROTOCOL := "tcp(localhost:3306)" // db:3306
	USER := os.Getenv("DB_ROLE")       // データベース
	PASS := os.Getenv("DB_PASSWORD")
	DBNAME := os.Getenv("DB_NAME")

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME //+ "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("DB接続成功")
	}
	return db
}
