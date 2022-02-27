package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type JyxUserRouter struct {
}

// InitJyxUserRouter 初始化 JyxUser 路由信息
func (s *JyxUserRouter) InitJyxUserRouter(Router *gin.RouterGroup) {
	//jyxUserRouter := Router.Group("jyxUser").Use(middleware.OperationRecord())
	jyxUserRouter := Router.Group("jyxUser")
	jyxUserRouterWithoutRecord := Router.Group("jyxUser")
	var jyxUserApi = v1.ApiGroupApp.AutoCodeApiGroup.JyxUserApi
	{
		jyxUserRouter.POST("createJyxUser", jyxUserApi.CreateJyxUser)   // 新建JyxUser
		jyxUserRouter.DELETE("deleteJyxUser", jyxUserApi.DeleteJyxUser) // 删除JyxUser
		jyxUserRouter.DELETE("deleteJyxUserByIds", jyxUserApi.DeleteJyxUserByIds) // 批量删除JyxUser
		jyxUserRouter.PUT("updateJyxUser", jyxUserApi.UpdateJyxUser)    // 更新JyxUser
	}
	{
		jyxUserRouterWithoutRecord.GET("findJyxUser", jyxUserApi.FindJyxUser)        // 根据ID获取JyxUser
		jyxUserRouterWithoutRecord.GET("getJyxUserList", jyxUserApi.GetJyxUserList)  // 获取JyxUser列表
	}
}
