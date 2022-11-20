package main

import (
	"GROM_demo/test/conn"
	"GROM_demo/test/user"
	"log"
)

func main() {
	//操作数据库test
	//获取数据库连接实例
	db := conn.GetConn()
	//自动迁移表结构，当数据库中没有该表结构时，会自动生成，正式环境中不推荐使用
	// db.AutoMigrate(&user.User{})
	var err error
	//增加记录
	log.Println("insert one record")
	err = db.Create(&user.User{Name: "ysl", Age: 24}).Error
	if err != nil {
		log.Fatalf("Create error: %v", err)
	}

	//列出user表中5条以内的数据
	user.PrintUsers(db)

	ysl := &user.User{}
	//查找记录
	log.Println("find one record")
	err = db.Where("name= ?", "ysl").First(&ysl).Error
	if err != nil {
		log.Fatalf("Find user error: %v", err)
	}
	log.Printf("find data: %v", ysl)
	//修改记录，如果不存在该主键，则新增
	log.Println("update one record")
	ysl.Age = 25
	err = db.Save(ysl).Error
	if err != nil {
		log.Fatalf("Update user error: %v", err)
	}
	user.PrintUsers(db)
	//删除记录
	log.Println("delte one record")
	err = db.Where("age= ?", "25").Delete(&user.User{}).Error
	if err != nil {
		log.Fatalf("Delete user error: %v", err)
	}
	user.PrintUsers(db)

}
