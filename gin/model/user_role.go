package model

import (
	"fmt"

	"github.com/ohimma/odemo/gin/middleware"
)

// example
type UserRole struct {
	Id uint `json:"id" gorm:"primaryKey; autoIncrement; unique"`

	Name    string `json:"name" gorm:"size:32; unique; uniqueIndex; not null; comment:角色名"`
	AuthIds string `json:"auth_ids" gorm:"size:64; comment:权限因子列表"`

	Detail   string `json:"detail" gorm:"size:32; comment:备注信息"`
	Status   bool   `json:"status" gorm:"type:bool;default:1; comment:状态1-正常，0-删除"`
	CreateId int    `json:"create_id" gorm:"size:64; not null; comment:创建者id"`
	UpdateId int    `json:"update_id" gorm:"size:32; not null; comment:修改者id"`
	CreateAt string `json:"create_at" gorm:"not null; default:2021-01-01"`
	UpdateAt string `json:"update_at" gorm:"not null; default:2021-01-01"`
	DeleteAt string `json:"delete_at" gorm:"default: null; comment:value不为空时表示删除"`
}

func RoleCreate(db *UserRole) (*UserRole, error) {
	result := DB.Create(&db)

	middleware.Logger.Info("CreateUsername = ", *db, result.RowsAffected, result.Error)
	return db, result.Error
}

func RoleUpdate(db *UserRole) (*UserRole, error) {
	sql := DB.Where("id = ? ", db.Id)
	result := sql.Updates(&db)

	middleware.Logger.Info("UpdateUsername = ", *db, result, result.Error)
	return db, result.Error
}

// 暂时没用，不真正删除
func RoleDelete(db *UserRole) (*UserRole, error) {
	sql := DB.Where("id = ? ", db.Id)
	result := sql.Delete(&db)

	middleware.Logger.Info("DeleteUsername result = ", *db, result.RowsAffected, result.Error)
	return db, result.Error
}

func RoleExits(id uint) (*UserRole, int64, error) {
	Info := &UserRole{} // 使用地址传参，比值传参要快、节省资源
	// where := "1=1"
	sql := DB.Where("id = ?", id)
	result := sql.Find(&Info)

	middleware.Logger.Info("AuthExits result = ", sql, Info, result.RowsAffected, result.Error)
	return Info, result.RowsAffected, result.Error
}

func RoleList(offset, limit int, orderby, ordersort string, name string) ([]*UserRole, int64, error) {
	Info := make([]*UserRole, 0)

	where := "delete_at is null"
	if name != "" {
		where = fmt.Sprintf("delete_at is null and name like '%%%s%%'", name)
	}

	sql := DB.Where(where).Offset(offset).Limit(limit).Order(fmt.Sprintf("%s %s", orderby, ordersort))
	result := sql.Find(&Info)
	middleware.Logger.Info("List = ", where, sql, result)
	return Info, result.RowsAffected, result.Error
}

// 关联查询，通过用户id，查询该用户有的权限，用户id -> role id -> 可以了
func RoleListbyInId(id string) ([]*UserRole, int64, error) {
	Info := make([]*UserRole, 0)

	where := fmt.Sprintf("delete_at is null and id in ('%s') ", id)

	sql := DB.Where(where)
	result := sql.Find(&Info)
	middleware.Logger.Info("xxxxxx", where)
	return Info, result.RowsAffected, result.Error
}
