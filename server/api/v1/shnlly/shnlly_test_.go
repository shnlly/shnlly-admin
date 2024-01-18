package shnlly

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/shnlly"
	shnllyReq "github.com/flipped-aurora/gin-vue-admin/server/model/shnlly/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ShnllyTestApi struct {
}

var shnllyTestService = service.ServiceGroupApp.ShnllyServiceGroup.ShnllyTestService

// CreateShnllyTest 创建shnlly的test模型
// @Tags ShnllyTest
// @Summary 创建shnlly的test模型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shnlly.ShnllyTest true "创建shnlly的test模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /shnllyTest/createShnllyTest [post]
func (shnllyTestApi *ShnllyTestApi) CreateShnllyTest(c *gin.Context) {
	var shnllyTest shnlly.ShnllyTest
	err := c.ShouldBindJSON(&shnllyTest)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := shnllyTestService.CreateShnllyTest(&shnllyTest); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteShnllyTest 删除shnlly的test模型
// @Tags ShnllyTest
// @Summary 删除shnlly的test模型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shnlly.ShnllyTest true "删除shnlly的test模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /shnllyTest/deleteShnllyTest [delete]
func (shnllyTestApi *ShnllyTestApi) DeleteShnllyTest(c *gin.Context) {
	var shnllyTest shnlly.ShnllyTest
	err := c.ShouldBindJSON(&shnllyTest)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := shnllyTestService.DeleteShnllyTest(shnllyTest); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteShnllyTestByIds 批量删除shnlly的test模型
// @Tags ShnllyTest
// @Summary 批量删除shnlly的test模型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除shnlly的test模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /shnllyTest/deleteShnllyTestByIds [delete]
func (shnllyTestApi *ShnllyTestApi) DeleteShnllyTestByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := shnllyTestService.DeleteShnllyTestByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateShnllyTest 更新shnlly的test模型
// @Tags ShnllyTest
// @Summary 更新shnlly的test模型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body shnlly.ShnllyTest true "更新shnlly的test模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /shnllyTest/updateShnllyTest [put]
func (shnllyTestApi *ShnllyTestApi) UpdateShnllyTest(c *gin.Context) {
	var shnllyTest shnlly.ShnllyTest
	err := c.ShouldBindJSON(&shnllyTest)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := shnllyTestService.UpdateShnllyTest(shnllyTest); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindShnllyTest 用id查询shnlly的test模型
// @Tags ShnllyTest
// @Summary 用id查询shnlly的test模型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query shnlly.ShnllyTest true "用id查询shnlly的test模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /shnllyTest/findShnllyTest [get]
func (shnllyTestApi *ShnllyTestApi) FindShnllyTest(c *gin.Context) {
	var shnllyTest shnlly.ShnllyTest
	err := c.ShouldBindQuery(&shnllyTest)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reshnllyTest, err := shnllyTestService.GetShnllyTest(shnllyTest.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reshnllyTest": reshnllyTest}, c)
	}
}

// GetShnllyTestList 分页获取shnlly的test模型列表
// @Tags ShnllyTest
// @Summary 分页获取shnlly的test模型列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query shnllyReq.ShnllyTestSearch true "分页获取shnlly的test模型列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /shnllyTest/getShnllyTestList [get]
func (shnllyTestApi *ShnllyTestApi) GetShnllyTestList(c *gin.Context) {
	var pageInfo shnllyReq.ShnllyTestSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := shnllyTestService.GetShnllyTestInfoList(pageInfo); err != nil {
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
