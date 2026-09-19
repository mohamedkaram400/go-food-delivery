package entity

import "time"


type User struct {
	ID                      int    `gorm:"primaryKey;autoIncrement" json:"id"`
    Name                    string `json:"name"`
    Email                   string `json:"email"`
    Phone                   string `json:"phone"`
    Status                  string `json:"status"`
    PasswordHash            string `json:"password_hash"`
    RoleID                  int       `json:"role_id"`
    EmailVerifiedAt         *time.Time    `json:"email_verified_at"`
	CreatedAt               time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt               time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}   