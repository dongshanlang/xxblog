package model

import "database/sql"

var Models = []interface{}{&User{}, &Tag{}, &Article{}, &ArticleTag{}}

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

// 标签
type Tag struct {
	Model
	Name        string `gorm:"size:32;unique;not null" json:"name" form:"name"`
	Description string `gorm:"size:1024" json:"description" form:"description"`
	Status      int    `gorm:"index:idx_tag_status;not null" json:"status" form:"status"`
	CreateTime  int64  `json:"createTime" form:"createTime"`
	UpdateTime  int64  `json:"updateTime" form:"updateTime"`
}

// 文章
type Article struct {
	Model
	UserId      int64  `gorm:"index:idx_article_user_id" json:"userId" form:"userId"`             // 所属用户编号
	Title       string `gorm:"size:128;not null;" json:"title" form:"title"`                      // 标题
	Summary     string `gorm:"type:text" json:"summary" form:"summary"`                           // 摘要
	Content     string `gorm:"type:longtext;not null;" json:"content" form:"content"`             // 内容
	ContentType string `gorm:"type:varchar(32);not null" json:"contentType" form:"contentType"`   // 内容类型：markdown、html
	Status      int    `gorm:"int;not null;index:idx_article_status" json:"status" form:"status"` // 状态
	Share       bool   `gorm:"not null" json:"share" form:"share"`                                // 是否是分享的文章，如果是这里只会显示文章摘要，原文需要跳往原链接查看
	SourceUrl   string `gorm:"type:text" json:"sourceUrl" form:"sourceUrl"`                       // 原文链接
	ViewCount   int64  `gorm:"not null;index:idx_view_count;" json:"viewCount" form:"viewCount"`  // 查看数量
	CreateTime  int64  `gorm:"index:idx_article_create_time" json:"createTime" form:"createTime"` // 创建时间
	UpdateTime  int64  `json:"updateTime" form:"updateTime"`                                      // 更新时间
}

// 文章标签
type ArticleTag struct {
	Model
	ArticleId  int64 `gorm:"not null;index:idx_article_id;" json:"articleId" form:"articleId"`  // 文章编号
	TagId      int64 `gorm:"not null;index:idx_article_tag_tag_id;" json:"tagId" form:"tagId"`  // 标签编号
	Status     int64 `gorm:"not null;index:idx_article_tag_status" json:"status" form:"status"` // 状态：正常、删除
	CreateTime int64 `json:"createTime" form:"createTime"`                                      // 创建时间
}
