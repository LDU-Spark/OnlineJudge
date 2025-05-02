package user

import (
	"context"
	"spoj/api/user/v1"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) Register(ctx context.Context, req *v1.RegisterReq) (res *v1.RegisterRes, err error) {
	g.Log().Info(ctx, req.Username, req.Password)

	// 校验参数
	err = g.Validator().Data(req).Run(ctx)
	if err != nil {
		g.Log().Errorf(ctx, "注册信息校验失败：%v", err)
		return
	}

	// 注册用户
	err = c.user.Register(ctx, req.Username, req.Password)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}

	return
}
