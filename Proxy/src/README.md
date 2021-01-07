
# Proxy 反向代理服务器



# 语言及核心框架版本
1. Golang >= 1.3.x
2. LollipopGo == 2.8.x

# 代码组织及分层
![代码组织](imgs/tower_01.jpg)

说明：

1. 所有消息入口，流量是所有服务器的总和，内网与外网隔离，降低内网服务器被攻击的风险。
2. 可以理解为注册中心，针对于客户端及内网服务器是服务器角色。
3. 不涉及到数据存储，仅仅是消息转发**不跨服务调用其它服务存储层**

# Client API 接口通用协议
## Url规范
全部接口采用`Websocket`请求，URL格式:

> `ws://ProxyServer.ip/BaBaLiuLiu?parameters'

`Content-Type`为`application/json`。ws首先要通过握手协议，onopen成功后，可以发消息onsend,消息监听onmessage。

## `请求`、返回结构

### 【消息注册】Request定义：

```go
协议号： 
C2Proxy_ConnDataProto == 7  客户端连接协议

数据结构：
type C2Proxy_ConnData struct {
	Protocol  int
	Protocol2 int
	OpenID    string // 客户端第一次发空
}
```

### 【消息注册】Response定义：

```go
协议号：
Proxy2C_ConnDataProto == 8  

数据结构：
type Proxy2C_ConnData struct {
	Protocol  int
	Protocol2 int
	OpenID    string
}
```

------

### 【发送数据】Request定义：

```go
协议号：C2Proxy_SendDataProto == 1  客户端发送协议

数据结构：
type C2Proxy_SendData struct {
	Protocol  int
	Protocol2 int
	ServerID  string
	Data      interface{}
}
```

### 【发送数据】Response定义：

```go
协议号：Proxy2C_SendDataProto == 2

数据结构：
type Proxy2C_SendData struct {
	Protocol  int
	Protocol2 int
	ServerID  string
	Data      interface{}
}

```

------

### 数据加密

request和response内容全部需要AES加密传输，AES加密使用ECB模式（ECB模式不需要IV），填充算法使用PKCS7（Java称之为PKCS5Padding）即，如使用Java开发，加密算法使用"AES/ECB/PKCS5Padding"。

### `aes_key`和`sign_key`交换
通过Aes.ProxyYun.Com交换加密和签名key，具体密钥分配原则由服务端决定，比如可以根据chn，或者user_id来分配。每次游戏启动时刷新一次key。

## 单身依赖
各业务service层不要调用其它业务的对象。