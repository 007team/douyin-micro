package redis

import (
	"github.com/007team/douyin-micro/user/models"
	"log"
)

// IsFollowUser 我是否关注了这个用户
func IsFollowUser(user *models.User, myUserId int64) (ok bool, err error) {
	KeyMyFollowSet := getKeyUserFollowSet(myUserId)
	ok, err = rdb.SIsMember(KeyMyFollowSet, user.Id).Result()
	if err != nil {
		log.Println("rdb.SIsMember(KeyMyFollowSet,user.Id).Result() failed", err)
		return false, err
	}
	return ok, err
}
