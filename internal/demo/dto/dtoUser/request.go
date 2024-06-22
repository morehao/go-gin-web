package dtoUser

type GetUserReq struct {
	ID uint64 `json:"id" form:"id"` // 用户ID
}
