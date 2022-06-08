package mysql

import (
	"github.com/007team/douyin-micro/user/models"
	"log"


)

func UserInfo(user *models.User) (err error) {
	// 根据 userid 查询数据库
	if err = Db.First(user, user.Id).Error; err != nil {
		log.Fatalln("mysql.UserInfo 查询错误", err)
		return err
	}

	return nil
}

// CheckUserExist 查询用户是否存在
func CheckUserExist(user *models.User) (err error) {
	var count int64
	err = Db.Model(&models.User{}).Where("name = ?", user.Name).Count(&count).Error
	if err != nil {
		log.Fatalln("mysql.CeckUserExist  用户数据查询失败", err)
		return err
	}
	if count > 0 {
		return ErrorUserExist // 用户已存在
	}

	return
}

// CreateNewUser 创建新的用户
func CreateNewUser(user *models.User) (err error) {
	if err = Db.Select("Name", "Password", "Salt").Create(user).Error; err != nil {
		log.Fatalln("mysql.CreateNewUser  创建新用户失败", err)
		return err
	}
	return
}

func Login(user *models.User) (err error) {
	return Db.Where("name = ?", user.Name).First(&user).Error
}


