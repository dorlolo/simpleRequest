# simpleRequest -- 基于http库封装的request请求方案
## 特点
适用于基于http协议的iot数据采集。 功能简介、易于上手。只要了解postman工具，就能使用这个库。极大地减少了代码量。

## 如何使用?

### 模块导入

```go
import simpleReuqest2 "go-simple/simpleRequest/v2"
```

### 实例化

```go
var r = simpleReuqest2.NewRequest()
```

### 添加请求头

#### 单个赋值
```go
r.Headers().Set("token", "d+jfdji*D%1=")
r.Headers().Set("Content-Type", "application/json")
```

#### map赋值
```go
mapHeaders:= map[string]string{
    "token":      "d+jfdji*D%1=",
    "Content-Type": "application/json",
}
r.Headers().Sets(mapHeaders)
```

#### 链式赋值
```go
r.Headers().Set("token", "d+jfdji*D%1=").Set("Content-Type", "application/json")
```
#### 使用预设的key
```go
r.Headers().SetConentLength("2")
//r.Headers().Set("Content-Length", "2")

r.Headers().SetConentEncoding("gzip, deflate, br")
//r.Headers().Set("Content-Encoding", "gzip, deflate, br")

r.Headers().SetUserAgent("Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/535.1 (KHTML, like Gecko) Chrome/14.0.835.163 Safari/535.1"")
//r.Headers().Set("User-Agent",Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/535.1 (KHTML, like Gecko) Chrome/14.0.835.163 Safari/535.1")
```
#### 使用预设的key-value
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

### 添加queryParams
#### 单个赋值
```go
r.QueryParams().Set("user", "dorlolo")
```
#### map赋值
不会覆盖上面之前填充过的参数
```go
pamarsBulid := make(map[string]interface{})
pamarsBulid["passwd"] = "123456"
pamarsBulid["action"] = "login"
r.QueryParams().Sets(pamarsBulid)
```

#### 链式赋值
```go
r.QueryParams().Set("user", "dorlolo").Set("passwd","123456")
```

#### 获取url.Values对象进行赋值
对象类型为`*url.Values`,取到地址后，可以使用`url.Values`中的方法继续进行赋值
```go
qpData:=r.QueryParams().Gets()
qpData.Add("age","18")
```

### 添加请求体body

#### 单个赋值
```go
r.Body().Set("beginDate", "2022-03-01").Set("endDate", "2022-03-03")

```

#### map赋值
```go
bodyBulid := map[string]interface{
    "beginDate":"2022-03-01",
	"endDate":"2022-03-03",
}
r.Body().Sets(bodyBulid)
```

#### 链式赋值
```go
r.Body().Set("beginDate", "2022-03-01").Set("endDate", "2022-03-03")
```

#### 字符串赋值
通过此方法直接赋值json数据
```go
bodydata:=`{"devSn":"230000000008","type":"day"}`
r.Body().SetString(bodydata)
```

###  其它请求参数

#### 设置超时时间
```go
r.TimeOut(time.Second * 30)
```

#### 跳过证书验证
```go
r.SkipCertVerify()
```

### 发送请求
#### post请求
```go
res, err :=r.Post("https://127.0.0.1:80/excample")
if err != nil {
    ftm.Println( "error occured", err)
} else {
    fmt.Println(res)
}
```

#### get请求
```go
res, err :=r.Get("https://127.0.0.1:80/excample")
if err != nil {
ftm.Println( "error occured", err)
} else {
fmt.Println(res)
}
```
#### 其它请求
后续支持...敬请期待


### 获取上下文
请注意，需要完成请求后才能获得上下文数据！
#### 获取请求的上下文
```go
requestContext:=r.Request
```

#### 获取返回的上下文
```go
responseContext:=r.Response
```

## 使用示例
[simpleRequest_test.go](./test/simpleRequest_test.go)