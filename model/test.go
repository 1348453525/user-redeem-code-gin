package model

import (
	"github.com/1348453525/user-redeem-code-gin/global"
)

// Test 模型
type Test struct {
	ID    uint64 `json:"id" gorm:"column:id"`
	Key   string `json:"key" gorm:"column:key"`
	Value string `json:"value" gorm:"column:value"`
}

func (t *Test) TableName() string {
	return "test"
}

// GetByID 根据ID获取Test信息
func (t *Test) GetByID(id uint64) error {
	result := global.DB.Where("id = ?", id).First(t)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
