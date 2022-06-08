package handlers

import (
	"context"
	"fmt"
	"github.com/007team/douyin-micro/gateway/pkg/jwt"
	"github.com/007team/douyin-micro/gateway/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
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
	if userResp.StatusCode == 0{
		userResp.Token=token
	}
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
		log.Println("Login error:",err)
		return
	}
	// 生成 token
	token, _, err := jwt.GenToken(userResp.UserId)
	if err != nil {
		log.Fatalln("jwt.GenToken  生成token失败", err)
		return
	}
	if userResp.StatusCode == 0{
		userResp.Token=token
	}

	c.JSON(http.StatusOK,userResp)
}

// 用户信息
func UserInfo(c *gin.Context){
	var userInfoReq services.UserRequest
	token := c .Query("token")
	userIdStr := c.Query("user_id") // 获取用户id
	// myUserId, _ := c.Get("user_id")
	userid, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		log.Println("UserInfo: user_id invalied", err)
		c.JSON(http.StatusOK, gin.H{
			"status_code": 1,
			"status_msg":  "userinfo parseInt failed",
		})
		return
	}


	userInfoReq.UserId = userid
	userInfoReq.Token = token
	// 从gin.Key中取出服务实例
	userService := c.Keys["userService"].(services.UserService)
	userResp,err := userService.UserInfo(context.Background(),&userInfoReq)
	if err!=nil{
		log.Println("UserInfo error:",err)
		return
	}

	c.JSON(http.StatusOK,userResp)
}