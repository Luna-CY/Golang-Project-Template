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

## 错误定义

Golang本身提供的`error`接口，`error`接口太过简略，无法满足系统层面以及业务层面对于错误信息描述的需求，因此此项目模板内错误不直接使用`error`接口

模板内将发生的所有错误归于两个类别：系统内部错误、业务边界错误。业务遍及错误对象可以包含一个或多个系统内部错误对象，它们的区别在使用场景

### 内部错误定义

系统内部错误用于定义发生在系统内部，且只在系统内部流转的错误对象，该对象的行为由`errors.Error`接口定义

系统内部错误具体定义了发生错误的详尽位置（必须硬编码在发生错误的地方，以便于全局搜索）以及错误发生的原因

通常情况下系统内部错误不适合直接返回给任何外部API（HTTP/RPC等），仅应该用于内部调用链之间的传递

#### 错误编码硬编码格式说明

在多人并行开发时，集中化的错误管理容易引入额外的协同复杂度（比如代码定义冲突，文件编辑冲突等），因此硬编码错误应该采用一种非集中化、具有约定格式的标准编码方式

此模板中，按照如下格式约定错误编码

`{全部路径首字母}_{最后一级路径的尾2尾字母}[.{结构体首字母}_{结构体尾2位字母}.]{方法首字母}_{方法尾2为字母}.{首次定义错误发生时的行号}`

**其中结构体部分是可选地，如果文件中直接定义了工具方法，则没有结构体名称部分**

#### 示例

- 例1：以路径`server/http/gateway/web/handler/example`目录下`impl_create.go`文件为例，该文件内定义了`Example`处理器的`Create`接口

该接口内产生的错误，其编码为: `{SHGWHE}_{LE}.{E}_{LE}.{C}_{TE}.{初始行号}`

#### 特别说明

代码是动态的，错误定义时的初始行号可能会随着代码变动与其所在行号不符，通常情况下不需要随时矫正错误所在的行号，错误编码的目的是为了能够方便的定位某一个错误发生的位置，因此其行号正确与否并不重要

但是存在一种情况：当旧的错误向下移动，原始行号的位置定了一个新的错误，此时会发生两个错误拥有相同编码的冲突，为了解决此问题，定义新的错误之后需要使用错误编码在当前文件或全局范围内检索，如果编码不唯一，应调整旧的错误行号

### 业务边界错误

业务边界错误包含若干个系统内部错误和具体业务、边界相匹配的国际化输出信息，该对象行为由`errors.I18nError`接口定义

通常业务边界（HTTP/RPC等API）需要应对较为复杂的格式化输出需求，比如根据不同语言输出不同内容，根据不同业务错误类型返回特定错误码，这种输出复杂性不应该传递给系统内部的其它部分

因此所有的边界逻辑在调用系统内部方法获得到任何系统内部错误之后，都应该将其转换为业务边界错误，将最后处理完成的错误信息返回给外部调用方

## 使用方法

- `git clone github.com/Luna-CY/Golang-Project-Template`克隆项目到本地
- 更改git仓库为私有仓库
- 全局替换`github.com/Luna-CY/Golang-Project-Template`为私有项目的module名称
- `go run ./cmd/main/main.go`运行项目
