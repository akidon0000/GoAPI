
package main
// https://qiita.com/fukumone/items/0313004d60ddb4d92d55

import (
	"fmt"
	"net/http"
	"log"
	// jsonデータを作成
	"encoding/json"
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
	uuid                         string `json:"uuid"`
	my_association        string `json:"my_association"`
	partner_association string `json:"partner_association"`
	quadkey                  string `json:"quadkey"`
	status                      int `json:"status"`
}

// type User struct {
// 	uuid                         string `gorm:"uuid"`
// 	my_association        string `gorm:"my_association"`
// 	partner_association string `gorm:"partner_association"`
// 	quadkey                  string `gorm:"quadkey"`
// 	status                      int `gorm:"status"`
// }

// var UserEX []User

func main() {
	// DBに接続
	// db := gormConnect()
	// // main関数が終わる際にDBの接続を切る
	// defer db.Close()


	r := mux.NewRouter()
	// localhost:8080一覧を取得
	r.HandleFunc("/v1/user/all", getUser).Methods("GET")
	// r.HandleFunc("/opening/", showOpeningIndex)
	log.Fatal(http.ListenAndServe(":8080", r))

	// UserEX = append(UserEX, User{ my_association: "matsuyama", partner_association: "akihiro", quadkey: "1234", status: 1})


	// テーブル名の複数形化を無効化します。trueにすると`User`のテーブル名は`UserEX`になります
	// db.SingularTable(true)
}

// Get All Books
func getUser(w http.ResponseWriter, r *http.Request) {
	// DBに接続
	db := gormConnect()
	// db.SingularTable(true)
	// main関数が終わる際にDBの接続を切る
	defer db.Close()

	// var result User
	UserEx := User{}
	// UserEx.uuid = "123"
	db.First(&UserEx)


	// db.Row("SELECT quadkey FROM User").Scan(&UserEX)
	// fmt.Println(UserEX)
	// rows, err := db.Row("SELECT * FROM testlist")
	// if err != nil {
	// 		fmt.Println("Err2")
	// 		panic(err.Error())
	// }
	//データを格納する変数を定義
	//複数レコード
	// UsersEx := []User{}
	// var User User
	//全取得
	// db.Find(&UsersEx)
	// db.First(&User)
	// fmt.Println("-")
	// fmt.Printf(UserEx[0].my_association)
	// fmt.Println("-")
	//表示
	// for _, emp := range UsersEx {
	// 	// fmt.Println("%t\n", emp.uuid)
	// 	// fmt.Println("%t\n", true)
	// 	fmt.Sprint(emp.uuid)
	// 	fmt.Println("1")
	// 	fmt.Println(emp.my_association)
	// 	fmt.Println("2")
	// 	fmt.Println(emp.partner_association)
	// 	fmt.Println("3")
	// 	fmt.Println(emp.quadkey)
	// 	fmt.Println("4")
	// 	fmt.Println(emp.status)
	// }


	w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(UserEx)
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
