# 环境部署

## 目标

- 安装go运行环境: 能运行go程序[hellow world](../HelloWorld/README.md)
- 配置go开发环境: 能方便的书写代码,调试go程序

## 步骤

### win10
https://golang.google.cn/dl/

直接下载安装

**注意配置环境变量**


### mac

同上

## linux

## 结果

## 总结

## 其他

### go 运行程序的流程
- go build xxx.go
- ./xxx

或者直接运行
- go run xxx.go

### go install 慢
使用代理  
[goproxy](https://github.com/goproxy/goproxy.cn/blob/master/README.zh-CN.md)
```shell
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
```


### go项目模块安装步骤
应该先保证
>GO111MODULE=on  
 GOPROXY=https://goproxy.cn,direct


- 进入到项目根目录
- go mod init your_project_name         # 这一步是初始化项目
- go mod tidy                           # 这一步是检测依赖
- go mod download                       # 这一步是下载依赖 

### go的版本
- Go团队每年发布两次大版本,一般在二月和8月
- Go团队承诺对最新的两个Go稳定版提供支持

建议使用最新版