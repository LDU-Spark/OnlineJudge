package logic

import (
	"context"
	"spoj/internal/consts"
	"spoj/internal/dao"
	"spoj/internal/model/do"
	"spoj/internal/model/entity"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
)

type User struct{}

func New() *User {
	return &User{}
}

// Login 登录
func (u *User) Login(ctx context.Context, username, password string) (err error) {
	md := dao.UserBase.Ctx(ctx)

	// 检查用户是否存在
	cnt, err := md.Where("username", username).Count()
	if err != nil {
		return
	}
	if cnt == 0 {
		err = consts.ErrUserNotExist
		return
	}

	// 检查密码是否正确
	user := &entity.UserBase{}
	err = md.Where("username", username).Scan(&user)
	if err != nil {
		return
	}
	md5Password, err := gmd5.Encrypt(password)
	if err != nil {
		return
	}
	if user.Password != md5Password {
		err = consts.ErrInvalidPassword
		g.Log().Error(ctx, err, user)
		return
	}

	return
}

// Register 注册
func (u *User) Register(ctx context.Context, username, password string) (err error) {
	md := dao.UserBase.Ctx(ctx)

	// 检查用户是否存在
	cnt, err := md.Where("username", username).Count()
	if err != nil {
		return
	}
	if cnt > 0 {
		err = consts.ErrUserExist
		return
	}
	g.Log().Info(ctx, cnt)

	// 往数据库添加用户
	md5Password, err := gmd5.Encrypt(password)
	if err != nil {
		return
	}
	res, err := md.Data(do.UserBase{
		Username: username,
		Password: md5Password,
		Nickname: username,
	}).Insert()
	if err != nil {
		return
	}
	g.Log().Info(ctx, res)

	return
}

// GetUserInfo 获取用户全部信息
func (u *User) GetUserInfo(ctx context.Context, username string) (userInfo *entity.UserBase, err error) {
	md := dao.UserBase.Ctx(ctx)

	// 获取用户信息
	userInfo = &entity.UserBase{}
	err = md.Where("username", username).Scan(&userInfo)
	if err != nil {
		return
	}

	return
}
