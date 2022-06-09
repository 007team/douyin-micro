package logic

import (
	"bytes"
	"context"
	"fmt"
	"github.com/007team/douyin-micro/video/dao/mysql"
	"github.com/007team/douyin-micro/video/dao/qiniu"
	"github.com/007team/douyin-micro/video/dao/redis"
	"github.com/007team/douyin-micro/video/models"
	"github.com/007team/douyin-micro/video/pkg/jwt"
	"github.com/007team/douyin-micro/video/services"
	"github.com/disintegration/imaging"
	ffmpeg_go "github.com/u2takey/ffmpeg-go"
	"io"
	"log"
	"mime/multipart"
	"os"
	"strconv"
	"sync"
	"time"
)

var (
	VideoPath      string = "D:\\GO_WORK\\src\\douyinapp\\public\\video" // 保存视频的路径
	ImgPath        string = "D:\\GO_WORK\\src\\douyinapp\\public\\img"   // 保存图片的路径
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
	// 返回30个视频
	videos, err := mysql.FindVideo()
	if err != nil {
		log.Println("mysql.FindVideo() failed")
		resp.StatusCode = 1
		resp.StatusMsg = "服务器繁忙 请稍后再试"
		return nil
	}
	videolist := make([]*services.Video, 0, len(videos))

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
		for i, video := range videos {
			// 是否被用户点赞
			ok, err := redis.IsFavoriteVideo(userId, video.Id)
			if err != nil {
				log.Println("redis.IsFavoriteVideo failed", err)
				resp.StatusCode = 1
				resp.StatusMsg = "服务器繁忙 请稍后再试"
				return nil
			}
			if ok {
				videos[i].IsFavorite = true
			}
			// 展示视频的点赞数
			videos[i].FavoriteCount, err = redis.VideoFavoriteCount(video.Id)
			if err != nil {
				log.Println("redis.VideoFavoriteCount(video.Id) failed", err)
				resp.StatusCode = 1
				resp.StatusMsg = "服务器繁忙 请稍后再试"
				return nil
			}
			// 展示视频的评论数
			videos[i].CommentCount, err = redis.VideoCommentCount(video.Id)
			if err != nil {
				log.Println("redis.VideoCommentCount(video.Id) failed", err)
				resp.StatusCode = 1
				resp.StatusMsg = "服务器繁忙 请稍后再试"
				return nil
			}
			// 脱敏处理
			videos[i].Author.Salt = ""
			videos[i].Author.Password = ""
			videolist = append(videolist, BuildVideo(videos[i]))
		}

	} else {

		// 未登录状态
		for i, video := range videos {
			// 展示视频的点赞数
			videos[i].FavoriteCount, err = redis.VideoFavoriteCount(video.Id)
			if err != nil {
				log.Println("redis.VideoFavoriteCount(video.Id) failed", err)
				resp.StatusCode = 1
				resp.StatusMsg = "服务器繁忙 请稍后再试"
				return nil
			}
			// 展示视频的评论数
			videos[i].CommentCount, err = redis.VideoCommentCount(video.Id)
			if err != nil {
				log.Println("redis.VideoCommentCount(video.Id) failed", err)
				resp.StatusCode = 1
				resp.StatusMsg = "服务器繁忙 请稍后再试"
				return nil
			}
			// 脱敏处理
			videos[i].Author.Salt = ""
			videos[i].Author.Password = ""
			videolist = append(videolist, BuildVideo(videos[i]))
		}
	}

	resp.StatusCode = 0
	resp.StatusMsg = "操作成功"
	resp.VideoList = videolist
	resp.NextTime = time.Now().Unix()
	return nil
}

func (*VideoService) PublishAction(ctx context.Context, req *services.VideoPublishActionRequest, resp *services.VideoPublishActionResponse) error {
	var video models.Video
	// 并发不安全，加锁
	GetLastIdMutex.Lock()
	video.Id = mysql.GetLastId(&models.Video{}) + 1
	GetLastIdMutex.Unlock()

	//将视频保存到本地
	if err = c.SaveUploadedFile(data, VideoPath+"\\"+strconv.Itoa(int(video.Id))+".mp4"); err != nil {
		fmt.Println("c.SaveUploadedFile failed", err)
		return err
	}
	fmt.Println("保存视频完成")

	/*
		生成缩略图 （视频封面）
	*/
	snapshotName, err := GetSnapshot(VideoPath+`\`+strconv.Itoa(int(video.Id))+".mp4", ImgPath, 5, video.Id)
	if err != nil {
		fmt.Println("缩略图生成失败", err)
		return err
	}
	fmt.Println("生成缩略图完成")

	/*
		上传视频到七牛云
	*/
	_, fileUrl, err := qiniu.UploadVideoToQiNiu(data, video.Id)
	if err != nil {
		fmt.Println("qiniu upload video failed")
		return
	}
	fmt.Println("上传视频完成")
	video.PlayUrl = fileUrl

	//上传封面到七牛云
	coverUrl := qiniu.UploadImgToQiNiu(snapshotName, ImgPath, video.Id)
	video.CoverUrl = coverUrl
	fmt.Println("上传封面到七牛云完成")

	// 添加videoId到 视频数zset 和 评论数zset
	err = redis.Publish(video.Id)
	if err != nil {
		fmt.Println("redis.Publish(video.Id) failed")
		return
	}
	go func() {
		/*
			删除本地视频
		*/
		VideoFilePath := VideoPath + "\\" + strconv.Itoa(int(video.Id)) + ".mp4"
		err = os.Remove(VideoFilePath)
		if err != nil {
			fmt.Println("本地视频文件删除失败", err)
			return
		}
		fmt.Println("本地视频文件删除完成")

		/*
			将本地封面删除
		*/
		CoverFilePath := ImgPath + "\\" + strconv.Itoa(int(video.Id)) + ".jpeg"
		err = os.Remove(CoverFilePath)
		if err != nil {
			fmt.Println("本地封面缩略图删除失败")
			return
		}
		fmt.Println("本地封面缩略图删除完成")
	}()

	return mysql.CreateNewVideo(video)
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
	// mysql查询视频数据
	if len(es) == 0 {
		resp.StatusCode = 0
		resp.StatusMsg = "操作成功"
		return nil
	}

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

func SaveUploadedFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}
