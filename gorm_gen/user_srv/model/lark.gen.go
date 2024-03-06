// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const TableNameLark = "lark"

// Lark 用户表
type Lark struct {
	ID            uuid.UUID      `gorm:"column:id;type:uuid;primaryKey;default:uuid_generate_v4();comment:自然主键" json:"id"`                                      // 自然主键
	CreatedAt     time.Time      `gorm:"column:created_at;type:datetime;not null;comment:创建时间" json:"created_at"`                                               // 创建时间
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at;type:datetime;comment:删除时间（软删除）" json:"deleted_at"`                                                   // 删除时间（软删除）
	UpdatedAt     time.Time      `gorm:"column:updated_at;type:datetime;comment:更新时间" json:"updated_at"`                                                        // 更新时间
	StuNum        string         `gorm:"column:stu_num;type:char(8);not null;comment:学号" json:"stu_num"`                                                        // 学号
	Password      string         `gorm:"column:password;type:char(255);not null;comment:密码" json:"password"`                                                    // 密码
	Name          string         `gorm:"column:name;type:varchar(30);not null;comment:姓名" json:"name"`                                                          // 姓名
	Gender        *string        `gorm:"column:gender;type:varchar(10);not null;default:保密;comment:用户性别：女，男，其他，保密" json:"gender"`                               // 用户性别：女，男，其他，保密
	College       string         `gorm:"column:college;type:varchar(30);not null;comment:用户所在学院" json:"college"`                                                // 用户所在学院
	Major         string         `gorm:"column:major;type:varchar(30);not null;comment:用户专业" json:"major"`                                                      // 用户专业
	Grade         string         `gorm:"column:grade;type:varchar(10);not null;comment:用户年级：大一，大二，大三，大四，研究生,毕业生" json:"grade"`                                  // 用户年级：大一，大二，大三，大四，研究生,毕业生
	StuCardURL    string         `gorm:"column:stu_card_url;type:varchar(255);not null;comment:学生证照片url" json:"stu_card_url"`                                   // 学生证照片url
	Phone         string         `gorm:"column:phone;type:char(11);not null;comment:用户手机号" json:"phone"`                                                        // 用户手机号
	Province      string         `gorm:"column:province;type:varchar(10);comment:用户家乡省份" json:"province"`                                                       // 用户家乡省份
	Age           int32          `gorm:"column:age;type:tinyint(4);comment:用户年龄" json:"age"`                                                                    // 用户年龄
	PhotoURL      string         `gorm:"column:photo_url;type:varchar(255);comment:照片url" json:"photo_url"`                                                     // 照片url
	Email         string         `gorm:"column:email;type:varchar(255);comment:邮箱地址" json:"email"`                                                              // 邮箱地址
	Introduction  string         `gorm:"column:introduction;type:text;comment:用户个人介绍" json:"introduction"`                                                      // 用户个人介绍
	Avatar        *string        `gorm:"column:avatar;type:varchar(255);default:https://static.skylab.org.cn/default/avatar.png;comment:用户头像url" json:"avatar"` // 用户头像url
	QqUnionID     string         `gorm:"column:qq_union_id;type:varchar(255);comment:qq社会化登录" json:"qq_union_id"`                                               // qq社会化登录
	WechatUnionID string         `gorm:"column:wechat_union_id;type:varchar(255);comment:微信社会化登录" json:"wechat_union_id"`                                       // 微信社会化登录
	State         int32          `gorm:"column:state;type:tinyint(4);not null;comment:用户状态：0禁用、1审核中、2启用、3其他" json:"state"`                                      // 用户状态：0禁用、1审核中、2启用、3其他
}

// TableName Lark's table name
func (*Lark) TableName() string {
	return TableNameLark
}
