package redis

import (
	"github.com/007team/douyin-micro/video/models"
	"github.com/go-redis/redis"
	"log"
	"strconv"
)

// IsFavoriteVideo 判断视频是否被用户点赞
func IsFavoriteVideo(userId, videoId int64) (ok bool, err error) {
	ok, err = rdb.SIsMember(getKeyUserFavoriteSet(userId), videoId).Result()

	if err != nil {
		log.Println("rdb.SIsMember failed", err)
		return false, err
	}
	return ok, nil
}

func FavoriteList(userId int64) (es []string, err error) {
	KeyUserFavoriteSet := getKeyUserFavoriteSet(userId)
	es, err = rdb.SMembers(KeyUserFavoriteSet).Result()
	if err != nil {
		log.Println("rdb.SMembers failed", err)
	}
	return es, err

}

// 添加videoId 到点赞数zset 和 评论数zset
func Publish(videoId int64) (err error) {
	KeyAllVideoZSet := getKeyAllVideoZSet()                         //点赞数zset
	KeyAllVideoCommentCountZSet := getKeyAllVideoCommentCountZSet() // 评论数zset
	err = rdb.ZAdd(KeyAllVideoZSet, redis.Z{0, videoId}).Err()
	if err != nil {
		log.Println("rdb.ZAdd(KeyAllVideoZSet, redis.Z{0, videoId}) failed", err)
		return err
	}
	err = rdb.ZAdd(KeyAllVideoCommentCountZSet, redis.Z{0, videoId}).Err()
	if err != nil {
		log.Println("rdb.ZAdd(KeyAllVideoCommentCountZSet,redis.Z{0,videoId}) failed", err)
		return err
	}
	return nil
}

func FavoriteAction(userId int64, videoId int64) (err error) {
	KeyAllVideoZSet := getKeyAllVideoZSet()                  // 视频点赞数集合
	KeyUserFavoriteVideoSet := getKeyUserFavoriteSet(userId) //用户的点赞视频集合
	// 开启事务
	pipe := rdb.TxPipeline()
	// 对视频点赞数++
	videoIdstr := strconv.Itoa(int(videoId))
	err = pipe.ZIncrBy(KeyAllVideoZSet, 1, videoIdstr).Err()
	if err != nil {
		log.Println("pipe.ZIncrBy failed", err)
	}

	// 在用户的点赞视频列表里写入videoId
	err = pipe.SAdd(KeyUserFavoriteVideoSet, videoId).Err()
	if err != nil {
		log.Println("pipe.SAdd failed", err)
	}
	_, err = pipe.Exec()
	if err != nil {
		log.Println("pipe failed", err)
	}
	return err
}

func UnFavoriteAction(userId, videoId int64) (err error) {
	KeyAllVideoZSet := getKeyAllVideoZSet()                  // 视频点赞数集合
	KeyUserFavoriteVideoSet := getKeyUserFavoriteSet(userId) //用户的点赞视频集合

	//开启事务
	pipe := rdb.TxPipeline()
	// 对视频点赞数--
	videoIdStr := strconv.Itoa(int(videoId))
	err = pipe.ZIncrBy(KeyAllVideoZSet, -1, videoIdStr).Err()
	if err != nil {
		log.Println("pipe.ZIncrBy failed", err)
	}
	// 在用户的点赞视频列表里删除videoId
	err = pipe.SRem(KeyUserFavoriteVideoSet, videoId).Err()
	if err != nil {
		log.Println("pipe.SRem failed", err)
	}
	_, err = pipe.Exec()
	if err != nil {
		log.Println("pipe failed", err)
	}
	return err
}

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
