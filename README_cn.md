[English](./README.md) | [简体中文](./README_cn.md)

# 项目简介
`go-gin-web` 是一个从零开始构建的 Go 工程化实践项目，基于[gin](https://github.com/gin-gonic/gin)框架，旨在搭建一个分层相对合理、开发相对规范、可维护性高、方便扩展的 Go 服务端项目。
# 项目特点

- 清晰的项目结构：参考了[project-layout](https://github.com/golang-standards/project-layout)，遵循分层架构思想，目录组织合理，便于团队协作与长期维护。
- 常用组件集成：MySQL、Redis、ES。
- 全链路日志追踪：基于`zap`封装的日志组件`glog`,支持链路 ID 贯穿 MySQL、Redis、ES、Http 调用。
- 代码生成工具：提供`gocli`命令行工具，支持根据配置快速生成标准代码（包括 model、dao、object、dto、code、service、controller、router层代码）。
- `Swagger`文档支持：使用`swaggo`自动生成 API 文档，方便前后端联调与测试。
- Docker 支持： 提供基础的`Dockerfile`，实现容器化部署
- 丰富的`Makefile`工具链：支持`make`命令快速构建、运行、代码生成、接口文档生成、docker 部署等基础操作。
- 逐渐丰富的`golib`库：对常用组件封装，使用更友好。

# 项目结构

参考了[project-layout](https://github.com/golang-standards/project-layout)。当前项目结构如下：
``` bash
.
├── apps
│   ├── demoapp
│   │   ├── cmd
│   │   ├── client
│   │   │   └── httpbingo
│   │   ├── config
│   │   ├── dao
│   │   │   └── daouser
│   │   ├── docs
│   │   ├── internal
│   │   │   ├── controller
│   │   │   │   ├── ctrexample
│   │   │   │   └── ctruser
│   │   │   ├── dto
│   │   │   │   ├── dtoexample
│   │   │   │   └── dtouser
│   │   │   └── service
│   │   │       ├── svcexample
│   │   │       └── svcuser
│   │   ├── middleware
│   │   ├── model
│   │   ├── object
│   │   │   ├── objcommon
│   │   │   └── objuser
│   │   ├── router
│   │   └── scripts
│   └── newapp
├── log
├── output
│   └── build
├── pkg
│   ├── code
│   ├── storages
│   ├── testutil
│   └── utils
└── scripts
    └── sql
```

# 基础功能

## 代码生成

安装命令行终端
```bash
go install github.com/morehao/gocli@latest
```
确保项目应用目录下有代码生成配置文件，示例：`go-gin-web/apps/demoapp/config/code_gen.yaml`。代码生成命令如下：
```bash
# 基于表生成整个功能模块
make codegen MODE=module APP=demoapp
# 生成model代码
make codegen MODE=model APP=demoapp
# 生成单个接口代码
make codegen MODE=api APP=demoapp
```
代码生成详细说明文档见[generate](https://github.com/morehao/gocli?tab=readme-ov-file#generate)。

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
安装`cutter`
```shell
go install github.com/morehao/gocli@latest
```
在 **当前项目根目录下（即`./`）** 执行命令
```shell
gocli cutter -d /goProject/yourAppName
```
执行后，会以当前项目为模板项目，在`/goProject`目录下生成一个名为`yourAppName`的项目。

`go-cutter`是一个快速生成项目代码的命令行工具，可以基于现有项目快速生成一个新的项目，具体使用方法请参考 [cutter](https://github.com/morehao/gocli)。


## 相关组件
相关组件均在[golib](https://github.com/morehao/golib)包中实现。

