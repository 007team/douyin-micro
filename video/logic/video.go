package logic

import (
	"bytes"
	"context"
	"fmt"
	"github.com/007team/douyin-micro/video/dao/mysql"
	"github.com/007team/douyin-micro/video/dao/redis"
	"github.com/007team/douyin-micro/video/models"
	"github.com/007team/douyin-micro/video/pkg/jwt"
	"github.com/007team/douyin-micro/video/services"
	"github.com/disintegration/imaging"
	ffmpeg_go "github.com/u2takey/ffmpeg-go"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

var (
	GetLastIdMutex sync.Mutex
)

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

func BuildVideo(item models.Video) *services.Video {
	videoModel := services.Video{
		Id:            item.Id,
		Author:        BuildUser(item.Author),
		PlayUrl:       item.PlayUrl,
		CoverUrl:      item.CoverUrl,
		FavoriteCount: item.FavoriteCount,
		CommentCount:  item.CommentCount,
		IsFavorite:    item.IsFavorite,
		Title:         item.Title,
	}
	return &videoModel
}

func (*VideoService) Feed(ctx context.Context, req *services.VideoFeedRequest, resp *services.VideoFeedResponse) error {
	// 查询视频
	videos, err := mysql.FindVideo()
	if err != nil {
		log.Println("mysql.FindVideo() failed")
		resp.StatusCode = 1
		resp.StatusMsg = "服务器繁忙 请稍后再试"
		return nil
	}
	videolist := make([]*services.Video, len(videos))

	wg := sync.WaitGroup{}
	wg.Add(len(videos))
	errChan := make(chan error, 100)

	if len(req.Token) != 0 {
		// 登录状态
		m, err := jwt.ParseToken(req.Token)
		userId := m.UserID
		if err != nil {
			log.Println("jwt.ParseToken(req.Token) failed")
			resp.StatusCode = 1
			resp.StatusMsg = "服务器繁忙 请稍后再试"
			return nil
		}
		for i := range videos {
			go func(i int) {
				// 是否被用户点赞
				ok, err := redis.IsFavoriteVideo(userId, videos[i].Id)
				if err != nil {
					log.Println("redis.IsFavoriteVideo failed", err)
					errChan <- err
				}
				if ok {
					videos[i].IsFavorite = true
				}

				// “我”是否关注了这个用户
				videos[i].Author.IsFollow, err = redis.IsFollowUser(&videos[i].Author, userId)
				if err != nil {
					log.Println("redis.IsFollowUser(user, myUserId) failed", err)
					errChan <- err
				}

				videolist[i] = BuildVideo(videos[i])
				wg.Done()
			}(i)
		}
	} else {
		// 未登录状态
		for i, _ := range videos {
			go func(i int) {
				videolist[i] = BuildVideo(videos[i])
				wg.Done()
			}(i)
		}

	}

	wg.Wait()
	if len(errChan) != 0 {
		resp.StatusCode = 1
		resp.StatusMsg = "服务器繁忙 请稍后再试"
		return nil
	}
	resp.StatusCode = 0
	resp.StatusMsg = "操作成功"
	resp.VideoList = videolist
	resp.NextTime = time.Now().Unix()
	return nil
}

func (*VideoService) PublishAction(ctx context.Context, req *services.VideoPublishActionRequest, resp *services.VideoPublishActionResponse) error {
	video := models.Video{
		UserId:   req.UserId,
		PlayUrl:  req.Video.PlayUrl,
		CoverUrl: req.Video.CoverUrl,
		Title:    req.Video.Title,
	}

	err := mysql.CreateNewVideo(&video)
	if err != nil {
		log.Println("mysql.CreateNewVideo(&video)")
		resp.StatusCode = 1
		resp.StatusMsg = "服务繁忙 请稍后再试"
		return nil
	}

	// 添加videoId到 视频数zset 和 评论数zset
	err = redis.Publish(video.Id)
	if err != nil {
		log.Println("redis.Publish(video.Id) failed")
		resp.StatusCode = 1
		resp.StatusMsg = "服务繁忙 请稍后再试"
		return nil
	}

	resp.StatusCode = 0
	resp.StatusMsg = "发布成功"
	return nil
}

func (*VideoService) PublishList(ctx context.Context, req *services.VideoPublishListRequest, resp *services.VideoPublishListResponse) error {
	VideoArr := mysql.GetVideoArr(req.UserId)
	videolist := make([]*services.Video, 0, len(VideoArr))
	for i, _ := range VideoArr {
		videolist = append(videolist, BuildVideo(VideoArr[i]))
	}
	resp.StatusCode = 0
	resp.StatusMsg = "操作成功"
	resp.VideoList = videolist
	return nil
}

func (*VideoService) FavoriteAction(ctx context.Context, req *services.VideoFavoriteActionRequest, resp *services.VideoFavoriteActionResponse) error {
	if req.ActionType == 1 {
		// 赞
		err := redis.FavoriteAction(req.UserId, req.VideoId)
		if err != nil {
			log.Println("redis.FavoriteAction(req.UserId, req.VideoId) failed")
			resp.StatusCode = 1
			resp.StatusMsg = "操作失败"
			return nil
		}
	} else if req.ActionType == 2 {
		// 取消赞
		err := redis.UnFavoriteAction(req.UserId, req.VideoId)
		if err != nil {
			log.Println("redis.UnFavoriteAction(req.UserId, req.VideoId) failed")
			resp.StatusCode = 1
			resp.StatusMsg = "操作失败"
			return nil
		}
	}

	resp.StatusCode = 0
	resp.StatusMsg = "操作成功"
	return nil
}

func (*VideoService) FavoriteList(ctx context.Context, req *services.VideoFavoriteListRequest, resp *services.VideoFavoriteListResponse) error {
	// 从redis获取用户的点赞视频列表
	es, err := redis.FavoriteList(req.UserId)
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = "服务繁忙 请稍后再试"
		return nil
	}

	if len(es) == 0 {
		resp.StatusCode = 0
		resp.StatusMsg = "操作成功"
		return nil
	}
	// mysql查询视频数据
	videos, err := mysql.FavoriteList(es)
	if err != nil {
		log.Println("mysql.FavoriteList(es) failed")
		resp.StatusCode = 1
		resp.StatusMsg = "服务繁忙 请稍后再试"
		return nil
	}
	videolist := make([]*services.Video, 0, len(videos))
	for i, _ := range videos {
		videolist = append(videolist, BuildVideo(videos[i]))
	}

	resp.StatusCode = 0
	resp.StatusMsg = "操作成功"
	resp.VideoList = videolist
	return nil
}

//  生成缩略图
func GetSnapshot(videoPath, snapshotPath string, frameNum int, video_id int64) (snapshotName string, err error) {
	buf := bytes.NewBuffer(nil)

	err = ffmpeg_go.Input(videoPath).
		Filter("select", ffmpeg_go.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg_go.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()

	if err != nil {
		fmt.Println("生成缩略图失败：", err)
		return "", err
	}

	img, err := imaging.Decode(buf)
	if err != nil {
		fmt.Println("生成缩略图失败：", err)
		return "", err
	}

	err = imaging.Save(img, snapshotPath+`\`+strconv.Itoa(int(video_id))+".jpeg")
	if err != nil {
		fmt.Println("生成缩略图失败：", err)
		return "", err
	}

	// 成功则返回生成的缩略图名
	//names := strings.Split(snapshotPath, `\`)
	snapshotName = strconv.Itoa(int(video_id)) + ".jpeg"
	fmt.Println("缩略图名是：", snapshotName)
	return
}
