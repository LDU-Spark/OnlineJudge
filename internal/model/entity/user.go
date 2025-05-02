// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Uuid        string      `json:"uuid"        orm:"uuid"        ` //
	Username    string      `json:"username"    orm:"username"    ` //
	Password    string      `json:"password"    orm:"password"    ` //
	Solved      int         `json:"solved"      orm:"solved"      ` //
	Submitted   int         `json:"submitted"   orm:"submitted"   ` //
	Locked      bool        `json:"locked"      orm:"locked"      ` //
	Sex         string      `json:"sex"         orm:"sex"         ` //
	Nickname    string      `json:"nickname"    orm:"nickname"    ` //
	School      string      `json:"school"      orm:"school"      ` //
	Department  string      `json:"department"  orm:"department"  ` //
	Major       string      `json:"major"       orm:"major"       ` //
	Email       string      `json:"email"       orm:"email"       ` //
	CreateAt    *gtime.Time `json:"create_at"   orm:"create_at"   ` //
	UpdateAt    *gtime.Time `json:"update_at"   orm:"update_at"   ` //
	DeleteAt    *gtime.Time `json:"delete_at"   orm:"delete_at"   ` //
	Description string      `json:"description" orm:"description" ` //
}
