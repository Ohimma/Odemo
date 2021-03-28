package user

import (
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ohimma/odemo/gin/handler"
	"github.com/ohimma/odemo/gin/middleware"
	"github.com/ohimma/odemo/gin/model"
	"github.com/ohimma/odemo/gin/utils"
)

type RoleRequest struct {
	Id      uint   `json:"id" form:"id"`
	Name    string `json:"name" form:"name"`
	AuthIds []int  `json:"auth_ids" form:"auth_ids"`
	Detail  string `json:"detail" form:"detail"`
	Status  bool   `json:"status" form:"status"`
}

type RoleResponse struct {
	Name    string `json:"name"`
	AuthIds string `json:"auth_ids" form:"auth_ids"`
	Detail  string `json:"detail" form:"detail"`
	Status  bool   `json:"status" form:"status"`
}

type RoleListResponse struct {
	Total     int64             `json:"total"`
	List      []*model.UserRole `json:"list"`
	UserAuths []int             `json:"user_auths"`
}

func RoleList(c *gin.Context) {
	// 1. 解析 query 翻页排序数据
	offset1 := c.DefaultQuery("offset", "0")
	limit1 := c.DefaultQuery("limit", "0")
	order_by := c.DefaultQuery("order_by", "id")
	order_sort := c.DefaultQuery("order_sort", "asc")
	name := c.Query("name")
	role_ids := c.Query("role_ids")
	offset, _ := strconv.Atoi(offset1)
	limit, _ := strconv.Atoi(limit1)

	// 2. 查询数据库
	Info, count, err := model.RoleList(offset, limit, order_by, order_sort, name)
	if err != nil {
		middleware.Logger.Error("列表查询失败: ", Info, count, err)
		handler.ResponseError(c, 500, err, "列表查询失败")
		return
	}

	// 3. 同时把 auth 列表转换为 数组形式，发给客户端
	// 6. 构造返回函数
	// 获取用户基本信息，上边验证时已经获取过了
	// 获取用户有的权限列表，关联查询
	role_ids = strings.Replace(role_ids, ",", "','", -1)
	middleware.Logger.Error("获取该用户权限列表 = ", role_ids)

	Auths, count, err := model.RoleListbyInId(role_ids)

	authids_tmp := ""
	for _, v := range Auths {
		authids_tmp = authids_tmp + v.AuthIds + ","
	}
	// 字符串转 int 列表
	var userAuths []int
	userAuths = utils.StrToArr(authids_tmp)

	// 4. 构造返回函数
	res := RoleListResponse{
		Total:     count,
		List:      Info,
		UserAuths: userAuths,
	}

	middleware.Logger.Info("success: ", res)
	handler.ResponseOK(c, 200, res, "success")
}

func RoleCreate(c *gin.Context) {
	// 1. 验证数据格式
	var req RoleRequest
	if err := c.Bind(&req); err != nil {
		handler.ResponseError(c, 500, err, "解析请求参数失败")
		return
	}
	middleware.Logger.Info("req = ", req)
	if req.Name == "" {
		handler.ResponseError(c, 500, req, "请求参数错误")
		return
	}

	// 3. 合并要插入的数据
	db := model.UserRole{
		Name:    req.Name,
		AuthIds: utils.ArrToStr(req.AuthIds),
		Detail:  req.Detail,
		Status:  req.Status,
	}
	db.CreateAt = time.Now().Format("2006-01-02 15:04:05")
	db.UpdateAt = time.Now().Format("2006-01-02 15:04:05")

	// 4. 将用户插入数据库
	Info, err := model.RoleCreate(&db)
	if err != nil {
		middleware.Logger.Info("创建失败 = ", err)
		handler.ResponseError(c, 500, err, "创建失败")
		return
	}

	middleware.Logger.Info("success = ", Info)
	handler.ResponseOK(c, 200, Info, "success")
}

func RoleDelete(c *gin.Context) {
	// 1. 验证请求参数
	var req RoleRequest
	if err := c.Bind(&req); err != nil {
		handler.ResponseError(c, 500, err, "解析请求参数失败")
		return
	}
	middleware.Logger.Info("req = ", req)

	// 2. 验证记录是否存在
	_, count, err := model.RoleExits(req.Id)
	if count == 0 {
		middleware.Logger.Warn("记录不存在: ", err)
		handler.ResponseError(c, 500, err, "记录不存在")
		return
	}

	// 3. 组合数据库字段

	db := model.UserRole{
		Id:      req.Id,
		Name:    "del_" + req.Name,
		AuthIds: utils.ArrToStr(req.AuthIds),
		Detail:  req.Detail,
		Status:  req.Status,
	}
	db.UpdateAt = time.Now().Format("2006-01-02 15:04:05")
	db.DeleteAt = time.Now().Format("2006-01-02 15:04:05")

	// 4. 删除记录
	Info, err := model.RoleDelete(&db)
	if err != nil {
		middleware.Logger.Error("删除失败: ", err)
		handler.ResponseError(c, 500, err, "删除失败")
		return
	}
	// 5. 构造返回函数
	middleware.Logger.Info("success ", Info)
	handler.ResponseOK(c, 200, Info, "success")
}

func RoleUpdate(c *gin.Context) {
	// 1. 验证传入参数
	var req RoleRequest
	if err := c.Bind(&req); err != nil {
		handler.ResponseError(c, 500, err, "解析请求参数失败")
		return
	}
	middleware.Logger.Info("req = ", req)

	if req.Name == "" {
		middleware.Logger.Warn("参数错误 ")
		handler.ResponseError(c, 500, "", "参数错误")
		return
	}

	// 2. 验证记录是否存在
	Info, count, err := model.RoleExits(req.Id)
	if count == 0 {
		middleware.Logger.Warn("记录不存在: ", err)
		handler.ResponseError(c, 500, "", "记录不存在")
		return
	}

	// 3. 验证请求数据格式
	ids := utils.ArrToStr(req.AuthIds)
	if req.Name == Info.Name &&
		ids == Info.AuthIds &&
		req.Detail == Info.Detail &&
		req.Status == Info.Status {
		handler.ResponseError(c, 200, req, "用户没有更新内容")
		return
	}

	// 4. 加密用户密码
	db := model.UserRole{
		Id:      req.Id,
		Name:    req.Name,
		AuthIds: ids,
		Detail:  req.Detail,
		Status:  req.Status,
	}
	db.UpdateAt = time.Now().Format("2006-01-02 15:04:05")

	// 4. 更新数据库
	Info2, err := model.RoleUpdate(&db)
	if err != nil {
		middleware.Logger.Info("更新失败 = ", err)
		handler.ResponseError(c, 500, err, "更新失败")
		return
	}

	// 5. 构造返回数据
	middleware.Logger.Info("success = ", Info2)
	handler.ResponseOK(c, 200, Info2, "success")
}
