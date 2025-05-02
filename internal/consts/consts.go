package consts

import (
	"github.com/gogf/gf/v2/errors/gerror"
)

// Error Message
var (
	ErrUserExist       = gerror.New("用户已存在")
	ErrUserNotExist    = gerror.New("用户不存在")
	ErrInvalidUsername = gerror.New("用户名格式错误")
	ErrInvalidPassword = gerror.New("密码错误")
)
