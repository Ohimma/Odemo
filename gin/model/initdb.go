package model

import (
	"fmt"

	"github.com/ohimma/odemo/gin/config"
	"github.com/ohimma/odemo/gin/middleware"

	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var Redis *redis.Client

func InitMysql() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
		config.Conf.Mysql.User,
		config.Conf.Mysql.Password,
		config.Conf.Mysql.Host,
		config.Conf.Mysql.Db,
		config.Conf.Mysql.Config,
	)
	fmt.Println("initmysql dsn = ", dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("connect mysql succuss", db)

	// 迁移 添加测试数据
	if config.Conf.Mysql.AutoMigrate {
		db.AutoMigrate(new(UserUser), new(UserAuth), new(UserRole))

		// 1. 验证该用户是否存在
		result := db.Where("name = 'admin'").Find(&UserUser{})
		if result.RowsAffected == 0 {
			middleware.Logger.Error("未初始化数据库 = ", result.Error, result.RowsAffected)

			auth := []UserAuth{
				{Pid: 0, Level: 1, Name: "用户管理", Url: "/user", Method: "GET", Sort: 1, Icon: "el-icon-user-solid", Detail: "用户管理", Status: true},

				{Pid: 1, Level: 2, Name: "用户", Url: "/user/user", Method: "GET", Sort: 5, Status: true},
				{Pid: 2, Level: 3, Name: "用户新增", Url: "/user/user", Method: "POST", Sort: 5, Status: true},
				{Pid: 2, Level: 3, Name: "用户编辑", Url: "/user/user", Method: "PUT", Sort: 10, Status: true},
				{Pid: 2, Level: 3, Name: "用户删除", Url: "/user/user", Method: "DELETE", Sort: 20, Status: true},

				{Pid: 1, Level: 2, Name: "角色", Url: "/user/role", Method: "GET", Sort: 10, Status: true},
				{Pid: 6, Level: 3, Name: "新增", Url: "/user/role", Method: "POST", Sort: 5, Status: true},
				{Pid: 6, Level: 3, Name: "编辑", Url: "/user/role", Method: "PUT", Sort: 10, Status: true},
				{Pid: 6, Level: 3, Name: "删除", Url: "/user/role", Method: "DELETE", Sort: 20, Status: true},

				{Pid: 1, Level: 2, Name: "权限因子", Url: "/user/auth", Method: "GET", Sort: 20, Status: true},
				{Pid: 10, Level: 3, Name: "新增", Url: "/user/auth", Method: "POST", Sort: 20, Status: true},
				{Pid: 10, Level: 3, Name: "编辑", Url: "/user/auth", Method: "PUT", Sort: 20, Status: true},
				{Pid: 10, Level: 3, Name: "删除", Url: "/user/auth", Method: "DELETE", Sort: 20, Status: true},

				{Pid: 0, Level: 1, Name: "Element", Url: "/element", Method: "GET", Sort: 1, Icon: "el-icon-eleme", Status: true},
				{Pid: 14, Level: 2, Name: "图片类", Url: "/element/images", Method: "GET", Sort: 5, Icon: "el-icon-eleme", Status: true},
				{Pid: 14, Level: 2, Name: "步骤类", Url: "/element/steps", Method: "GET", Sort: 10, Icon: "el-icon-eleme", Status: true},

				{Pid: 0, Level: 1, Name: "错误页", Url: "/error", Method: "GET", Sort: 1, Icon: "el-icon-question", Status: true},
				{Pid: 17, Level: 2, Name: "404", Url: "/404", Method: "GET", Sort: 1, Icon: "el-icon-question", Status: true},
			}
			result_auth := db.Create(&auth) // 通过数据的指针来创建
			fmt.Println("result2 = ", &result_auth, result_auth.Error)

			role := []UserRole{
				{Name: "管理员", AuthIds: "0", Detail: "拥有所有权限", Status: true},
				{Name: "用户管理员", AuthIds: "3,4,5", Detail: "有用户管理权限", Status: true},
				{Name: "角色权限管理员", AuthIds: "6,7,8,10,11,12,13,14,15,16", Detail: "有角色管理权限", Status: false},
			}
			result_role := db.Create(&role) // 通过数据的指针来创建
			fmt.Println("result2 = ", &result_role, result_role.Error)

			user := []UserUser{
				{Name: "admin", Phone: "1234", RoleIds: "0", Password: "d59dd9b2dd80ac58", Salt: "c15d52b3-d71d-413e-b671-dc1adab56c78", Avatar: "admin"},
				{Name: "user", Phone: "12345", RoleIds: "2", Password: "a151b466e68e9724", Salt: "14e35c41-b8d3-437e-b05d-792e9ba418d0", Avatar: "test"},
				{Name: "view", Phone: "view123", RoleIds: "3", Password: "cd4bc6ec70f6178f", Salt: "be2f459d-2d7f-460d-8670-0ff5ffaf11c7", Avatar: "view"},
			}
			result_user := db.Create(&user) // 通过数据的指针来创建
			fmt.Println("result2 = ", &result_user, result_user.Error)

		}
	}

	DB = db
}

func InitRedis() {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Conf.Redis.Host,
		Password: config.Conf.Redis.Password,
		DB:       config.Conf.Redis.Db,
	})
	pong, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("connect redis succuss", pong)
	Redis = client
}
