# iris-admin
IRIS（go语言web框架）脚手架项目


### 功能介绍
- [ ] 权限控制
- [x] mongoDB操作
- [x] redis 操作
- [ ] redis 集群连接
- [x] Mysql 操作
- [x] 日志系统
- [x] 定时任务
- [ ] 文件上传
- [ ] csv文件处理
- [x] 配置文件读取
- [x] 令牌桶限流
- [x] 开发调试热加载
- [ ] 日常开发utils集成
- [ ] 邮件发送模块
- [ ] 权限管理
- [ ] 登录校验


### 使用说明
#### 1.项目名配置
在go.mod 中修改第一行中的项目名称，再使用ctrl+shift+r(goland环境下)进行全局替换


### 相关主要依赖库

- Mongo连接：gopkg.in/mgo.v2
- config配置读取 ：github.com/olebedev/config 


### 开发模式热加载
使用：[gowatch](https://github.com/silenceper/gowatch)进行程序热部署

安装：`go get github.com/silenceper/gowatch`
安装完成之后，直接使用`gowatch`命令即可

gowatch -p ./main.go


### 注释插件
Goanno



### 配置prometheus
docker run \
    -p 8797:9090 \
    -v /root/soft/prometheus/prometheus.yml:/etc/prometheus/prometheus.yml \
    prom/prometheus


docker run -d --name grafana  -p 8795:3000 grafana/grafana grafana
120.46.173.6:8795