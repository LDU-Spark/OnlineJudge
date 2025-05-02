package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type RegisterReq struct {
	g.Meta     `path:"/register" method:"post"`
	Username   string `v:"required|length:3,12" json:"username"`
	Password   string `v:"required|length:6,16" json:"password"`
	RePassword string `v:"required|same:Password" json:"re_password"`
}

type RegisterRes struct{}
