package shnlly

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ShnllyTestRouter struct {
}

// InitShnllyTestRouter 初始化 shnlly的test模型 路由信息
func (s *ShnllyTestRouter) InitShnllyTestRouter(Router *gin.RouterGroup) {
	shnllyTestRouter := Router.Group("shnllyTest").Use(middleware.OperationRecord())
	shnllyTestRouterWithoutRecord := Router.Group("shnllyTest")
	var shnllyTestApi = v1.ApiGroupApp.ShnllyApiGroup.ShnllyTestApi
	{
		shnllyTestRouter.POST("createShnllyTest", shnllyTestApi.CreateShnllyTest)             // 新建shnlly的test模型
		shnllyTestRouter.DELETE("deleteShnllyTest", shnllyTestApi.DeleteShnllyTest)           // 删除shnlly的test模型
		shnllyTestRouter.DELETE("deleteShnllyTestByIds", shnllyTestApi.DeleteShnllyTestByIds) // 批量删除shnlly的test模型
		shnllyTestRouter.PUT("updateShnllyTest", shnllyTestApi.UpdateShnllyTest)              // 更新shnlly的test模型
	}
	{
		shnllyTestRouterWithoutRecord.GET("findShnllyTest", shnllyTestApi.FindShnllyTest)       // 根据ID获取shnlly的test模型
		shnllyTestRouterWithoutRecord.GET("getShnllyTestList", shnllyTestApi.GetShnllyTestList) // 获取shnlly的test模型列表
	}
}
