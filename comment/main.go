package main

//protoc --proto_path=. --micro_out=. --go_out=. userService.proto
func main() {
	//// 配置信息初始化
	//if err := settings.Init(); err != nil {
	//	fmt.Println("settings init failed", err)
	//	return
	//}
	//// mysql 初始化
	//if err := mysql.Init(settings.Conf.MySQLConfig); err != nil {
	//	fmt.Println("mysql init failed", err)
	//	return
	//}
	//// redis 初始化
	//if err := redis.Init(settings.Conf.RedisConfig); err != nil {
	//	fmt.Println("redis init failed", err)
	//	return
	//}
	//defer redis.Close()
	//
	//
	//// etcd注册件
	//etcdReg := etcd.NewRegistry(
	//	registry.Addrs("127.0.0.1:2379"),
	//)
	//// 得到一个微服务实例
	//microService := micro.NewService(
	//	micro.Name("rpcUserService"), // 微服务名字
	//	micro.Address("127.0.0.1:8082"),
	//	micro.Registry(etcdReg), // etcd注册件
	//)
	//// 结构命令行参数，初始化
	//microService.Init()
	//// 服务注册
	//_ = services.RegisterUserServiceHandler(microService.Server(), new(logic.UserService))
	//// 启动微服务
	//_ = microService.Run()
}