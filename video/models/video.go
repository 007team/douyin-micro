package models

import "time"

var LastVideoId int64 = 0

type Video struct {
	Id            int64  `json:"id,omitempty"              gorm:"primaryKey; type:bigint(20) AUTO_INCREMENT"`
	UserId        int64  `json:"-"                         gorm:"type:bigint(20)  NOT NULL"`
	Author        User   `json:"author"                    gorm:"foreignKey:UserId"`
	PlayUrl       string `json:"play_url,omitempty"        gorm:"type:varchar(255) NOT NULL"`
	CoverUrl      string `json:"cover_url,omitempty"       gorm:"type:varchar(255) NOT NULL"`
	FavoriteCount int64  `json:"favorite_count,omitempty"  gorm:"type:int  NOT NULL DEFAULT 0"`
	CommentCount  int64  `json:"comment_count,omitempty"   gorm:"type:int  NOT NULL DEFAULT 0"`
	IsFavorite    bool   `json:"is_favorite,omitempty"     gorm:"-"`
	Title         string `json:"title,omitempty"           gorm:"type:varchar(255)  collate utf8mb4_general_ci NOT NULL DEFAULT ''  "`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (Video) tableName() string {
	return "videos"
}

type User struct {
	Id            int64  `json:"id,omitempty"             gorm:"primaryKey; type:bigint(20) AUTO_INCREMENT;"`
	Name          string `json:"name,omitempty"           gorm:"uniqueIndex:idx_name; type:varchar(64) UNIQUE collate utf8mb4_general_ci not null" `
	FollowCount   int64  `json:"follow_count,omitempty"   gorm:"column:follow_count; type:INT NOT NULL DEFAULT 0 "`
	FollowerCount int64  `json:"follower_count,omitempty" gorm:"column:follower_count; type:INT NOT NULL DEFAULT 0"`
	Password      string `json:"-"                        gorm:"column:password; type:varchar(200) NOT NULL"`
	IsFollow      bool   `json:"is_follow,omitempty"      gorm:"-"`
	Salt          string `json:"-"                        gorm:"column:salt;  type:varchar(255) NOT NULL"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
