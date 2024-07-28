package objCommon

type OperatorBaseInfo struct {
	CreatedBy   uint64 `json:"createdBy" form:"createdBy"`     // 创建人id
	UpdatedBy   uint64 `json:"updatedBy" form:"updatedBy"`     // 更新人id
	CreatedTime int64  `json:"createdTime" form:"createdTime"` // 创建时间
	UpdatedTime int64  `json:"updatedTime" form:"updatedTime"` // 更新时间
}
