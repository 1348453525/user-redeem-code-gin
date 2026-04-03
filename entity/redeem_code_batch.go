package entity

import (
	"github.com/1348453525/user-redeem-code-gin/model"
)

type CreateRedeemCodeBatchDto struct {
	Title       string `json:"title"`       // 批次名称
	Description string `json:"description"` // 批次描述
	UsageLimit  int32  `json:"usage_limit"` // 可使用次数
	TotalCount  int32  `json:"total_count"` // 生成数量
	StartedAt   string `json:"started_at"`  // 开始时间
	EndedAt     string `json:"ended_at"`    // 结束时间
	Status      int32  `json:"status"`      // 状态：0，默认；1，启用；2，停用；
}

// CreateRedeemCodeBatchDvo &model.RedeemCodeBatch{}

// RedeemCodeBatchDetailDto ID64

// RedeemCodeBatchDetailDvo &model.RedeemCodeBatch{}

type GetRedeemCodeBatchListDto struct {
	Page     int32 `form:"page" validate:"required,gte=1"`
	PageSize int32 `form:"page_size" validate:"required,gte=10"`
}

type GetRedeemCodeBatchListDvo struct {
	Page     int32                    `json:"page"`
	PageSize int32                    `json:"page_size"`
	Total    int64                    `json:"total"`
	Data     []*model.RedeemCodeBatch `json:"data"`
}

type UpdateRedeemCodeBatchDto struct {
	ID          int64  `json:"id" validate:"required,gte=1"`
	Title       string `json:"title"`       // 批次名称
	Description string `json:"description"` // 批次描述
	// UsageLimit  int32  `json:"usage_limit"` // 可使用次数
	// TotalCount  int32  `json:"total_count"` // 生成数量
	// UsedCount int32  `json:"used_count"`
	StartedAt string `json:"started_at"` // 开始时间
	EndedAt   string `json:"ended_at"`   // 结束时间
	Status    int32  `json:"status"`     // 状态：0，默认；1，启用；2，停用；
}

// UpdateRedeemCodeBatchDvo &model.RedeemCodeBatch{}

// DeleteRedeemCodeBatchDto ID64
