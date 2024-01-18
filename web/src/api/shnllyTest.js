import service from '@/utils/request'

// @Tags ShnllyTest
// @Summary 创建shnlly的test模型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ShnllyTest true "创建shnlly的test模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /shnllyTest/createShnllyTest [post]
export const createShnllyTest = (data) => {
  return service({
    url: '/shnllyTest/createShnllyTest',
    method: 'post',
    data
  })
}

// @Tags ShnllyTest
// @Summary 删除shnlly的test模型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ShnllyTest true "删除shnlly的test模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /shnllyTest/deleteShnllyTest [delete]
export const deleteShnllyTest = (data) => {
  return service({
    url: '/shnllyTest/deleteShnllyTest',
    method: 'delete',
    data
  })
}

// @Tags ShnllyTest
// @Summary 批量删除shnlly的test模型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除shnlly的test模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /shnllyTest/deleteShnllyTest [delete]
export const deleteShnllyTestByIds = (data) => {
  return service({
    url: '/shnllyTest/deleteShnllyTestByIds',
    method: 'delete',
    data
  })
}

// @Tags ShnllyTest
// @Summary 更新shnlly的test模型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ShnllyTest true "更新shnlly的test模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /shnllyTest/updateShnllyTest [put]
export const updateShnllyTest = (data) => {
  return service({
    url: '/shnllyTest/updateShnllyTest',
    method: 'put',
    data
  })
}

// @Tags ShnllyTest
// @Summary 用id查询shnlly的test模型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ShnllyTest true "用id查询shnlly的test模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /shnllyTest/findShnllyTest [get]
export const findShnllyTest = (params) => {
  return service({
    url: '/shnllyTest/findShnllyTest',
    method: 'get',
    params
  })
}

// @Tags ShnllyTest
// @Summary 分页获取shnlly的test模型列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取shnlly的test模型列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /shnllyTest/getShnllyTestList [get]
export const getShnllyTestList = (params) => {
  return service({
    url: '/shnllyTest/getShnllyTestList',
    method: 'get',
    params
  })
}
