
# 道聚城自动签到工具----GitHub Actions版

### 功能：
- [x] 集成Github Actions
- [x] 多账号配置
- [x] Cookies信息加密保存
- [x] 协程并发执行
- [x] 自动签到、兑换(NBA2KOL2礼包)、完成每日任务
- [ ] 微信消息推送日志

### 使用方法：
- 先Fork本项目，再去`Setting`->`Secrets and variables`->`Actions`中添加`Secret`,Key为`COOKIE`，Value的格式为如下
```JSON
[
    {
        "cookie": "",
        "ptk": "",
        "openId": ""
    },
    {
        "cookie": "",
        "ptk": "",
        "openId": ""
    },
    {
        "cookie": "",
        "ptk": "",
        "openId": ""
    }
]
```

