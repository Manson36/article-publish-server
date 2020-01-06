package datamodels

import "time"

type AdminUser struct {
	ID           int64        `json:"id,string"` //???string,所有int64类型都带了;2.type如何选择
	NickName     string       `json:"nickName" gorm:"type:varchar(50);not null;default:'';pq_comment:用户昵称"`
	Email        string       `json:"email" gorm:"type:varchar(200);not null;unique;pq_comment:用户邮箱"`
	Password     string       `json:"password" gorm:"type:varchar(200);not null;unique;pq_comment:用户密码"`
	Salt         string       `json:"salt" gorm:"type:varchar(5);not null;pq_comment:加密密码的盐"`
	AdminType    int          `json:"adminType" gorm:"type:smallint;not null;default:2;pq_comment:管理端用户类型，1表示管理员，2表示普通管理者"`
	PlatformType PlatformType `json:"platformType" gorm:"type:smallint;pq_comment:管理员所处的平台类型"`
	CreateAt     *time.Time   `json:"createAt" gorm:"type:timestamptz;not null;default:now();pq_comment:该用户的创建时间"`
	UpdateAt     *time.Time   `json:"updateAt" gorm:"type:timestamptz;default:now();pq_comment:该用户的更新时间"`
	RemovedAt    *time.Time   `json:"removedAt" gorm:"type:timestamptz;pq_comment:用户的移除时间"`
	Removed      bool         `json:"removed" gorm:"pq_comment:该用户是否被移除"`
	Disabled     bool         `json:"disabled" gorm:"pq_comment:该用户是否被禁用"`
	DisabledAt   *time.Time   `json:"disabledAt" gorm:"type:timestamptz;pq_comment:该用户被禁用的时间"`
}
