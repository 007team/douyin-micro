package mysql

import (
	"github.com/007team/douyin-micro/comment/models"
	"log"
)

// 获取评论列表
func GetCommentList(videoId int64) (CommentList []models.Comment) {

	err := Db.Preload("User").Where("video_id = ?", videoId).Order("updated_at DESC").Find(&CommentList).Error
	if err != nil {
		log.Println("dao.GetCommentList error:", err)
	}
	return
}

// 增加评论
func AddComment(comment *models.Comment) (err error) {
	if err = Db.Preload("User").Create(comment).Error; err != nil {
		log.Println("mysql.comment.Addcomment error", err)
		return err
	}
	return nil
}

// 增加视频评论数
//func AddVideoCommentCount(videoId int64) (err error) {
//
//	var video models.Video
//	db.Preload("User").Where("id = ?", videoId).First(&video)
//	video.CommentCount++
//
//	db.Save(&video)
//
//	return
//}

// 删除评论
func DelComment(comment *models.Comment) (err error) {
	if err = Db.Where("id = ?", comment.Id).Delete(comment).Error; err != nil {
		log.Println("mysql.comment.DelComment error", err)
		return err
	}

	return nil
}


// 减少视频评论数
//func SubVideoCommentCount(videoId int64) (err error) {
//
//	var video models.Video
//
//	db.Preload("User").Where("id = ?", videoId).First(&video)
//	video.CommentCount--
//
//	db.Save(&video)
//
//	return
//}



//// 增加评论
//func AddComment(comment *models.Comment)(err error){
//	// 增加评论
//	tx := db.Begin()
//	if err = tx.Preload("Author").Create(comment).Error; err!=nil{
//		log.Println("mysql.comment.Addcomment error",err)
//		tx.Rollback()
//		return err
//	}
//	// 增加视频评论数
//	var video models.Video
//	if err = tx.Preload("Author").Where("id = ?", comment.VideoId).First(&video).Error; err!=nil{
//		log.Println("mysql.comment.Addcomment error",err)
//		tx.Rollback()
//		return err
//	}
//	video.CommentCount++
//
//	if err = tx.Save(&video).Error; err!=nil{
//		log.Println("mysql.comment.Addcomment error",err)
//		tx.Rollback()
//		return err
//	}
//	tx.Commit()
//
//	return nil
//}
//
//// 删除评论
//func DelComment(comment *models.Comment,videoId int64)(err error){
//	tx := db.Begin()
//	// 减少视频评论数
//	var video models.Video
//
//	if err = tx.Preload("Author").Where("id = ?", videoId).First(&video).Error; err!=nil{
//		log.Println("mysql.comment.DelComment error",err)
//		tx.Rollback()
//		return err
//	}
//	video.CommentCount--
//
//	if err = tx.Save(&video).Error; err!=nil{
//		log.Println("mysql.comment.DelComment error",err)
//		tx.Rollback()
//		return err
//	}
//
//	// 删除评论
//	if err = tx.Where("id = ?",comment.Id).Delete(comment).Error; err!=nil{
//		log.Println("mysql.comment.DelComment error",err)
//		tx.Rollback()
//		return err
//	}
//
//	tx.Commit()
//	return nil
//}
