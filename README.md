# iris-admin
IRIS（go语言web框架）脚手架项目


### 热部署
使用：[gowatch](https://github.com/silenceper/gowatch)进行程序热部署

安装：`go get github.com/silenceper/gowatch`
安装完成之后，直接使用`gowatch`命令即可


- [ ] 权限控制
- [ ] mongoDB操作
- [ ] redis 操作
- [ ] Mysql 操作
- [ ] 日志系统
- [ ] REST风格
- [ ] 模板引擎风格
- [ ] 定时任务
- [ ] 文件上传
- [ ] csv文件处理
- [ ] 配置文件读取
- [x] 令牌桶限流

### 使用说明
#### 1.项目名配置
在go.mod 中修改第一行中的项目名称，再使用ctrl+shift+r(goland环境下)进行全局替换


### 相关主要依赖库

- Mongo连接：gopkg.in/mgo.v2
- config配置读取 ：github.com/olebedev/config 