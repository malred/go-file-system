package go_file_web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	// db "go_file_db"
	utils "go_file_util"
	// "net/http"
)

// 定义路由组
// 组中组(嵌套路由组)
func DefineRouteGroup(fatherGroup *gin.RouterGroup, groupName string, r *gin.Engine) *gin.RouterGroup {
	var group *gin.RouterGroup
	// 如果有指定父级路由组 如 /v1
	if fatherGroup != nil {
		// v1/groupName
		group = fatherGroup.Group(groupName)
	} else {
		// /groupName
		group = r.Group(groupName)
	}
	// 返回路由组
	return group
}

// 存放 token (不同ip不同token)
var TokenMap = make(map[string]string, 10)

// 定时销毁token
func timeDT() {
	// 两小时后销毁
	t := utils.NewMyTimer(2*60*60, func() error {
		utils.DestoryTokenMap(TokenMap)
		return nil
	})
	t.Start()
	fmt.Println(TokenMap)
} 

// 路由器
// 启动默认的路由
var r = gin.Default()

// api版本路由组
var v1 *gin.RouterGroup

func Run() {
	// 使用中间件
	// 日志
	r.Use(gin.Logger())
	// 错误恢复
	r.Use(gin.Recovery())
	// 跨域
	r.Use(Core())
	// 验证
	r.Use(TokenValid())
	// 阻止缓存响应
	r.Use(NoCache())
	// 安全设置
	r.Use(Secure())
	// 创建路由组v1
	v1 = DefineRouteGroup(nil, "v1", r)
	// v1.POST("login", Login)
	// 注册file的路由
	registerFile(Core())
	registerUser(Core())
	// 开启静态资源
	r.Static("/files", "./tmp")
	// 启动webserver,监听本地127.0.0.1(默认)端口
	r.Run(":10101")
}
