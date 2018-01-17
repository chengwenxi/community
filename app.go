package main

import (
	"github.com/gin-gonic/gin"
	rest "github.com/community/rest"
	"github.com/community/models"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)



func main() {
	r := gin.Default()
	models.InitDB()
	//user
	rest.RegisterAll(r.Group("/user"))

	r.Run() // listen and serve on 0.0.0.0:8080
}

