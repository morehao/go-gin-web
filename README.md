[toc]

# 训练营代码地址
https://github.com/cncamp

# 作业1

## 内容

内容：编写一个 HTTP 服务器，大家视个人不同情况决定完成到哪个环节，但尽量把1都做完

1.接收客户端 request，并将 request 中带的 header 写入 response header
2.读取当前系统的环境变量中的 VERSION 配置，并写入 response header
3.Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
4.当访问 localhost/health 时，应返回200
提交链接🔗：https://jinshuju.net/f/PlZ3xg
补交截止时间：10月10日晚23:59前


## go项目初始化命令

```shell
go mod init go-web
go mod tidy
```

## 作业示范
https://github.com/Julian-Chu/cncamp/blob/main/module2_homework/main.go#L18

## 自己作业的参考
参考文章：[深入剖析Go Web服务器实现原理](https://studygolang.com/articles/25849)

# 作业2

## 内容
- 构建本地镜像。
- 编写 Dockerfile 将练习 2.2 编写的 httpserver 容器化（请思考有哪些最佳实践可以引入到 Dockerfile 中来）。
- 将镜像推送至 Docker 官方镜像仓库。
- 通过 Docker 命令本地启动 httpserver。
- 通过 nsenter 进入容器查看 IP 配置。

作业需编写并提交 Dockerfile 及源代码。
提交链接：https://jinshuju.net/f/rxeJhn
截止日期：10月17日晚23:59之前
提示💡：
1、自行选择做作业的地址，只要提交的链接能让助教老师打开即可
2、自己所在的助教答疑群是几组，提交作业就选几组

## 操作
### 上传docker镜像
```shell
docker build -t go-web-docker .
docker tag go-web-docker:latest morehao/go-web-docker:latest
docker push morehao/go-web-docker
```
### 运行docker镜像
```shell
# 运行docker镜像
docker run -it -p 9090:9090 -d morehao/go-web-docker:latest /bin/bash
# 确认镜像是否成功运行
curl 127.0.0.1:9090/health
```

### nsenter进入docker容器的网络命名空间
```shell
docker inspect -f {{.State.Pid}} containerId
sudo nsenter -n -t{pid}
> ip addr
```
### docker exec进入容器
```shell
docker exec -it containerId sh
# pwd看到项目目录/src/go-web
pwd
```

# kubectl


## simple pod demo
### run nginx as webserver
```
$ kubectl run --image=nginx nginx
// or
$ kubectl run --image=nginx nginx --restart='Always'
```
### show running pod
```
$ kubectl get po --show-labels -owide -w
```
### expose svc
```
$ touch nginx-deploy.yaml
$ vim nginx-deploy.yaml
$ kubectl apply -f nginx-deploy.yaml
$ kubectl describe deployment
$ kubectl expose deployment nginx-deployment --port=80 --target-port=80 --type=NodePort
```
nginx-deploy.yaml
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
spec:
  replicas: 1
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
        - name: nginx
          image: nginx
```
### check svc detail
```
$ kubectl get svc
```
### access service
```
$ curl 192.168.34.2:<nodeport>
```
