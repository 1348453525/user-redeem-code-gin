package router

import (
	"github.com/1348453525/user-redeem-code-gin/api/test"
	"github.com/gin-gonic/gin"
)

func RouterGroup(r *gin.RouterGroup) {
	// test
	testGroup := r.Group("/test")
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

	// 兑换码
}
