package controller

import (
	"chuxin0816/SE/common"
	"chuxin0816/SE/models"
	"chuxin0816/SE/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")
	if len(telephone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422, "msg": "手机号必须为11位",
		})
		return
	}
	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422, "msg": "密码不能少于6位",
		})
		return
	}
	if len(name) == 0 {
		name = util.RandomString(10)
	}
	if isTelephoneExist(telephone) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422, "msg": "用户已存在",
		})
		return
	}
	user := models.User{
		Name:      name,
		Telephone: telephone,
		Password:  password,
	}
	err:=models.Register(user)
	if err!=nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "注册成功",
	})
}

func isTelephoneExist(telephone string) bool {
	var user models.User
	common.DB.Where("telephone = ?", telephone).First(&user)
	return user.ID != 0
}
