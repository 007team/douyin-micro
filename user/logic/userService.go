package logic

import (
	"context"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"log"

	"github.com/007team/douyin-micro/user/dao/mysql"
	"github.com/007team/douyin-micro/user/dao/redis"
	"github.com/007team/douyin-micro/user/models"
	"github.com/007team/douyin-micro/user/pkg/jwt"
	"github.com/007team/douyin-micro/user/services"
)

var letters = []byte("abcdefghjkmnpqrstuvwxyz123456789")

func BuildUser(item models.User) *services.User {
	userModel := services.User{
		Id:            item.Id,
		Name:          item.Name,
		FollowCount:   item.FollowCount,
		FollowerCount: item.FollowerCount,
		IsFollow:      item.IsFollow,
	}
	return &userModel
}

func BuildUserList(item []models.User) []*services.User {
	userlist := []*services.User{}
	for _, user := range item {
		userlist = append(userlist, BuildUser(user))
	}
	return userlist
}

func (*UserService) Login(ctx context.Context, req *services.UserLoginRequest, resp *services.UserLoginResponse) error {
	oPassword := req.Password // 未加密的原密码
	user := models.User{
		Name:     req.Username,
		Password: req.Password,
	}

	// 查询用户
	if err := mysql.Login(&user); err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = "登录失败"
		return nil
	}

	// 进行密码校验
	salt := user.Salt
	newPassword := encryptPassword(oPassword, salt) // 将原密码加密
	if newPassword != user.Password {
		resp.StatusCode = 1
		resp.StatusMsg = "密码错误"
		return nil
	}

	// 生成token
	token, _, err := jwt.GenToken(user.Id)
	if err != nil {
		log.Println("jwt,GenToken 生成token失败")
		resp.StatusCode = 1
		resp.StatusMsg = "登录失败 生成token失败"
		return nil
	}

	resp.StatusCode = 0
	resp.StatusMsg = "登录成功"
	resp.UserId = user.Id
	resp.Token = token
	return nil
}

func (s *UserService) Register(ctx context.Context, request *services.UserRegisterRequest, response *services.UserRegisterResponse) error {

	user := &models.User{
		Name:     request.Username,
		Password: request.Password,
	}

	//先查询该用户是否存在， 如存在则直接返回错误
	if err := mysql.CheckUserExist(user); err != nil {
		response.StatusCode = 1
		response.StatusMsg = "用户名已被注册"
		log.Println(err)
		return nil
	}

	// 对用户密码进行加密
	salt := RandLow()          //生成 salt
	oPassword := user.Password // 旧的密码
	newPassword := encryptPassword(oPassword, string(salt))
	user.Salt = string(salt) // 保存salt
	user.Password = newPassword

	// 插入新用户
	if err := mysql.CreateNewUser(user); err != nil {
		return err
	}

	response.StatusCode = 0
	response.StatusMsg = "注册成功"
	response.UserId = user.Id
	//response.Token = token

	return nil
}

func (s *UserService) UserInfo(ctx context.Context, request *services.UserRequest, response *services.UserResponse) error {
	m, _ := jwt.ParseToken(request.Token)
	// 我的id
	myId := m.UserID

	user := models.User{
		Id: request.UserId,
	}
	// mysql查询用户具体信息
	if err := mysql.UserInfo(&user); err != nil {
		log.Fatalln("mysql.UserInfo failed", err)
		response.StatusCode = 1
		response.StatusMsg = "服务器繁忙，请稍后再试"
		return nil
	}

	// redis查询用户的粉丝与关注数
	var err error
	user.FollowCount, err = redis.UserFollowCount(user.Id)
	if err != nil {
		log.Println("redis.UserFollowCount(user.Id) failed", err)
		response.StatusCode = 1
		response.StatusMsg = "服务器繁忙，请稍后再试"
		return nil
	}
	user.FollowerCount, err = redis.UserFollowerCount(user.Id)
	if err != nil {
		log.Println("redis.UserFollowerCount(user.Id) failed", err)
		response.StatusCode = 1
		response.StatusMsg = "服务器繁忙，请稍后再试"
		return nil
	}

	// “我”是否关注了这个用户
	user.IsFollow, err = redis.IsFollowUser(&user, myId)
	if err != nil {
		log.Println("redis.IsFollowUser(user, myUserId) failed", err)
		response.StatusCode = 1
		response.StatusMsg = "服务器繁忙，请稍后再试"
		return nil
	}

	response.StatusCode = 0
	response.StatusMsg = "操作成功"
	response.User = BuildUser(user)

	return nil
}

func (s *UserService) RelationAction(ctx context.Context, request *services.RelationActionRequest, response *services.RelationActionResponse) error {
	return nil
}

func encryptPassword(oPassword string, salt string) string {
	h := md5.New()
	h.Write([]byte(salt))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
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
