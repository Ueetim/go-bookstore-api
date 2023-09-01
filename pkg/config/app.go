// FOR DATABASE CONNECTION

package config

import (
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
)

// db variable
var (
	db *gorm.DB
)

// OPEN CONNECTION WITH DB
func Connect() {
	d, err := gorm.Open("mysql", "uduak:quixote4@/bookstore?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}