# 项目简介
`demo`项目展示了如何从零开始搭建一个基于Go和Gin框架的工程化实践，包含项目结构、日志组件、数据库连接、错误码设计、代码生成等内容。



# 功能实现
## 项目结构设计

采用标准的Go项目结构，包含cmd、internal、pkg、log等目录。目录结构如下：
``` bash
.
├── config
├── docs
├── internal
│   ├── app
│   │   ├── controller
│   │   ├── dto
│   │   ├── middleware
│   │   ├── model
│   │   ├── object
│   │   ├── router
│   │   └── service
│   └── pkg
│       ├── context
│       ├── errorCode
│       ├── helper
│       └── test
├── log
├── pkg
│   └── utils
└── sql
```


## 日志组件设计

支持链路追踪功能，确保日志可以完整地记录请求的全链路信息。

## 数据库client设计

支持MySQL和Redis，均支持日志链路追踪。

## 错误码设计

设计统一的错误码管理，并对错误进行封装处理。

## HTTP响应封装
GinRender 组件用于统一处理 HTTP 响应，确保响应格式一致。
- 标准化响应格式：所有HTTP响应均采用统一的格式，包含状态码、消息和数据。
- 错误处理：自动捕捉和处理错误，将错误信息以标准格式返回给客户端。

## 代码生成

安装命令行终端
```bash
go install github.com/morehao/gcli@latest
```
确保配置文件有中有代码生成所需要的配置项`code_gen`，在`项目根目录下`使用`gcli`生成代码，示例如下：
```bash
# 基于表生成整个功能模块
gcli generate -m module
# 生成model代码
gcli generate -m model
# 生成单个接口代码
gcli generate -m api
```

## 接口文档

安装swag工具
```shell
go install github.com/swaggo/swag/cmd/swag@latest
```
生成接口文档
``` shell
chmod +x scripts/swag.sh
scripts/swag.sh demo
```
访问接口文档
访问 `http://localhost:8099/demo/swagger/index.html` 即可查看接口文档。

## 项目部署
构建镜像
``` bash
make APP=demo docker-build
```
运行容器
``` bash
 make APP=demo docker-run
```

## 快速生成新项目
安装`go-cutter`
```shell
go install github.com/morehao/go-cutter@latest
```
在 **当前项目根目录下（即`./`）** 执行命令
```shell
go-cutter -d /goProject/yourAppName
```
执行后，会以当前项目为模板项目，在`/goProject`目录下生成一个名为`yourAppName`的项目。

`go-cutter`是我实现的一个快速生成项目代码的命令行工具，可以基于本项目快速生成一个新的项目，具体使用方法请参考 [https://github.com/morehao/go-cutter](https://github.com/morehao/go-cutter)。


## 后续功能

- httpClient封装（支持日志链路追踪）
- Makefile文件

## 相关组件
相关组件均在[go-tools](https://github.com/morehao/go-tools)包中实现。

