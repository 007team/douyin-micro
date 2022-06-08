package main

import (
	"fmt"
	"github.com/007team/douyin-micro/mq-server/dao/mysql"
	"github.com/007team/douyin-micro/mq-server/dao/redis"
	"github.com/007team/douyin-micro/mq-server/dao/rmq"
	"github.com/007team/douyin-micro/mq-server/settings"
)

func main() {
	// 配置信息初始化
	if err := settings.Init(); err != nil {
		fmt.Println("settings init failed", err)
		return
	}

	if err := mysql.Init(settings.Conf.MySQLConfig); err != nil {
		fmt.Println("mysql init failed", err)
		return
	}

	if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		fmt.Println("redis init failed", err)
		return
	}

	if err := rmq.Init(settings.Conf.RabbitmqConfig); err != nil {
		fmt.Println("rabbitmq init failed", err)
		return
	}



}