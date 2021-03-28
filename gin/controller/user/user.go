package user

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ohimma/odemo/gin/handler"
	"github.com/ohimma/odemo/gin/middleware"
	"github.com/ohimma/odemo/gin/model"
	"github.com/ohimma/odemo/gin/utils"
)

type UserRequest struct {
	Id       uint   `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Password string `json:"password" form:"password"`
	Phone    string `json:"phone" form:"phone"`
	RoleIds  []int  `json:"role_ids" form:"role_ids"`
	Avatar   string `json:"avatar" form:"avatar"`
	Detail   string `json:"detail" form:"detail"`
	Status   bool   `json:"status" form:"status"`
}

type UserResponse struct {
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	RoleIds string `json:"role_ids" form:"role_ids"`
	Detail  string `json:"detail" form:"detail"`
	Status  bool   `json:"status" form:"status"`
}

// Auth AuthRequest `json:"auth"`
// Limit 指定获取记录的最大数量
// Offset 指定在开始返回记录之前要跳过的记录数量
// OrderBy 根据哪个字段排序
// OrderSort 升序还是降序 desc/asc

type UserListRequest struct {
	Name      string `json:"name"`
	Limit     int    `json:"limit" form:"limit"`
	Offset    int    `json:"offset" form:"limit"`
	OrderBy   string `json:"orderby"`
	OrderSort string `json:"ordersort"`
}
type UserListResponse struct {
	Total int64             `json:"total"`
	List  []*model.UserUser `json:"list"`
}

func UserList(c *gin.Context) {
	// var req UserListRequest
	// if err := c.Bind(&req); err != nil {
	// 	handler.ResponseError(c, 500, err, "解析请求参数失败")
	// 	return
	// }
	// middleware.Logger.Info("req = ", req)

	// 1. 解析 query 请求参数
	offset1 := c.DefaultQuery("offset", "0")
	limit1 := c.DefaultQuery("limit", "0")
	order_by := c.DefaultQuery("order_by", "id")
	order_sort := c.DefaultQuery("order_sort", "asc")
	name := c.Query("name")
	offset, _ := strconv.Atoi(offset1)
	limit, _ := strconv.Atoi(limit1)

	// 2. 查询数据库
	Info, count, err := model.UserList(offset, limit, order_by, order_sort, name)
	if err != nil {
		middleware.Logger.Error("列表查询失败: ", Info, count, err)
		handler.ResponseError(c, 500, err, "列表查询失败")
		return
	}
	// 3. 关联查询该id 所拥有的权限列表

	// 4. 构造返回函数
	res := UserListResponse{
		Total: count,
		List:  Info,
	}

	middleware.Logger.Info("success: ", res)
	handler.ResponseOK(c, 200, res, "success")
}

func UserCreate(c *gin.Context) {
	// 1. 验证数据格式
	var req UserRequest
	if err := c.Bind(&req); err != nil {
		handler.ResponseError(c, 500, err, "解析请求参数失败")
		return
	}
	middleware.Logger.Info("req = ", req)
	if req.Name == "" || req.Password == "" || req.Phone == "" {
		handler.ResponseError(c, 500, req, "请求参数错误")
		return
	}

	// 3. 合并要插入的数据
	db := model.UserUser{
		Id:      req.Id,
		Name:    req.Name,
		Phone:   req.Phone,
		RoleIds: utils.ArrToStr(req.RoleIds),
		Avatar:  req.Avatar,
		Detail:  req.Detail,
		Status:  req.Status,
	}
	db.CreateAt = time.Now().Format("2006-01-02 15:04:05")
	db.UpdateAt = time.Now().Format("2006-01-02 15:04:05")
	user_id := c.Request.Header.Get("user_id")
	db.CreateId, _ = strconv.Atoi(user_id)
	db.UpdateId, _ = strconv.Atoi(user_id)

	db.Salt = utils.GenerateSalt()
	db.Password = utils.MD5([]byte(req.Password + db.Salt))

	// 4. 将用户插入数据库
	if err := model.UserCreate(&db); err != nil {
		middleware.Logger.Info("创建失败 = ", err)
		handler.ResponseError(c, 500, err, "创建失败")
		return
	}

	middleware.Logger.Info("success = ", req)
	handler.ResponseOK(c, 200, req, "success")
}

func UserDelete(c *gin.Context) {
	// 1. 验证请求参数

	var req UserRequest
	if err := c.Bind(&req); err != nil {
		handler.ResponseError(c, 500, err, "解析请求参数失败")
		return
	}
	middleware.Logger.Info("req = ", req)

	// 2. 验证记录是否存在
	_, count, err := model.UserExits(req.Id)
	if count == 0 {
		middleware.Logger.Warn("记录不存在: ", err)
		handler.ResponseError(c, 500, err, "记录不存在")
		return
	}

	// 3. 组合数据库字段
	db := model.UserUser{
		Id:      req.Id,
		Name:    "del_" + req.Name,
		Phone:   "del_" + req.Phone,
		RoleIds: utils.ArrToStr(req.RoleIds),
		Avatar:  req.Avatar,
		Detail:  req.Detail,
		Status:  req.Status,
	}
	db.UpdateAt = time.Now().Format("2006-01-02 15:04:05")
	db.DeleteAt = time.Now().Format("2006-01-02 15:04:05")
	user_id := c.Request.Header.Get("user_id")
	db.UpdateId, _ = strconv.Atoi(user_id)

	// 4. 删除记录
	if err = model.UserDelete(&db); err != nil {
		middleware.Logger.Error("删除失败: ", err)
		handler.ResponseError(c, 500, err, "删除失败")
		return
	}
	// 5. 构造返回函数
	middleware.Logger.Info("success ", req)
	handler.ResponseOK(c, 200, req, "success")
}

func UserUpdate(c *gin.Context) {
	// 1. 验证传入参数
	var req UserRequest
	if err := c.ShouldBind(&req); err != nil {
		handler.ResponseError(c, 500, err, "解析请求参数失败")
		return
	}
	middleware.Logger.Info("req = ", req)
	middleware.Logger.Info("req2 = ", req.RoleIds, &req.RoleIds)

	if req.Id == 0 || req.Name == "" || req.Password == "" || req.Phone == "" {
		middleware.Logger.Warn("参数错误 ")
		handler.ResponseError(c, 500, "", "参数错误")
		return
	}

	// 2. 验证记录是否存在
	Info, count, err := model.UserExits(req.Id)
	if count == 0 {
		middleware.Logger.Warn("记录不存在: ", err)
		handler.ResponseError(c, 500, "", "记录不存在")
		return
	}

	// 4. 加密用户密码
	roleids := utils.ArrToStr(req.RoleIds)
	db := model.UserUser{
		Id:      req.Id,
		Name:    req.Name,
		Phone:   req.Phone,
		RoleIds: roleids,
		Avatar:  req.Avatar,
		Detail:  req.Detail,
		Status:  req.Status,
	}
	db.UpdateAt = time.Now().Format("2006-01-02 15:04:05")
	user_id := c.Request.Header.Get("user_id")
	db.UpdateId, _ = strconv.Atoi(user_id)

	// 3. 验证请求数据格式
	// RoleIds_tmp := fmt.Sprintf("%s", req.RoleIds)
	if req.Password == Info.Password {
		if req.Name == Info.Name &&
			req.Phone == Info.Phone &&
			roleids == Info.RoleIds &&
			req.Avatar == Info.Avatar &&
			req.Detail == Info.Detail &&
			req.Status == Info.Status {
			handler.ResponseError(c, 200, req, "用户没有更新内容")
			return
		}
	} else {
		db.Password = utils.MD5([]byte(req.Password + Info.Salt))
	}

	// 4. 更新数据库
	if err := model.UserUpdate(&db); err != nil {
		middleware.Logger.Info("更新失败 = ", err)
		handler.ResponseError(c, 500, err, "更新失败")
		return
	}

	// 5. 构造返回数据
	middleware.Logger.Info("success = ", req)
	handler.ResponseOK(c, 200, req, "success")
}
