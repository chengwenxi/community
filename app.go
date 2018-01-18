package main

import (
	"github.com/gin-gonic/gin"
	"github.com/community/rest"
	"github.com/community/models"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"io"
	"os"
	"log"
)

func main() {

	//log
	r := gin.New()
	f, _ := os.Create("app.log")
	gin.DefaultWriter = io.MultiWriter(f)
	r.Use(gin.Logger())
	log.SetOutput(gin.DefaultWriter) // You may need this

	//db
	models.InitDB()

	//rest
	rest.UserRegisterAll(r.Group("/user"))//user


	r.Run() // listen and serve on 0.0.0.0:8080
	log.Println("server start")
}

