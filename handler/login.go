package handler

import (
	"blog/database"
	"blog/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginResponse struct {
	Code  int    `json:"code"` //前后端分离,前端根据code向用户展示对应的话术.若需要改话术，后端代码不用动
	Msg   string `json:"msg"`  //msg仅仅用于研发人员内部排查，不会展示给用户
	Uid   int    `json:"uid"`
	Token string `json:"token"`
}

func Login(ctx *gin.Context) {
	name := ctx.PostForm("user") //从post formz中获取数据
	pass := ctx.PostForm("pass")
	if len(name) == 0 {
		ctx.JSON(http.StatusBadRequest, LoginResponse{1, "must indicate user name", 0, ""})
		return
	}
	if len(pass) != 32 {
		ctx.JSON(http.StatusBadRequest, LoginResponse{2, "invalid password", 0, ""})
		return
	}
	user := database.GetUserByName(name)
	if user == nil {
		ctx.JSON(http.StatusForbidden, LoginResponse{3, "用户名不存在", 0, ""})
		return
	}
	if user.PassWd != pass {
		ctx.JSON(http.StatusForbidden, LoginResponse{4, "密码错误", 0, ""})
		return
	}
	util.LogRus.Infof("user %s(%d) login", name, user.Id)
}
