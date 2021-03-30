
package main

import (
	"fmt"
	// "log"
	"net/http"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	// jsonデータを作成
	// "encoding/json"
	// "gen"
	// "github.com/labstack/echo"

	// _ "github.com/go-sql-driver/mysql"
	// Json形式で返す
	// "github.com/gorilla/mux"

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



func findByID(db *gorm.DB, id int) User {
	var user User
	db.First(&user, id)
	return user
}

func insert(users []User, db *gorm.DB) {
	for _, user := range users {
			db.NewRecord(user)
			db.Create(&user)
	}
}


func main() {
	// DBに接続
	// db := gormConnect()
	// // main関数が終わる際にDBの接続を切る
	// defer db.Close()
	// fmt.Println(findAll(db))

  // user1 := User{
	// 					Uuid: "12",
	// 					My_association: "test1",
	// 					Partner_association: "test1",
	// 					Quadkey: "1234",
	// 					Status: 1,}
  // insertUsers := []User{user1}
  // insert(insertUsers, db)
	e := echo.New()
	e.Use(middleware.Logger())
  e.Use(middleware.Recover())
	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, World!")})
	// e.GET("/user/all",getUsers)
	e.POST("/user",baseAPI_POSTUser())
	e.GET("/user",baseAPI_GETUser)


	// e.Logger.Fatal(e.Start(":1323"))
	fmt.Println("サーバー始動")
	e.Start(":8080")

	// r := mux.NewRouter()
	// // localhost:8080一覧を取得
	// r.HandleFunc("/v1/user/all", getUsers).Methods("GET")
	// r.HandleFunc("/v1/user/{id}", getUser).Methods("GET")
	// r.HandleFunc("/v1/user", postUser).Methods("POST")
	// // r.HandleFunc("/opening/", showOpeningIndex)
	// log.Fatal(http.ListenAndServe(":8080", r))
	// db.Set("gorm:table_options", "ENGINE = InnoDB").AutoMigrate(&User{})
}
func MainPage() echo.HandlerFunc {
	return func(c echo.Context) error {     //c をいじって Request, Responseを色々する
			return c.String(http.StatusOK, "Hello World")
	}
}

func baseAPI_POSTUser() echo.HandlerFunc{
	return func(c echo.Context) error {
		var jsonMap map[string]interface{} = make(map[string]interface{})
		var errors = make([]map[string]interface{}, 0)
		var httpStatus = 200
		db := gormConnect()
		defer db.Close()

		uuid := c.FormValue("Uuid")
		myAssoci := c.FormValue("My_association")
		parAssoci := c.FormValue("Partner_association")
		quadkey := c.FormValue("Quadkey")

		user1 := User{
							Uuid: uuid,
							My_association: myAssoci,
							Partner_association: parAssoci,
							Quadkey: quadkey,
							Status: 1,}

		if uuid == ""{
			errors = append(errors, map[string]interface{}{
				"status":  400,
				"param":   "username",
				"message": "invalid",
			})
			jsonMap["errors"] = errors
			httpStatus = 400
		}else{
			insertUsers := []User{user1}
			insert(insertUsers, db)
			errors = append(errors, map[string]interface{}{
				"status":  200,
				"param":   "OK",
				"message": "OK",
			})
			jsonMap["errors"] = errors
			httpStatus = 200
		}

		return c.JSON(httpStatus, jsonMap)
	}
}

func baseAPI_GETUser(c echo.Context) error {
	db := gormConnect()
	defer db.Close()
	var alUsers []User
	db.Find(&alUsers)
	fmt.Println(alUsers)
	allUsers := &alUsers
	if err := c.Bind(allUsers); err != nil {
			return err
		}
	return c.JSON(http.StatusOK, allUsers)
}

// func getUsers(c echo.Context) error {
// 		db := gormConnect()
// 		defer db.Close()
// 		var allUsers []*User
// 		db.Find(&allUsers)
// 		// allUsers := &alUsers
// 		// allUsers := json.NewEncoder(w).Encode(allUsers)

// 		// allUsers := &User{
// 		// 						Uuid: "12",
// 		// 						My_association: "test1",
// 		// 						Partner_association: "test1",
// 		// 						Quadkey: "1234",
// 		// 						Status: 1,
// 		// }
// 		if err := c.Bind(allUsers); err != nil {
// 				return err
// 			}
// 		return c.JSON(http.StatusOK, allUsers)
// }

// Get All User
// func getUsers(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	db := gormConnect()
// 	defer db.Close()
// 	var allUsers []User
// 	db.Find(&allUsers)

// 	json.NewEncoder(w).Encode(allUsers)
// }
// // Get Single Book
// func getUser(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	db := gormConnect()
// 	defer db.Close()

// 	params := mux.Vars(r)

// 	var user User
// 	db.Find(&user, params["id"])

// 	json.NewEncoder(w).Encode(user)
// }

// // post Single Book
// func postUser(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	db := gormConnect()
// 	defer db.Close()

// 	var user User
// 	_ = json.NewDecoder(r.Body).Decode(&user)
// 	// user1 := User{
// 	// 				Uuid: "12",
// 	// 				My_association: "test1",
// 	// 				Partner_association: "test1",
// 	// 				Quadkey: "1234",
// 	// 				Status: 1,}
// 	// insertUsers := []User{user1}
// 	// insert(insertUsers, db)


// 	json.NewEncoder(w).Encode(user)



// 	// book.ID = strconv.Itoa(rand.Intn(10000)) // Mock ID - not safe in production
// 	// books = append(books, book)
// 	// json.NewEncoder(w).Encode(book)
// }



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
