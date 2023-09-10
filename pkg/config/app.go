package config

import(
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect(){
	d, err := gorm.Open("mysql", "root:Mstunter@1/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err != nil{
		panic(err)
	}
	db = db
}

func GetDB() * =gorm.DB{
	return db
}