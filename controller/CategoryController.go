package controller

import (
	"chuxin0816/SE/common"
	"chuxin0816/SE/models"
	"chuxin0816/SE/response"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ICategoryController interface {
	RestController
}

type CategoryController struct{}

func NewCategoryController() ICategoryController {
	common.DB.AutoMigrate(&models.Category{})
	return CategoryController{}
}

func (c CategoryController) Create(ctx *gin.Context) {
	var requestCategory = models.Category{}
	ctx.ShouldBind(&requestCategory)
	if requestCategory.Name == "" {
		response.Fail(ctx, nil, "数据验证错误，分类名称必填")
		return
	}
	err := common.DB.Create(&requestCategory).Error
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "创建失败")
		log.Println("category create error: ", err)
		return
	}
	response.Success(ctx, gin.H{"category": requestCategory}, "创建成功")
}
func (c CategoryController) Update(ctx *gin.Context) {
	var requestCategory = models.Category{}
	ctx.ShouldBind(&requestCategory)
	if requestCategory.Name == "" {
		response.Fail(ctx, nil, "数据验证错误，分类名称必填")
		return
	}
	categoryID, _ := strconv.Atoi(ctx.Param("id"))
	var updateCategory = models.Category{}
	err := common.DB.First(&updateCategory, categoryID).Error
	if err != nil {
		response.Fail(ctx, nil, "分类不存在")
		return
	}
	common.DB.Model(&updateCategory).Update("name", requestCategory.Name)
	response.Success(ctx, gin.H{"category": updateCategory}, "更新成功")
}
func (c CategoryController) Show(ctx *gin.Context) {
	categoryID, _ := strconv.Atoi(ctx.Param("id"))
	var category = models.Category{}
	err := common.DB.First(&category, categoryID).Error
	if err != nil {
		response.Fail(ctx, nil, "分类不存在")
		return
	}
	response.Success(ctx, gin.H{"category": category}, "查询成功")
}
func (c CategoryController) Delete(ctx *gin.Context) {
	categoryID, _ := strconv.Atoi(ctx.Param("id"))
	err := common.DB.Delete(&models.Category{}, categoryID).Error
	if err != nil {
		response.Fail(ctx, nil, "删除失败")
		return
	}
	response.Success(ctx, nil, "删除成功")
}
