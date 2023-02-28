# rj97807_work_serve

## 预期实现

账号注册 完成  
账号登录 完成  
jwt 完成  
OSCHINA的爬虫内容获取 完成
根据权限显示不同菜单--需要验证token 完成  
七牛云分片上传新文件--需要验证token 完成方式1
> ```
> 方式1:存入本地 通过sdk 一步到位，如果有错误，将错误返回给前端
> 方式2:存入本地 通过Object接口，启动另一个业务。分布完成
> 暂时采取方式1 先开发完再迭代
>```
七牛云已存在文件秒传--需要验证token 完成

## 二阶段

jwt过期刷新 完成  
通过rpc调取验证码进行注册 完成  
校验验证码 完成  
添加日志 完成-只有用户登录写了一个Info  
websocket广播 查到有文章标识服务端只能发送ping 无法接收。另js无法发送ping帧，所以解决方案分位前端不断发送 后端直接返回pong内容。或者通过nginx设置无信息过期timeout

# Tips:

其他环境运行注意：

- docker-compose中需要在容器根路径创建一个files_template文件夹(可以直接通过数据映射，如果方便查看文件的话)
- 需要在utils自定义`privateConfig.go`文件夹  
  EmailPwd 邮箱密码
- rpc/api文件下缺少etc/x.yaml文件
- 需要手动配置容器中的yam配置

# 后续

- websocket
- resp模版修改[参考](https://go-zero.dev/cn/docs/advance/template/)

# 打包命令

- `CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o serve api.go`