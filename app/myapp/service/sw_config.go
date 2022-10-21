package service

import (
	"errors"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"go-admin/app/myapp/models"
	"go-admin/app/myapp/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
)

type SwConfig struct {
	service.Service
}

// GetPage 获取SwConfig列表
func (e *SwConfig) GetPage(c *dto.SwConfigGetPageReq, p *actions.DataPermission, list *[]models.SwConfig, count *int64) error {
	var err error
	var data models.SwConfig

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("SwConfigService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取SwConfig对象
func (e *SwConfig) Get(d *dto.SwConfigGetReq, p *actions.DataPermission, model *models.SwConfig) error {
	var data models.SwConfig

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetSwConfig error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建SwConfig对象
func (e *SwConfig) Insert(c *dto.SwConfigInsertReq) error {
	var err error
	var data models.SwConfig
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("SwConfigService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改SwConfig对象
func (e *SwConfig) Update(c *dto.SwConfigUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.SwConfig{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("SwConfigService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除SwConfig
func (e *SwConfig) Remove(d *dto.SwConfigDeleteReq, p *actions.DataPermission) error {
	var data models.SwConfig

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveSwConfig error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
