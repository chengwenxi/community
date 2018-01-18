package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/community/models"
	"strconv"
	"net/http"
)

func UserRegisterAll(g *gin.RouterGroup) {
	g.GET("", List)
	g.GET("/:id", Find)
	g.POST("", Create)
	g.PUT("", Update)
	g.DELETE("/:id", Delete)
}

func List(c *gin.Context) {
	skip, err := strconv.Atoi(c.DefaultQuery("skip", "0"))
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "20"))
	users, err := models.UserList(skip, limit)
	if err != nil {

	}
	c.JSON(200, users)
}

func Find(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user := &models.User{
		Id: id,
	}
	user.First()
	c.JSON(http.StatusOK, user)
}

func Create(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err == nil {
		if len(user.Email) == 0 {
			c.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
			return
		}
		if dbErr := user.Create(); dbErr == nil {
			c.JSON(http.StatusOK, user)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": dbErr.Error()})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user := &models.User{
		Id: id,
	}
	if err := user.Delete(); err == nil {
		c.JSON(http.StatusOK, user)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

func Update(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err == nil {
		if len(user.Email) == 0 {
			c.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
			return
		}
		if dbErr := user.Update(); dbErr == nil {
			c.JSON(http.StatusOK, user)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": dbErr.Error()})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
