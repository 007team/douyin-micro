package logic

import (
	"context"
	"github.com/007team/douyin-micro/user/services"
)

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
	panic("implement me")
}

func (s *UserService) Login(ctx context.Context, request *services.UserLoginRequest, response *services.UserLoginResponse) error {
	panic("implement me")
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