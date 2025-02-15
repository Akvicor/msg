# MSG

消息发送平台

- [x] [SMS](https://github.com/Akvicor/sms)
- [x] WeChat 企业应用消息
- [x] Email SMTP/IMAP
- [x] Telegram Bot

## 配套工具

- 命令行快速发消息工具 [https://github.com/Akvicor/msg-cmd](https://github.com/Akvicor/msg-cmd)
- Go快速调用Mod: [https://github.com/Akvicor/gmsg](https://github.com/Akvicor/gmsg)

## 编译

- Go 22
- Node 20

```bash
# 进入前端目录,编译
cd frontend && make build
# 将编译出来的build文件夹移动到后端目录
cp -r frontend/build backend/cmd/app/server/web/build
# 编译后端
cd backend && make build
```

## 搭建

生成默认配置文件

```bash
# 数据目录为设置为 /data/
./msg example -p "/data/" -c > /data/config.toml
# 根据自身情况修改config.toml

# 初始化数据库
./msg migrate -c /data/config.toml
# 运行
./msg server -c /data/config.toml

# 默认用户名: admin
# 默认密码: password
```

默认情况下使用的是sqlite数据库,如果想要使用postgres请将`config.toml`中的配置项`database.type`修改为`postgres`, 然后重启应用程序或者docker

- 程序中默认`bot.maid`支持`sms`,`mail`,`telegram`,`wechat`, 所需配置需要填写到配置文件中才能生效
- 程序中默认`bot.reminder`支持`wechat`, 这是一个企业微信应用的ToDo功能, 支持在微信中设定定时发送的信息

## 自定义功能

本项目支持的功能非常有限, 默认情况下仅支持从miad发送消息, 以及reminder的定时提醒. 如果需要配置收到信息后如何处理, 或绑定多个通知渠道, 请自行在代码中编写对应bot的功能或添加新bot


