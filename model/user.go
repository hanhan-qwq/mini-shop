package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:50;uniqueIndex;not null" json:"name"` // 用户名：长度50、唯一索引、非空
	Password string `gorm:"size:100;not null" json:"password"`        // 密码：长度100（存储哈希后的值）、非空
	Email    string `gorm:"size:100;uniqueIndex" json:"email"`        // 邮箱：唯一索引（允许为空，可选填）
	Phone    string `gorm:"size:20;uniqueIndex" json:"phone"`         // 手机号：唯一索引（允许为空，可选填）
	Avatar   string `gorm:"size:255" json:"avatar"`                   // 头像URL：长度255（URL可能较长）
	Role     string `gorm:"default:user" json:"role"`                 // 是否管理员：默认值user
}
