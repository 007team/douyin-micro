package redis

import (
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

// VideoFavoriteCount 视频的点赞数
func VideoFavoriteCount(videoId int64) (int64, error) {
	KeyAllVideoZSet := getKeyAllVideoZSet()
	videoIdStr := strconv.Itoa(int(videoId))
	countF, err := rdb.ZScore(KeyAllVideoZSet, videoIdStr).Result()
	if err != nil {
		log.Println("rdb.ZScore(KeyAllVideoZSet, videoIdStr) failed", err)
		return 0, err
	}
	return int64(countF), nil
}

// VideoCommentCount 获取视频的评论数
func VideoCommentCount(videoId int64) (int64, error) {
	KeyAllVideoCommentCountZSet := getKeyAllVideoCommentCountZSet()
	videoIdStr := strconv.Itoa(int(videoId))
	countF, err := rdb.ZScore(KeyAllVideoCommentCountZSet, videoIdStr).Result()
	if err != nil {
		log.Println("rdb.ZScore(KeyAllVideoCommentCountZSet, videoIdStr) failed", err)
		return 0, err
	}
	return int64(countF), nil
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
