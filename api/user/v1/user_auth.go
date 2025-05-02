package v1

import (
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

type ProtectedReq struct {
	g.Meta `path:"/protected" method:"get" tags:"Auth"`
}

type ProtectedRes struct {
	Username string    `json:"username"`
	Time     time.Time `json:"time"`
}
