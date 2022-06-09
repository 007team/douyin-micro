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
	//apiRouter.GET("/feed/", middlewares.JWTAuthMiddlewareForFeed(), handlers.Feed)
	apiRouter.POST("/user/register/", handlers.Register)
	apiRouter.POST("/user/login/", handlers.Login)
	apiRouter.GET("/user/", middlewares.JWTAuthMiddleware(), handlers.UserInfo)
	//apiRouter.POST("/publish/action/", middlewares.JWTAuthMiddlewareForPublish(), handlers.Publish)
	//apiRouter.GET("/publish/list/", middlewares.JWTAuthMiddleware(), handlers.PublishList)
	//
	//// extra apis - I
	//apiRouter.POST("/favorite/action/", middlewares.JWTAuthMiddleware(), handlers.FavoriteAction)
	//apiRouter.GET("/favorite/list/", middlewares.JWTAuthMiddleware(), handlers.FavoriteList)
	apiRouter.POST("/comment/action/", middlewares.JWTAuthMiddleware(), handlers.CommentAction)
	apiRouter.GET("/comment/list/", middlewares.JWTAuthMiddleware(), handlers.CommentList)
	//
	//// extra apis - II
	apiRouter.POST("/relation/action/", middlewares.JWTAuthMiddleware(), handlers.RelationAction)
	apiRouter.GET("/relation/follow/list/", middlewares.JWTAuthMiddleware(), handlers.FollowList)
	apiRouter.GET("/relation/follower/list/", middlewares.JWTAuthMiddleware(), handlers.FollowerList)

	return r
}