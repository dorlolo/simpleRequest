# simpleRequest 
## 1. 说明
[simpleRequest](www.github.com/dorlolo/simpleRequest) 是基于golang原生http库的封装，适合用来对第三方接口进行快速地对接和开发。
它具备以下特点：
- 相对于其它请求库，更易于理解和使用。极大减少了开发过程中的代码量和资料查询时间。
- 适合对接一些未遵循restful规范的接口。

## 2. 如何使用?

### 2.1 模块导入

```go
import  "github.com/dorlolo/simpleRequest"
```

### 2.2 实例化

```go
var r = simpleRequest.NewRequest()
```

### 2.3 添加请求头

#### 2.3.1 单个赋值
```go
r.Headers().Set("token", "d+jfdji*D%1=")
r.Headers().Set("Content-Type", "application/json")
```

#### 2.3.2 map赋值
```go
mapHeaders:= map[string]string{
    "token":      "d+jfdji*D%1=",
    "Content-Type": "application/json",
}
r.Headers().Sets(mapHeaders)
```

#### 2.3.3 链式赋值
```go
r.Headers().Set("token", "d+jfdji*D%1=").Set("Content-Type", "application/json")
```

#### 2.3.4 添加多值
对单个key添加多只不要使用`.set`,因为原先的值会被覆盖
```go
r.Headers().Set("Accept", "text/html")
r.Headers().Add("Accept","application/xhtml+xml")
r.Headers().Add("Accept","application/xml;q=0.8")
r.Headers().Add("Accept","image/webp")
r.Headers().Add("Accept","*/*;q=0.8")
```

#### 2.3.4 使用预设的key
```go
r.Headers().SetConentType("application/json")
//r.Headers().Set("Content-Type", "application/json")

r.Headers().SetUserAgent("Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/535.1 (KHTML, like Gecko) Chrome/14.0.835.163 Safari/535.1"")
//r.Headers().Set("User-Agent",Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/535.1 (KHTML, like Gecko) Chrome/14.0.835.163 Safari/535.1")

r.Headers().SetConentEncoding("gzip, deflate, br")
//r.Headers().Set("Content-Encoding", "gzip, deflate, br")
```

#### 2.3.5 使用预设的key-value
```go
//随机user-agent
r.Headers().SetRandomUerAgent()
//r.Headers().Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/535.1 (KHTML, like Gecko) Chrome/14.0.835.163 Safari/535.1")

r.Headers().ConentType_json()
//r.Headers().Set("Content-Type", "application/json")

r.Headers()ConentType_formData()
//r.Headers().Set("Content-Type","multipart/form-data")

r.Headers()ConentType_formUrlencoded()
//r.Headers().Set("Content-Type","application/x-www-form-urlencoded")

r.Headers()ConentType_textPlain()
//r.Headers().Set("Content-Type","text/plain; charset=utf-8")
```

### 2.4 添加queryParams
#### 2.4.1 单个赋值
```go
r.QueryParams().Set("user", "dorlolo")
```
#### 2.4.2 map赋值
不会覆盖上面之前填充过的参数
```go
pamarsBulid := make(map[string]interface{})
pamarsBulid["passwd"] = "123456"
pamarsBulid["action"] = "login"
r.QueryParams().Sets(pamarsBulid)
```

#### 2.4.3 链式赋值
```go
r.QueryParams().Set("user", "dorlolo").Set("passwd","123456")
```

#### 2.4.4 获取url.Values对象进行赋值
对象类型为`*url.Values`,取到地址后，可以使用`url.Values`中的方法继续进行赋值
```go
qpData:=r.QueryParams().Gets()
qpData.Add("age","18")
```

### 2.5 添加请求体body

#### 2.5.1 单个赋值
```go
r.Body().Set("beginDate", "2022-03-01").Set("endDate", "2022-03-03")

```

#### 2.5.2 map赋值
```go
bodyBulid := map[string]interface{}{
    "beginDate":"2022-03-01",
	"endDate":"2022-03-03",
}
r.Body().Sets(bodyBulid)
```

#### 2.5.3 链式赋值
```go
r.Body().Set("beginDate", "2022-03-01").Set("endDate", "2022-03-03")
```

#### 2.5.4 字符串赋值
json格式不要使用此方法
```go
bodydata:=`{"devSn":"230000000008","type":"day"}`
r.Body().SetString(bodydata)
```

###  2.6 其它请求参数

#### 2.6.1 设置超时时间
```go
r.TimeOut(time.Second * 30)
```

#### 2.6.2 跳过证书验证
```go
r.SkipCertVerify()
```

### 2.7 发送请求
#### 2.7.1 post请求
```go
res, err :=r.Post("https://127.0.0.1:80/excample")
if err != nil {
    ftm.Println( "error occured", err)
} else {
    fmt.Println(res)
}
```

#### 2.7.2 get请求
```go
res, err :=r.Get("https://127.0.0.1:80/excample")
if err != nil {
ftm.Println( "error occured", err)
} else {
fmt.Println(res)
}
```
#### 2.7.3 其它类型的请求
后续支持...敬请期待


### 2.8 获取上下文
请注意，需要完成请求后才能获得上下文数据！
#### 2.8.1 获取请求的上下文对象
```go
requestContext:=r.Request
```

#### 2.8.2 获取返回的上下文对象
```go
responseContext:=r.Response
```

## 3. 使用示例
[simpleRequest_test.go](./test/simpleRequest_test.go)