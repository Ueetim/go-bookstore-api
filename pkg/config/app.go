package config

import (
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
)

// db variable
var (
	db * gorm.DB
)

// OPEN CONNECTION WITH DB
func Connect() {

}