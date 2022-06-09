package main

import (
	"github.com/007team/douyin-micro/gateway/services"
	//"github.com/007team/douyin-micro/gateway/wrappers"
	"github.com/007team/douyin-micro/gateway/wrblib/routers"
	"github.com/micro/go-micro/v2"

	//"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/web"
	//"microDouyinapp/gateway/wrappers"
	"time"
)

func main() {

	// etcd 服务注册
	etcdReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)
	// user
	userMicroService := micro.NewService(
		micro.Name("userService.client"),
		//micro.WrapClient(wrappers.NewUserWrapper),
	)

	// 用户服务调用实例
	userService := services.NewUserService("rpcUserService",userMicroService.Client())

	// video
	//videoMicroService := micro.NewService(
	//	micro.Name("taskService.client"),
	//	micro.WrapClient(wrappers.NewVideoWrapper),
	//)
	// video服务调用实例
	//taskService := services.NewTaskService("rpcTaskService",taskMicroService.Client())

	//comment
	commentMicroService := micro.NewService(
		micro.Name("commentService.client"),
		//micro.WrapClient(wrappers.NewCommentWrapper),
	)
	//comment调用实例
	commentService := services.NewCommentService("rpcCommentService",commentMicroService.Client())


	//创建微服务实例，使用gin暴露http接口并注册到etcd
	server := web.NewService(
		web.Name("httpService"),
		//web.Address("10.54.12.125:4000"),
		web.Address("127.0.0.1:4000"),
		//将服务调用实例使用gin处理
		web.Handler(routers.NewRouter(userService,commentService)),
		//web.Handler(routers.NewRouter(userService,videoService,commentService)),
		web.Registry(etcdReg),
		web.RegisterTTL(time.Second*30),
		web.RegisterInterval(time.Second*15),
		web.Metadata(map[string]string{"protocol": "http"}),
	)
	//接收命令行参数
	_ = server.Init()
	_ = server.Run()
}
