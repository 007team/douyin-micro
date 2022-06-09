package mysql

import (
	"github.com/007team/douyin-micro/video/models"
	"log"
)

func FindVideo() (videos []models.Video, err error) {

	//

	err = db.Preload("Author").Order("updated_at DESC").Limit(30).Find(&videos).Error

	if err != nil {
		log.Println(err)
		return nil, err
	}
	return videos, nil

}

func GetVideoArr(user_id int64) (VideoArr []models.Video) {
	db.Preload("Author").Find(&VideoArr, "user_id = ?", user_id)

	return
}

func FavoriteList(es []string) (videos []models.Video, err error) {
	err = db.Find(&videos, es).Error
	if err != nil {
		log.Println("db.Find failed", err)
	}
	return videos, err
}

// GetLastId 获取最后一位视频id
func GetLastId(video *models.Video) (id int64) {
	db.Last(&video)
	return video.Id
}

// CreateNewVideo
func CreateNewVideo(video *models.Video) (err error) {
	if err = db.Select("user_id", "play_url", "cover_url", "title").Create(video).Error; err != nil {
		log.Fatalln("mysql.CreateNewVideo failed", err)
		return
	}

	return nil
}
