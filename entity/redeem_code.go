package entity

import (
	"github.com/1348453525/user-redeem-code-gin/model"
)

// RedeemCodeDetailDto ID64

// RedeemCodeDetailDvo &model.RedeemCode{}

type GetRedeemCodeListDto struct {
	// ID       int64 `json:"id" validate:"required,gte=1"` // batch id
	Page     int32 `form:"page" validate:"required,gte=1"`
	PageSize int32 `form:"page_size" validate:"required,gte=10"`
}

type GetRedeemCodeListDvo struct {
	Page     int32               `json:"page"`
	PageSize int32               `json:"page_size"`
	Total    int64               `json:"total"`
	Data     []*model.RedeemCode `json:"data"`
}

type UpdateRedeemCodeDto struct {
	ID    int64  `json:"id" validate:"required,gte=1"`
	Title string `json:"title"` // 标题
	// UsedCount    int32  `json:"used_count"`    // 已使用数量
	ExpirationAt string `json:"expiration_at"` // 过期时间
	IsDel        int32  `json:"is_del"`        // 是否删除：0，默认；1，已删除；2，未删除；
}

// UpdateRedeemCodeDvo &model.RedeemCode{}

// DeleteRedeemCodeDto ID64

type UseRedeemCodeDto struct {
	RedeemCodeID int64 `json:"redeem_code_id" validate:"required,gte=1"`
	UserID       int64 `json:"user_id" validate:"required,gte=1"`
}
