package models

import "github.com/jinzhu/gorm"

var DB *gorm.DB

func InitDB()  {
	db, err  := gorm.Open("postgres", "host=127.0.0.1 user=com dbname=com sslmode=disable password=com")
	if err != nil {
		print(err)
	}

	//use singular table by default，else table name will add 's'
	//db.SingularTable(true)

	//连接池上限
	db.DB().SetMaxIdleConns(100)
	DB = db

	//defer db.Close()
}
