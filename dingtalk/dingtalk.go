package dingtalk

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

