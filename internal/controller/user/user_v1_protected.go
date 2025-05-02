package user

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"spoj/api/user/v1"
)

func (c *ControllerV1) Protected(ctx context.Context, req *v1.ProtectedReq) (res *v1.ProtectedRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
