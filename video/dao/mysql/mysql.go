// Package mysql
// 操作mysql数据库 （增删改查）
//
package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"

	"github.com/007team/douyin-micro/video/settings"
)

// 对mysql进行操作时，用db这个变量来操作数据库

var db *gorm.DB

func Init(cfg *settings.MySQLConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DbName,
	)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		fmt.Println("mysql Open failed", err)
	}

	sqlDB, err := db.DB()

	if err != nil {
		log.Println(err)
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)

	//db.AutoMigrate(&models.Video{})
	return err
}
func DB() *gorm.DB {
	return db
}
