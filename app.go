package main

import (
	"github.com/gin-gonic/gin"
	rest "github.com/community/rest"
	"github.com/community/models"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"io"
	"os"
	"log"
)



func main() {
	r := gin.New()
	models.InitDB()
	f, _ := os.Create("app.log")
	gin.DefaultWriter = io.MultiWriter(f)
	r.Use(gin.Logger())
	//user
	rest.RegisterAll(r.Group("/user"))
	log.SetOutput(gin.DefaultWriter) // You may need this
	log.Println("server start")
	r.Run() // listen and serve on 0.0.0.0:8080
}

