package my

import (
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Migrate program.
func Migrate() {
	CONNECT := "root:secret@tcp(localhost:3306)/mysql_go?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	db, er := gorm.Open("mysql", CONNECT)
	if er != nil {
		fmt.Println(er)
		return
	}
	defer db.Close()

	db.AutoMigrate(&User{}, &Group{}, &Post{}, &Comment{})

}