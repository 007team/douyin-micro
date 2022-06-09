package handlers

import (
	"github.com/007team/douyin-micro/gateway/pkg/qiniu"
	"mime/multipart"
)

type videofile struct {
	data     *multipart.FileHeader
	filename string
}

type VideoUpload struct {
	VideoChan chan videofile
}

type ImgUpload struct {
	ImgChan chan string
}

var (
	VideoProcess VideoUpload
	ImgProcess   ImgUpload
)

// 上传视频到七牛云

func VideoUploadFunc() {
	VideoProcess.VideoChan = make(chan videofile, 100)
	for videofile := range VideoProcess.VideoChan {
		qiniu.UploadVideoToQiNiu(videofile.data, videofile.filename)
	}

}
func ImgUploadFunc() {
	//上传封面到七牛云
	ImgProcess.ImgChan = make(chan string, 100)
	for snapshotName := range ImgProcess.ImgChan {
		qiniu.UploadImgToQiNiu(snapshotName, ImgPath)
	}

}
