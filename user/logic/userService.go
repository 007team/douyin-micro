package logic

import (
	"context"
	"douyin-micro/user/dao/mysql"
	"douyin-micro/user/models"
	"douyin-micro/user/services"
)

func BuildUser(item models.User) *services.UserModel {
	userModel := services.UserModel{
		Id:            item.Id,
		Name:          item.Name,
		FollowCount:   item.FollowCount,
		FollowerCount: item.FollowerCount,
		Password:      item.Password,
		IsFollow:      item.IsFollow,
		Salt:          item.Salt,
		CreatedAt:     item.CreatedAt.Unix(),
		UpdatedAt:     item.UpdatedAt.Unix(),
	}
	return &userModel
}

func (*UserService) Login(ctx context.Context, req *services.UserRequest, resp *services.UserDetailResponse) error {
	var user models.User
	resp.Code = 200
	if err := mysql.Db.Where("user_name=?", req.UserName).First(&user).Error; err != nil {
		if err != nil {
			resp.Code = 400
			return nil
		}
		resp.Code = 500
		return nil
	}
	//if user.CheckPassword(req.Password)==false {
	//	resp.Code=400
	//	return nil
	//}
	resp.UserDetail = BuildUser(user)
	return nil
}

func (s *UserService) Register(ctx context.Context, request *services.UserRequest, response *services.UserDetailResponse) error {

}

func (s *UserService) UserInfo(ctx context.Context, request *services.UserRequest, response *services.UserDetailResponse) error {
	panic("implement me")
}

func (s *UserService) RelationAction(ctx context.Context, request *services.UserRequest, response *services.UserDetailResponse) error {
	panic("implement me")
}

func (s *UserService) FollowList(ctx context.Context, request *services.UserRequest, response *services.UserDetailResponse) error {
	panic("implement me")
}

func (s *UserService) FollowerList(ctx context.Context, request *services.UserRequest, response *services.UserDetailResponse) error {
	panic("implement me")
}
