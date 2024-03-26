package model

import "time"

type Collect struct {
	ID        uint      `json:"CollectId" gorm:"column:Collect_id;"`
	Type      string    `json:"CollectType"` //收藏类型 1 商品 2 文章 3 视频
	User      User      `gorm:"ForeignKey:UserID"`
	UserID    uint      `gorm:"not null"` //收藏人ID
	Boss      User      `gorm:"ForeignKey:BossID"`
	BossID    uint      `gorm:"not null"` //被收藏人ID
	Product   Product   `gorm:"ForeignKey:ProductID"`
	ProductID uint      `gorm:"not null"`   //商品ID
	CreatedAt time.Time `json:"createTime"` //字段使用CreatedAt，不是CreatedTime
	UpdatedAt time.Time `json:"updateTime"`
}
