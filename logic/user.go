package logic

import (
	"crypto/sha512"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/1348453525/user-redeem-code-gin/entity"
	"github.com/1348453525/user-redeem-code-gin/global"
	"github.com/1348453525/user-redeem-code-gin/model"
	"github.com/1348453525/user-redeem-code-gin/pkg/helper"
	"github.com/1348453525/user-redeem-code-gin/pkg/jwt"
	"github.com/anaskhan96/go-password-encoder"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserLogic struct{}

func NewUserLogic() *UserLogic {
	return &UserLogic{}
}

func (l *UserLogic) Register(c *gin.Context, r *entity.RegisterDto) (*entity.RegisterDvo, error) {
	// 验证密码是否一致
	if r.Password != r.ConfirmPassword {
		return nil, entity.ErrPasswordNotMatch
	}

	// 查询用户是否存在
	var user model.User
	result := global.DB.Where("username = ?", r.Username).Or("mobile = ?", r.Mobile).First(&user)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		zap.L().Error("查询用户失败：", zap.Error(result.Error))
		return nil, entity.ErrInternal
	}
	if result.RowsAffected != 0 {
		return nil, entity.ErrUserExisted
	}

	// 生成密码
	options := &password.Options{
		SaltLen:      16,
		Iterations:   100,
		KeyLen:       32,
		HashFunction: sha512.New,
	}
	salt, encodedPwd := password.Encode(r.Password, options)
	pwd := fmt.Sprintf("pbkdf2-sha512$%s$%s", salt, encodedPwd)

	// 处理生日
	var birthday *time.Time
	if r.Birthday != "" {
		if t, err := time.Parse("2006-01-02", r.Birthday); err == nil {
			birthday = &t
		}
	}

	// 创建用户
	user = model.User{
		Username: r.Username,
		Password: pwd,
		Nickname: r.Nickname,
		Mobile:   r.Mobile,
		Gender:   r.Gender,
		Birthday: birthday,
	}
	result = global.DB.Create(&user)
	if result.Error != nil {
		return nil, entity.ErrOperationFailed
	}

	// 返回数据
	resp := &entity.RegisterDvo{
		ID:       user.ID,
		Username: user.Username,
		Nickname: user.Nickname,
		Mobile:   user.Mobile,
		Gender:   user.Gender,
		Birthday: "",
	}
	resp.Birthday = helper.FormatDate(user.Birthday)
	return resp, nil
}

func (l *UserLogic) Login(c *gin.Context, r *entity.LoginDto) (*entity.LoginDvo, error) {
	// 查询用户
	var user model.User
	result := global.DB.Where("username = ?", r.Username).First(&user)
	if result.RowsAffected == 0 {
		return nil, entity.ErrParam
	}
	if result.Error != nil {
		zap.L().Error("查询用户失败：", zap.Error(result.Error))
	}

	// 检验状态
	if user.IsDel == 1 {
		return nil, entity.ErrUserDisabled
	}

	// 校验密码
	options := &password.Options{
		SaltLen:      16,
		Iterations:   100,
		KeyLen:       32,
		HashFunction: sha512.New,
	}
	encrypted := strings.Split(user.Password, "$")
	if !password.Verify(r.Password, encrypted[1], encrypted[2], options) {
		return nil, entity.ErrPasswordError
	}

	// 生成 token
	token, err := jwt.GenerateToken(user.ID)
	if err != nil {
		zap.L().Error("生成token失败：", zap.Error(err))
		return nil, entity.ErrInternal
	}

	// 返回数据
	birthday := helper.FormatDate(user.Birthday)
	resp := &entity.LoginDvo{
		Info: entity.UserInfoDvo{
			ID:       user.ID,
			Username: user.Username,
			Nickname: user.Nickname,
			Mobile:   user.Mobile,
			Gender:   user.Gender,
			Birthday: birthday,
		},
		Token: token,
	}
	return resp, nil
}

func (l *UserLogic) Info(c *gin.Context, id int64) (*entity.UserInfoDvo, error) {
	var user model.User
	if err := user.GetByID(id); err != nil {
		return nil, err
	}

	resp := &entity.UserInfoDvo{
		ID:       user.ID,
		Username: user.Username,
		Nickname: user.Nickname,
		Mobile:   user.Mobile,
		Gender:   user.Gender,
		Birthday: "",
	}
	resp.Birthday = helper.FormatDate(user.Birthday)
	return resp, nil
}

func (l *UserLogic) GetList(c *gin.Context, r *entity.GetUserListDto) (*entity.GetUserListDvo, error) {
	var user model.User
	list, count := user.GetList(r.Page, r.PageSize)
	resp := &entity.GetUserListDvo{
		Page:     r.Page,
		PageSize: r.PageSize,
		Total:    count,
	}
	for _, v := range list {
		userInfoDvo := entity.UserInfoDvo{
			ID:       v.ID,
			Username: v.Username,
			Nickname: v.Nickname,
			Mobile:   v.Mobile,
			Gender:   v.Gender,
			Birthday: "",
		}
		userInfoDvo.Birthday = helper.FormatDate(v.Birthday)
		resp.Data = append(resp.Data, userInfoDvo)
	}
	return resp, nil
}

func (l *UserLogic) Update(c *gin.Context, r *entity.UpdateUserDto) error {
	// 查询用户是否存在
	var user model.User
	result := global.DB.Where("id = ?", r.ID).First(&user)
	if result.RowsAffected == 0 {
		return entity.ErrParam
	}
	if result.Error != nil {
		zap.L().Error("查询用户失败：", zap.Error(result.Error))
		return entity.ErrInternal
	}

	// 处理生日
	var birthday *time.Time
	if r.Birthday != "" {
		if t, err := time.Parse("2006-01-02", r.Birthday); err == nil {
			birthday = &t
		}
	}
	user = model.User{
		Username: r.Username,
		Nickname: r.Nickname,
		Mobile:   r.Mobile,
		Gender:   r.Gender,
		Birthday: birthday,
	}
	result = global.DB.Model(&model.User{}).Where("id=?", r.ID).Updates(&user)
	if result.Error != nil {
		zap.L().Error("更新用户失败：", zap.Error(result.Error))
		return entity.ErrInternal
	}
	return nil
}

func (l *UserLogic) Delete(c *gin.Context, id int64) error {
	// 删除用户
	user := model.User{
		IsDel: 1,
	}
	result := global.DB.Model(&model.User{}).Where("id=?", id).Updates(&user)
	if result.Error != nil {
		zap.L().Error("删除用户失败：", zap.Error(result.Error))
		return entity.ErrInternal
	}
	return nil
}
