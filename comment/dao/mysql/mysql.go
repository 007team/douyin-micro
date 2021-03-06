// Package mysql
// 操作mysql数据库 （增删改查）
//
package mysql

import (
	"fmt"
	"github.com/007team/douyin-micro/comment/settings"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 对mysql进行操作时，用db这个变量来操作数据库

var Db *gorm.DB

func Init(cfg *settings.MySQLConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DbName,
	)
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		fmt.Println("mysql Open failed", err)
	}

	sqlDB, err := Db.DB()

	//err = Db.AutoMigrate(&models.Comment{})
	//if err != nil {
	//	log.Println(err)
	//}

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)

	return err
}
