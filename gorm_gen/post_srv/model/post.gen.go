// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const TableNamePost = "post"

// Post 博文
type Post struct {
	ID           uuid.UUID      `gorm:"column:id;type:uuid;primaryKey;default:uuid_generate_v4();comment:自然主键" json:"id"`             // 自然主键
	CreatedAt    time.Time      `gorm:"column:created_at;type:datetime;not null;comment:创建时间" json:"created_at"`                      // 创建时间
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at;type:datetime;comment:删除时间（软删除）" json:"deleted_at"`                          // 删除时间（软删除）
	UpdatedAt    time.Time      `gorm:"column:updated_at;type:datetime;comment:更新时间" json:"updated_at"`                               // 更新时间
	Title        string         `gorm:"column:title;type:varchar(200);comment:博文标题" json:"title"`                                     // 博文标题
	CoverImage   string         `gorm:"column:cover_image;type:varchar(255);comment:博文标题配图" json:"cover_image"`                       // 博文标题配图
	UserID       string         `gorm:"column:user_id;type:char(36);not null;comment:作者id" json:"user_id"`                            // 作者id
	Summary      string         `gorm:"column:summary;type:text;not null;comment:博文概览" json:"summary"`                                // 博文概览
	Content      string         `gorm:"column:content;type:text;not null;comment:博文内容" json:"content"`                                // 博文内容
	CategoryID   string         `gorm:"column:category_id;type:char(36);not null;comment:隶属哪个归档" json:"category_id"`                  // 隶属哪个归档
	Temperature  int64          `gorm:"column:temperature;type:bigint(20) unsigned;not null;comment:博文热度（排序文章时用）" json:"temperature"` // 博文热度（排序文章时用）
	LikeCount    int64          `gorm:"column:like_count;type:bigint(20) unsigned;not null;comment:博文点赞量" json:"like_count"`          // 博文点赞量
	ViewCount    int64          `gorm:"column:view_count;type:bigint(20) unsigned;not null;comment:观看量" json:"view_count"`            // 观看量
	StarCount    int64          `gorm:"column:star_count;type:bigint(20) unsigned;not null;comment:收藏数量" json:"star_count"`           // 收藏数量
	CommentCount int32          `gorm:"column:comment_count;type:int(11);not null" json:"comment_count"`
	ShareCount   int32          `gorm:"column:share_count;type:int(11);not null;comment:分享数量" json:"share_count"`            // 分享数量
	State        int32          `gorm:"column:state;type:tinyint(4);not null;comment:文章状态：0审核中、1通过、2被举报、3热点文章" json:"state"` // 文章状态：0审核中、1通过、2被举报、3热点文章
	LinkURL      string         `gorm:"column:link_url;type:varchar(255);comment:文章外部链接" json:"link_url"`                    // 文章外部链接
}

// TableName Post's table name
func (*Post) TableName() string {
	return TableNamePost
}