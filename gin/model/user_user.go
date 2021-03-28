package model

import (
	"fmt"
	_ "time"

	"github.com/ohimma/odemo/gin/middleware"
)

// example
type UserUser struct {
	Id uint `json:"id" gorm:"primaryKey; autoIncrement; unique"`

	Name      string `json:"name" gorm:"size:32; uniqueIndex; not null; comment:用户名"`
	Phone     string `json:"phone" gorm:"size:32; unique; not null"`
	RoleIds   string `json:"role_ids" gorm:"size:32; comment:角色列表字符串，如: 2,3,4"`
	Password  string `json:"password" gorm:"type:char(64); not null; comment:密码"`
	Salt      string `json:"salt" gorm:"size:64; not null; comment:密码salt "`
	Avatar    string `json:"avatar" gorm:"size:32" comment:"头像"`
	LoginTime string `json:"login_time" gorm:"size:32; comment:最后登录时间"`
	LoginIp   string `json:"login_ip" gorm:"size:16; comment:最后登录ip"`

	Detail   string `json:"detail" gorm:"size:32; comment:备注信息"`
	Status   bool   `json:"status" gorm:"type:boolean;not null; default:1; comment:状态1-启用，0-禁用"`
	CreateId int    `json:"create_id" gorm:"size:64;not null; default:0; comment:创建者id"`
	UpdateId int    `json:"update_id" gorm:"size:32;not null; default:0; comment:修改者id"`
	CreateAt string `json:"create_at" gorm:"not null; default:2021-01-01"`
	UpdateAt string `json:"update_at" gorm:"not null; default:2021-01-01"`
	DeleteAt string `json:"delete_at" gorm:"default: null; comment:value不为空时表示删除"`
}

func UserCreate(db *UserUser) error {
	result := DB.Create(&db)

	middleware.Logger.Info("CreateUsername = ", *db, result.RowsAffected, result.Error)
	return result.Error
}

func UserUpdate(db *UserUser) error {
	sql := DB.Where("id = ? ", db.Id)
	result := sql.Updates(&db)

	middleware.Logger.Info("UpdateUsername = ", *db, result, result.Error)
	return result.Error
}

func UserDelete(db *UserUser) error {
	sql := DB.Where("id = ? ", db.Id)
	result := sql.Delete(&db)
	// result := sql.Updates(&db)

	middleware.Logger.Info("UpdateUsername = ", *db, result, result.Error)
	return result.Error
}

func UserExits(id uint) (*UserUser, int64, error) {
	Info := &UserUser{} // 使用地址传参，比值传参要快、节省资源
	sql := DB.Where("id = ?", id)
	result := sql.Find(&Info)

	middleware.Logger.Info("AuthExits result = ", sql, Info, result.RowsAffected, result.Error)
	return Info, result.RowsAffected, result.Error
}

func UserExitsByName(name string) (*UserUser, int64, error) {
	Info := &UserUser{} // 使用地址传参，比值传参要快、节省资源
	sql := DB.Where("name = ?", name)
	result := sql.Find(&Info)

	middleware.Logger.Info("AuthExits result = ", sql, Info, result.RowsAffected, result.Error)
	return Info, result.RowsAffected, result.Error
}

func UserList(offset, limit int, order_by, order_sort string, name string) ([]*UserUser, int64, error) {
	Info := make([]*UserUser, 0)

	where := "delete_at is null"
	if name != "" {
		where = fmt.Sprintf("delete_at is null and name like '%%%s%%'", name)
	}

	sql := DB.Where(where).Offset(offset).Limit(limit).Order(fmt.Sprintf("%s %s", order_by, order_sort))
	result := sql.Find(&Info)
	middleware.Logger.Info("List = ", where, sql, result)
	return Info, result.RowsAffected, result.Error
}
