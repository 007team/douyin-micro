package conf

import (
	"fmt"
	"github.com/007team/douyin-micro/video/dao/mysql"
	"github.com/007team/douyin-micro/video/dao/redis"
	"github.com/007team/douyin-micro/video/settings"
)

func Init() {
	// 配置信息初始化
	if err := settings.Init(); err != nil {
		fmt.Println("settings init failed", err)
		return
	}
	// mysql 初始化
	if err := mysql.Init(settings.Conf.MySQLConfig); err != nil {
		fmt.Println("mysql init failed", err)
		return
	}
	// redis 初始化
	if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		fmt.Println("redis init failed", err)
		return
	}
}
