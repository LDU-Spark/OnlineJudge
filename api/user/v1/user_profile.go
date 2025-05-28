package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type ProfileReq struct {
	g.Meta `path:"/profile/{username}" method:"get"`
}

type ProfileRes struct {
	Nickname string `json:"nickname"`
	Profile  string `json:"profile"`
}
