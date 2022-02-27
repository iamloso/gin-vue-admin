package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autocodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type JyxUserApi struct {
}

var jyxUserService = service.ServiceGroupApp.AutoCodeServiceGroup.JyxUserService


// CreateJyxUser 创建JyxUser
// @Tags JyxUser
// @Summary 创建JyxUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.JyxUser true "创建JyxUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /jyxUser/createJyxUser [post]
func (jyxUserApi *JyxUserApi) CreateJyxUser(c *gin.Context) {
	var jyxUser autocode.JyxUser
	_ = c.ShouldBindJSON(&jyxUser)
	if err := jyxUserService.CreateJyxUser(jyxUser); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteJyxUser 删除JyxUser
// @Tags JyxUser
// @Summary 删除JyxUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.JyxUser true "删除JyxUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /jyxUser/deleteJyxUser [delete]
func (jyxUserApi *JyxUserApi) DeleteJyxUser(c *gin.Context) {
	var jyxUser autocode.JyxUser
	_ = c.ShouldBindJSON(&jyxUser)
	if err := jyxUserService.DeleteJyxUser(jyxUser); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteJyxUserByIds 批量删除JyxUser
// @Tags JyxUser
// @Summary 批量删除JyxUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除JyxUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /jyxUser/deleteJyxUserByIds [delete]
func (jyxUserApi *JyxUserApi) DeleteJyxUserByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := jyxUserService.DeleteJyxUserByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateJyxUser 更新JyxUser
// @Tags JyxUser
// @Summary 更新JyxUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.JyxUser true "更新JyxUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /jyxUser/updateJyxUser [put]
func (jyxUserApi *JyxUserApi) UpdateJyxUser(c *gin.Context) {
	var jyxUser autocode.JyxUser
	_ = c.ShouldBindJSON(&jyxUser)
	if err := jyxUserService.UpdateJyxUser(jyxUser); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindJyxUser 用id查询JyxUser
// @Tags JyxUser
// @Summary 用id查询JyxUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.JyxUser true "用id查询JyxUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /jyxUser/findJyxUser [get]
func (jyxUserApi *JyxUserApi) FindJyxUser(c *gin.Context) {
	var jyxUser autocode.JyxUser
	_ = c.ShouldBindQuery(&jyxUser)
	if err, rejyxUser := jyxUserService.GetJyxUser(jyxUser.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rejyxUser": rejyxUser}, c)
	}
}

// GetJyxUserList 分页获取JyxUser列表
// @Tags JyxUser
// @Summary 分页获取JyxUser列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.JyxUserSearch true "分页获取JyxUser列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /jyxUser/getJyxUserList [get]
func (jyxUserApi *JyxUserApi) GetJyxUserList(c *gin.Context) {
	var pageInfo autocodeReq.JyxUserSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := jyxUserService.GetJyxUserInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
