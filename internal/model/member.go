package model

import (
	"gorm.io/gorm"
	"time"
)

type Members struct {
	gorm.Model
	Status       uint8     `json:"status,omitempty" gorm:"default:1;comment: Status"`     // http method
	Username     string    `json:"username,omitempty" gorm:"comment: Username"`           // Username
	Password     string    `json:"password,omitempty" gorm:"comment: password"`           // password
	Nickname     string    `json:"nickname,omitempty" gorm:"comment: Nickname"`           // Nickname
	RankID       uint64    `json:"rank_id,omitempty" gorm:"comment: Status"`              // RankID
	Mobile       string    `json:"mobile,omitempty" gorm:"comment: Mobile"`               // Mobile
	Email        string    `json:"email,omitempty" gorm:"comment: Email"`                 // Email
	Avatar       string    `json:"avatar,omitempty" gorm:"comment: Avatar"`               // Avatar
	WechatOpenID string    `json:"wechat_open_id,omitempty" gorm:"comment: WechatOpenID"` // WechatOpenID
	ExpiredAt    time.Time `json:"expired_at,omitempty" gorm:"comment: ExpiredAt"`        // ExpiredAt
}
