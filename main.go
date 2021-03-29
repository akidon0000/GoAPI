
package main
// https://qiita.com/fukumone/items/0313004d60ddb4d92d55

import (
	"fmt"
	"net/http"
	"log"
	// jsonデータを作成
	// "encoding/json"
	// "gen"
	// "github.com/labstack/echo"

	// _ "github.com/go-sql-driver/mysql"
	"os"
	"github.com/joho/godotenv"
	// Json形式で返す
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Uuid                string
	My_association      string `gorm:"size:255"`
	Partner_association string
	Quadkey             string
	Status              int
}

func (u User) String() string {
	return fmt.Sprintf("%s(%d)", u.My_association, u.Status)
}

func insert(users []User, db *gorm.DB) {
	for _, user := range users {
			db.NewRecord(user)
			db.Create(&user)
	}
}

func findAll(db *gorm.DB) []User {
	var allUsers []User
	db.Find(&allUsers)
	return allUsers
}

// var uss User
// Get All User
// func getUser(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	db := gormConnect()
// 	defer db.Close()

	// db.Find(&uss)
	// var utest []User
	// UserEx := User{}
	// // UserEx.uuid = "123"
	// db.First(&UserEx)

	// json.NewEncoder(w).Encode(UserEx.my_association)
  // json.NewEncoder(w).Encode(uss)
// }

func main() {
	// DBに接続
	db := gormConnect()
	// main関数が終わる際にDBの接続を切る
	defer db.Close()
	fmt.Println(findAll(db))
  user1 := User{
						Uuid: "12",
						My_association: "test1",
						Partner_association: "test1",
						Quadkey: "1234",
						Status: 1,}
  insertUsers := []User{user1}
  insert(insertUsers, db)

	// rows := db.Query("")
	// fmt.Println("DB接続成功")



	r := mux.NewRouter()
	// localhost:8080一覧を取得
	// r.HandleFunc("/v1/user/all", getUser).Methods("GET")
	// r.HandleFunc("/opening/", showOpeningIndex)
	log.Fatal(http.ListenAndServe(":8080", r))
	db.Set("gorm:table_options", "ENGINE = InnoDB").AutoMigrate(&User{})
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
