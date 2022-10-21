package apis

import (
	"fmt"
	"go-admin/app/myapp/lib/sw"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"go-admin/app/myapp/models"
	"go-admin/app/myapp/service"
	"go-admin/app/myapp/service/dto"
	"go-admin/common/actions"
)

type SwConfig struct {
	api.Api
}

// GetPage 获取sw配置列表
// @Summary 获取sw配置列表
// @Description 获取sw配置列表
// @Tags sw配置
// @Param code query string false "配置标识"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.SwConfig}} "{"code": 200, "data": [...]}"
// @Router /api/v1/sw-config [get]
// @Security Bearer
func (e SwConfig) GetPage(c *gin.Context) {
	req := dto.SwConfigGetPageReq{}
	s := service.SwConfig{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	p := actions.GetPermissionFromContext(c)
	list := make([]models.SwConfig, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取sw配置失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取sw配置
// @Summary 获取sw配置
// @Description 获取sw配置
// @Tags sw配置
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.SwConfig} "{"code": 200, "data": [...]}"
// @Router /api/v1/sw-config/{id} [get]
// @Security Bearer
func (e SwConfig) Get(c *gin.Context) {
	req := dto.SwConfigGetReq{}
	s := service.SwConfig{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	var object models.SwConfig

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取sw配置失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建sw配置
// @Summary 创建sw配置
// @Description 创建sw配置
// @Tags sw配置
// @Accept application/json
// @Product application/json
// @Param data body dto.SwConfigInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/sw-config [post]
// @Security Bearer
func (e SwConfig) Insert(c *gin.Context) {
	req := dto.SwConfigInsertReq{}
	s := service.SwConfig{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	// 设置创建人
	req.SetCreateBy(user.GetUserId(c))

	err = s.Insert(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("创建sw配置失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改sw配置
// @Summary 修改sw配置
// @Description 修改sw配置
// @Tags sw配置
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.SwConfigUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/sw-config/{id} [put]
// @Security Bearer
func (e SwConfig) Update(c *gin.Context) {
	req := dto.SwConfigUpdateReq{}
	s := service.SwConfig{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = sw.Cli.DelConfigCache(req.Code)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("删除 sw 缓存是吧，\r\n失败信息 %s", err.Error()))
		return
	}
	err = s.Update(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("修改sw配置失败，\r\n失败信息 %s", err.Error()))
		return
	}
	time.Sleep(2 * time.Second) //延迟双删
	err = sw.Cli.DelConfigCache(req.Code)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("删除 sw 缓存是吧，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除sw配置
// @Summary 删除sw配置
// @Description 删除sw配置
// @Tags sw配置
// @Param data body dto.SwConfigDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/sw-config [delete]
// @Security Bearer
func (e SwConfig) Delete(c *gin.Context) {
	s := service.SwConfig{}
	req := dto.SwConfigDeleteReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	// req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Remove(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("删除sw配置失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
