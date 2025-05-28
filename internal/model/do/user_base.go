// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserBase is the golang structure of table user_base for DAO operations like Where/Data.
type UserBase struct {
	g.Meta    `orm:"table:user_base, do:true"`
	Uuid      interface{} //
	Username  interface{} //
	Password  interface{} //
	Nickname  interface{} //
	Profile   interface{} //
	Role      interface{} //
	Solved    interface{} //
	Submitted interface{} //
	Rating    interface{} //
	CreateAt  *gtime.Time //
	UpdateAt  *gtime.Time //
	DeleteAt  *gtime.Time //
}
