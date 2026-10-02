package entity

import "time"


type RefreshToken struct {
	ID                      int       `gorm:"primaryKey;autoIncrement" json:"id"`
    UserID                  int       `json:"user_id"    gorm:"not null"`
    TokenHash 				string    `json:"token_hash" gorm:"column:token_hash;not null"`
	RevokedAt 				*time.Time `json:"revoked_at" gorm:"column:revoked_at;default:null"`
	ExpiresAt               time.Time `json:"expires_at" gorm:"not null"`
	CreatedAt               time.Time `json:"created_at" gorm:"autoCreateTime"`
}   

func (RefreshToken) TableName() string {
    return "refresh_tokens"
}