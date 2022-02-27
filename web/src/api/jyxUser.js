import service from '@/utils/request'

// @Tags JyxUser
// @Summary 创建JyxUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.JyxUser true "创建JyxUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /jyxUser/createJyxUser [post]
export const createJyxUser = (data) => {
  return service({
    url: '/jyxUser/createJyxUser',
    method: 'post',
    data
  })
}

// @Tags JyxUser
// @Summary 删除JyxUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.JyxUser true "删除JyxUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /jyxUser/deleteJyxUser [delete]
export const deleteJyxUser = (data) => {
  return service({
    url: '/jyxUser/deleteJyxUser',
    method: 'delete',
    data
  })
}

// @Tags JyxUser
// @Summary 删除JyxUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除JyxUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /jyxUser/deleteJyxUser [delete]
export const deleteJyxUserByIds = (data) => {
  return service({
    url: '/jyxUser/deleteJyxUserByIds',
    method: 'delete',
    data
  })
}

// @Tags JyxUser
// @Summary 更新JyxUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.JyxUser true "更新JyxUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /jyxUser/updateJyxUser [put]
export const updateJyxUser = (data) => {
  return service({
    url: '/jyxUser/updateJyxUser',
    method: 'put',
    data
  })
}

// @Tags JyxUser
// @Summary 用id查询JyxUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.JyxUser true "用id查询JyxUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /jyxUser/findJyxUser [get]
export const findJyxUser = (params) => {
  return service({
    url: '/jyxUser/findJyxUser',
    method: 'get',
    params
  })
}

// @Tags JyxUser
// @Summary 分页获取JyxUser列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取JyxUser列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /jyxUser/getJyxUserList [get]
export const getJyxUserList = (params) => {
  return service({
    url: '/jyxUser/getJyxUserList',
    method: 'get',
    params
  })
}
