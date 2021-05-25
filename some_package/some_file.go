package some_package

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)
// Solve the parameter structure passed by the request
type User struct {
	UserName string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Age      int    `form:"age" json:"age" binding:"required"`
}
// This is the processing function for login
func LoginHandler(c *gin.Context) {
	req := &User{}
	if err := c.Bind(req); err != nil {
		log.Printf("err:%v", err)
		c.JSON(http.StatusOK, gin.H{
			"errno":  "1",
			"errmsg": err.Error(),
		})
		return
	}

	// judge the password and username
	if req.UserName != "Valiben" || req.Password != "123456" {
		c.JSON(http.StatusOK, gin.H{
			"errno":  "2",
			"errmsg": "password or username is wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"errno":  "0",
		"errmsg": "login success",
	})
}