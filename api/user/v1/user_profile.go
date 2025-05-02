package v1

import (
	"spoj/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type ProfileReq struct {
	g.Meta `path:"/profile/{username}" method:"get"`
}

type ProfileRes = entity.User
