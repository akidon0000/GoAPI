
package main

import (
	"fmt"
	"net/http"
	"github.com/labstack/echo"
	// "github.com/labstack/echo/middleware"

	"os"
	"github.com/joho/godotenv"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// type User struct {
// 	gorm.Model
// 	Uuid                string
// 	My_association      string
// 	Partner_association string
// 	Quadkey             string
// 	Status              int
// }

type User struct {
	gorm.Model
	Uuid                string `json:"uuid"`
}

// func (u User) String() string {
// 	return fmt.Sprintf("Uuid:%s \n My_association:%s \n Partner_association:%s \n Quadkey:%s \n Status:%d \n \n",
// 		u.Uuid,
// 		u.My_association,
// 		u.Partner_association,
// 		u.Quadkey,
// 		u.Status)
// }

func (u User) String() string {
	return fmt.Sprintf("Uuid:%s \n",u.Uuid)
}

func insert(users []User, db *gorm.DB) {
	for _, user := range users {
			db.NewRecord(user)
			db.Create(&user)
	}
}


func main() {
	e := echo.New()
	// e.Use(middleware.Logger())
  // e.Use(middleware.Recover())
	// e.Use(middleware.BodyDump(bodyDumpHandler))
	e.POST("/user",baseAPI_POSTUser())
	e.POST("/hello", HandleHelloPost)
	e.POST("/api/hello", HandleAPIHelloPost)



	// e.Logger.Fatal(e.Start(":1323"))
	e.Start(":8080")

}
func MainPage() echo.HandlerFunc {
	return func(c echo.Context) error {     //c をいじって Request, Responseを色々する
			return c.String(http.StatusOK, "Hello World")
	}
}

func bodyDumpHandler(c echo.Context, reqBody, resBody []byte) {
  fmt.Printf("Request Body: %v\n", string(reqBody))
  fmt.Printf("Response Body: %v\n", string(resBody))
}

func HandleHelloPost(c echo.Context) error {
	greetingto := c.FormValue("greetingto")
	return c.Render(http.StatusOK, "hello", greetingto)
}
// HelloParam は /api/hello が受けとるJSONパラメータを定義します。
type HelloParam struct {
	GreetingTo string `json:"greetingto"`
}

// HandleAPIHelloPost は /api/hello のPost時のJSONデータ生成処理を行います。
func HandleAPIHelloPost(c echo.Context) error {
	param := new(HelloParam)

	if err := c.Bind(param); err != nil {
			return err
	}
	fmt.Println(param)
	return c.JSON(http.StatusOK, map[string]interface{}{"hello": param.GreetingTo})
}

func baseAPI_POSTUser() echo.HandlerFunc{
	return func(c echo.Context) error {

		var jsonMap map[string]interface{} = make(map[string]interface{})
		var errors = make([]map[string]interface{}, 0)
		var httpStatus = 200

		db := gormConnect()
		defer db.Close()
		// fmt.Print(c.FormValue("Uuid"))
		// fmt.Print("aa")
		uuid := c.FormValue("Uuid")
		// gree := c.FormValue("uuid")
		gree := new(User)
		if err := c.Bind(gree); err != nil {
			return err
		}

		fmt.Println(uuid)
		fmt.Println(gree.Uuid)
		// myAssoci := c.FormValue("My_association")
		// parAssoci := c.FormValue("Partner_association")
		// quadkey := c.FormValue("Quadkey")
		user1 := User{Uuid: uuid}

		// user1 := User{
		// 					Uuid: uuid,
		// 					My_association: myAssoci,
		// 					Partner_association: parAssoci,
		// 					Quadkey: quadkey,
		// 					Status: 1,}

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

// func findAll(db *gorm.DB) []User {
// 	var allUsers []User
// 	db.Find(&allUsers)
// 	return allUsers
// }



// func findByID(db *gorm.DB, id int) User {
// 	var user User
// 	db.First(&user, id)
// 	return user
// }
