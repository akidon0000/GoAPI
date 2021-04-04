package routing

import (
	"fmt"
	"math"
	"strconv"
	"github.com/labstack/echo"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"goAPI/databases"
	// "goAPI/algorithm"
	"github.com/masatana/go-textdistance"

	"github.com/google/uuid"
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
	Uuid     string
	Status   string
	Affinity string
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
func BaseAPI_user() echo.HandlerFunc {
	return func(c echo.Context) error {
		db := databases.GormConnect()
		defer db.Close()
		var result Result

		//追加・更新
		user := new(User)
		if err := c.Bind(user); err != nil {
			return err
		}
		var uu = user.Uuid

		if user.Uuid == "" {
			// uuid生成
			u, err := uuid.NewRandom()
			if err != nil {
				fmt.Println(err)
				// return
			}
			uu = u.String()
		}

		user1 := User{
			Uuid:                uu,
			My_association:      user.My_association,
			Partner_association: user.Partner_association,
			Quadkey:             user.Quadkey,
			Status:              user.Status}

		if user.Uuid == "" {
			insertUsers := []User{user1}
			insert(insertUsers, db)

			result.Uuid = uu
		} else {
			result.Uuid = user.Uuid
			update(user1, db)
		}

		result.Status = "0"

		// 相性取得
		var affinity = search(user.Partner_association, user.My_association, user.Quadkey, user.Status, db)

		result.Affinity = strconv.FormatFloat(affinity, 'f', 2, 64)
		// strconv.Itoa(affinity)

		// return c.JSON(200, result)
		return c.JSON(200, result)
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

//
func search(partner string, my string, quadkey string, status int, db *gorm.DB) float64 {
	var user []User
	db.Raw("SELECT * FROM users WHERE quadkey = ? AND status = ?", quadkey, status).Scan(&user)

	affinity := 0.0
	for i,s := range user{
		num := 0.0
		fmt.Println("No:",i)
		fmt.Println("自分：",partner, "相手：",s.My_association)
		fmt.Println("相性", textdistance.JaroWinklerDistance(partner, s.My_association))
		fmt.Println("自分：",my,"相手：", s.Partner_association)
		fmt.Println("相性", textdistance.JaroWinklerDistance(my, s.Partner_association))
		num += textdistance.JaroWinklerDistance(partner, s.My_association)
		num += textdistance.JaroWinklerDistance(my, s.Partner_association)
		num /= 2
		fmt.Println("総合相性", num)
		affinity = math.Max(affinity,num)
		fmt.Println(" ")
	}
	fmt.Println("Max総合相性", strconv.FormatFloat(affinity, 'f', 2, 64))

	return affinity
}

// func search(partner string,my string,quadkey string,status int, db *gorm.DB) ([]User){
// 	var count int
// 	var user []User
// 	db.Raw("SELECT * FROM users WHERE quadkey = ? AND status = ?", quadkey, status).Scan(&user)

// 	db.Model(&User{}).Where("my_association = ? AND partner_association = ? AND quadkey = ? AND status = ?",partner,my,quadkey,status).Count(&count)

// 	fmt.Println("完全一致検索件数：" , count)
// 	return count
// }
