package user

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ohimma/odemo/gin/config"
	"github.com/ohimma/odemo/gin/handler"
	"github.com/ohimma/odemo/gin/middleware"
	"github.com/ohimma/odemo/gin/model"
	"github.com/ohimma/odemo/gin/utils"
)

type LoginRequest struct {
	Name     string `json:"name" form:"name"`
	Password string `json:"password" form:"password"`
}
type LoginResponse struct {
	User      UserResponse `json:"user"`
	Token     string       `json:"token"`
	ExpiresAt int64        `json:"expiresAt"`
	// UserAuths []int        `json:"user_auths"`
}

func UserLogin(c *gin.Context) {
	// 1. 验证请求参数
	var req LoginRequest
	if err := c.Bind(&req); err != nil {
		middleware.Logger.Error("解析请求参数失败 = ", err)
		handler.ResponseUnauthorized(c, 500, err, "解析请求参数失败")
		return
	}
	middleware.Logger.Info("req = ", req)

	// 2. 验证 用户名密码是否为空
	if req.Name == "" || req.Password == "" {
		middleware.Logger.Error("请输入用户名和密码")
		handler.ResponseUnauthorized(c, 414, req, "请输入用户名和密码")
		return
	}

	// 3. 验证该用户是否存在
	Info, count, err := model.UserExitsByName(req.Name)
	if count == 0 || err != nil {
		middleware.Logger.Error("该用户不存在 = ", err)
		handler.ResponseUnauthorized(c, 401, err, "该用户不存在")
		return
	}

	// 4. 验证密码是否正确
	password := utils.MD5([]byte(req.Password + Info.Salt))
	if password != Info.Password {
		middleware.Logger.Error("密码不正确 = ", err)
		handler.ResponseUnauthorized(c, 403, err, "密码不正确")
		return
	}

	// 5. 密码匹配成功，生成 token
	j := &middleware.JWT{
		SigningKey: []byte(config.Conf.Jwt.Signkey), // 唯一签名
	}

	token, err := j.GenerateToken(req.Name)
	if err != nil {
		middleware.Logger.Error("获取token失败", err)
		handler.ResponseUnauthorized(c, 405, err, "获取token失败")
		return
	}

	// 5. token成功，即登录成功，把最新的登录ip写到数据库
	db := model.UserUser{
		Id:        Info.Id,
		LoginIp:   c.ClientIP(),
		LoginTime: time.Now().Format("2006-01-02 15:04:05"),
	}
	if err := model.UserUpdate(&db); err != nil {
		middleware.Logger.Info("更新失败 = ", err)
		handler.ResponseError(c, 500, err, "更新失败")
		return
	}

	res := LoginResponse{
		User: UserResponse{
			Id:      Info.Id,
			Name:    Info.Name,
			Phone:   Info.Phone,
			RoleIds: Info.RoleIds,
			Detail:  Info.Detail,
			Status:  Info.Status,
		},
		Token:     token,
		ExpiresAt: time.Now().Unix() + 60*60*2,
	}

	middleware.Logger.Info("success", res)
	handler.ResponseOK(c, 200, res, "success")
}

// 解析前端参数方式
// ShouldBindQuery  Only Bind Query String
// ShouldBind  Bind Query String or Post Data
// ShouldBindUri  Bind Uri
// ShouldBindHeader  Bind Header
// ShouldBind  Bind HTML checkboxes
// Bind(&objA)    Bind form-data request with custom struct
// ShouldBind(&objA) Try to bind body into different structs
// ShouldBindBodyWith(&objB, binding.JSON)
// ShouldBindBodyWith(&objB, binding.XML)

// 汇总 解析单个参数
// c.Param()
// c.Query
// c.DefaultQuery
// c.PostForm
// c.DefaultPostForm
// c.QueryMap
// c.PostFormMap
// c.FormFile
// c.MultipartForm

// 汇总 解析绑定方法
// c.Bind
// c.BindJSON
// c.BindXML
// c.BindQuery
// c.BindYAML
// c.ShouldBind
// c.ShouldBindJSON
// c.ShouldBindXML
// c.ShouldBindQuery
// c.ShouldBindYAML
