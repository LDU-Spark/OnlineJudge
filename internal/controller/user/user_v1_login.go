package user

import (
	"context"
	"spoj/api/user/v1"
	"spoj/internal/middleware"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	// 校验请求
	err = g.Validator().Data(req).Run(ctx)
	if err != nil {
		g.Log().Error(ctx, "登录信息校验失败", err)
		return
	}

	// 校验密码
	err = c.user.Login(ctx, req.Username, req.Password)
	if err != nil {
		g.Log().Error(ctx, "用户名或密码错误", err)
		return
	}

	// 生成 token
	token, err := middleware.GenToken(ctx, req.Username)
	if err != nil {
		g.Log().Error(ctx, "生成 Token 失败", err)
		return
	}
	res = &v1.LoginRes{
		Token: token,
	}

	return
}
