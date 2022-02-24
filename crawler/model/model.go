package model

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB       *gorm.DB
	err      error
	username string = "root"
	password string = "root"
	dbName   string = "crawler_core"
)

func init() {
	// DB, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", username, password, dbName))
	DB, err = gorm.Open(
		"mysql", 
		fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", username, password, dbName),
	)

	if err != nil {
		log.Fatalf("gorm.Open.err: %v", err)
	} else {
		log.Println("gorm.Open.success")
	}

	DB.SingularTable(true)
	// gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	// 	return "crawler_" + defaultTableName
	// }
}
