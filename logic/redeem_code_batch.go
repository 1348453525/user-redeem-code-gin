package logic

import (
	"github.com/1348453525/user-redeem-code-gin/entity"
	"github.com/1348453525/user-redeem-code-gin/global"
	"github.com/1348453525/user-redeem-code-gin/model"
	"github.com/1348453525/user-redeem-code-gin/pkg/helper"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid/v5"
	"go.uber.org/zap"
)

type RedeemCodeBatchLogic struct{}

func NewRedeemCodeBatchLogic() *RedeemCodeBatchLogic {
	return &RedeemCodeBatchLogic{}
}

func (l *RedeemCodeBatchLogic) Create(c *gin.Context, userID int64, r *entity.CreateRedeemCodeBatchDto) (*model.RedeemCodeBatch, error) {
	// 获取用户信息
	var userModel model.User
	if err := userModel.GetByID(userID); err != nil {
		return nil, entity.ErrInternal
	}

	// 创建兑换码批次
	startedAt, err := helper.ParseDatetime(r.StartedAt)
	if err != nil {
		return nil, entity.ErrParam
	}
	endedAt, err := helper.ParseDatetime(r.EndedAt)
	if err != nil {
		return nil, entity.ErrParam
	}
	redeemCodeBatchModel := model.RedeemCodeBatch{
		Title:       r.Title,
		Description: r.Description,
		UsageLimit:  r.UsageLimit,
		TotalCount:  r.TotalCount,
		StartedAt:   *startedAt,
		EndedAt:     *endedAt,
		Status:      1,
		CreatorID:   userModel.ID,
		CreatorName: userModel.Nickname,
	}
	if result := global.DB.Create(&redeemCodeBatchModel); result.Error != nil {
		return nil, entity.ErrOperationFailed
	}

	// 批量创建兑换码
	var redeemCodesModel []model.RedeemCode
	for i := 0; i < int(r.TotalCount); i++ {
		redeemCode := model.RedeemCode{
			RedeemCodeBatchID: redeemCodeBatchModel.ID,
			Title:             r.Title,
			Value:             uuid.Must(uuid.NewV7()).String(),
			UsageLimit:        r.UsageLimit,
			ExpirationAt:      *endedAt,
			IsDel:             2,
		}
		redeemCodesModel = append(redeemCodesModel, redeemCode)
	}
	if result := global.DB.Create(&redeemCodesModel); result.Error != nil {
		return nil, entity.ErrOperationFailed
	}

	// 返回数据
	return &redeemCodeBatchModel, nil
}

func (l *RedeemCodeBatchLogic) Detail(c *gin.Context, id int64) (*model.RedeemCodeBatch, error) {
	var redeemCodeBatchModel model.RedeemCodeBatch
	if err := redeemCodeBatchModel.GetByID(id); err != nil {
		return nil, entity.ErrInternal
	}
	return &redeemCodeBatchModel, nil
}

func (l *RedeemCodeBatchLogic) GetList(c *gin.Context, r *entity.GetRedeemCodeBatchListDto) (*entity.GetRedeemCodeBatchListDvo, error) {
	var redeemCodeBatchModel model.RedeemCodeBatch
	list, count := redeemCodeBatchModel.GetList(r.Page, r.PageSize)
	resp := &entity.GetRedeemCodeBatchListDvo{
		Page:     r.Page,
		PageSize: r.PageSize,
		Total:    count,
		Data:     list,
	}
	return resp, nil
}

func (l *RedeemCodeBatchLogic) Update(c *gin.Context, r *entity.UpdateRedeemCodeBatchDto) error {
	startedAt, err := helper.ParseDatetime(r.StartedAt)
	if err != nil {
		return entity.ErrParam
	}
	endedAt, err := helper.ParseDatetime(r.EndedAt)
	if err != nil {
		return entity.ErrParam
	}
	redeemCodeBatchModel := model.RedeemCodeBatch{
		ID:          r.ID,
		Title:       r.Title,
		Description: r.Description,
		// UsageLimit:  r.UsageLimit,
		// TotalCount:  r.TotalCount,
		StartedAt: *startedAt,
		EndedAt:   *endedAt,
		Status:    r.Status,
	}
	result := global.DB.Model(&model.RedeemCodeBatch{}).Where("id=?", r.ID).Updates(&redeemCodeBatchModel)
	if result.Error != nil {
		zap.L().Error("更新兑换码批次失败：", zap.Error(result.Error))
		return entity.ErrInternal
	}
	return nil
}

func (l *RedeemCodeBatchLogic) Delete(c *gin.Context, id int64) error {
	result := global.DB.Model(&model.RedeemCodeBatch{}).Where("id=?", id).Update("status", 2)
	if result.Error != nil {
		zap.L().Error("删除兑换码批次失败：", zap.Error(result.Error))
		return entity.ErrInternal
	}
	return nil
}
