## go入门学习之简易聊天室加弹幕

注：本项目基于 [chitchat](https://github.com/nonfu/chitchat)项目做的二次开发(扩展一点小功能)，适合初学go的小白上手，业务简单，逻辑简单
### 环境
**go 1.13**

### 演示地址
+ http://shop.aabbccm.com/
+   用户名test 密码123456

### 安装
**拉取项目**
```php
git clone https://github.com/wuqinqiang/go-barrage.git
```
**修改运行环境**
config.json 文件 ps:这种配置文件不应该提交到版本库中
```json
{
  "App": {  //应用配置
    "Address": "0.0.0.0:8080", //服务地址
    "Static": "public",  //静态文件存放目录
    "Log": "logs"     //日志目录
  },
  "Db": {          //数据库配置
    "Driver": "mysql",
    "Address": "192.168.10.10:3306",  //地址
    "Database": "chitchat",  //数据库名
    "User": "homestead",     //账户
    "Password": "secret"     //密码
  }
}
```

### 运行
**在项目根目录下运行**
```go
go run main.go
```

**访问页面，注册账户登录之后，来到Room**

​    <img src="https://github.com/wuqinqiang/go-barrage/master/chat.png">






