package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// username:password@tcp(serveraddress:port)/dbname?charset=utf8&parseTime=true&loc=local
//const connStr = "root:rootmac098@tcp(localhost:3306)/student?charset=utf8&parseTime=true&loc=Local"

// DB is a connection variable.
var DB *gorm.DB

// Init initiates db
func Init(connStr string) {
	str := os.Getenv("db_connection_str")
	if str == "" {
		str = connStr
	}
	db, err := gorm.Open("mysql", str)
	if err != nil {
		fmt.Println("Error connecting to DB", err)
	}
	db.LogMode(true)
	DB = db
	// fmt.Println(db)
}

// GetInstance retruns instanc eof DB
func GetInstance() *gorm.DB {
	return DB
}
