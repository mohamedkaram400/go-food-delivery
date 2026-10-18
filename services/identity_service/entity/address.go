package entity

import "time"


type Address struct {
	ID                      int       `gorm:"primaryKey;autoIncrement" json:"id"`
    Address                 string    `json:"name"`
    UserID                  int       `json:"user_id"`
	CreatedAt               time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt               time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}   