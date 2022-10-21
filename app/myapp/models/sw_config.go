package models

import (
	"go-admin/common/models"
)

type SwConfig struct {
	models.Model

	Code    string `json:"code" gorm:"type:varchar(128);comment:配置标识"`
	Content string `json:"content" gorm:"type:json;comment:配置内容"`
	models.ModelTime
	models.ControlBy
}

func (SwConfig) TableName() string {
	return "sw_config"
}

func (e *SwConfig) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *SwConfig) GetId() interface{} {
	return e.Id
}
