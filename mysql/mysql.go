package mysql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Db *gorm.DB

func init()  {
	var err error
	Db,err = gorm.Open("mysql","root:root@tcp(127.0.0.1:3306)/gotest")
	if err != nil{
		panic(err.Error())
	}
}
