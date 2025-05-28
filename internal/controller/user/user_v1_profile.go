package user

import (
	"context"
	"spoj/api/user/v1"
	"spoj/internal/consts"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) Profile(ctx context.Context, req *v1.ProfileReq) (res *v1.ProfileRes, err error) {
	// 提取 URL 中的 username
	r := g.RequestFromCtx(ctx)
	username := r.Get("username").Val().(string)

	// 判断非空
	if username == "" {
		err = consts.ErrInvalidUsername
		g.Log().Error(ctx, err)
		return
	}

	// 获取用户信息
	user, err := c.user.GetUserInfo(ctx, username)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	if user == nil {
		err = consts.ErrUserNotExist
		g.Log().Error(ctx, err)
		return
	}

	// 绑定返回
	res = &v1.ProfileRes{
		Nickname: user.Nickname,
		Profile:  user.Profile,
	}

	//g.Log().Debug(ctx, "user info:", user)

	return
}
