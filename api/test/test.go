package test

import (
	"errors"
	"strconv"
	"time"

	"github.com/1348453525/user-redeem-code-gin/logic"
	"github.com/1348453525/user-redeem-code-gin/pkg/result"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
	id, _ := strconv.Atoi(c.Query("id"))
	if id <= 0 {
		result.Error(c, 400, "参数错误")
		return
	}

	res, err := logic.Test.Db(uint64(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			result.Success(c)
		} else {
			result.Error(c, 500, "查询失败："+err.Error())
		}
		return
	}
	result.Success(c, res)
}

func Redis(c *gin.Context) {
	value := logic.Test.Redis()
	result.Success(c, gin.H{
		"value": value,
	})
}
