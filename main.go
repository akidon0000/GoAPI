
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

type apiHandler struct {
}

type User struct {
	MyAssociation      string
	PartnerAssociation string
	Quadkey            string
	Status             int
}

func (u User) String() string {
	return fmt.Sprintf("%s(%d)[%s]", u.MyAssociation, u.PartnerAssociation, u.Quadkey)
}


func main() {
	// DBに接続
	db := gormConnect()
	// main関数が終わる際にDBの接続を切る
	defer db.Close()
	// テーブル名の複数形化を無効化します。trueにすると`User`のテーブル名は`user`になります
	db.SingularTable(true)

	// db.NewRecord(User{Name:"AA", Age:10})
	// db.NewRecord(User{Name:"BB", Age:20})
	// // 構造体のインスタンス化
	// user := User{}
	// // 挿入したい情報を構造体に与える
	// user.MyAssociation = "AAA"
	// user.PartnerAssociation = "BBB"
	// // INSERTを実行
	// db.Create(&user)

	// e := echo.New()
	// handler := apiHandler{}
	// // 定義した struct を登録
	// openapi.RegisterHandlers(e, handler)
	// //e.Logger.Fatal(e.Start(fmt.Sprintf("0.0.0.0:%d", *port)))
	// e.Logger.Fatal(e.Start(":1323"))
}

// SQLConnect DB接続
func gormConnect() (database *gorm.DB) {
	err := godotenv.Load(fmt.Sprintf("../%s.env", os.Getenv("GO_ENV")))
    if err != nil {
			panic(err.Error())
        // .env読めなかった場合の処理
    }
	DBMS := "mysql"
	PROTOCOL := "tcp(localhost:3306)" // db:3306 <<< docker-compose.ymlで定義したMySQLのサービス名:ポート
	USER := os.Getenv("DB_ROLE")       // データベース
	PASS := os.Getenv("DB_PASSWORD")
	DBNAME := os.Getenv("DB_NAME")
	fmt.Println(DBNAME)
	fmt.Println("成功")

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME //+ "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("DB接続成功")
	}
	return db
}





// // Userを追加する。ログイン認証なし、idとtokenをkeyChainに保存
// // (POST /user)
// func (si apiHandler) AddUser(ctx echo.Context) error {
// 	//	b := NewUser{}
// 	// リクエスト Body は echo の APIを利用
// 	//	ctx.Bind(&b)

// 	return ctx.String(http.StatusOK, "こんにちは! AddUser")
// }




// // Userを更新する
// // (PUT /user)
// func (si apiHandler) UpdateUser(ctx echo.Context, params openapi.UpdateUserParams) error {
// 	// ビジネスロジック
// 	// var err error
// 	// return err
// 	return ctx.String(http.StatusOK, "UpdateUser")
// }

// // Userへメッセージを送る
// // (POST /user/message)
// func (si apiHandler) PostUserMessage(ctx echo.Context, params openapi.PostUserMessageParams) error {
// 	// ビジネスロジック
// 	var err error
// 	return err
// }

// // Userを取得する
// // (GET /user/{userId})
// func (si apiHandler) GetUserId(ctx echo.Context, userId int64, params openapi.GetUserIdParams) error {
// 	// ビジネスロジック
// 	var err error
// 	return err
// }

// // Userの相性を取得する
// // (GET /user/{userId}/affinity)
// func (si apiHandler) GetUserIdUnbrella(ctx echo.Context, userId int64, params openapi.GetUserIdUnbrellaParams) error {
// 	// ビジネスロジック
// 	var err error
// 	return err
// }
