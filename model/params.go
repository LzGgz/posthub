package model

type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type ParamVote struct {
	PostID    string `json:"post_id" binding:"required"`
	Direction int8   `json:"direction" binding:"oneof=1 0 -1"`
}
type ParamPostList struct {
	Page  int    `form:"page"`
	Order string `form:"order"`
}
