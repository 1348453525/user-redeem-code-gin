package main

import (
	"fmt"
	"strconv"

	"github.com/1348453525/user-redeem-code-gin/entity"
	"github.com/1348453525/user-redeem-code-gin/initialize"
	"github.com/1348453525/user-redeem-code-gin/logic"
	"github.com/gin-gonic/gin"
)

func TestRegister() {
	ctx := &gin.Context{}
	mobile := 1301301300
	for i := 0; i < 100; i++ {
		mobile++
		r := &entity.RegisterDto{
			Username:        fmt.Sprintf("test%d", i),
			Nickname:        fmt.Sprintf("test%d", i),
			Mobile:          strconv.Itoa(mobile),
			Gender:          1,
			Birthday:        "1990-01-01",
			Password:        "123456",
			ConfirmPassword: "123456",
		}
		_, err := logic.User.Register(ctx, r)
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	cfgFile := "../../../config.yaml"
	// 初始化配置
	initialize.InitConfig(cfgFile)
	// 初始化日志
	initialize.InitLogger()
	// 初始化数据库
	initialize.InitDB()

	TestRegister()
}
