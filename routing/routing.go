package routing

import (
	"fmt"
	"github.com/labstack/echo"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"goAPI/databases"
)

type User struct {
	gorm.Model
	Uuid                string `json:"uuid"`
	My_association      string `json:"myAssociation"`
	Partner_association string `json:"partnerAssociation"`
	Quadkey             string `json:"quadkey"`
	Status              int    `json:"status"`
}

type Result struct {
	Status              string
	Affinity            string
}

func (u User) String() string {
	return fmt.Sprintf("Uuid:%s \n My_association:%s \n Partner_association:%s \n Quadkey:%s \n Status:%d \n ",
		u.Uuid,
		u.My_association,
		u.Partner_association,
		u.Quadkey,
		u.Status)
}

// ユーザーを登録，更新
func BaseAPI_user() echo.HandlerFunc{
	return func(c echo.Context) error {
		db := databases.GormConnect()
		defer db.Close()

		user := new(User)
		if err := c.Bind(user); err != nil {
			return err
		}

		user1 := User{
							Uuid: user.Uuid,
							My_association: user.My_association,
							Partner_association: user.Partner_association,
							Quadkey: user.Quadkey,
							Status: user.Status,}

		if user.Uuid == ""{
			insertUsers := []User{user1}
			insert(insertUsers, db)
		}else{
			update(user1, db)
		}

		var newuser = search(user.Partner_association, user.My_association ,user.Quadkey, db)

		var result Result
		result.Status = "0"
		if newuser.Uuid == ""{
			// 一致する条件が見当たらなかった場合
			result.Affinity = "0"
		}else{
			// 完全一致が見つかった場合
			result.Affinity = "100"
		}

		return c.JSON(200, result)
	}
}

// 相性を取得
func BaseAPI_affinity() echo.HandlerFunc{
	return func(c echo.Context) error {
		db := databases.GormConnect()
		defer db.Close()

		user := new(User)
		if err := c.Bind(user); err != nil {
			return err
		}
		var newuser = search(user.Partner_association, user.My_association ,user.Quadkey, db)
		var httpStatus = 200
		return c.JSON(httpStatus, newuser)
	}
}

func insert(users []User, db *gorm.DB) {
	for _, user := range users {
			db.NewRecord(user)
			db.Create(&user)
	}
}

func update(users User, db *gorm.DB) {
	var user User
	db.Model(&user).Where("uuid = ?", users.Uuid).Update(map[string]interface{}{"my_association": users.My_association, "partner_association": users.Partner_association, "quadkey": users.Quadkey, "status": users.Status})
}

func search(partner string,my string,quadkey string, db *gorm.DB) (User){
	var user User
	db.Raw("SELECT * FROM users WHERE my_association = ? AND partner_association = ? AND quadkey = ? ", partner, my ,quadkey).Scan(&user)
	return user
}
