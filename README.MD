_# simpleRequest 
## 1. 说明
[simpleRequest](www.github.com/dorlolo/simpleRequest) 是一款面向对象开发的http请求库。他是基于Go原生http库。开发这个模块的主要目的是为了更快的对接http协议的IOT设备。
它具备以下特点：
- 极简化的代码，大大提高了代码编写速度。
- 更易于理解和使用，减少了资料查询的时间。
- 更适合对接一些未遵循restful规范的接口。

这是它的一个示例:
```go
var r = simpleRequest.NewRequest()
//设置请求头
r.Headers().Set("x-token", "d+jfdji*D%1=")
// 添加query
r.QueryParams().Set("user", "JJXu").Set("job", "developer")
//--发送请求
res, err := r.GET("http://www.webSite.com/end/point")
if err != nil {
    t.Error(err)
} else {
    fmt.Println(res)
}
```
## 2. 如何使用?
### 2.2 安装
```bash
go get github.com/dorlolo/simpleRequest
```
### 2.2 模块导入

```go
import  "github.com/dorlolo/simpleRequest"
```

### 2.3 实例化

```go
var r = simpleRequest.NewRequest()
```

### 2.4 添加请求头

#### 2.4.1 单个赋值
```go
r.Headers().Set("token", "d+jfdji*D%1=")
r.Headers().Set("Content-Type", "application/json")
```

#### 2.4.2 map赋值
```go
mapHeaders:= map[string]string{
    "token":      "d+jfdji*D%1=",
    "Content-Type": "application/json",
}
r.Headers().Sets(mapHeaders)
```

#### 2.4.3 链式赋值
```go
r.Headers().Set("token", "d+jfdji*D%1=").Set("Content-Type", "application/json")
```

#### 2.4.4 添加多值
对单个key添加多只不要使用`.set`,因为原先的值会被覆盖
```go
r.Headers().Set("Accept", "text/html")
r.Headers().Add("Accept","application/xhtml+xml")
r.Headers().Add("Accept","application/xml;q=0.8")
r.Headers().Add("Accept","image/webp")
r.Headers().Add("Accept","*/*;q=0.8")
```

#### 2.4.4 使用预设的key
```go
r.Headers().SetConentType("application/json")
//r.Headers().Set("Content-Type", "application/json")

r.Headers().SetUserAgent("Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/535.1 (KHTML, like Gecko) Chrome/14.0.835.163 Safari/535.1"")
//r.Headers().Set("User-Agent",Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/535.1 (KHTML, like Gecko) Chrome/14.0.835.163 Safari/535.1")

r.Headers().SetConentEncoding("gzip, deflate, br")
//r.Headers().Set("Content-Encoding", "gzip, deflate, br")
```

#### 2.4.5 使用预设的key-value
```go
//随机user-agent
r.Headers().SetRandomUerAgent()
//r.Headers().Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/535.1 (KHTML, like Gecko) Chrome/14.0.835.163 Safari/535.1")

r.Headers().ConentType_json()
//r.Headers().Set("Content-Type", "application/json")

r.Headers().ConentType_formData()
//r.Headers().Set("Content-Type","multipart/form-data")

r.Headers().ConentType_formUrlencoded()
//r.Headers().Set("Content-Type","application/x-www-form-urlencoded")

r.Headers().ConentType_textPlain()
//r.Headers().Set("Content-Type","text/plain; charset=utf-8")
```

### 2.5 添加queryParams
#### 2.5.1 单个赋值
```go
r.QueryParams().Set("user", "dorlolo")
```
#### 2.5.2 map赋值
不会覆盖上面之前填充过的参数
```go
pamarsBulid := make(map[string]any)
pamarsBulid["passwd"] = "123456"
pamarsBulid["action"] = "login"
r.QueryParams().Sets(pamarsBulid)
```

#### 2.5.3 链式赋值
```go
r.QueryParams().Set("user", "dorlolo").Set("passwd","123456")
```

#### 2.5.4 获取url.Values对象进行赋值
对象类型为`*url.Values`,取到地址后，可以使用`url.Values`中的方法继续进行赋值
```go
qpData:=r.QueryParams().Gets()
qpData.Add("age","18")
```

### 2.6 添加请求体body

#### 2.6.1 单个参数赋值
支持和map赋值同时使用
```go
r.Body().Set("beginDate", "2022-03-01").Set("endDate", "2022-03-03")

```

