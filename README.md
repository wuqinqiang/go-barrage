## go入门学习之简易聊天室加弹幕

注：本项目基于 [chitchat](https://github.com/nonfu/chitchat)项目做的扩展开发(扩展功能)，适合初学go的小白上手，业务简单，逻辑简单
### 环境
**go环境自行安装**

**项目未使用框架**

### 演示地址

+ http://room.aabbccm.com/
可以自己注册 或者使用下面
+   用户名curry@qq.com 密码123456
+   用户名test(1到9都行)@qq.com 密码123456

### 增加配置
根目录下增加 config.json 文件 
```json
{
  "App": {  //应用配置
    "Address": "0.0.0.0:8080", //服务地址 自行设置
    "Static": "public",  //静态文件存放目录 
    "Log": "logs",     //日志目录
    "StaticPath": "http://127.0.0.1/static/"  //本地上传图片访问地址 自行设置
  },
  "Db": {          //数据库配置
    "Driver": "mysql",
    "Address": "192.168.10.10:3306",  //地址 自行设置
    "Database": "chitchat",  //数据库名 自行设置
    "User": "homestead",     //账户 自行设置
    "Password": "secret"     //密码  自行设置
  },
 "OSS": {
//true表示默认图片上传至本地 false表示上传至阿里oss 自行配置Key
    "IsSaveLocal": true, 
    "AccessKeyID": "xxx",
    "AccessKeySecret": "xxxx",
    "Bucket": "xxx",
    "BucketUrl": "xxx"
  }
}
```

**数据库**

**数据库文件放在根目录下chitchat_2020-04-28.sql**
**对于数据库的操作未使用orm包**

### 运行
**在项目根目录下运行**
```go
go run main.go
```

**或者根目录下已将应用程序打包成二进制可执行文件,可以在根目录下直接执行**
```go
./chitchat
```


### demo截图

**访问页面，注册账户登录之后，来到弹幕室**

​    <img src="https://github.com/wuqinqiang/go-barrage/blob/master/chat.png">

**聊天室(单聊)**
​    <img src="https://github.com/wuqinqiang/go-barrage/blob/master/room.png">

**加好友**
​    <img src="https://github.com/wuqinqiang/go-barrage/blob/master/user.png">

**处理申请**
​    <img src="https://github.com/wuqinqiang/go-barrage/blob/master/handle.png">


### 应用部署
**只是很简单的用了nginx 做了反向代理,通过 Supervisor 维护应用守护进程,后续学习docker的时候会将此项目用docker部署 ps.学啥用啥**


### 主要功能模块

- [x] 弹幕
- [x] 单聊
- [x] 加好友,审核
- [x] 用户列表
- [ ] 群组
    - [ ] 创建群
    - [ ] 加群
    - [ ] 踢人
    - [ ] 禁言
- [ ] 消息类型
    - [x] 文本
    - [x] 文件文件暂时存放在本地
    - [ ] 语音
    - [ ] 视频(先写着把)
....

**其他可以自行看代码,有问题提交issue,后面代码也会慢慢优化**








