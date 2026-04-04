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
	// 参数验证
	if r.Title == "" || r.TotalCount <= 0 || r.UsageLimit <= 0 {
		return nil, entity.ErrParam
	}

	// 获取用户信息
	var user model.User
	if err := user.GetByID(userID); err != nil {
		zap.L().Error("获取用户信息失败：", zap.Error(err), zap.Int64("userID", userID))
		return nil, entity.ErrInternal
	}

	// 解析日期
	startedAt, err := helper.ParseDatetime(r.StartedAt)
	if err != nil {
		zap.L().Error("解析开始时间失败：", zap.Error(err), zap.String("startedAt", r.StartedAt))
		return nil, entity.ErrParam
	}
	endedAt, err := helper.ParseDatetime(r.EndedAt)
	if err != nil {
		zap.L().Error("解析结束时间失败：", zap.Error(err), zap.String("endedAt", r.EndedAt))
		return nil, entity.ErrParam
	}

	// 验证时间范围
	if endedAt.Before(*startedAt) {
		zap.L().Error("结束时间早于开始时间", zap.String("startedAt", r.StartedAt), zap.String("endedAt", r.EndedAt))
		return nil, entity.ErrParam
	}

	// 使用事务确保原子性
	tx := global.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			zap.L().Error("事务panic", zap.Any("panic", r))
		}
	}()

	// 创建兑换码批次
	newRedeemCodeBatch := model.RedeemCodeBatch{
		Title:       r.Title,
		Description: r.Description,
		UsageLimit:  r.UsageLimit,
		TotalCount:  r.TotalCount,
		StartedAt:   *startedAt,
		EndedAt:     *endedAt,
		Status:      1,
		CreatorID:   user.ID,
		CreatorName: user.Nickname,
	}
	if result := tx.Create(&newRedeemCodeBatch); result.Error != nil {
		tx.Rollback()
		zap.L().Error("创建兑换码批次失败：", zap.Error(result.Error))
		return nil, entity.ErrOperationFailed
	}

	// 批量创建兑换码（分批次处理，避免内存问题）
	const batchSize = 1000
	totalCount := int(r.TotalCount)
	for i := 0; i < totalCount; i += batchSize {
		end := i + batchSize
		if end > totalCount {
			end = totalCount
		}

		var newRedeemCodes []model.RedeemCode
		for j := i; j < end; j++ {
			redeemCode := model.RedeemCode{
				RedeemCodeBatchID: newRedeemCodeBatch.ID,
				Title:             r.Title,
				Value:             uuid.Must(uuid.NewV7()).String(),
				UsageLimit:        r.UsageLimit,
				ExpirationAt:      *endedAt,
				IsDel:             2,
			}
			newRedeemCodes = append(newRedeemCodes, redeemCode)
		}

		if result := tx.Create(&newRedeemCodes); result.Error != nil {
			tx.Rollback()
			zap.L().Error("批量创建兑换码失败：", zap.Error(result.Error))
			return nil, entity.ErrOperationFailed
		}
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		zap.L().Error("提交事务失败：", zap.Error(err))
		return nil, entity.ErrOperationFailed
	}

	// 返回数据
	return &newRedeemCodeBatch, nil
}

func (l *RedeemCodeBatchLogic) Detail(c *gin.Context, id int64) (*model.RedeemCodeBatch, error) {
	// 参数验证
	if id <= 0 {
		return nil, entity.ErrParam
	}

	var redeemCodeBatch model.RedeemCodeBatch
	if err := redeemCodeBatch.GetByID(id); err != nil {
		zap.L().Error("获取兑换码批次详情失败：", zap.Error(err), zap.Int64("id", id))
		return nil, entity.ErrInternal
	}
	return &redeemCodeBatch, nil
}

func (l *RedeemCodeBatchLogic) GetList(c *gin.Context, r *entity.GetRedeemCodeBatchListDto) (*entity.GetRedeemCodeBatchListDvo, error) {
	// 验证分页参数
	if r.Page <= 0 {
		r.Page = 1
	}
	if r.PageSize <= 0 || r.PageSize > 100 {
		r.PageSize = 10
	}

	var redeemCodeBatch model.RedeemCodeBatch
	list, count := redeemCodeBatch.GetList(r.Page, r.PageSize)
	resp := &entity.GetRedeemCodeBatchListDvo{
		Page:     r.Page,
		PageSize: r.PageSize,
		Total:    count,
		Data:     list,
	}
	return resp, nil
}

