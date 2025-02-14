package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// DingDingMessage 定义钉钉消息结构体
type DingDingMessage struct {
	MsgType  string `json:"msgtype"`
	Markdown struct {
		Title string `json:"title"`
		Text  string `json:"text"`
	} `json:"markdown"`
	At struct {
		AtMobiles []string `json:"atMobiles"`
		IsAtAll   bool     `json:"isAtAll"`
	} `json:"at"`
}

// SendDingDingNotification 发送钉钉通知
func SendDingDingNotification(webhookURL, secret, title, text string, atMobiles []string, isAtAll bool) error {
	// 构造消息体
	message := DingDingMessage{
		MsgType: "markdown",
	}
	message.Markdown.Title = title
	message.Markdown.Text = text
	message.At.AtMobiles = atMobiles
	message.At.IsAtAll = isAtAll

	// 序列化消息体为JSON
	payload, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %v", err)
	}

	// 如果有密钥，则进行加签
	if secret != "" {
		timestamp := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
		signature := generateSignature(timestamp, secret)
		webhookURL += "&timestamp=" + timestamp + "&sign=" + url.QueryEscape(signature)
	}

	// 发送HTTP POST请求
	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(body))
	}

	return nil
}

// generateSignature 生成签名
func generateSignature(timestamp, secret string) string {
	stringToSign := timestamp + "\n" + secret
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(stringToSign))
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return signature
}

func main() {
	// 钉钉机器人的Webhook URL
	webhookURL := "https://oapi.dingtalk.com/robot/send?access_token=your_dingtalk_bot_url"

	// 密钥（如果使用加签方式）
	secret := "SEC7e7bc810e21YOUR_SECRET"

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
	err := SendDingDingNotification(webhookURL, secret, title, text, atMobiles, isAtAll)
	if err != nil {
		fmt.Printf("Error sending notification: %v\n", err)
	} else {
		fmt.Println("Notification sent successfully!")
	}
}
