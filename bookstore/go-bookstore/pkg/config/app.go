package config

import(
	"fmt"
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
)

var 
db * gorm.DB


func Connect(){

	dsn := "yami:-camuenga!-@tcp(127.0.0.1:3306)/simplerest?charset=utf8mb4&parseTime=True&loc=Local"	
	var err error
	db, err = gorm.Open("mysql", dsn)
	//d, err := gorm.Open("mysql","yami:-camuenga!-@tcp(127.0.0.1:3306)/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err != nil{
		panic(err)
	}
	fmt.Println("Database connected...")
	
	sqlDB := db.DB()
    sqlDB.SetMaxIdleConns(10)
    sqlDB.SetMaxOpenConns(100)

}

func GetDB() *gorm.DB{
	return db
}