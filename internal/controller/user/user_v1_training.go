package user

import (
	"context"
	"spoj/api/user/v1"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

func (c *ControllerV1) Training(ctx context.Context, req *v1.TrainingReq) (res *v1.TrainingRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
