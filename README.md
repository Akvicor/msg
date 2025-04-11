# MSG

消息发送平台

- [x] [SMS](https://github.com/Akvicor/sms)
- [x] WeChat 企业应用消息
- [x] Email SMTP/IMAP
- [x] Telegram Bot

[博客链接 https://blog.akvicor.com/posts/project/msg/](https://blog.akvicor.com/posts/project/msg/)

## 配套工具

- 命令行快速发消息工具 [https://github.com/Akvicor/msg-cmd](https://github.com/Akvicor/msg-cmd)
- Go快速调用Mod: [https://github.com/Akvicor/gmsg](https://github.com/Akvicor/gmsg)

## Schedule 定时发送

### TypeSecond

每秒产生一次通知

### TypeMinute

每分钟产生一次通知

- `second`: `[0,59]` 正数指定秒, `[-1,-60]` 倒数指定秒

### TypeHour

每小时产生一次通知

- `second`: `[0,59]` 正数指定秒, `[-1,-60]` 倒数指定秒
- `minute`: `[0,59]` 正数指定分, `[-1,-60]` 倒数指定分

### TypeDaily

每天产生一次通知

- `second`: `[0,59]` 正数指定秒, `[-1,-60]` 倒数指定秒
- `minute`: `[0,59]` 正数指定分, `[-1,-60]` 倒数指定分
- `hour`:   `[0,23]` 正数指定时, `[-1,-24]` 倒数指定时

### TypeWeekly

每周产生一次通知

- `second`: `[0,59]` 正数指定秒, `[-1,-60]` 倒数指定秒
- `minute`: `[0,59]` 正数指定分, `[-1,-60]` 倒数指定分
- `hour`:   `[0,23]` 正数指定时, `[-1,-24]` 倒数指定时
- `day`:    `[0,6]` 正数周内指定天, `[-1,-7]` 倒数周内指定天, 周一为第0天

### TypeMonthly

每月产生一次通知

- `second`: `[0,59]` 正数指定秒, `[-1,-60]` 倒数指定秒
- `minute`: `[0,59]` 正数指定分, `[-1,-60]` 倒数指定分
- `hour`:   `[0,23]` 正数指定时, `[-1,-24]` 倒数指定时
- `day`:   `[1,max_day_of_month]` 正数月内指定天, `[-1,-max_day_of_month]` 倒数月内指定天, 每月1号为第1天

### TypeQuarterly

每季产生一次通知

- `second`: `[0,59]` 正数指定秒, `[-1,-60]` 倒数指定秒
- `minute`: `[0,59]` 正数指定分, `[-1,-60]` 倒数指定分
- `hour`:   `[0,23]` 正数指定时, `[-1,-24]` 倒数指定时
- `day`:    `[1,max_day_of_month]` 正数月内指定天, `[-1,-max_day_of_month]` 倒数月内指定天, 每月1号为第1天
- `month`:  `[0,2]` 正数季度内指定月, `[-1,-3]` 倒数季度内指定月, 每季度首个月为第0月

### TypeYearly

每年产生一次通知

- `second`: `[0,59]` 正数指定秒, `[-1,-60]` 倒数指定秒
- `minute`: `[0,59]` 正数指定分, `[-1,-60]` 倒数指定分
- `hour`:   `[0,23]` 正数指定时, `[-1,-24]` 倒数指定时
- `day`:   `[1,max_day_of_month]` 正数月内指定天, `[-1,-max_day_of_month]` 倒数月内指定天, 每月1号为第1天
- `month`:  `[1,12]` 正数指定月, `[-1,-12]` 倒数指定月, 每年首个月为第1月

### TypeSecondInterval

间隔秒产生一次通知

- `second`: 间隔指定秒

### TypeMinuteInterval

间隔分钟产生一次通知

- `second`: `[0,59]` 正数指定秒, `[-1,-60]` 倒数指定秒
- `minute`: 间隔指定分


### TypeHourInterval

间隔小时产生一次通知

- `second`: `[0,59]` 正数指定秒, `[-1,-60]` 倒数指定秒
- `minute`: `[0,59]` 正数指定分, `[-1,-60]` 倒数指定分
- `hour`: 间隔指定时

### TypeDailyInterval

间隔天产生一次通知

- `second`: `[0,59]` 正数指定秒, `[-1,-60]` 倒数指定秒
- `minute`: `[0,59]` 正数指定分, `[-1,-60]` 倒数指定分
- `hour`:   `[0,23]` 正数指定时, `[-1,-24]` 倒数指定时
- `day`: 间隔指定天

### TypeWeeklyInterval

间隔周产生一次通知

- `second`: `[0,59]` 正数指定秒, `[-1,-60]` 倒数指定秒
- `minute`: `[0,59]` 正数指定分, `[-1,-60]` 倒数指定分
- `hour`:   `[0,23]` 正数指定时, `[-1,-24]` 倒数指定时
- `day`:    `[0,6]` 正数周内指定天, `[-1,-7]` 倒数周内指定天, 周一为第0天
- `week`: 间隔指定周

### TypeMonthlyInterval

间隔月产生一次通知

- `second`: `[0,59]` 正数指定秒, `[-1,-60]` 倒数指定秒
- `minute`: `[0,59]` 正数指定分, `[-1,-60]` 倒数指定分
- `hour`:   `[0,23]` 正数指定时, `[-1,-24]` 倒数指定时
- `day`:   `[1,max_day_of_month]` 正数月内指定天, `[-1,-max_day_of_month]` 倒数月内指定天, 每月1号为第1天
- `month`: 间隔指定月

### TypeQuarterlyInterval

间隔季产生一次通知

- `second`: `[0,59]` 正数指定秒, `[-1,-60]` 倒数指定秒
- `minute`: `[0,59]` 正数指定分, `[-1,-60]` 倒数指定分
- `hour`:   `[0,23]` 正数指定时, `[-1,-24]` 倒数指定时
- `day`:    `[1,max_day_of_month]` 正数月内指定天, `[-1,-max_day_of_month]` 倒数月内指定天, 每月1号为第1天
- `month`:  `[0,2]` 正数季度内指定月, `[-1,-3]` 倒数季度内指定月, 每季度首个月为第0月
- `quarter`: 间隔指定季

### TypeYearlyInterval

间隔年产生一次通知

- `second`: `[0,59]` 正数指定秒, `[-1,-60]` 倒数指定秒
- `minute`: `[0,59]` 正数指定分, `[-1,-60]` 倒数指定分
- `hour`:   `[0,23]` 正数指定时, `[-1,-24]` 倒数指定时
- `day`:   `[1,max_day_of_month]` 正数月内指定天, `[-1,-max_day_of_month]` 倒数月内指定天, 每月1号为第1天
- `month`:  `[1,12]` 正数指定月, `[-1,-12]` 倒数指定月, 每年首个月为第1月
- `year`: 间隔指定年

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


