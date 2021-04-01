package databases

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// SQLConnect DB接続
func GormConnect() (database *gorm.DB) {
	// パスワード等を.envファイルから読み取る
	// program > go > .env
	err := godotenv.Load(fmt.Sprintf("../%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("env読み取り成功")
	}

	DBMS := "mysql"                   // MySQL
	PROTOCOL := "tcp(localhost:3306)" // db:3306
	DBNAME := os.Getenv("DB_NAME")    // テーブル名
	USER := os.Getenv("DB_ROLE")      // MySQLユーザー名
	PASS := os.Getenv("DB_PASSWORD")  // パスワード

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("DB接続成功")
	}
	return db
}
