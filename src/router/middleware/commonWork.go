package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/log"
	"siteol.com/smart/src/common/model"
	"siteol.com/smart/src/common/mysql/platDb"
	"siteol.com/smart/src/common/redis"
	"siteol.com/smart/src/common/utils/security"
	"siteol.com/smart/src/config"
	"strings"
	"time"
)

// 设置语言上下文
func setLang(c *gin.Context) {
	lang := c.GetHeader(constant.ContextLang)
	if lang == "" || lang == "null" {
		lang = "zh-CN"
	}
	c.Set(constant.ContextLang, lang)
}

// 设置路由上下文
func setRouter(c *gin.Context, url, traceID string) (router *model.CacheRouter, ng bool) {
	router = model.CacheRouterNormal
	// 尝试读取缓存
	cacheGet := false
	str, err := redis.Get(constant.CacheRouters)
	if err == nil {
		resCacheMap := make(map[string]*model.CacheRouter, 0)
		err = json.Unmarshal([]byte(str), &resCacheMap)
		if err == nil {
			// 取得实际的配置
			routerGet, ok := resCacheMap[url]
			if !ok {
				log.WarnTF(traceID, "Get %s RouterCache Empty", url)
			} else {
				cacheGet = true
				router = routerGet
			}
		} else {
			log.ErrorTF(traceID, "Unmarshal RouterCache Fail . Err Is : %v", err)
		}
	} else {
		log.ErrorTF(traceID, "Get RouterCache Fail . Err Is : %v", err)
	}

	c.Set(constant.ContextRouterC, router)
	// 缓存未取得且未启动Debug
	if !cacheGet && !config.JsonConfig.Server.Debug {
		// 非Debug直接退出
		log.ErrorTF(traceID, "Get %s RouterCache NG", url, err)
		ng = true
		return
	}
	return
}

// setReq 处理请求 请求日志、入库等
func setReq(c *gin.Context, router *model.CacheRouter, url, traceID string) (ng bool) {
	origReq, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.ErrorTF(traceID, "Req Read Fail . Err Is : %v", err)
		ng = true
		return
	}
	// COPY 请求给 Body
	runReq := io.NopCloser(bytes.NewBuffer(origReq))
	c.Request.Body = runReq
	// 处理日志打印
	printStr := "Req Set Not Print"
	// 请求字符为空
	// JSON序列化 以及安全处理
	printSafeStr := security.SafeJson(string(origReq), router.ReqSecure)
	// 打印日志，尝试脱敏
	if router.ReqPrint {
		printStr = printSafeStr
	}
	log.InfoTF(traceID, "Req body : %s", printStr)
	// 日志入库
	if router.LogInDb {
		now := time.Now()
		logInDb := &platDb.RouterLog{
			Id:         0,
			AppName:    config.JsonConfig.Server.Name,
			AppNode:    config.SysNode,
			AppTraceId: traceID,
			ReqIp:      getClientIP(c),
			ReqUrl:     url,
			ReqBody:    printSafeStr,
			ReqAt:      &now,
			ResStatus:  0,
			ResBody:    "",
			ResTime:    nil,
		}
		// 传入入库日志对象
		c.Set(constant.ContextRouterI, logInDb)
	}
	return
}

// getClientIP 获取客户端IP
func getClientIP(c *gin.Context) string {
	req := c.Request
	// 尝试从 X-Forwarded-For 头部获取 IP 地址（如果使用了反向代理）
	forwarded := req.Header.Get("X-Forwarded-For")
	if forwarded != "" {
		// X-Forwarded-For 可能包含多个 IP 地址，通常由逗号分隔
		// 第一个通常是客户端的真实 IP 地址
		ips := strings.Split(forwarded, ",")
		clientIP := strings.TrimSpace(ips[0])
		return clientIP
	}
	// 如果 X-Forwarded-For 头部不存在，则使用 RemoteAddr
	return req.RemoteAddr
}
