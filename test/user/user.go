package user

import (
	"log"

	"gorm.io/gorm"
)

//定义数据库中user表的列
type User struct {
	gorm.Model
	Name string `gorm:"name"`
	Age  int    `gorm:"age"`
}

func PrintUsers(db *gorm.DB) {
	users := make([]*User, 0)
	var count int64
	d := db.Where("").Offset(0).Limit(5).Order("id desc").Find(&users).Offset(-1).Limit(-1).Count(&count)
	if d.Error != nil {
		log.Fatalf("List products error: %v", d.Error)
	}

	log.Printf("totalcount: %d", count)
	for _, u := range users {
		log.Printf("\tName: %s, age: %d\n", u.Name, u.Age)
	}
}
