# 一、使用
## 1.1 根据不同的环境启动程序
`
go run ./cmd/web/main.go -env=test
`

这个env就是指环境名称，会根据不同环境读取不同的配置。

其中配置文件存放在：configs/web/config-[env].yml。

## 1.2 Demo数据库配置
修改配置文件中：configs/web/config-[env].yml 中mysql相关的配置。
创建表sql：

~~~ sql
create table user
(
    pk_id        bigint unsigned auto_increment comment '主键' primary key,
    id           varchar(36)                                   not null comment '业务id',
    name         varchar(100)                                  not null comment '姓名',
    sex          tinyint                                       null comment '性别',
    birthday     date                                          null comment '生日',
    is_deleted   tinyint(2)  default 0                         not null comment '是否删除，0:否，1：是',
    gmt_created  datetime                                      not null comment '创建时间',
    gmt_modified datetime                                      not null comment '修改时间',
    gmt_deleted  datetime(3) default '9999-12-31 23:59:59.000' null comment '删除时间',
    index idx_id(id)
) comment '用户';
~~~

# 二、程序说明
## 2.1 目录文件说明
```
- api: 可以存放swagger文档，因为个人不喜欢swagger，所以并不使用

- cmd: 启动程序入口，各组件启动入口
-- web: web便是一个组件，一个程序

- configs: 配置文件信息
-- web: web这个组件相关的配置信息
--- autoload: mysql,日志,缓存，系统配置等等相关的自定义配置参数信息
--- config.go: 读取配置的程序
--- config.yml: 默认设置的配置文件
--- config-local.yml: 本地开发环境配置文件
--- config-test.yml: 测试环境配置文件

- internal: 内部使用的程序
-- web: web这个组件内部使用的程序
--- api: 接口，类似java的controller
---- useapi: user模块相关的接口、模型、路由
----- api: user api相关函数
----- model: user模型，包括“参数，返回信息结构等模型
----- router: 路由信息
--- bll: 中间业务程序，类似java的service
---- usebll: user模块相关的业务代码
----- bll: 业务代码
----- model: 业务代码涉及的相关模型
--- dal: 持久层相关的代码，类似java的dao
---- userdao: user相关数据库操作
----- dao: 数据库操作的函数
----- model: 数据库操作的模型
--- bootstrap: 启动这个web组件的启动程序，初始化日志，初始化数据库，初始化缓存等等操作
--- model: 一些通用的model，比如数据库表都会包含pk_id,id,gmt_created,gmt_modified等
--- pkg: web这个组件内部程序使用的共享工具包，如：环境变量，日志等
---- content: 环境变量信息，比如数据库，缓存等
---- loginfo: 日志信息
--- routers: 配置路由信息，配置跨域，健康检测以及各模块的api接口

- pkg: 可以打包公用的程序/工具
-- binding: 获取binding中各种tag的数据
-- commfunc: 公共的一些操作，比如三元操作等公共的函数
-- httpcode: http请求返回的编码，200成功，500失败等
-- idgenerater: id生成工具
-- json: json相关的操作
-- model: 一些公共的模型，比如分页操作，分页结果，时间操作，序列化反序列化等
-- mysql: mysql数据库初始化的程序，需要的时候可以调用
-- response: 返回response相关的操作，比如返回成功，返回失败
-- sglog: 日志初始化程序，需要的时候可以调用

```

## 2.2 Demo 依赖
本程序比较简单，比较基础，依赖的也比较简单

（1）gin: web组件

（2）sqlx：喜欢类似于mybatis一样手写sql，不喜欢orm

（3）ulid：生成ulid的id

（4）viper：装配配置信息

（5）zapcore：日志信息




