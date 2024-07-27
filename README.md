# 项目简介
`go-gin-web`项目展示了如何从零开始搭建一个基于Go和Gin框架的工程化实践，包含项目结构、日志组件、数据库连接、错误码设计、代码生成等内容。



# 功能实现
## 项目结构设计

采用标准的Go项目结构，包含cmd、internal、pkg、log等目录。目录结构如下：
``` bash
go-gin-web
├── cmd
│   ├── app
│   └── genCode
├── internal
│   ├── app
│   ├── genCode
│   ├── pkg
│   └── resource
├── log
├── pkg
├── .gitignore
├── Dockerfile
├── go.mod
├── LICENSE
└── README.md
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

支持按模块生成、只生成Model层代码、按API接口生成等。

开发者可以通过调整配置文件来生成代码，生成后各层级代码自动补全，只需关心业务代码。 

## 后续功能

- httpClient封装（支持日志链路追踪）
- Makefile文件
- Dockerfile文件
- 如何基于本项目快速生成一个新的项目

## 项目部署
### 镜像编译
``` bash
docker buildx build --progress=plain -t go-gin-web .
```

## 相关组件
相关组件均在[go-tools](https://github.com/morehao/go-tools)包中实现。

