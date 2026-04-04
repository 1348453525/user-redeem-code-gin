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
	idStr := c.Query("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || id <= 0 {
		result.Error(c, 400, entity.ErrParam.Error())
		return
	}

	// 处理逻辑
	resp, err := logic.NewRedeemCodeLogic().Detail(c, id)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			result.Error(c, 500, entity.ErrInternal.Error())
			return
		}
		result.Success(c)
		return
	}
	result.Success(c, resp)
}

func GetList(c *gin.Context) {
	// 接收参数
	var dto entity.GetRedeemCodeListDto
	if err := c.ShouldBindQuery(&dto); err != nil {
		result.Error(c, 400, entity.ErrParam.Error())
		return
	}

	// 验证参数
	if ok := util.Validate(c, dto); !ok {
		return
	}

	// 处理逻辑
	resp, err := logic.NewRedeemCodeLogic().GetList(c, &dto)
	if err != nil {
		result.Error(c, 500, err.Error())
		return
	}
	result.Success(c, resp)
}

func Update(c *gin.Context) {
	// 接收参数
	var dto entity.UpdateRedeemCodeDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		result.Error(c, 400, entity.ErrParam.Error())
		return
	}

	// 验证参数
	if ok := util.Validate(c, dto); !ok {
		return
	}

	// 处理逻辑
	err := logic.NewRedeemCodeLogic().Update(c, &dto)
	if err != nil {
		result.Error(c, 500, err.Error())
		return
	}
	result.Success(c)
}

func Delete(c *gin.Context) {
	// 接收参数
	idStr := c.Query("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || id <= 0 {
		result.Error(c, 400, entity.ErrParam.Error())
		return
	}

	// 处理逻辑
	err = logic.NewRedeemCodeLogic().Delete(c, id)
	if err != nil {
		result.Error(c, 500, err.Error())
		return
	}
	result.Success(c)
}

func Use(c *gin.Context) {
	// 接收参数
	var dto entity.UseRedeemCodeDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		result.Error(c, 400, entity.ErrParam.Error())
		return
	}

	// 验证参数
	if ok := util.Validate(c, dto); !ok {
		return
	}

	// 处理逻辑
	err := logic.NewRedeemCodeLogic().Use(c, &dto)
	if err != nil {
		result.Error(c, 500, err.Error())
		return
	}
	result.Success(c)
}
