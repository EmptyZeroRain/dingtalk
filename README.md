# proper_dingtalk_bot

> 基于 [CatchZeng/dingtalk](https://github.com/CatchZeng/dingtalk/) 修改，支持专有钉钉开放平台域名。

## 特性

- 支持多种消息类型：Text、Markdown、Link、ActionCard、FeedCard
- 支持专有钉钉开放平台域名
- 支持命令行工具和 Go SDK 两种使用方式
- 支持配置文件和环境变量

## 安装

### Go SDK

```bash
go get github.com/EmptyZeroRain/proper_dingtalk_bot
```

### 命令行工具

```bash
# 源码安装
go install github.com/EmptyZeroRain/proper_dingtalk_bot@latest

# 或编译二进制
git clone https://github.com/EmptyZeroRain/proper_dingtalk_bot.git
cd proper_dingtalk_bot
make build
```

## 使用方式

### 命令行工具

#### 全局参数

| 参数 | 简写 | 说明 |
|------|------|------|
| `--access_token` | `-t` | 机器人 access_token |
| `--secret` | `-s` | 加签密钥（可选） |
| `--domain` | `-d` | 钉钉开放平台域名（专有钉钉必填） |
| `--isAtAll` | `-a` | 是否 @所有人 |
| `--atMobiles` | `-m` | @指定手机号（多个用逗号分隔） |
| `--debug` | `-D` | 调试模式，输出请求详情 |

#### 发送 Text 消息

```bash
dingtalk text -d 你的域名 -t your_token -s your_secret -c "Hello World" -m "13800138000"
```

#### 发送 Markdown 消息

```bash
dingtalk markdown -d 你的域名 -t your_token -s your_secret -i "标题" -e "# Hello\n\n**Bold**"
```

#### 发送 Link 消息

```bash
dingtalk link -d 你的域名 -t your_token -i "标题" -e "描述" -p "图片URL" -u "跳转链接"
```

#### 发送 ActionCard 消息

整体跳转：

```bash
dingtalk actionCard -d 你的域名 -t your_token -i "标题" -e "内容" -n "查看详情" -u "https://example.com"
```

独立跳转：

```bash
dingtalk actionCard -d 你的域名 -t your_token -i "标题" -e "内容" -b "按钮1,按钮2" -c "https://a.com,https://b.com"
```

#### 发送 FeedCard 消息

```bash
dingtalk feedCard -d 你的域名 -t your_token \
  -i "标题1,标题2" \
  -p "图片1URL,图片2URL" \
  -u "链接1URL,链接2URL"
```

### Go SDK

```go
package main

import (
    "log"
    "github.com/EmptyZeroRain/proper_dingtalk_bot/pkg/dingtalk"
)

func main() {
    accessToken := "your_access_token"
    secret := "your_secret"
    domain := "你的专有域名" // 公共钉钉传 "oapi.dingtalk.com"

    client := dingtalk.NewClient(accessToken, secret, domain)

    // Text 消息
    msg := dingtalk.NewTextMessage().
        SetContent("Hello World").
        SetAt([]string{"13800138000"}, false)

    _, resp, err := client.Send(msg)
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("errcode: %d, errmsg: %s", resp.ErrCode, resp.ErrMsg)
}
```

更多消息类型：

```go
// Markdown 消息
msg := dingtalk.NewMarkdownMessage().
    SetMarkdown("标题", "# Hello\n\n**Bold**").
    SetAt([]string{}, false)

// Link 消息
msg := dingtalk.NewLinkMessage().
    SetLink("标题", "描述", "图片URL", "跳转链接")

// ActionCard 消息（整体跳转）
msg := dingtalk.NewActionCardMessage().
    SetOverallJump("标题", "内容", "查看详情", "https://example.com", "0", "0")

// ActionCard 消息（独立跳转）
btns := []dingtalk.Btn{
    {Title: "按钮1", ActionURL: "https://a.com"},
    {Title: "按钮2", ActionURL: "https://b.com"},
}
msg := dingtalk.NewActionCardMessage().
    SetIndependentJump("标题", "内容", btns, "0", "0")

// FeedCard 消息
msg := dingtalk.NewFeedCardMessage()
msg.AppendLink("标题1", "https://a.com", "图片1URL")
msg.AppendLink("标题2", "https://b.com", "图片2URL")
```

## 配置

支持配置文件和环境变量，命令行参数优先级最高。

### 配置文件

路径：`~/.dingtalk/config.toml`

```toml
access_token = "your_access_token"
secret = "your_secret"
domain = "你的专有域名"
```

### 环境变量

支持自定义前缀，默认无前缀：

```bash
export DINGTALK_ENV_PREFIX="DINGTALK_"  # 可选，设置前缀
export DINGTALK_ACCESS_TOKEN="your_access_token"
export DINGTALK_SECRET="your_secret"
export DINGTALK_DOMAIN="你的专有域名"
```

## License

[MIT](LICENSE)
