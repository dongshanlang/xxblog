package model

import "database/sql"
type Model struct {
	Id int64 `gorm:"PRIMARY_KEY;AUTO_INCREMENT" json:"id" form:"id"`
}

type User struct {
	Model
	Username     sql.NullString `gorm:"size:32;unique;" json:"username" form:"username"`            // 用户名
	Email        sql.NullString `gorm:"size:128;unique;" json:"email" form:"email"`                 // 邮箱
	Nickname     string         `gorm:"size:16;" json:"nickname" form:"nickname"`                   // 昵称
	Avatar       string         `gorm:"type:text" json:"avatar" form:"avatar"`                      // 头像
	Password     string         `gorm:"size:512" json:"password" form:"password"`                   // 密码
	HomePage     string         `gorm:"size:1024" json:"homePage" form:"homePage"`                  // 个人主页
	Description  string         `gorm:"type:text" json:"description" form:"description"`            // 个人描述
	Status       int            `gorm:"index:idx_user_status;not null" json:"status" form:"status"` // 状态
	TopicCount   int            `gorm:"not null" json:"topicCount" form:"topicCount"`               // 帖子数量
	CommentCount int            `gorm:"not null" json:"commentCount" form:"commentCount"`           // 跟帖数量
	Roles        string         `gorm:"type:text" json:"roles" form:"roles"`                        // 角色
	Type         int            `gorm:"not null" json:"type" form:"type"`                           // 用户类型
	CreateTime   int64          `json:"createTime" form:"createTime"`                               // 创建时间
	UpdateTime   int64          `json:"updateTime" form:"updateTime"`                               // 更新时间
}