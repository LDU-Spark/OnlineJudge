package user

import (
	"spoj/internal/logic"
)

type ControllerV1 struct {
	user *logic.User
}

func NewV1() *ControllerV1 {
	return &ControllerV1{
		user: logic.New(),
	}
}
