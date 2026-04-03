package user

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

func Register(c *gin.Context) {
	// 接收参数
	var dto entity.RegisterDto
	_ = c.ShouldBindJSON(&dto)

	// 验证参数
	if ok := util.Validate(c, dto); !ok {
		return
	}

	// 处理逻辑
	resp, err := logic.User.Register(c, &dto)
	if err != nil {
		result.Error(c, 500, err.Error())
		return
	}
	result.Success(c, resp)
}

func Login(c *gin.Context) {
	// 接收参数
	var dto entity.LoginDto
	_ = c.ShouldBindJSON(&dto)

	// 验证参数
	if ok := util.Validate(c, dto); !ok {
		return
	}

	// 处理逻辑
	resp, err := logic.User.Login(c, &dto)
	if err != nil {
		result.Error(c, 500, err.Error())
		return
	}
	result.Success(c, resp)
}

func Logout(c *gin.Context) {
	result.Success(c)
}

func Info(c *gin.Context) {
	// 接收参数
	id, _ := strconv.Atoi(c.Query("id"))
	if id <= 0 {
		result.Error(c, 400, entity.ErrParam.Error())
		return
	}

	// 处理逻辑
	resp, err := logic.User.Info(c, int64(id))
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
	var dto entity.GetUserListDto
	_ = c.ShouldBindQuery(&dto)

	// 验证参数
	if ok := util.Validate(c, dto); !ok {
		return
	}

	// 处理逻辑
	resp, err := logic.User.GetList(c, &dto)
	if err != nil {
		result.Error(c, 500, err.Error())
		return
	}
	result.Success(c, resp)
}

func Update(c *gin.Context) {
	// 接收参数
	var dto entity.UpdateUserDto
	_ = c.ShouldBindJSON(&dto)

	// 验证参数
	if ok := util.Validate(c, dto); !ok {
		return
	}

	// 处理逻辑
	err := logic.User.Update(c, &dto)
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
	err := logic.User.Delete(c, int64(id))
	if err != nil {
		result.Error(c, 500, err.Error())
		return
	}
	result.Success(c)
}
