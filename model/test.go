package model

import (
	"github.com/1348453525/user-redeem-code-gin/global"
)

const TableNameTest = "test"

// Test mapped from table <test>
type Test struct {
	ID    uint64 `json:"id"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

// TableName Test's table name
func (*Test) TableName() string {
	return TableNameTest
}

// GetByID 根据ID获取Test信息
func (t *Test) GetByID(id uint64) error {
	result := global.DB.Where("id = ?", id).First(t)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
