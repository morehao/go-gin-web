# 项目简介
`demo`项目展示了如何从零开始搭建一个基于Go和Gin框架的工程化实践，包含项目结构、日志组件、数据库连接、错误码设计、代码生成等内容。



# 功能实现
## 项目结构设计

参考了[project-layout](https://github.com/golang-standards/project-layout)。当前项目结构如下：
``` bash
.
├── cmd
│   └── demoapp
├── internal
│   └── apps
│       └── demoapp
│           ├── code
│           ├── config
│           ├── controller
│           │   ├── ctrexample
│           │   └── ctruser
│           ├── dao
│           │   └── daouser
│           ├── docs
│           ├── dto
│           │   ├── dtoexample
│           │   └── dtouser
│           ├── middleware
│           ├── model
│           ├── object
│           │   ├── objcommon
│           │   └── objuser
│           ├── router
│           ├── scripts
│           └── service
│               ├── svcexample
│               └── svcuser
├── log
├── output
│   └── build
├── pkg
│   ├── storages
│   ├── test
│   └── utils
└── scripts
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
go install github.com/morehao/gocli@latest
```
确保项目应用目录下有代码生成配置文件，示例：`go-gin-web/internal/apps/demoapp/config/code_gen.yaml`。代码生成命令如下：
```bash
# 基于表生成整个功能模块
make codegen MODE=module APP=demoapp
# 生成model代码
make codegen MODE=model APP=demoapp
# 生成单个接口代码
make codegen MODE=api APP=demoapp
```

## 接口文档

安装swag工具
```shell
go install github.com/swaggo/swag/cmd/swag@latest
```
生成接口文档
``` shell
make swag APP=demoapp
```
访问接口文档
访问 `http://localhost:8099/demoapp/swagger/index.html` 即可查看接口文档。

## 项目部署
构建镜像
``` bash
make docker-build APP=demoapp
```
运行容器
``` bash
 make docker-run APP=demoapp
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
相关组件均在[golib](https://github.com/morehao/golib)包中实现。

