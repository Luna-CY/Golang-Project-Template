# Golang-Project-Template

## 简介

具备基础代码抽象的通用项目框架，框架的目标是满足Controller、Service、DAO三者基础抽象的情况下，简化项目从0到1构建的成本 ，在降低成本的同时还能够满足项目自身的扩展性需求，如支持多协议、多入口等

## 目录说明

```
Golang-Project-Template
-- cmd                                                  # 所有命令的存放位置，标准go工程目录约定
-- -- main                                              # 主命令
-- -- -- command                                        # 子命令目录
-- config                                               # 配置文件存放目录
-- -- i18n                                              # i18n国际化语言配置文件目录
-- internal                                             # 内部代码，标准go工程目录约定
-- -- build                                             # 自定义编译参数
-- -- configuration                                     # 配置对象定义
-- -- context                                           # 内部上下文定义
-- -- -- contextutil                                    # 上下文工具
-- -- dao                                               # dao实现
-- -- docs                                              # 文档存放目录
-- -- errors                                            # 内部错误定义
-- -- i18n                                              # i18n实现
-- -- interface                                         # 接口定义
-- -- -- dao                                            # dao接口定义
-- -- -- service                                        # service接口定义
-- -- -- transactional                                  # 事务接口定义
-- -- itype                                             # 内部自定义通用类型
-- -- language                                          # 语言定义
-- -- logger                                            # 日志记录器实现
-- -- runtime                                           # 运行时环境定义
-- -- service                                           # service实现
-- -- transactional                                     # 事务实现
-- -- util                                              # 工具集
-- migration                                            # 数据库迁移脚本
-- model                                                # 模型定义
-- server                                               # 协议及入口定义
-- http                                                 # http协议实现
-- -- gateway                                           # http入口定义
-- -- -- web                                            # 入口一
-- -- middleware                                        # http通用中间件定义
-- -- request                                           # http通用请求方法定义
-- -- response                                          # http通用相应方法定义
-- -- router                                            # http通用路由相关方法定义
```

## 使用方法

- `git clone github.com/Luna-CY/Golang-Project-Template`克隆项目到本地
- 更改git仓库为私有仓库
- 全局替换`github.com/Luna-CY/Golang-Project-Template`为私有项目的module名称
- `go run ./cmd/main/main.go`运行项目
