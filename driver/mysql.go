package driver

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"movie/model"
)

var DB *gorm.DB

func InitDB() (err error) {
	dsn := "root:sdf@df%%$65#fdsbXT@tcp(127.0.0.1:3306)/movie?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {

		return err
	}
	DB = db
	return nil
}

func init()  {
	InitDB()
	DB.AutoMigrate(&model.User{})

}



