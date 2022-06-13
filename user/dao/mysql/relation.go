package mysql

import (
	"log"

	"github.com/007team/douyin-micro/user/models"
)

func FollowListPointer(es []string) (users []*models.User, err error) {
	err = db.Find(&users, es).Error
	if err != nil {
		log.Println("FollowList db.Find failed", err)
		return nil, err
	}
	return users, nil
}

func FollowerListPointer(es []string) (users []*models.User, err error) {
	err = db.Find(&users, es).Error
	if err != nil {
		log.Println("FollowerList db.Find failed", err)
		return nil, err
	}
	return users, nil
}
