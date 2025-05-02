// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure of table user for DAO operations like Where/Data.
type User struct {
	g.Meta      `orm:"table:user, do:true"`
	Uuid        interface{} //
	Username    interface{} //
	Password    interface{} //
	Solved      interface{} //
	Submitted   interface{} //
	Locked      interface{} //
	Sex         interface{} //
	Nickname    interface{} //
	School      interface{} //
	Department  interface{} //
	Major       interface{} //
	Email       interface{} //
	CreateAt    *gtime.Time //
	UpdateAt    *gtime.Time //
	DeleteAt    *gtime.Time //
	Description interface{} //
}
