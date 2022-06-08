package routers

import (
	"github.com/007team/douyin-micro/gateway/wrblib/handlers"
	"github.com/007team/douyin-micro/gateway/wrblib/middlewares"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func NewRouter(service ...interface{}) *gin.Engine{

	r := gin.Default()
	r.Static("/static", "./public")
	r.Use(middlewares.Cors(), middlewares.InitMiddleware(service), middlewares.ErrorMiddleware())
	store := cookie.NewStore([]byte("something-very-secret"))
	r.Use(sessions.Sessions("mysession", store))
	apiRouter := r.Group("/douyin")

	//// basic apis
	//apiRouter.GET("/feed/", middlewares.JWTAuthMiddlewareForFeed(), controller.Feed)
	apiRouter.POST("/user/register/", handlers.Register)
	//apiRouter.POST("/user/login/", controller.Login)
	//apiRouter.GET("/user/", middlewares.JWTAuthMiddleware(), controller.UserInfo)
	//apiRouter.POST("/publish/action/", middlewares.JWTAuthMiddlewareForPublish(), controller.Publish)
	//apiRouter.GET("/publish/list/", middlewares.JWTAuthMiddleware(), controller.PublishList)
	//
	//// extra apis - I
	//apiRouter.POST("/favorite/action/", middlewares.JWTAuthMiddleware(), controller.FavoriteAction)
	//apiRouter.GET("/favorite/list/", middlewares.JWTAuthMiddleware(), controller.FavoriteList)
	//apiRouter.POST("/comment/action/", middlewares.JWTAuthMiddleware(), controller.CommentAction)
	//apiRouter.GET("/comment/list/", middlewares.JWTAuthMiddleware(), controller.CommentList)
	//
	//// extra apis - II
	//apiRouter.POST("/relation/action/", middlewares.JWTAuthMiddleware(), controller.RelationAction)
	//apiRouter.GET("/relation/follow/list/", middlewares.JWTAuthMiddleware(), controller.FollowList)
	//apiRouter.GET("/relation/follower/list/", middlewares.JWTAuthMiddleware(), controller.FollowerList)

	return r
}