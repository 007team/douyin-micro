package handlers

import (
	"context"
	"github.com/007team/douyin-micro/gateway/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

//用户注册
func Register(c *gin.Context){
	var userRegisterReq services.UserRegisterRequest
	userRegisterReq.Username = c.Param("username")
	userRegisterReq.Password = c.Param("password")

	// 从gin.Key中取出服务实例
	userService := c.Keys["userService"].(services.UserService)
	userResp,err := userService.Register(context.Background(),&userRegisterReq)
	if err!=nil{
		log.Println("Register error:",err)
		return
	}

	c.JSON(http.StatusOK,userResp)
}