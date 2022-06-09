package handlers

import (
	"context"
	"fmt"
	"github.com/007team/douyin-micro/gateway/pkg/util"
	"github.com/007team/douyin-micro/gateway/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

var (
	VideoPath string = "D:\\GO_WORK\\src\\douyin-micro\\video\\public\\video" // 保存视频的路径
	ImgPath   string = "D:\\GO_WORK\\src\\douyin-micro\\video\\public\\img"   // 保存图片的路径
	Url              = "rchgbnnln.hn-bkt.clouddn.com"                         // 域名
)

// feed流
func Feed(c *gin.Context) {
	var videoReq services.VideoFeedRequest
	videoReq.Token = c.Query("token")

	// 从gin.key 中取出服务实例
	videoService := c.Keys["videoService"].(services.VideoService)
	videoResp, err := videoService.Feed(context.Background(), &videoReq)
	PanicIfVideoError(err)
	c.JSON(http.StatusOK, videoResp)

}

// 发布视频
func Publish(c *gin.Context) {

	userId, ok := c.Get("user_id")
	if !ok {
		log.Println("c.Get(\"user_id\")")
		return
	}
	title := c.PostForm("title")

	// 获取视频数据
	data, err := c.FormFile("data")
	if err != nil {
		log.Println("c.FormFile(\"data\") failed")
		return
	}

	//	将视频保存到本地
	fileName := strconv.Itoa(int(userId.(int64))) + data.Filename + strconv.FormatInt(time.Now().Unix(), 10)
	videoFileName := "video==" + fileName + ".mp4"
	if err := c.SaveUploadedFile(data, VideoPath+"\\"+videoFileName); err != nil {
		fmt.Println("c.SaveUploadedFile failed", err)
		return
	}
	//  生成缩略图 （视频封面）
	coverFileName := "cover==" + fileName + ".jpeg"
	snapshotName, err := util.GetSnapshot(VideoPath, fileName, ImgPath, 1)
	if err != nil {
		fmt.Println("缩略图生成失败", err)
		return
	}

	v := videofile{
		data:     data,
		filename: videoFileName,
	}

	VideoProcess.VideoChan <- v
	ImgProcess.ImgChan <- snapshotName

	video := services.Video{
		Author:   &services.User{Id: userId.(int64)},
		PlayUrl:  "http://" + Url + "/videos/" + videoFileName,
		CoverUrl: "http://" + Url + "/cover/" + coverFileName,
		Title:    title,
	}

	var videoReq services.VideoPublishActionRequest
	videoReq.UserId = userId.(int64)
	videoReq.Video = &video

	videoService := c.Keys["videoService"].(services.VideoService)
	videoResp, err := videoService.PublishAction(c, &videoReq)
	PanicIfVideoError(err)
	c.JSON(http.StatusOK, videoResp)
}

// 用的发布视频列表
func PublishList(c *gin.Context) {
	var videoReq services.VideoPublishListRequest
	userIdStr := c.Query("user_id")
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		log.Println("strconv.ParseInt failed")
		return
	}
	videoReq.UserId = userId
	videoService := c.Keys["videoService"].(services.VideoService)
	videoResp, err := videoService.PublishList(context.Background(), &videoReq)
	PanicIfVideoError(err)
	c.JSON(http.StatusOK, videoResp)

}

// 点赞视频列表
func FavoriteList(c *gin.Context) {
	var videoReq services.VideoFavoriteListRequest
	userId, ok := c.Get("user_id")
	if !ok {
		log.Println("c.Get(\"user_id\") failed")
	}
	videoReq.UserId = userId.(int64)
	videoService := c.Keys["videoService"].(services.VideoService)
	videoResp, err := videoService.FavoriteList(context.Background(), &videoReq)
	PanicIfVideoError(err)
	c.JSON(http.StatusOK, videoResp)

}
