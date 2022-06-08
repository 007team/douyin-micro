package logic

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"github.com/007team/douyin-micro/user/dao/mysql"
	"github.com/007team/douyin-micro/user/models"
	"github.com/007team/douyin-micro/user/services"
	"log"
	"math/rand"
)

var letters = []byte("abcdefghjkmnpqrstuvwxyz123456789")

//func BuildUser(item models.User) *services.UserModel {
//	userModel := services.UserModel{
//		Id: 				item.Id,
//		Name:				item.Name,
//		FollowCount:		item.FollowCount,
//		FollowerCount:		item.FollowerCount,
//		Password:			item.Password,
//		IsFollow:			item.IsFollow,
//		Salt:				item.Salt,
//		CreatedAt:			item.CreatedAt.Unix(),
//		UpdatedAt:			item.UpdatedAt.Unix(),
//	}
//	return &userModel
//}


func (s *UserService) Register(ctx context.Context, request *services.UserRegisterRequest, response *services.UserRegisterResponse) error {

	user := &models.User{
		Name: request.Username,
		Password: request.Password,
	}

	//先查询该用户是否存在， 如存在则直接返回错误
	if err := mysql.CheckUserExist(user); err != nil {
		return  err
	}

	// 对用户密码进行加密
	salt := RandLow()          //生成 salt
	oPassword := user.Password // 旧的密码
	newPassword := encryptPassword(oPassword, string(salt))
	user.Salt = string(salt) // 保存salt
	user.Password = newPassword

	// 插入新用户
	if err := mysql.CreateNewUser(user); err != nil {
		return  err
	}

	//// 生成 token
	//token, _, err := jwt.GenToken(user.Id)
	//if err != nil {
	//	log.Fatalln("jwt.GenToken  生成token失败", err)
	//	return  err
	//}
	response.StatusCode = 0
	response.StatusMsg = "注册成功"
	response.UserId = user.Id
	//response.Token = token

	return nil
}

func (s *UserService) Login(ctx context.Context, request *services.UserLoginRequest, response *services.UserLoginResponse) error {
	// 获取未加密的原密码
	oPassword := request.Password

	user := &models.User{
		Name: request.Username,
		Password: request.Password,
	}
	// 查询用户
	if err := mysql.Login(user); err != nil {
		response.StatusCode = 1
		response.StatusMsg = "用户不存在"
		log.Println(err)
		return nil
	}
	// 进行密码校验
	salt := user.Salt
	newPassword := encryptPassword(oPassword, salt) // 将原密码加密
	if newPassword != user.Password {
		response.StatusCode = 1
		response.StatusMsg = "密码错误"
		return mysql.ErrorInvalidUserPassword // 用户密码错误
	}
	// 生成token
	//token, _, err := jwt.GenToken(user.Id)
	//if err != nil {
	//	log.Fatalln("jwt,GenToken 生成token失败")
	//	return err
	//}

	response.StatusCode = 0
	response.StatusMsg = "登录成功"
	response.UserId = user.Id
	//response.Token = token

	return nil
}

func (s *UserService) UserInfo(ctx context.Context, request *services.UserRequest, response *services.UserResponse) error {
	panic("implement me")
}

func (s *UserService) RelationAction(ctx context.Context, request *services.RelationActionRequest, response *services.RelationActionResponse) error {
	panic("implement me")
}

func (s *UserService) FollowList(ctx context.Context, request *services.FollowListRequest, response *services.FollowListResponse) error {
	panic("implement me")
}

func (s *UserService) FollowerList(ctx context.Context, request *services.FollowerListRequest, request2 *services.FollowerListRequest) error {
	panic("implement me")
}






// RandLow 生成加密密码用的随机字符串  salt
func RandLow() []byte {
	n := 15
	if n <= 0 {
		return []byte{}
	}
	b := make([]byte, n)
	arc := uint8(0)
	if _, err := rand.Read(b[:]); err != nil {
		return []byte{}
	}
	for i, x := range b {
		arc = x & 31
		b[i] = letters[arc]
	}
	return b
}
func encryptPassword(oPassword string, salt string) string {
	h := md5.New()
	h.Write([]byte(salt))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}