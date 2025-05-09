package model

import "time"

type UserModel struct {
	UserID        int32  `gorm:"primaryKey;colum:user_id"`
	DisplayUserID string `gorm:"column:display_user_id;unique;not null"`
	UserName      string `gorm:"column:user_name;not null"`
	ProfileText   string `gorm:"column:profile_text;not null"`
	CreatedAt     time.Time `gorm:"column:created_at;not null"`
}

func (UserModel) TableName() string {
	return "user"
}
