package main

import (
	"github.com/gin-gonic/gin"
	"github.com/community/rest"
	"github.com/community/models"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"io"
	"os"
	"log"
	"github.com/community/config"
)

func main() {

	//init config
	config.LoadConfiguration("./config.yml")

	r := gin.New()

	//log
	f, _ := os.Create("app.log")
	gin.DefaultWriter = io.MultiWriter(f)
	r.Use(gin.Logger())
	log.SetOutput(gin.DefaultWriter) // You may need this

	//authorizer
	//e := casbin.NewEnforcer("authz_model.conf", "authz_policy.csv")
	//r.Use(authz.NewAuthorizer(e))

	//static source
	r.Static("/static", config.Config.StaticPath)

	//db
	models.InitDB()

	//rest
	rest.UserRegisterAll(r.Group("/user")) //user

	r.Run(config.Config.Server) // listen and serve on 0.0.0.0:8080
	log.Println("server start")
}
