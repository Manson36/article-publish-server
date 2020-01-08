package datamodels

import (
	"database/sql/driver"
	"encoding/json"
	"github.com/jinzhu/gorm"
	pg "github.com/jinzhu/gorm/dialects/postgres"
)

type PlatformType int8

const (
	ZingglobalPlatform   PlatformType = 1
	ZhidreamPlatform     PlatformType = 2
	HealthEnginePlatform PlatformType = 3
)

type JsonNumArray []int64

//???What are doing?
func (t JsonNumArray) Value() (driver.Value, error) {
	buf, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	j := pg.Jsonb{RawMessage: buf}
	return j.Value()
}
func (t *JsonNumArray) Scan(value interface{}) error {
	j := &pg.Jsonb{}
	err := j.Scan(value)
	if err != nil {
		return err
	}
	return json.Unmarshal(j.RawMessage, t)
}

func GetModelList() []interface{} {
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "t_ow_" + defaultTableName
	}

	return []interface{}{
		&AdminUser{},
		&Article{},
		&File{},
		&Image{},
		&Tag{},
	}
}
