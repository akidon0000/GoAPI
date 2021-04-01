
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

type User struct {
	gorm.Model
	Uuid                string `json:"uuid"`
	My_association      string `json:"myAssociation"`
	Partner_association string `json:"partnerAssociation"`
	Quadkey             string `json:"quadkey"`
	Status              int    `json:"status"`
}


func (u User) String() string {
	return fmt.Sprintf("Uuid:%s \n My_association:%s \n Partner_association:%s \n Quadkey:%s \n Status:%d \n \n",
		u.Uuid,
		u.My_association,
		u.Partner_association,
		u.Quadkey,
		u.Status)
}


func main() {
	e := echo.New()
	// e.Use(middleware.Logger())
  // e.Use(middleware.Recover())
	// e.Use(middleware.BodyDump(bodyDumpHandler))
	e.POST("/user",baseAPI_user())
	e.POST("/affinity",baseAPI_affinity())

	// e.Logger.Fatal(e.Start(":1323"))
	e.Start(":8080")
}

func MainPage() echo.HandlerFunc {
	return func(c echo.Context) error {     //c をいじって Request, Responseを色々する
			return c.String(http.StatusOK, "Hello World")
	}
}

func insert(users []User, db *gorm.DB) {
	for _, user := range users {
			db.NewRecord(user)
			db.Create(&user)
	}
}

func updateWhereID(users User, db *gorm.DB) {
	var user User
	db.Model(&user).Where("uuid = ?", users.Uuid).Update("my_association", users.My_association)
	db.Model(&user).Where("uuid = ?", users.Uuid).Update("partner_association",users.Partner_association)
	db.Model(&user).Where("uuid = ?", users.Uuid).Update("quadkey",users.Quadkey)
	db.Model(&user).Where("uuid = ?", users.Uuid).Update("status",users.Status)
}

func search(partner string, db *gorm.DB) (User){
	var user User
	db.Raw("SELECT * FROM users WHERE my_association = ?", partner).Scan(&user)
	return user
}

func baseAPI_user() echo.HandlerFunc{
	return func(c echo.Context) error {
		db := gormConnect()
		defer db.Close()

		var jsonMap map[string]interface{} = make(map[string]interface{})
		var errors = make([]map[string]interface{}, 0)
		var httpStatus = 200

		// gree := c.FormValue("uuid")
		user := new(User)
		if err := c.Bind(user); err != nil {
			return err
		}

		fmt.Println(user)


		user1 := User{
							Uuid: user.Uuid,
							My_association: user.My_association,
							Partner_association: user.Partner_association,
							Quadkey: user.Quadkey,
							Status: user.Status,}

		if user.Uuid == ""{
			insertUsers := []User{user1}
			insert(insertUsers, db)
			errors = append(errors, map[string]interface{}{
				"status":  200,
				"param":   "OK",
				"message": "追加しました",
			})
			jsonMap["sucsess"] = errors
			httpStatus = 200
		}else{
			// insertUsers := []User{user1}

			// insert(insertUsers, db)
			updateWhereID(user1, db)
			errors = append(errors, map[string]interface{}{
				"status":  200,
				"param":   "OK",
				"message": "更新しました",
			})
			jsonMap["sucsess"] = errors
			httpStatus = 200
		}

		return c.JSON(httpStatus, jsonMap)
	}
}



func baseAPI_affinity() echo.HandlerFunc{
	return func(c echo.Context) error {
		db := gormConnect()
		defer db.Close()

		// var jsonMap map[string]interface{} = make(map[string]interface{})
		// var errors = make([]map[string]interface{}, 0)
		var httpStatus = 200

		user := new(User)
		if err := c.Bind(user); err != nil {
			return err
		}

		fmt.Println(user)

		// var newuser User

		var newuser = search(user.Partner_association, db)

		// if user.Uuid == ""{
		// 	errors = append(errors, map[string]interface{}{
		// 		"status":  400,
		// 		"param":   "username",
		// 		"message": "invalid",
		// 	})
		// 	jsonMap["errors"] = errors
		// 	httpStatus = 400
		// }else{
		// 	// insertUsers := []User{user1}
		// 	// insert(insertUsers, db)
		// 	errors = append(errors, map[string]interface{}{
		// 		"status":  200,
		// 		"param":   "OK",
		// 		"message": "OK",
		// 	})
		// 	jsonMap["sucsess"] = errors
		// 	httpStatus = 200
		// }
		// return c.JSON(httpStatus, jsonMap)
		return c.JSON(httpStatus, newuser)
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

// func baseAPI_affinity(db *gorm.DB, id int) User {
// 	var user User
// 	db.First(&user, id)
// 	return user
// }


