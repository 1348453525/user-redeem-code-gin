package handler

import (
	"errors"
	"strconv"
	"time"

	"github.com/1348453525/user-redeem-code-gin/logic"
	"github.com/1348453525/user-redeem-code-gin/pkg/result"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Test struct{}

func NewTest() *Test {
	return &Test{}
}

func (h *Test) Test(c *gin.Context) {
	result.Success(c)
}

func (h *Test) TestData(c *gin.Context) {
	result.Success(
		c,
		gin.H{
			"test": "test/test",
		},
		429,
		"too many request",
	)
}

func (h *Test) TestError(c *gin.Context) {
	result.Error(c)
}

func (h *Test) TestErrorData(c *gin.Context) {
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

func (h *Test) Shutdown(c *gin.Context) {
	time.Sleep(10 * time.Second)
	result.Success(c, gin.H{}, 200, "ok")
}

func (h *Test) Db(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	if id <= 0 {
		result.Error(c, 400, "参数错误")
		return
	}

	res, err := logic.NewTestLogic().Db(uint64(id))
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

func (h *Test) Redis(c *gin.Context) {
	value := logic.NewTestLogic().Redis()
	result.Success(c, gin.H{
		"value": value,
	})
}
