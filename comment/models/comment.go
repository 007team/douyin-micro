package models

type Comment struct {
	Id        int64    `json:"id,omitempty" gorm:"primaryKey; type:bigint(20) AUTO_INCREMENT"`
	UserId    int64    `json:"user_id" gorm:"type:bigint(20) not null"`
	VideoId   int64    `json:"video_id"  gorm:"type:bigint(20)"`
	User      User     `json:"user" gorm:"foreignKey:UserId"`
	Content   string   `json:"content,omitempty" gorm:"type:mediumtext collate utf8mb4_general_ci NOT NULL"`
	CreatedAt JSONTime `json:"create_date"`
	UpdatedAt JSONTime `json:"update_date"`
}

func (Comment) tableName() string {
	return "comments"

}
