package util

import (
	"bytes"
	"fmt"
	"github.com/disintegration/imaging"
	ffmpeg_go "github.com/u2takey/ffmpeg-go"
	"os"
)

//  生成缩略图
func GetSnapshot(videoPath, FileName, snapshotPath string, frameNum int) (snapshotName string, err error) {
	buf := bytes.NewBuffer(nil)

	err = ffmpeg_go.Input(videoPath+"\\video=="+FileName+".mp4").
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

	err = imaging.Save(img, snapshotPath+`\`+"cover=="+FileName+".jpeg")
	if err != nil {
		fmt.Println("生成缩略图失败：", err)
		return "", err
	}

	// 成功则返回生成的缩略图名
	//names := strings.Split(snapshotPath, `\`)
	snapshotName = "cover==" + FileName + ".jpeg"
	fmt.Println("缩略图名是：", snapshotName)
	return
}
