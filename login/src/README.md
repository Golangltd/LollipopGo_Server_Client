# Login 登录服务器

# 语言及核心框架版本

1. golang >= 1.3.x
2. LollipopGo == 2.8.x

# 代码组织及分层

![img](file:///D:/GONGSI/tw/tw/login/src/imgs/login.png?stamp=0)

说明：

1. 接口（API）层不要写得太重，不要做太多业务逻辑处理。主要责任：接口请求参数检验，权限控制，组合下层核心服务接口（RPC或者直接进程内方法调用），返回约定数据结构。
2. 中间的核心业务层，按业务模块组织分服，前期一般也是单进程内运行，可相互直接调用。但是从编码原则上，应该支持后面微服务架构升级，独立部署，RPC调用。
3. 最底层是持久化数据存储或者高速缓存。每个核心业务服务只能操作相应的存储层数据，不要**跨服务调用其它服务存储层**

# Client API 接口通用协议

## URL规范

全部接口采用`GET`请求，URL格式:

> `http://loginserver.ip/BaBaLiuLiu/route?parameters'

`Content-Type`为`application/json`。HTTP状态码统一返回200，具体业务异常码定义在Response结构中定义，当HTTP返回非200时，则表示当前接口服务不可用。

## 请求、返回结构

### Request定义：

```protobuf
http://loginserver.ip/BaBaLiuLiu/client/login?AccountName='001'&AccountPw=‘123456’
```

### Response定义：

```go
type LoginData struct {
	Token  string                           // 检验码
	Url    string                           // 最近登录的服务器
	ServerList []*twlib_user.ServerList     // 服务器列表
	AesKey string                           // 消息加密key，随版本更换
}
```

### 数据加密

request和response内容全部需要AES加密传输，AES加密使用ECB模式（ECB模式不需要IV），填充算法使用PKCS7（Java称之为PKCS5Padding）即，如使用Java开发，加密算法使用"AES/ECB/PKCS5Padding"。

### `aes_key`和`sign_key`交换

通过Aes.ProxyYun.Com交换加密和签名key，具体密钥分配原则由服务端决定，比如可以根据chn，或者user_id来分配。每次游戏启动时刷新一次key。

# 代码规范

## 时区配置及时间使用

LollipopGo2.8.x配置 验证登录：

```go
http://IP:8867/BaBaLiuLiu/client/login?AccountName=001&AccountPw=123456
```

## 单身依赖

各业务service层不要调用其它业务的对象。