package domain

type ReadyResponse struct {
	Message string
}

type ReadyRequest struct {
	BaseReq
	Shout string `json:"shout" binding:"required"`
}

type BaseReq struct{}

// Entityなどのドメイン層の構造体もこのファイルに記述する
