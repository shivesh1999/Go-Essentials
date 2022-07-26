package config

import "gorm.io/gorm"

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "root:Shivesh@102/simplerest?charset=utf8&parseTime-True&loc-Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
