package database

import(
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/sqlite"

)

var(
	DbConn *gorm.DB // Helps gollang to interact with the database
)