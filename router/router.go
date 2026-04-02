package router

import (
	"github.com/1348453525/user-redeem-code-gin/api/test"
	"github.com/1348453525/user-redeem-code-gin/api/user"
	"github.com/gin-gonic/gin"
)

func RouterGroup(r *gin.RouterGroup) {
	// test
	testGroup := r.Group("/Test")
	{
		testGroup.GET("/Test", test.Test)
		testGroup.GET("/TestData", test.TestData)
		testGroup.GET("/TestError", test.TestError)
		testGroup.GET("/TestErrorData", test.TestErrorData)
		testGroup.GET("/Db", test.Db)
		testGroup.GET("/Redis", test.Redis)
		testGroup.GET("/Shutdown", test.Shutdown)
	}

	// 用户
	userGroup := r.Group("/User")
	{
		userGroup.POST("/Register", user.Register) // 注册
		userGroup.POST("/Login", user.Login)       // 登录
		userGroup.GET("/Logout", user.Logout)      // 退出
		userGroup.GET("/Info", user.Info)          // 获取用户信息
		userGroup.GET("/GetList", user.GetList)    // 获取用户列表
		userGroup.PUT("/Update", user.Update)      // 更新用户信息
		userGroup.DELETE("/Delete", user.Delete)   // 删除用户
	}

	// 兑换码
}
