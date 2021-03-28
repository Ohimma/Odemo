package user

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ohimma/odemo/gin/handler"
	"github.com/ohimma/odemo/gin/middleware"
	"github.com/ohimma/odemo/gin/model"
)

type AuthRequest struct {
	Id   int    `json:"id" form:"id"`
	Pid  int16  `json:"pid" form:"pid"`
	Url  string `json:"url" form:"url"`
	Name string `json:"name" form:"name"`

	Icon   string `json:"icon" form:"icon"`
	Method string `json:"method" form:"method"`
	Level  int16  `json:"level" form:"level"`
	Sort   int32  `json:"sort" form:"sort"`
	Status bool   `json:"status" form:"status"`
	Detail string `json:"detail" form:"detail"`
}
type AuthListResponse struct {
	Total int64             `json:"total"`
	List  []*model.UserAuth `json:"list"`
}

func AuthList(c *gin.Context) {
	// 1. 解析 query 请求参数
	limit1 := c.DefaultQuery("limit", "0")
	offset1 := c.DefaultQuery("offset", "0")
	order_by := c.DefaultQuery("order_by", "id")
	order_sort := c.DefaultQuery("order_sort", "asc")
	url := c.Query("url")
	method := c.Query("method")
	offset, _ := strconv.Atoi(offset1)
	limit, _ := strconv.Atoi(limit1)

	// 2. 查询数据库
	middleware.Logger.Info("11111", offset, limit, order_by, order_sort, url, method)
	Info, count, err := model.AuthList(offset, limit, order_by, order_sort, url, method)
	if err != nil {
		middleware.Logger.Error("列表查询失败: ", Info, count, err)
		handler.ResponseError(c, 500, err, "列表查询失败")
		return
	}

	// 3. 构造返回函数
	res := AuthListResponse{
		Total: count,
		List:  Info,
	}

	middleware.Logger.Info("success: ", res)
	handler.ResponseOK(c, 200, res, "success")
}

func AuthCreate(c *gin.Context) {
	// 1. 解析数据格式 并验证
	var req AuthRequest
	if err := c.Bind(&req); err != nil {
		handler.ResponseError(c, 500, err, "解析请求参数失败")
		return
	}
	middleware.Logger.Info("req = ", req)

	if req.Url == "" || req.Method == "" {
		middleware.Logger.Warn("参数错误 ")
		handler.ResponseError(c, 500, "", "参数错误")
		return
	}

	// 2. 得判断唯一性，method + url 确定
	_, count, err := model.AuthExitsByUrlMethod(req.Url, req.Method)
	if count != 0 {
		middleware.Logger.Warn("记录已经存在: ", err)
		handler.ResponseError(c, 500, "", "记录已经存在")
		return
	}

	// 3. 组合数据库字段
	db := &model.UserAuth{
		Pid:    req.Pid,
		Level:  req.Level,
		Url:    req.Url,
		Method: req.Method,
		Name:   req.Name,
		Icon:   req.Icon,
		Sort:   req.Sort,
		Detail: req.Detail,
		Status: req.Status,
	}
	db.CreateAt = time.Now().Format("2006-01-02 15:04:05")
	db.UpdateAt = time.Now().Format("2006-01-02 15:04:05")

	// 4. 将记录插入数据库
	Info, count, err := model.AuthCreate(db)
	if err != nil {
		middleware.Logger.Info("创建失败 = ", err)
		handler.ResponseError(c, 500, err, "创建失败")
		return
	}
	// 5. 构造返回数据
	middleware.Logger.Info("success = ", Info)
	handler.ResponseOK(c, 200, Info, "success")
}

func AuthDelete(c *gin.Context) {
	// 1. 验证请求参数
	var req AuthRequest
	if err := c.Bind(&req); err != nil {
		handler.ResponseError(c, 500, err, "解析请求参数失败")
		return
	}
	middleware.Logger.Info("req = ", req)

	// 2. 验证记录是否存在
	_, count, err := model.AuthExits(req.Id)
	if count == 0 {
		middleware.Logger.Warn("记录不存在: ", err)
		handler.ResponseError(c, 500, "", "记录不存在")
		return
	}

	// 3. 组合数据库字段
	db := model.UserAuth{
		Id:     uint(req.Id),
		Pid:    req.Pid,
		Level:  req.Level,
		Url:    req.Url,
		Method: req.Method,
		Name:   req.Name,
		Icon:   req.Icon,
		Sort:   req.Sort,
		Detail: req.Detail,
		Status: req.Status,
	}
	// 编辑用户id 需要动态获取
	db.UpdateAt = time.Now().Format("2006-01-02 15:04:05")
	db.DeleteAt = time.Now().Format("2006-01-02 15:04:05")

	// 4. 删除记录
	Info, count, err := model.AuthDelete(&db)
	if err != nil {
		middleware.Logger.Error("删除失败: ", err)
		handler.ResponseError(c, 500, err, "删除失败")
		return
	}
	// 5. 构造返回函数
	middleware.Logger.Info("success ", Info)
	handler.ResponseOK(c, 200, Info, "success")
}

func AuthUpdate(c *gin.Context) {
	// 1. 验证传入参数
	var req AuthRequest
	if err := c.Bind(&req); err != nil {
		handler.ResponseError(c, 500, err, "解析请求参数失败")
		return
	}
	middleware.Logger.Info("req = ", req)

	if req.Url == "" || req.Method == "" {
		middleware.Logger.Warn("参数错误 ")
		handler.ResponseError(c, 500, "", "参数错误")
		return
	}

	// 2. 验证记录是否存在
	Info, count, err := model.AuthExits(req.Id)
	if count == 0 {
		middleware.Logger.Warn("记录不存在: ", err)
		handler.ResponseError(c, 500, "", "记录不存在")
		return
	}

	// 3. 组合数据库字段
	db := model.UserAuth{
		Id:     uint(req.Id),
		Pid:    req.Pid,
		Level:  req.Level,
		Url:    req.Url,
		Method: req.Method,
		Name:   req.Name,
		Icon:   req.Icon,
		Sort:   req.Sort,
		Detail: req.Detail,
		Status: req.Status,
	}
	// 最后编辑用户id 需要动态获取
	db.UpdateAt = time.Now().Format("2006-01-02 15:04:05")

	// 插曲，多一步验证是否有更新数据
	if req.Pid == Info.Pid &&
		req.Url == Info.Url &&
		req.Level == Info.Level &&
		req.Method == Info.Method &&
		req.Name == Info.Name &&
		req.Icon == Info.Icon &&
		req.Sort == Info.Sort &&
		req.Detail == Info.Detail &&
		req.Status == Info.Status {
		handler.ResponseError(c, 200, req, "用户没有更新内容")
		return
	}
	// 4. 将记录插入数据库
	Info2, count, err := model.AuthUpdate(&db)
	if err != nil {
		middleware.Logger.Info("更新失败 = ", err)
		handler.ResponseError(c, 500, err, "更新失败")
		return
	}

	// 5. 构造返回数据
	middleware.Logger.Info("success = ", Info2)
	handler.ResponseOK(c, 200, Info2, "success")
}
