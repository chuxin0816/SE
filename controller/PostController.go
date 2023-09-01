package controller

import (
	"chuxin0816/SE/common"
	"chuxin0816/SE/models"
	"chuxin0816/SE/response"
	"chuxin0816/SE/vo"
	"strconv"

	"github.com/gin-gonic/gin"
)

type IPostController interface {
	RestController
	PageList(ctx *gin.Context)
}

type PostController struct {
}

func NewPostController() IPostController {
	common.DB.AutoMigrate(&models.Post{})
	return PostController{}
}

func (p PostController) Create(ctx *gin.Context) {
	var requestPost vo.CreatePostRequest
	// 数据验证
	err := ctx.ShouldBind(&requestPost)
	if err != nil {
		response.Fail(ctx, nil, "数据验证错误，分类名称必填")
		return
	}

	// 获取登录用户
	user, _ := ctx.Get("user")

	// 创建文章
	post := models.Post{
		UserId:     user.(models.User).ID,
		CategoryId: requestPost.CategoryId,
		Title:      requestPost.Title,
		HeadImg:    requestPost.HeadImg,
		Content:    requestPost.Content,
	}
	err = common.DB.Create(&post).Error
	if err != nil {
		panic(err)
	}

	// 返回响应
	response.Success(ctx, gin.H{"post": post}, "创建成功")
}

func (p PostController) Update(ctx *gin.Context) {
	var requestPost vo.CreatePostRequest
	// 数据验证
	err := ctx.ShouldBind(&requestPost)
	if err != nil {
		response.Fail(ctx, nil, "数据验证错误")
		return
	}

	// 获取path中的post_id
	postId := ctx.Param("id")

	// 查询文章
	var post models.Post
	err = common.DB.Where("id = ?", postId).First(&post).Error
	if err != nil {
		response.Fail(ctx, nil, "文章不存在")
		return
	}

	// 获取登录用户
	user, _ := ctx.Get("user")

	// 判断当前用户是否为文章的作者
	if user.(models.User).ID != post.UserId {
		response.Fail(ctx, nil, "文章不属于您，无法修改")
	}

	// 更新文章
	err = common.DB.Model(&post).Updates(requestPost).Error
	if err != nil {
		response.Fail(ctx, nil, "更新失败")
		return
	}

	response.Success(ctx, gin.H{"post": post}, "更新成功")

}
func (p PostController) Show(ctx *gin.Context) {
	// 获取path中的post_id
	postId := ctx.Param("id")

	// 查询文章
	var post models.Post
	err := common.DB.Preload("Category").Where("id = ?", postId).First(&post).Error
	if err != nil {
		response.Fail(ctx, nil, "文章不存在")
		return
	}

	response.Success(ctx, gin.H{"post": post}, "成功")
}
func (p PostController) Delete(ctx *gin.Context) {
	// 获取path中的post_id
	postId := ctx.Param("id")

	// 查询文章
	var post models.Post
	err := common.DB.Where("id = ?", postId).First(&post).Error
	if err != nil {
		response.Fail(ctx, nil, "文章不存在")
		return
	}

	// 获取登录用户
	user, _ := ctx.Get("user")

	// 判断当前用户是否为文章的作者
	if user.(models.User).ID != post.UserId {
		response.Fail(ctx, nil, "文章不属于您，无法删除")
	}

	// 删除文章
	err = common.DB.Delete(&post).Error
	if err != nil {
		response.Fail(ctx, nil, "删除失败")
		return
	}

	response.Success(ctx, nil, "删除成功")
}

func (p PostController) PageList(ctx *gin.Context) {
	// 获取分页参数
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))

	// 分页
	var posts []models.Post
	common.DB.Order("created_at desc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&posts)

	// 总数
	var total int64
	common.DB.Model(&models.Post{}).Count(&total)

	response.Success(ctx, gin.H{"posts": posts, "total": total}, "成功")
}
