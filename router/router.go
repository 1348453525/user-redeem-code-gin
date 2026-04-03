package router

import (
	"github.com/1348453525/user-redeem-code-gin/api/redeem_code"
	"github.com/1348453525/user-redeem-code-gin/api/redeem_code_batch"
	"github.com/1348453525/user-redeem-code-gin/api/test"
	"github.com/1348453525/user-redeem-code-gin/api/user"
	"github.com/1348453525/user-redeem-code-gin/middleware"
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

	// auth 注册、登录、退出
	r.POST("/Register", user.Register)                            // 注册
	r.POST("/Login", user.Login)                                  // 登录
	r.GET("/Logout", middleware.JWTAuthMiddleware(), user.Logout) // 退出

	// 用户
	userGroup := r.Group("/User")
	userGroup.Use(middleware.JWTAuthMiddleware())
	{
		userGroup.GET("/Logout", user.Logout)    // 退出
		userGroup.GET("/Info", user.Info)        // 获取用户信息
		userGroup.GET("/GetList", user.GetList)  // 获取用户列表
		userGroup.PUT("/Update", user.Update)    // 更新用户信息
		userGroup.DELETE("/Delete", user.Delete) // 删除用户
	}

	// 兑换码批次
	redeemCodeBatchGroup := r.Group("/RedeemCodeBatch")
	redeemCodeBatchGroup.Use(middleware.JWTAuthMiddleware())
	{
		redeemCodeBatchGroup.POST("/Create", redeem_code_batch.Create)   // 创建兑换码批次
		redeemCodeBatchGroup.GET("/Detail", redeem_code_batch.Detail)    // 获取兑换码批次详情
		redeemCodeBatchGroup.GET("/GetList", redeem_code_batch.GetList)  // 获取兑换码批次列表
		redeemCodeBatchGroup.PUT("/Update", redeem_code_batch.Update)    // 更新兑换码批次
		redeemCodeBatchGroup.DELETE("/Delete", redeem_code_batch.Delete) // 删除兑换码批次
	}

	// 兑换码
	redeemCodeGroup := r.Group("/RedeemCode")
	redeemCodeGroup.Use(middleware.JWTAuthMiddleware())
	{
		redeemCodeGroup.GET("/Detail", redeem_code.Detail)    // 获取兑换码详情
		redeemCodeGroup.GET("/GetList", redeem_code.GetList)  // 获取兑换码列表
		redeemCodeGroup.PUT("/Update", redeem_code.Update)    // 更新兑换码
		redeemCodeGroup.DELETE("/Delete", redeem_code.Delete) // 删除兑换码
		redeemCodeGroup.POST("/Use", redeem_code.Use)         // 使用兑换码
	}
}
