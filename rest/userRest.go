package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/community/models"
	"strconv"
)

func RegisterAll(g *gin.RouterGroup) {
	g.GET("", List)
	g.POST("", Create)
	g.PUT("", Update)
	g.DELETE("", Delete)
}

func List(c *gin.Context) {
	skip, err := strconv.Atoi(c.Query("skip"))
	limit, err := strconv.Atoi(c.Query("limit"))
	user, err := models.UserList(skip, limit)
	if err != nil {

	}
	c.JSON(200, user)
}

func Find(c *gin.Context) {

}

func Create(c *gin.Context) {
	email := c.PostForm("email")
	telephone := c.PostForm("telephone")
	user := &models.User{
		Email:     email,
		Telephone: telephone,
	}
	if len(user.Email) == 0 || len(user.Email) == 0 {
		c.JSON(400, gin.H{
			"message": "invalid parameter",
		})
		return
	}
	c.JSON(200, user.Create())
}

func Delete(c *gin.Context) {

}

func Update(c *gin.Context) {

}
