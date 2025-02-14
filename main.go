package main

import (
	"fmt"
	"github.com/youxihu/dingtalk/dingtalk"
	"time"
)

func main() {
	// 钉钉机器人的Webhook URL
	webhookURL := "https://oapi.dingtalk.com/robot/send?access_token=606800e0504d138c8ee6b764e57ccb1efca88a1d4d61dcac23b6413d83a0b42a"

	// 密钥（如果使用加签方式）
	secret := "SEC7e7bc8109e2113de0fc7f316f86d0bf1e70a8f523b3fe90d853d584ad311b2f6"

	// 消息标题和内容
	title := "测试通知"
	text := " 测试通知\n" +
		">  这是一条测试消息\n" +
		"> - 时间：" + time.Now().Format("2006-01-02 15:04:05") + "\n" +
		"> - 内容：测试内容"

	// 艾特设置
	atMobiles := []string{"19******502"} // 艾特指定手机号
	isAtAll := false                     // 是否艾特所有人 //true or false

	// 发送通知
	err := dingtalk.SendDingDingNotification(webhookURL, secret, title, text, atMobiles, isAtAll)
	if err != nil {
		fmt.Printf("Error sending notification: %v\n", err)
	} else {
		fmt.Println("Notification sent successfully!")
	}
}
