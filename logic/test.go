package logic

import (
	"context"
	"errors"
	"time"

	"github.com/1348453525/user-redeem-code-gin/global"
	"github.com/1348453525/user-redeem-code-gin/model"
	"github.com/redis/go-redis/v9"
)

type test struct {
}

var Test = &test{}

func (t *test) Db(id uint64) (*model.Test, error) {
	test := &model.Test{}
	if err := test.GetByID(id); err != nil {
		return nil, err
	}
	return test, nil
}

func (t *test) Redis() string {
	key := "test"
	value, err := global.Redis.Get(context.Background(), key).Result()
	if err != nil {
		// redis.Nil 表示 Key 不存在
		if errors.Is(err, redis.Nil) {
			// 设置key
			newValue := time.Now().Format("2006-01-02 15:04:05")
			if err := global.Redis.Set(context.Background(), key, newValue, 5*time.Second).Err(); err != nil {
				return err.Error()
			}
			return newValue
		}
		return err.Error()
	}
	return value
}
