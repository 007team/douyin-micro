package handlers

import (
	"context"
	"fmt"
	"github.com/007team/douyin-micro/gateway/pkg/jwt"
	"github.com/007team/douyin-micro/gateway/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// 用户注册
func Register(c *gin.Context){
	var userRegisterReq services.UserRegisterRequest
	userRegisterReq.Username = c.Query("username")
	userRegisterReq.Password = c.Query("password")

	fmt.Println(userRegisterReq)

	// 从gin.Key中取出服务实例
	userService := c.Keys["userService"].(services.UserService)
	userResp,err := userService.Register(context.Background(),&userRegisterReq)
	if err!=nil{
		log.Println("Register error:",err)
		return
	}
	// 生成 token
	token, _, err := jwt.GenToken(userResp.UserId)
	if err != nil {
		log.Fatalln("jwt.GenToken  生成token失败", err)
		return
	}
	userResp.Token=token
	c.JSON(http.StatusOK,userResp)
}

// 用户登录
func Login(c *gin.Context){
	var userLoginReq services.UserLoginRequest
	userLoginReq.Username = c.Query("username")
	userLoginReq.Password = c.Query("password")

	// 从gin.Key中取出服务实例
	userService := c.Keys["userService"].(services.UserService)
	userResp,err := userService.Login(context.Background(),&userLoginReq)
	if err!=nil{
		log.Println("Register error:",err)
		return
	}
	// 生成 token
	token, _, err := jwt.GenToken(userResp.UserId)
	if err != nil {
		log.Fatalln("jwt.GenToken  生成token失败", err)
		return
	}
	userResp.Token=token
	c.JSON(http.StatusOK,userResp)
}

// 用户信息
func UserInfo(c *gin.Context){

}