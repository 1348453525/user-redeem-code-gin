package redeem_code

import (
	"errors"
	"strconv"

	"github.com/1348453525/user-redeem-code-gin/entity"
	"github.com/1348453525/user-redeem-code-gin/logic"
	"github.com/1348453525/user-redeem-code-gin/pkg/result"
	"github.com/1348453525/user-redeem-code-gin/pkg/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Detail(c *gin.Context) {
	// 接收参数
	id, _ := strconv.Atoi(c.Query("id"))
	if id <= 0 {
		result.Error(c, 400, entity.ErrParam.Error())
		return
	}

	// 处理逻辑
	resp, err := logic.RedeemCodeLogic.Detail(c, int64(id))
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			result.Error(c, 500, entity.ErrInternal.Error())
		}
		result.Success(c)
		return
	}
	result.Success(c, resp)
}

func GetList(c *gin.Context) {
	// 接收参数
	var dto entity.GetRedeemCodeListDto
	_ = c.ShouldBindQuery(&dto)

	// 验证参数
	if ok := util.Validate(c, dto); !ok {
		return
	}

	// 处理逻辑
	resp, err := logic.RedeemCodeLogic.GetList(c, &dto)
	if err != nil {
		result.Error(c, 500, err.Error())
		return
	}
	result.Success(c, resp)
}

func Update(c *gin.Context) {
	// 接收参数
	var dto entity.UpdateRedeemCodeDto
	_ = c.ShouldBindJSON(&dto)

	// 验证参数
	if ok := util.Validate(c, dto); !ok {
		return
	}

	// 处理逻辑
	err := logic.RedeemCodeLogic.Update(c, &dto)
	if err != nil {
		result.Error(c, 500, err.Error())
		return
	}
	result.Success(c)
}

func Delete(c *gin.Context) {
	// 接收参数
	id, _ := strconv.Atoi(c.Query("id"))
	if id <= 0 {
		result.Error(c, 400, entity.ErrParam.Error())
		return
	}

	// 处理逻辑
	err := logic.RedeemCodeLogic.Delete(c, int64(id))
	if err != nil {
		result.Error(c, 500, err.Error())
		return
	}
	result.Success(c)
}

func Use(c *gin.Context) {
	// 接收参数
	var dto entity.UseRedeemCodeDto
	_ = c.ShouldBindJSON(&dto)

	// 验证参数
	if ok := util.Validate(c, dto); !ok {
		return
	}

	// 处理逻辑
	err := logic.RedeemCodeLogic.Use(c, &dto)
	if err != nil {
		result.Error(c, 500, err.Error())
		return
	}
	result.Success(c)
}
