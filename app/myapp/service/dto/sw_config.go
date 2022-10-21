package dto

import (
	"go-admin/app/myapp/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type SwConfigGetPageReq struct {
	dto.Pagination `search:"-"`
	Code           string `form:"code"  search:"type:exact;column:code;table:sw_config" comment:"配置标识"`
	SwConfigOrder
}

type SwConfigOrder struct {
	Id        string `form:"idOrder"  search:"type:order;column:id;table:sw_config"`
	Code      string `form:"codeOrder"  search:"type:order;column:code;table:sw_config"`
	Content   string `form:"contentOrder"  search:"type:order;column:content;table:sw_config"`
	CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:sw_config"`
	UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:sw_config"`
	DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:sw_config"`
	CreateBy  string `form:"createByOrder"  search:"type:order;column:create_by;table:sw_config"`
	UpdateBy  string `form:"updateByOrder"  search:"type:order;column:update_by;table:sw_config"`
}

func (m *SwConfigGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type SwConfigInsertReq struct {
	Id      int    `json:"-" comment:""` //
	Code    string `json:"code" comment:"配置标识"`
	Content string `json:"content" comment:"配置内容"`
	common.ControlBy
}

func (s *SwConfigInsertReq) Generate(model *models.SwConfig) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Code = s.Code
	model.Content = s.Content
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *SwConfigInsertReq) GetId() interface{} {
	return s.Id
}

type SwConfigUpdateReq struct {
	Id      int    `uri:"id" comment:""` //
	Code    string `json:"code" comment:"配置标识"`
	Content string `json:"content" comment:"配置内容"`
	common.ControlBy
}

func (s *SwConfigUpdateReq) Generate(model *models.SwConfig) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Code = s.Code
	model.Content = s.Content
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *SwConfigUpdateReq) GetId() interface{} {
	return s.Id
}

// SwConfigGetReq 功能获取请求参数
type SwConfigGetReq struct {
	Id int `uri:"id"`
}

func (s *SwConfigGetReq) GetId() interface{} {
	return s.Id
}

// SwConfigDeleteReq 功能删除请求参数
type SwConfigDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *SwConfigDeleteReq) GetId() interface{} {
	return s.Ids
}
