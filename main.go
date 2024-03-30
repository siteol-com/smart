package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"siteol.com/smart/src/service/plat/platServer"
	"syscall"
	"time"

	"siteol.com/smart/src/common/log"
	"siteol.com/smart/src/common/mysql"
	"siteol.com/smart/src/common/redis"
	"siteol.com/smart/src/common/utils"
	"siteol.com/smart/src/config"
	"siteol.com/smart/src/router"
)

// @title					smart
// @version					1.0
// @description.markdown
// @contact.name			smart
// @contact.url				https://smart.siteol.com
// @contact.email			dev@siteol.com
// @host					localhost:8000
// @BasePath				/
// @accept					json

// @securityDefinitions.apikey	Token
// @in							header
// @name						Token

// @x-logo	{"url" :"/docs/file/logo.png","altText":"smart"}

// @tag.name		开放接口
// @tag.description	基础开发接口

// @x-tagGroups	[{ "name": "基础", "tags": ["开放接口"]},{ "name": "平台", "tags": ["租户管理","集团部门","角色配置","登陆账号","访问权限","路由接口","响应配置","数据字典"]}]

// 主函数
func main() {
	traceId := fmt.Sprintf("%s%s", config.SysNode, "INIT")
	// 初始化数据库
	mysql.Init(traceId)
	// 初始化Redis
	redis.Init(traceId)
	// 业务初始化
	serviceInit(traceId)
	// 初始化路由
	newRouter := router.NewRouter()
	httpServer := &http.Server{Addr: config.JsonConfig.Server.Port, Handler: newRouter}
	// 启用HTTP服务 - 注册自定义路由
	go utils.RecoverWrap(func() {
		log.InfoTF(traceId, "Server Listening on port %s success .", config.JsonConfig.Server.Port)
		if err := httpServer.ListenAndServe(); err != nil {
			log.ErrorTF(traceId, "Server Listening on port %s . Err %v", config.JsonConfig.Server.Port, err)
			os.Exit(1)
		}
	})()
	// 优雅关机
	var cancelFunc func()
	defer cancelFunc()
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		sig := <-sigChan
		switch sig {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.ErrorTF(traceId, "Server Get a signal %s, Stop the consume process", sig.String())
			ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
			cancelFunc = cancel
			// gracefully shutdown with timeout
			_ = httpServer.Shutdown(ctx)
			return
		}
	}
}

// serviceInit 业务初始化
func serviceInit(traceId string) {
	// 主服务进行响应码初始化
	if config.SysNode == "APP01" {
		// 系统配置初始化
		err := platServer.SyncSysConfigCache(traceId)
		if err != nil {
			log.ErrorTF(traceId, "InitSysConfigCache Fail . Err is : %v", err)
			os.Exit(1)
		}
		log.InfoTF(traceId, "InitSysConfigCache success .")
		// 响应码配置初始化
		err = platServer.SyncResponseCache(traceId)
		if err != nil {
			log.ErrorTF(traceId, "InitResponseCache Fail . Err is : %v", err)
			os.Exit(1)
		}
		log.InfoTF(traceId, "InitResponseCache success .")
	}
}
