package test

import (
	"fmt"
	"time"

	"github.com/1348453525/user-redeem-code-gin/pkg/result"
	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {
	result.Success(c)
}

func TestData(c *gin.Context) {
	result.Success(
		c,
		gin.H{
			"test": "test/test",
		},
		429,
		"too many request",
	)
}

func TestError(c *gin.Context) {
	result.Error(c)
}

func TestErrorData(c *gin.Context) {
	result.Error(
		c,
		500,
		"test/test",
		gin.H{
			"data": "test/test",
		},
		gin.H{
			"error": "test/test",
		},
	)
}

func Shutdown(c *gin.Context) {
	time.Sleep(10 * time.Second)
	result.Success(c, gin.H{}, 200, "ok")
}

func Db(c *gin.Context) {
	fmt.Println("test")
}

func Redis(c *gin.Context) {
	fmt.Println("test")
}
