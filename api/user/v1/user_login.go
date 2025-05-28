package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type LoginReq struct {
	g.Meta   `path:"/login" method:"post"`
	Username string `v:"required|length:3,16" json:"username"`
	Password string `v:"required|length:3,16" json:"password"`
}

type LoginRes struct {
	Token string `json:"token"`
}
