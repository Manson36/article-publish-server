package datamodels

import "time"

type Tag struct {
	ID           int64        `json:"id,string" gorm:"type:bigint"`
	Name         string       `json:"name" gorm:"type:varchar(200);pq_comment:标签名称"`
	Uploader     int64        `json:"uploader" gorm:"type:bigint;pq_comment:上传者id，与admin_user表中id关联"`
	PlatformType PlatformType `json:"platformType" gorm:"type:smallint;pq_comment:标签所在的平台类型"`
	CreatedAt    *time.Time   `json:"createdAt" gorm:"type:timestamptz;default:now();pq_comment:标签的创建时间"`
	UpdatedAt    *time.Time   `json:"updatedAt" gorm:"type:timestamptz;default:now();pq_comment:标签的修改时间"`
	RemovedAt    *time.Time   `json:"removedAt" gorm:"type:timestamptz;default:NULL;pq_comment:标签的移除时间"`
	Removed      bool         `json:"removed" gorm:"type:boolean;default:FALSE;pq_comment:标签是否被移除"`
}
