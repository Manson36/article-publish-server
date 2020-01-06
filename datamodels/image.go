package datamodels

import "time"

type Image struct {
	ID           int64        `json:"id, string"`
	Name         string       `json:"name" gorm:"type:varchar(200);pq_comment:图片名称"`
	Link         int64        `json:"link, string" gorm:"type:bigint;pq_comment:图片关联file表中的id"`
	Path         string       `json:"path" gorm:"type:varchar(200);pq_comment:图片的存储路径"`
	Uploader     int64        `json:"uploader,string" gorm:"type:bigint;pq_comment:图片的上传者id，与admin_user表中id关联"`
	Width        int64        `json:"width" gorm:"type:bigint;pq_comment:图片的宽度"`
	Height       int64        `json:"height" gorm:"type:bigint;pq_comment:图片的高度"`
	PlatformType PlatformType `json:"platformType" gorm:"type:smallint;pq_comment:图片所在的平台类型"`
	CreatedAt    *time.Time   `json:"createdAt" gorm:"type:timestamptz;default:now();pq_comment:图片的创建时间"`
	UpdatedAt    *time.Time   `json:"updatedAt" gorm:"type:timestamptz;default:now();pq_comment:图片的更新时间"`
	RemovedAt    *time.Time   `json:"removedAt" gorm:"type:timestamptz;default:NULL;pq_comment:图片的移除时间"`
	Remove       bool         `json:"remove" gorm:"type:boolean;default:FALSE;pq_comment:图片是否被移除"`
}
