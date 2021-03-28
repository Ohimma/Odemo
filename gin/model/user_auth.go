package model

import (
	"fmt"

	"github.com/ohimma/odemo/gin/middleware"
)

// example
type UserAuth struct {
	Id uint `json:"id" gorm:"primaryKey; autoIncrement; unique"`

	Pid    int16  `json:"pid" gorm:"size:16; not null; default:0; comment:上级ID, 0为顶级"`
	Level  int16  `json:"level" gorm:"size:16; not null; default:1; comment:1-一级菜单 2二级菜单 3动作项"`
	Name   string `json:"name" gorm:"size:32; not null; comment:权限名"`
	Url    string `json:"url" gorm:"size:32; not null; comment:URL"`
	Method string `json:"method" gorm:"size:16; not null; default:GET; comment: POST/GET/PUT/DELETE"`
	Sort   int32  `json:"sort" gorm:"size:32; not null; default: 99"`

	Icon     string `json:"icon" gorm:"size:32; default:el-icon-menu"`
	Detail   string `json:"detail" gorm:"size:32; comment:备注信息"`
	Status   bool   `json:"status" gorm:"type:boolean;default:1; comment:状态1-显示，0-隐藏"`
	CreateId int    `json:"create_id" gorm:"size:64; not null; default:0; comment:创建者id"`
	UpdateId int    `json:"update_id" gorm:"size:32; not null; default:0; comment:修改者id"`
	CreateAt string `json:"create_at" gorm:"not null; default:2021-01-01"`
	UpdateAt string `json:"update_at" gorm:"not null; default:2021-01-01"`
	DeleteAt string `json:"delete_at" gorm:"default: null; comment:value不为空时表示删除"`
}

func AuthCreate(db *UserAuth) (*UserAuth, int64, error) {
	result := DB.Create(&db)

	middleware.Logger.Info("CreateAuth = ", *db, result.RowsAffected, result.Error)
	return db, result.RowsAffected, result.Error
}

func AuthUpdate(db *UserAuth) (*UserAuth, int64, error) {
	sql := DB.Where("id = ? ", db.Id)
	result := sql.Updates(&db)

	middleware.Logger.Info("UpdateAuth = ", *db, result, result.Error)
	return db, result.RowsAffected, result.Error
}

// 暂时没用
func AuthDelete(db *UserAuth) (*UserAuth, int64, error) {
	sql := DB.Where("id = ? ", db.Id)
	result := sql.Delete(&db)

	middleware.Logger.Info("DeleteAuth result = ", *db, result.RowsAffected, result.Error)
	return db, result.RowsAffected, result.Error
}

func AuthExits(id int) (*UserAuth, int64, error) {
	Info := &UserAuth{}
	sql := DB.Where("id = ?", id)
	result := sql.Find(&Info)

	middleware.Logger.Info("AuthExits result = ", sql, Info, result.RowsAffected, result.Error)
	return Info, result.RowsAffected, result.Error
}

// 因为没有唯一键值，用于创建时判断是否存在
func AuthExitsByUrlMethod(url, method string) (*UserAuth, int64, error) {
	Info := &UserAuth{}
	sql := DB.Where("url = ? and method = ?", url, method)
	result := sql.Find(&Info)

	middleware.Logger.Info("AuthExits result = ", sql, Info, result.RowsAffected, result.Error)
	return Info, result.RowsAffected, result.Error
}

func AuthList(offset, limit int, order_by, order_sort string, url, method string) ([]*UserAuth, int64, error) {
	// 定义返回的数据
	Info := make([]*UserAuth, 0)

	where := "delete_at is null"
	if url != "" || method != "" {
		where = fmt.Sprintf("delete_at is null and url like '%%%s%%' and method like '%%%s%%'", url, method)
	}

	middleware.Logger.Info("where = ", where)

	sql := DB.Where(where).Offset(offset).Limit(limit).Order(fmt.Sprintf("%s %s", order_by, order_sort))
	result := sql.Find(&Info)
	middleware.Logger.Info("List = ", where, sql, result)
	return Info, result.RowsAffected, result.Error
}
