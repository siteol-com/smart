package commHander

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/model"
	"siteol.com/smart/src/service"
	"strings"
)

// Sample godoc
// @id			 Sample示例
// @Summary      通用API示例
// @Description  系统API基本示例
// @Router       /docs/sample [post]
// @Tags         开放接口
// @Accept       json
// @Produce      json
// @Security	 Token
// @Param        req body model.DocsReq true "示例请求"
// @Success      200 {object} model.DocsOk "业务受理成功"
// @Failure      400 {object} model.DocsVail "数据校验失败"
// @Failure      401 {object} model.DocsAuthLg "当前尚未登陆"
// @Failure      403 {object} model.DocsAuthNg "权限校验失败"
// @Failure      500 {object} model.DocsErr "服务系统异常"
func Sample(c *gin.Context) {
	_, req, err := service.ValidateReqObj(c, &model.DocsReq{})
	if err != nil {
		c.JSON(http.StatusBadRequest, model.DocsVail{Code: constant.ValidErr, Msg: "参数非法"})
		return
	}
	DocsReq := req.(*model.DocsReq)
	switch DocsReq.HttpCode {
	case http.StatusInternalServerError: // 500
		c.JSON(http.StatusInternalServerError, model.DocsErr{Code: constant.Error, Msg: "系统异常"})
	case http.StatusBadRequest: // 400
		c.JSON(http.StatusBadRequest, model.DocsVail{Code: constant.ValidErr, Msg: "参数非法"})
	case http.StatusUnauthorized: // 401
		c.JSON(http.StatusBadRequest, model.DocsAuthLg{Code: constant.LoginErr, Msg: "当前尚未登陆"})
	case http.StatusForbidden: // 403
		c.JSON(http.StatusBadRequest, model.DocsAuthNg{Code: constant.AuthErr, Msg: "禁止访问"})
	default:
		c.JSON(http.StatusOK, model.DocsOk{Code: constant.Success, Msg: "业务请求成功"})
	}
	return
}

// ReDoc HTML加载
func ReDoc(c *gin.Context) {
	c.Data(http.StatusOK, "text/html", []byte(constant.DocHtml))
	return
}

// SwaggerDoc HTML加载
func SwaggerDoc(c *gin.Context) {
	c.Data(http.StatusOK, "text/html", []byte(constant.SwaggerHtml))
	return
}

// DocsFile 静态文件
func DocsFile(c *gin.Context) {
	url := c.Request.URL.Path
	fileInfo := url[strings.LastIndex(url, "/"):]
	fileEnd := fileInfo[strings.LastIndex(fileInfo, "."):]
	contextType := "application/octet-stream"
	switch fileEnd {
	case ".png":
		contextType = "image/png"
	case ".js":
		contextType = "application/javascript"
	case ".css":
		contextType = "text/css"
	}
	data, err := os.ReadFile("docs" + fileInfo)
	if err != nil {
		c.JSON(http.StatusNotFound, nil)
	} else {
		c.Data(http.StatusOK, contextType, data)
	}
	return
}
