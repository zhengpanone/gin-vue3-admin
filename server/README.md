## 项目结构

```js
├── api # 接口
├── config # 配置
├── core # 核心代码
├── global # 全局变量和常量
├── initialize # 初始化相关
├── logs # 日志目录
├── main.go # 启动文件
├── model # 实体
├── router # 路由
    └── middleware #中间件
├── test # 单元测试目录
└── utils # 工具包
```

### 初始化项目

```shell
mkdir gin-vue3-admin && cd gin-vue3-admin && go mod init gin-vue3-admin
```

### air运行项目

```shell
go install github.com/cosmtrek/air@latest
cd server && air
```

### gin swagger

```shell
go install github.com/swaggo/swag/cmd/swag@latest
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files
go get -u github.com/alecthomas/template
```

项目根目录下 命令行 执行命令 swag init

 http://localhost:8099/swagger/index.html 

https://www.cnblogs.com/you-men/p/14054348.html
