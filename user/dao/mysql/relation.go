package mysql

import (
	"github.com/007team/douyin-micro/user/models"
	"log"
)

func FollowList(es []string) (users []models.User, err error) {
	err = Db.Find(&users, es).Error
	if err != nil {
		log.Println("FollowList db.Find failed", err)
		return nil, err
	}
	return users, nil
}

func FollowerList(es []string) (users []models.User, err error) {
	err = Db.Find(&users, es).Error
	if err != nil {
		log.Println("FollowerList db.Find failed", err)
		return nil, err
	}
	return users, nil
}


func FollowListPointer(es []string) (users []*models.User, err error) {
	err = Db.Find(&users, es).Error
	if err != nil {
		log.Println("FollowList db.Find failed", err)
		return nil, err
	}
	return users, nil
}

func FollowerListPointer(es []string) (users []*models.User, err error) {
	err = Db.Find(&users, es).Error
	if err != nil {
		log.Println("FollowerList db.Find failed", err)
		return nil, err
	}
	return users, nil
}