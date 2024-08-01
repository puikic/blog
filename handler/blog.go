package handler

import (
	"blog/database"
	"blog/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func BlogList(ctx *gin.Context) {
	uid, err := strconv.Atoi(ctx.Param("uid")) //获取restful参数
	if err != nil {
		ctx.String(http.StatusBadRequest, "invalid uid")
		return
	}
	blogs := database.GetBlogByUserId(uid)
	ctx.HTML(http.StatusOK, "blog_list.html", blogs) // go template
}

// 获取某一篇博客的详情
func BlogDetail(ctx *gin.Context) {
	blogId := ctx.Param("bid") //获取restful参数
	bid, err := strconv.Atoi(blogId)
	if err != nil {
		ctx.String(http.StatusBadRequest, "invalid blog id")
		return
	}
	blog := database.GetBlogById(bid)
	if blog == nil {
		ctx.String(http.StatusNotFound, "博客不存在")
		return
	}
	util.LogRus.Debug(blog.Article)
	ctx.HTML(http.StatusOK, "blog.html", gin.H{"title": blog.Title, "article": blog.Article, "bid": blogId, "update_time": blog.UpdateTime.Format("2006-01-02 15:04:05")})
}

type UpdateRequest struct {
	BlogId  int    `form:"bid" binding:"gt=0"`
	Title   string `form:"title" binding:"gt=0"`
	Article string `form:"article" binding:"gt=0"`
}

// 更新博客
func BlogUpdate(ctx *gin.Context) {
	// blogId := ctx.PostForm("bid") //获取post form参数
	// title := ctx.PostForm("title")
	// article := ctx.PostForm("article")
	// bid, err := strconv.Atoi(blogId)
	// if err != nil {
	// 	ctx.String(http.StatusBadRequest, "invalid blog id")
	// 	return
	// }

	var request UpdateRequest
	err := ctx.ShouldBind(&request)
	if err != nil {
		ctx.String(http.StatusBadRequest, "invalid parameter")
		return
	}
	bid := request.BlogId
	title := request.Title
	article := request.Article
	blog := database.GetBlogById(bid)
	if blog == nil {
		ctx.String(http.StatusForbidden, "无权修改")
		return
	}
	err = database.UpdateBlogById(&database.Blog{Id: bid, Title: title, Article: article})
	if err != nil {
		util.LogRus.Errorf("update blog %d failed: %s", bid, err)
		ctx.String(http.StatusInternalServerError, "更新失败") //不要把原始的err发回给前端,否则用户通过查看页面源码能看到mysql表
		return
	}
	ctx.String(http.StatusOK, "success")
}
