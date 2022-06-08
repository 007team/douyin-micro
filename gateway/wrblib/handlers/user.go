package handlers

import (
	"context"
	"github.com/007team/douyin-micro/gateway/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func Register(ginCtx *gin.Context) {
	var userReq services.UserRegisterRequest
	PanicIfUserError(ginCtx.Bind(&userReq))

	// 从gin.key 中取出服务实例
	userService := ginCtx.Keys["userService"].(services.UserService)
	userResp, err := userService.Register(context.Background(), &userReq)
	PanicIfUserError(err)
	ginCtx.JSON(http.StatusOK, userResp)
}
func Login(ginCtx *gin.Context) {
	var userReq services.UserLoginRequest
	userReq.Username = ginCtx.Query("username")
	userReq.Password = ginCtx.Query("password")

	// 从gin.key 中取出服务实例
	userService := ginCtx.Keys["userService"].(services.UserService)
	userResp, err := userService.Login(context.Background(), &userReq)
	if err != nil {
		log.Println("Register error:", err)
		return
	}
	ginCtx.JSON(http.StatusOK, userResp)
}

// UserInfo handler
func UserInfo(c *gin.Context) {
	var userReq services.UserRequest
	userIdStr := c.Query("user_id") // 获取用户id
	userid, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		log.Fatalln("UserInfo: user_id invalied", err)
		return
	}

	token, ok := c.Get("token")
	if !ok {
		log.Fatalln("c.Get(\"token\") invalied", err)
		return
	}

	userReq.UserId = userid
	userReq.Token = token.(string)

	// 从gin.key 中取出服务实例
	userService := c.Keys["userService"].(services.UserService)
	userResp, err := userService.UserInfo(context.Background(), &userReq)
	if err != nil {
		log.Println("UserInfo error:", err)
		return
	}
	c.JSON(http.StatusOK, userResp)
}