func (l *RedeemCodeBatchLogic) Update(c *gin.Context, r *entity.UpdateRedeemCodeBatchDto) error {
	// 参数验证
	if r.ID <= 0 || r.Title == "" {
		return entity.ErrParam
	}

	// 解析日期
	startedAt, err := helper.ParseDatetime(r.StartedAt)
	if err != nil {
		zap.L().Error("解析开始时间失败：", zap.Error(err), zap.String("startedAt", r.StartedAt))
		return entity.ErrParam
	}
	endedAt, err := helper.ParseDatetime(r.EndedAt)
	if err != nil {
		zap.L().Error("解析结束时间失败：", zap.Error(err), zap.String("endedAt", r.EndedAt))
		return entity.ErrParam
	}

	// 验证时间范围
	if endedAt.Before(*startedAt) {
		zap.L().Error("结束时间早于开始时间", zap.String("startedAt", r.StartedAt), zap.String("endedAt", r.EndedAt))
		return entity.ErrParam
	}

	// 验证批次是否存在
	var existingBatch model.RedeemCodeBatch
	if err := existingBatch.GetByID(r.ID); err != nil {
		zap.L().Error("获取兑换码批次失败：", zap.Error(err), zap.Int64("id", r.ID))
		return entity.ErrInternal
	}

	redeemCodeBatch := model.RedeemCodeBatch{
		ID:          r.ID,
		Title:       r.Title,
		Description: r.Description,
		// UsageLimit:  r.UsageLimit, // 不允许修改使用限制
		// TotalCount:  r.TotalCount, // 不允许修改总数
		StartedAt: *startedAt,
		EndedAt:   *endedAt,
		Status:    r.Status,
	}

	// 更新批次信息
	result := global.DB.Model(&model.RedeemCodeBatch{}).Where("id=?", r.ID).Updates(&redeemCodeBatch)
	if result.Error != nil {
		zap.L().Error("更新兑换码批次失败：", zap.Error(result.Error), zap.Int64("id", r.ID))
		return entity.ErrInternal
	}

	// 如果结束时间发生变化，更新关联的兑换码过期时间
	if existingBatch.EndedAt != *endedAt {
		if err := global.DB.Model(&model.RedeemCode{}).Where("redeem_code_batch_id=?", r.ID).Update("expiration_at", *endedAt).Error; err != nil {
			zap.L().Error("更新兑换码过期时间失败：", zap.Error(err), zap.Int64("batchID", r.ID))
		}
	}

	return nil
}

func (l *RedeemCodeBatchLogic) Delete(c *gin.Context, id int64) error {
	// 参数验证
	if id <= 0 {
		return entity.ErrParam
	}

	// 验证批次是否存在
	var existingBatch model.RedeemCodeBatch
	if err := existingBatch.GetByID(id); err != nil {
		zap.L().Error("获取兑换码批次失败：", zap.Error(err), zap.Int64("id", id))
		return entity.ErrInternal
	}

	// 使用事务确保原子性
	tx := global.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			zap.L().Error("事务panic", zap.Any("panic", r))
		}
	}()

	// 更新批次状态为删除
	if err := tx.Model(&model.RedeemCodeBatch{}).Where("id=?", id).Update("status", 2).Error; err != nil {
		tx.Rollback()
		zap.L().Error("删除兑换码批次失败：", zap.Error(err), zap.Int64("id", id))
		return entity.ErrInternal
	}

	// 更新关联的兑换码状态
	if err := tx.Model(&model.RedeemCode{}).Where("redeem_code_batch_id=?", id).Update("is_del", 1).Error; err != nil {
		tx.Rollback()
		zap.L().Error("更新兑换码状态失败：", zap.Error(err), zap.Int64("batchID", id))
		return entity.ErrInternal
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		zap.L().Error("提交事务失败：", zap.Error(err))
		return entity.ErrOperationFailed
	}

	return nil
}
