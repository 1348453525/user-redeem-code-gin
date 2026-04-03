package helper

import (
	"time"

	"github.com/1348453525/user-redeem-code-gin/entity"
	"github.com/1348453525/user-redeem-code-gin/model"
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

	var id int64
	if id = c.GetInt64("userID"); id == 0 {
		return 0, entity.ErrUserNotLogin
	}
	return id, nil
}

func GetUserFromCtx(c *gin.Context) *model.User {
	if userObj, exists := c.Get("user"); exists {
		if user, ok := userObj.(*model.User); ok {
			return user
		}
	}
	return nil
}

func FormatDate(date *time.Time) string {
	var str string
	if date != nil {
		str = date.Format("2026-01-02")
	}
	return str
}

func FormatDatetime(datetime *time.Time) string {
	var str string
	if datetime != nil {
		str = datetime.Format("2026-01-02 15:04:05")
	}
	return str
}

func ParseDate(date string) (*time.Time, error) {

	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func ParseDatetime(datetime string) (*time.Time, error) {

	t, err := time.Parse("2006-01-02 15:04:05", datetime)
	if err != nil {
		return nil, err
	}
	return &t, nil
}
