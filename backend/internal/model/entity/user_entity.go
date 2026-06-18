package entity

import "time"

type UserEntity struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Username  string    `json:"username" gorm:"type:nvarchar(100)"`
	Email     string    `json:"email" gorm:"unique"`
	Url       string    `json:"url" gorm:"nvarchar(255)"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (UserEntity) TableName() string {
	return "users"
}
