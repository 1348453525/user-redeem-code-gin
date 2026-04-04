package logic

import (
	"time"

	"github.com/1348453525/user-redeem-code-gin/entity"
	"github.com/1348453525/user-redeem-code-gin/global"
	"github.com/1348453525/user-redeem-code-gin/model"
	"github.com/1348453525/user-redeem-code-gin/pkg/helper"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type redeemCodeLogic struct{}

var RedeemCodeLogic redeemCodeLogic

func (l *redeemCodeLogic) Detail(c *gin.Context, id int64) (*model.RedeemCode, error) {
	var redeemCode model.RedeemCode
	if err := redeemCode.GetByID(id); err != nil {
		return nil, err
	}
	return &redeemCode, nil
}

func (l *redeemCodeLogic) GetList(c *gin.Context, r *entity.GetRedeemCodeListDto) (*entity.GetRedeemCodeListDvo, error) {
	var redeemCodeModel model.RedeemCode
	list, count := redeemCodeModel.GetList(r.Page, r.PageSize)
	return &entity.GetRedeemCodeListDvo{
		Page:     r.Page,
		PageSize: r.PageSize,
		Total:    count,
		Data:     list,
	}, nil
}

func (l *redeemCodeLogic) Update(c *gin.Context, r *entity.UpdateRedeemCodeDto) error {
	redeemCodeModel := model.RedeemCode{
		ID: r.ID,
	}
	if r.Title != "" {
		redeemCodeModel.Title = r.Title
	}
	if r.ExpirationAt != "" {
		expirationAt, err := helper.ParseDatetime(r.ExpirationAt)
		if err != nil {
			return entity.ErrParam
		}
		redeemCodeModel.ExpirationAt = *expirationAt
	}
	if r.IsDel != 0 {
		redeemCodeModel.IsDel = r.IsDel
	}
	if result := global.DB.Model(&model.RedeemCode{}).Where("id=?", r.ID).Updates(&redeemCodeModel); result.Error != nil {
		zap.L().Error("更新兑换码失败：", zap.Error(result.Error))
		return entity.ErrInternal
	}
	return nil
}

func (l *redeemCodeLogic) Delete(c *gin.Context, id int64) error {
	if result := global.DB.Model(&model.RedeemCode{}).Where("id=?", id).Update("is_del", 1); result.Error != nil {
		zap.L().Error("更新兑换码失败：", zap.Error(result.Error))
		return entity.ErrInternal
	}
	return nil
}

func (l *redeemCodeLogic) Use(c *gin.Context, r *entity.UseRedeemCodeDto) error {
	// 开启事务
	tx := global.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			zap.L().Error("事务panic", zap.Any("panic", r))
		}
	}()

	// 获取兑换码信息
	var redeemCode model.RedeemCode
	if err := tx.Where("id = ?", r.RedeemCodeID).First(&redeemCode).Error; err != nil {
		tx.Rollback()
		zap.L().Error("获取兑换码失败：", zap.Error(err))
		return entity.ErrInternal
	}

	// 检查兑换码是否已过期
	if time.Now().After(redeemCode.ExpirationAt) {
		tx.Rollback()
		return entity.ErrRedeemCodeExpired
	}

	// 检查兑换码是否已达到使用上限
	if redeemCode.UsedCount >= redeemCode.UsageLimit {
		tx.Rollback()
		return entity.ErrRedeemCodeUsedUp
	}

	// 增加使用记录
	redeemCodeRecord := model.RedeemCodeRecord{
		UserID:       r.UserID,
		RedeemCodeID: r.RedeemCodeID,
	}
	if err := tx.Create(&redeemCodeRecord).Error; err != nil {
		tx.Rollback()
		zap.L().Error("创建使用记录失败：", zap.Error(err))
		return entity.ErrInternal
	}

	// 更新兑换码已使用数量（乐观锁：使用 updated_at 作为版本控制）
	result := tx.Model(&model.RedeemCode{}).
		Where("id = ? AND updated_at = ? AND used_count < usage_limit", r.RedeemCodeID, redeemCode.UpdatedAt).
		Update("used_count", gorm.Expr("used_count + 1"))
	if result.Error != nil {
		tx.Rollback()
		zap.L().Error("更新兑换码使用数量失败：", zap.Error(result.Error))
		return entity.ErrInternal
	}
	if result.RowsAffected == 0 {
		tx.Rollback()
		return entity.ErrRedeemCodeUsedUp
	}

	// 如果达到使用上限，更新批次已使用数量
	if redeemCode.UsedCount+1 >= redeemCode.UsageLimit {
		if err := tx.Model(&model.RedeemCodeBatch{}).Where("id = ?", redeemCode.RedeemCodeBatchID).Update("used_count", gorm.Expr("used_count + 1")).Error; err != nil {
			tx.Rollback()
			zap.L().Error("更新批次使用数量失败：", zap.Error(err))
			return entity.ErrInternal
		}
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		zap.L().Error("提交事务失败：", zap.Error(err))
		return entity.ErrInternal
	}

	return nil
}
