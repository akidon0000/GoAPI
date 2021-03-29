
package main

import (
	"fmt"
	"net/http"
	"log"
	// jsonデータを作成
	"encoding/json"
	// "gen"
	// "github.com/labstack/echo"

	// _ "github.com/go-sql-driver/mysql"
	// Json形式で返す
	"github.com/gorilla/mux"

	"os"
	"github.com/joho/godotenv"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Uuid                string
	My_association      string
	Partner_association string
	Quadkey             string
	Status              int
}

func (u User) String() string {
	return fmt.Sprintf("Uuid:%s \n My_association:%s \n Partner_association:%s \n Quadkey:%s \n Status:%d \n \n",
		u.Uuid,
		u.My_association,
		u.Partner_association,
		u.Quadkey,
		u.Status)
}

func findAll(db *gorm.DB) []User {
	var allUsers []User
	db.Find(&allUsers)
	return allUsers
}

func insert(users []User, db *gorm.DB) {
	for _, user := range users {
			db.NewRecord(user)
			db.Create(&user)
	}
}


func main() {
	// DBに接続
	db := gormConnect()
	// main関数が終わる際にDBの接続を切る
	defer db.Close()
	fmt.Println(findAll(db))

  // user1 := User{
	// 					Uuid: "12",
	// 					My_association: "test1",
	// 					Partner_association: "test1",
	// 					Quadkey: "1234",
	// 					Status: 1,}
  // insertUsers := []User{user1}
  // insert(insertUsers, db)


	r := mux.NewRouter()
	// localhost:8080一覧を取得
	r.HandleFunc("/v1/user/all", getUser).Methods("GET")
	// r.HandleFunc("/opening/", showOpeningIndex)
	log.Fatal(http.ListenAndServe(":8080", r))
	db.Set("gorm:table_options", "ENGINE = InnoDB").AutoMigrate(&User{})
}

// Get All User
func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db := gormConnect()
	defer db.Close()
	var allUsers []User
	db.Find(&allUsers)

	json.NewEncoder(w).Encode(allUsers)
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

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("DB接続成功")
	}
	return db
}
