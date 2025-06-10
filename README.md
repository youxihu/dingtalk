# dingtalk

一个简单的 Go 包，支持通过钉钉自定义机器人发送消息通知。

## 安装

```bash
go get github.com/youxihu/dingtalk/dingtalk
```

### SendDingDingNotification 参数：
- webhookURL (string): 您的钉钉机器人Webhook URL。这是必填项。
- secret (string): (可选) 加签消息的密钥。如果将其留空 ("")，消息将以不加签的方式发送。
- title (string): 通知标题。
- text (string): 通知的主要内容。此字段支持Markdown格式。
- atMobiles ([]string): 要艾特的手机号码切片。只有这些指定的手机号会被艾特。
- isAtAll (bool): 设置为 true 可以艾特群里的所有人，
- 如果为 false，则只会艾特 atMobiles 中指定的手机号（如果存在）。
