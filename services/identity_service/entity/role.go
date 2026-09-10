package entity

import "time"


type Role struct {
	ID                      int    `gorm:"primaryKey;autoIncrement" json:"id"`
    Name                    string `json:"name"`
	CreatedAt               time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt               time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}   