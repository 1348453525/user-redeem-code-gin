package main

import (
	"github.com/1348453525/user-redeem-code-gin/global"
	"github.com/1348453525/user-redeem-code-gin/initialize"
	"gorm.io/gen"
)

// GORM GEN 生成代码
func main() {
	cfgFile := "../../config.yaml"
	// 初始化配置
	initialize.InitConfig(cfgFile)
	// 初始化日志
	initialize.InitLogger()
	// 初始化数据库
	initialize.InitDB()

	g := gen.NewGenerator(gen.Config{
		OutPath:       "../../model/gen",
		Mode:          gen.WithoutContext,
		FieldNullable: true, // delete_at 是可以为空的
	})
	// *gorm.DB
	g.UseDB(global.DB)
	// sql to struct
	// g.ApplyBasic(g.GenerateAllTable()...)
	g.ApplyBasic(g.GenerateModel("user"))
	// 执行并生成代码
	g.Execute()
}
