package controller

import (
	"chuxin0816/SE/common"
	"chuxin0816/SE/models"
	"chuxin0816/SE/repository"
	"chuxin0816/SE/response"
	"chuxin0816/SE/vo"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ICategoryController interface {
	RestController
}

type CategoryController struct {
	Repository repository.CategoryRepository
}

func NewCategoryController() ICategoryController {
	repository := repository.NewCategoryRepository()
	common.DB.AutoMigrate(&models.Category{})
	return CategoryController{Repository: repository}
}

func (c CategoryController) Create(ctx *gin.Context) {
	var requestCategory vo.CreateCategoryRequest
	err := ctx.ShouldBind(&requestCategory)
	if err != nil {
		response.Fail(ctx, nil, "数据验证错误，分类名称必填")
		return
	}
	category, err := c.Repository.Create(requestCategory.Name)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "创建失败")
		panic(err)
	}
	response.Success(ctx, gin.H{"category": category}, "创建成功")
}

func (c CategoryController) Update(ctx *gin.Context) {
	var requestCategory vo.CreateCategoryRequest
	err := ctx.ShouldBind(&requestCategory)
	if err != nil {
		response.Fail(ctx, nil, "数据验证错误，分类名称必填")
		return
	}
	categoryID, _ := strconv.Atoi(ctx.Param("id"))
	updateCategory, err := c.Repository.SelectById(categoryID)
	if err != nil {
		response.Fail(ctx, nil, "分类不存在")
		return
	}
	category, err := c.Repository.Update(*updateCategory, requestCategory.Name)
	if err != nil {
		response.Fail(ctx, nil, "更新失败")
		panic(err)
	}
	response.Success(ctx, gin.H{"category": category}, "更新成功")
}

func (c CategoryController) Show(ctx *gin.Context) {
	categoryID, _ := strconv.Atoi(ctx.Param("id"))
	category, err := c.Repository.SelectById(categoryID)
	if err != nil {
		response.Fail(ctx, nil, "分类不存在")
		return
	}
	response.Success(ctx, gin.H{"category": category}, "查询成功")
}

func (c CategoryController) Delete(ctx *gin.Context) {
	categoryID, _ := strconv.Atoi(ctx.Param("id"))
	err := c.Repository.DeleteById(categoryID)
	if err != nil {
		response.Fail(ctx, nil, "删除失败")
		return
	}
	response.Success(ctx, nil, "删除成功")
}
