package main

import (
	"github.com/007team/douyin-micro/gateway/services"
	"github.com/007team/douyin-micro/gateway/wrblib/routers"
	"github.com/micro/go-micro/v2"

	"github.com/007team/douyin-micro/gateway/wrblib/handlers"
	//"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/web"
	//"microDouyinapp/gateway/wrappers"
	"time"
)

func main() {
	go handlers.VideoUploadFunc()
	go handlers.ImgUploadFunc()

	// etcd 服务注册
	etcdReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)
	//user
	userMicroService := micro.NewService(
		micro.Name("userService.client"),
	)

	// 用户服务调用实例
	userService := services.NewUserService("rpcUserService", userMicroService.Client())

	// video
	videoMicroService := micro.NewService(
		micro.Name("videoService.client"),

	)
	// video服务调用实例
	videoService := services.NewVideoService("rpcVideoService", videoMicroService.Client())

	// comment
	commentMicroService := micro.NewService(
		micro.Name("commentService.client"),
	)
	// comment调用实例
	commentService := services.NewCommentService("rpcCommentService", commentMicroService.Client())

	//创建微服务实例，使用gin暴露http接口并注册到etcd
	server := web.NewService(
		web.Name("httpService"),
		web.Address(""),
		//将服务调用实例使用gin处理
		web.Handler(routers.NewRouter(userService, videoService, commentService)),
		web.Registry(etcdReg),
		web.RegisterTTL(time.Second*30),
		web.RegisterInterval(time.Second*15),
		web.Metadata(map[string]string{"protocol": "http"}),
	)
	//接收命令行参数
	_ = server.Init()
	_ = server.Run()
}
