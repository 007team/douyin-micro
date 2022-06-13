package cron

import (
	"github.com/007team/douyin-micro/video/dao/mysql"
	"github.com/007team/douyin-micro/video/dao/redis"
	"github.com/007team/douyin-micro/video/models"
	"github.com/robfig/cron"
	"log"
	"strconv"
)

func taskUserFollowCount() {
	// 获取有哪些用户
	user := &models.User{}
	mysql.DB().Last(&user)

	// 取redis中的数据
	var i int64
	for i = 1; i <= user.Id; i++ {
		KeyUserFollowSet := redis.CronKeyUserFollowSet(i)
		rstFollow, _ := redis.RDB().SCard(KeyUserFollowSet).Result() // 关注数
		mysql.DB().Model(&models.User{}).Where("id = ?", i).Update("follow_count", rstFollow)
	}
}
func taskUserFollowerCount() {
	// 获取有哪些用户
	user := &models.User{}
	mysql.DB().Last(&user)

	// 取redis中的数据
	var i int64
	for i = 1; i <= user.Id; i++ {
		KeyUserFollowerSet := redis.CronKeyUserFollowerSet(i)
		rstFollower, _ := redis.RDB().SCard(KeyUserFollowerSet).Result() // 粉丝数
		mysql.DB().Model(&models.User{}).Where("id = ?", i).Update("follower_count", rstFollower)
	}
}

func taskVideoFavoriteCount() {
	// 查询有哪些视频
	video := &models.Video{}
	mysql.DB().Last(&video)

	// 查询视频对应的点赞数
	KeyAllVideoZSet := redis.CronKeyAllVideoZSet()
	var i int64
	for i = 1; i <= video.Id; i++ {
		videoFavoriteCount, _ := redis.RDB().ZScore(KeyAllVideoZSet, strconv.Itoa(int(i))).Result()
		mysql.DB().Model(&models.Video{}).Where("id = ?", i).Update("favorite_count", videoFavoriteCount)
	}
}

func taskVideoCommentCount() {
	// 查询有哪些视频
	video := &models.Video{}
	mysql.DB().Last(&video)

	// 查询视频对应的点赞数
	KeyAllVideoCommentCountZSet := redis.CronKeyAllVideoCommentCountZSet()
	var i int64
	for i = 1; i <= video.Id; i++ {
		VideoCommentCountZSet, _ := redis.RDB().ZScore(KeyAllVideoCommentCountZSet, strconv.Itoa(int(i))).Result()
		mysql.DB().Model(&models.Video{}).Where("id = ?", i).Update("comment_count", VideoCommentCountZSet)
	}
}

func CronTask() {
	log.Println("crontask start....")
	c := cron.New()
	c.AddFunc("*/5 * * * *", taskUserFollowCount)
	c.AddFunc("*/5 * * * *", taskUserFollowerCount)
	c.AddFunc("*/5 * * * *", taskVideoFavoriteCount)
	c.AddFunc("*/5 * * * *", taskVideoCommentCount)

	c.Start()

	select {}
}