#### 2.6.2 map赋值
支持和单个参数赋值同时使用
```go
bodyBulid := map[string]any{
    "beginDate":"2022-03-01",
	"endDate":"2022-03-03",
}
r.Body().Sets(bodyBulid)
```

#### 2.6.3 技巧：链式赋值
```go
r.Body().Set("beginDate", "2022-03-01").Set("endDate", "2022-03-03")
```

#### 2.6.4 字符串赋值
此方法为一次性赋值，不支持和其它赋值方法同时使用
```go
bodydata:=`{"devSn":"230000000008","type":"day"}`
r.Body().SetString(bodydata)
```

#### 2.6.5 字节赋值
此方法为一次性赋值，不支持和其它赋值方法同时使用
```go
bytesdata:=[]byte(`{"devSn":"230000000008","type":"day"}`)
r.Body().SetBytes(bytesdata)
```

#### 2.6.6 结构体赋值
此方法为一次性赋值，不支持和其它赋值方法同时使用
```go
type demo struct{
	DevSn string `json:"devSn"`
	Type string `json:"day"`
}
modeldata:=demo{
    DevSn:"230000000008"
    Type:"day"
}
r.Body().SetModel(&modeldata)
```

### 2.7 文件上传与转发
### 2.7.1 文件上传
```go
var req = sRequest.NewRequest()
req.Headers().ConentType_formData()
req.Body().
    SetFromDataFile("file", "C:\\Users\\lenovo\\Pictures\\Saved Pictures\\demo.jpg").
    Set("fromFormat", "jpg").
    Set("toFormat", "png")
req.TimeOut(15 * time.Second)
resp, err := req.POST("http://xxx/xxx")
if err != nil {
    fmt.Println(err.Error())
    return
}
fmt.Println(string(resp))
```
### 2.7.2 文件转发
下面示例中使用gin作为服务端，配合simpleRequest进行文件转发
1. 通过multipart.FileHeader对象进行转发
```go
func FileForwardUseMultipartFile(c *gin.Context){
    file,err:=c.FormFile("file")
    
	var req = sRequest.NewRequest()
    req.Headers().ConentType_formData()
    req.Body().
        SetFromDataMultipartFile("file", file).
        Set("fromFormat", "jpg").
        Set("toFormat", "png")
    req.TimeOut(15 * time.Second)
    resp, err := req.POST("http://xxx/xxx")
    if err != nil {
        fmt.Println(err.Error())
        return
    }
   // parse response and so on
   // ...
   // ...
```

2. 在一些小众场景下，可能已经在外部构建好了body，此时也可将body转为bytes传入simpleRequest进行请求
```go
func FileForwardUseBytesBody(c *gin.Context){
	file,err:=c.FormFile("file")
	
	// body data prepare
    vars (
		body = &bytes.Buffer{}
        writer = multipart.NewWriter(body)
    )
	// add file object
	filePart, _ := i.writer.CreateFormFile("file", file.Filename)
    src, err := file.Open()
    if err != nil {
        fmt.Println( err.Error())
        return
    }
    defer src.Close()
    _, err = io.Copy(filePart, src)
    if err != nil {
        fmt.Println(err.Error())
        return
    }
	// add other form data
	writer.WriteField("fromFormat", "jpg")
	writer.WriteField("toFormat","png")
	
	// post request
	_ = writer.close()
	var r = simpleRequest.NewRequest()
    req.Headers().SetConentType(writer.FormDataContentType())
    req.Body().SetBytes(body.Bytes())
    req.TimeOut(15 * time.Second)
    resp, err := req.POST("http://xxx/xxx")
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    // parse response and so on
	// ...
	// ...
}

```

###  2.8 其它请求参数

#### 2.8.1 设置超时时间
```go
r.TimeOut(time.Second * 30)
```

#### 2.7.2 跳过证书验证
```go
r.SkipCertVerify()
```

### 2.9 发送请求
#### 2.9.1 post请求
```go
res, err :=r.POST("https://127.0.0.1:80/excample")
if err != nil {
    ftm.Println( "error occured", err)
} else {
    fmt.Println(res)
}
```
**支持的请求类型有**
- POST
- GET
- PUT
- DELETE
- PATCH
- HEAD
- CONNECT
- OPTIONS
- TRACE


### 2.10 获取上下文
请注意，需要完成请求后才能获得上下文数据！
#### 2.10.1 获取请求的上下文对象
```go
requestContext:=r.Request
```

#### 2.10.2 获取返回的上下文对象
```go
responseContext:=r.Response
```

## 3. 使用示例
[simpleRequest_test.go](excample/simpleRequest_test.go)_