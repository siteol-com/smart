# smart

## 介绍

轻量易用的授权基座smart，提供一个RBAC授权模型的基础开箱即用中台管理服务。

支持响应码自定义国际化、高粒度权限分配、接口路由配置和实时生效，同时支持部门数据权限横向访问限制。

本系统的全部接口采用【POST】【application/json】方式传输数据。

除开放接口以外的其他接口均需要通过【ApiKeyAuth:请求头[Token]】完成鉴权。

系统技术栈：Golang、VueNext、MySQL、Redis、Gin、ArcoDesign

## 当前说明

工程尚在构建阶段，仓库未提供数据库初始化脚本。

请耐心等待1.0版本完成。

## 启动准备

### 依赖准备

1. golang 1.20
2. 工程目录执行`go mod tidy`更新依赖
3. 配置文件位置`config`目录下，配置注释`comment.js`

### 环境变量

应用启动需要添加以下环境变量，IDE（如：goLand）可临时添加。

生产部署，可放在启动脚本中，参考`restart.sh`

如：`ENV=test;NODE=APP01`

- ENV=test/prod
- NODE=APP01/APP0X

# SwaggerAPI文档

轻度依赖集成，通过静态HTML加载yaml文件进行打开接口文档。

## 生成说明

- 首次集成

go install github.com/swaggo/swag/cmd/swag@latest

- 格式化注释代码

swag fmt

- 初始化swagger.yaml文件

swag init

- 如果像本工程一样依赖了api.md文件描述项目

swag init --md .

- 删除多余的生成

rm .\docs\docs.go

删除docs.go是因为本项目中并未采用下述依赖来集成。

而是通过HTML模板+JS+YAML引入集成，移除docs.go用于避免依赖编译报错。

- github.com/swaggo/swag
- github.com/swaggo/gin-swagger
- github.com/swaggo/files

源帮助页：[https://github.com/swaggo/swag/blob/master/README_zh-CN.md](https://github.com/swaggo/swag/blob/master/README_zh-CN.md)

## API文档地址

该接口文档提供Swagger[支持调试]和ReDoc[阅读增强]两个版本。

[Swagger[支持调试]：http://localhost:8000/docs/swagger/index.html](http://localhost:8000/docs/swagger/index.html) 

[ReDoc[阅读增强]：http://localhost:8000/docs/redoc/index.html](http://localhost:8000/docs/redoc/index.html)

# 设计思想

## RBAC模型

账号 - 角色（多） - 权限（多） - 接口路由（多）。

## 数据权限

部门数据权限体系，上级可以看下级，特殊部门允许查看全体数据。

## 动态系统配置

系统配置入库，支持热配置生效。

## 响应码国际化

未配置的响应码翻译以默认成功和默认失败响应，支持响应码国际化翻译，支持变量注入。

## 动态路由配置

未配置的路由禁止访问，支持热处理，接口加入/移除、授权配置、日志记录、报文入库、脱敏加密等。

