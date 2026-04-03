package logic

import (
	"github.com/1348453525/user-redeem-code-gin/entity"
	"github.com/1348453525/user-redeem-code-gin/global"
	"github.com/1348453525/user-redeem-code-gin/model"
	"github.com/1348453525/user-redeem-code-gin/pkg/helper"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type redeemCode struct{}

var RedeemCode redeemCode

func (l *redeemCode) Detail(c *gin.Context, id int64) (*model.RedeemCode, error) {
	var redeemCode model.RedeemCode
	if err := redeemCode.GetByID(id); err != nil {
		return nil, err
	}
	return &redeemCode, nil
}

func (l *redeemCode) GetList(c *gin.Context, r *entity.GetRedeemCodeListDto) (*entity.GetRedeemCodeListDvo, error) {
	var redeemCodeModel model.RedeemCode
	list, count := redeemCodeModel.GetList(r.Page, r.PageSize)
	return &entity.GetRedeemCodeListDvo{
		Page:     r.Page,
		PageSize: r.PageSize,
		Total:    count,
		Data:     list,
	}, nil
}

func (l *redeemCode) Update(c *gin.Context, r *entity.UpdateRedeemCodeDto) error {
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
		zap.L().Error("更新用户失败：", zap.Error(result.Error))
		return entity.ErrInternal
	}
	return nil
}

func (l *redeemCode) Delete(c *gin.Context, id int64) error {
	if result := global.DB.Model(&model.RedeemCode{}).Where("id=?", id).Update("is_del", 1); result.Error != nil {
		zap.L().Error("删除用户失败：", zap.Error(result.Error))
		return entity.ErrInternal
	}
	return nil
}

func (l *redeemCode) Use(c *gin.Context, r *entity.UseRedeemCodeDto) error {
	// redeem_code_record 增加使用记录
	// redeem_code.used_count 已使用数量+1
	// 如果 used_count>=usage_limit 时，redeem_code_batch.used_count 已使用数量+1
	return nil
}
