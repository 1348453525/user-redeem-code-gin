package helper

import (
	"time"

	"github.com/1348453525/user-redeem-code-gin/entity"
	"github.com/gin-gonic/gin"
)

func GetUserID(c *gin.Context) (int64, error) {
	// userID, exists := c.Get("userID")
	// if !exists {
	// 	return 0, entity.ErrUserNotLogin
	// }

	// id, ok := userID.(int64)
	// if !ok {
	// 	return 0, entity.ErrUserNotLogin
	// }
	// return id, nil

	id := c.GetInt64("userID")
	if id == 0 {
		return 0, entity.ErrUserNotLogin
	}
	return id, nil
}

func FormatBirthday(birthday *time.Time) string {
	var result string
	if birthday != nil {
		result = birthday.Format("2026-01-02")
	}
	return result
}
