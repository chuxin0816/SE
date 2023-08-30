package controller

import (
	"chuxin0816/SE/common"
	"chuxin0816/SE/dto"
	"chuxin0816/SE/models"
	"chuxin0816/SE/response"
	"chuxin0816/SE/util"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(ctx *gin.Context) {
	//获取参数
	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	//数据验证
	if len(telephone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}
	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
		return
	}
	if len(name) == 0 {
		name = util.RandomString(10)
	}

	//判断手机号是否存在
	if isTelephoneExist(telephone) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户已存在")
		return
	}

	//加密密码
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "加密错误")
		return
	}

	//创建用户
	user := models.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hashPassword),
	}
	err = common.DB.Create(&user).Error
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "注册失败")
		log.Println("user create error: ", err)
		return
	}
	response.Success(ctx, nil, "注册成功")
}

func Login(ctx *gin.Context) {
	//获取参数
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	//数据验证
	if len(telephone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}
	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
		return
	}

	//判断手机号是否存在
	var user models.User
	common.DB.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户不存在")
		return
	}

	//判断密码是否正确
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		response.Response(ctx, http.StatusBadRequest, 400, nil, "密码错误")
		return
	}

	//发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "系统异常")
		log.Println("token generate error: ", err)
		return
	}

	//返回结果
	response.Success(ctx, gin.H{"token": token}, "登录成功")
}

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	ctx.JSON(http.StatusOK, gin.H{
		"user": dto.ToUserDto(user.(models.User)),
	})
}

func isTelephoneExist(telephone string) bool {
	var user models.User
	common.DB.Where("telephone = ?", telephone).First(&user)
	return user.ID != 0
}
