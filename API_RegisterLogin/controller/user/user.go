package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"hitgub.com/Mickeythitiwut/Api_registerlogin/orm"
)

func ReadAll(c *gin.Context) {
	var users []orm.User
	orm.Db.Find(&users)
	c.JSON(http.StatusOK, gin.H{
		"status ": "Ok",
		"message": "User Read Success", "users": users})

}

func Profile(c *gin.Context) {
	userId := c.MustGet("userId").(float64)
	var user []orm.User
	orm.Db.First(&user, userId)
	c.JSON(http.StatusOK, gin.H{
		"status ": "Ok",
		"message": "User Read Success", "users": user})

}
