/*
 * @FileName:   header.go
 * @Author:		JuneXu
 * @CreateTime:	2022/3/1 下午9:44
 * @Description:
 */

package simpleRequest

import (
	"bytes"
	"encoding/json"
	"math/rand"
	"net/http"
	"regexp"
	"sync"
	"time"
)

var (
	hdrUserAgentKey       = http.CanonicalHeaderKey("User-Agent")
	hdrAcceptKey          = http.CanonicalHeaderKey("Accept")
	hdrContentTypeKey     = http.CanonicalHeaderKey("Content-Type")
	hdrContentLengthKey   = http.CanonicalHeaderKey("Content-Length")
	hdrContentEncodingKey = http.CanonicalHeaderKey("Content-Encoding")
	hdrLocationKey        = http.CanonicalHeaderKey("Location")

	plainTextType      = "text/plain; charset=utf-8"
	jsonContentType    = "application/json"
	formUrlencodedType = "application/x-www-form-urlencoded"
	formDataType       = "multipart/form-data"
	xmlDataType        = "application/xml"
	textPlainType      = "text/plain"
	javaScriptType     = "javascript"

	jsonCheck = regexp.MustCompile(`(?i:(application|text)/(json|.*\+json|json\-.*)(;|$))`)
	xmlCheck  = regexp.MustCompile(`(?i:(application|text)/(xml|.*\+xml)(;|$))`)
	bufPool   = &sync.Pool{New: func() interface{} { return &bytes.Buffer{} }}
)

var userAgentList = [...]string{
	"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0;",
	"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; Trident/4.0; InfoPath.2; .NET4.0C; .NET4.0E; .NET CLR 2.0.50727; 360SE)",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/535.1 (KHTML, like Gecko) Chrome/14.0.835.163 Safari/535.1",
}

type HeadersConf struct {
	simpleReq *SimpleRequest
}

//-------------------------------------------------------------
// Common key settings

//batch settings
func (s *HeadersConf) Sets(headers map[string]string) *HeadersConf {
	for k, v := range headers {
		s.simpleReq.headers.Set(k, v)
	}
	return s
}

//single setting
func (s *HeadersConf) Set(header, value string) *HeadersConf {
	s.simpleReq.headers.Set(header, value)
	return s
}

func (s *HeadersConf) Add(header, value string) *HeadersConf {
	s.simpleReq.headers.Add(header, value)
	return s
}

//一般用不到
//func (s *HeadersConf) Values(keys string) *HeadersConf {
//	s.simpleReq.headers.Values(keys)
//	return s
//}

// SetHeaderMultiValues 支持多值传入，一般用不到
//func (s *HeadersConf) SetMultiValues(headers map[string][]string) *HeadersConf {
//	for key, values := range headers {
//		s.simpleReq.headers.Set(key, strings.Join(values, ", "))
//	}
//	return s
//}

//-------------------------------------------------------------
// base Key settings

//SetUserAgent
func (s *HeadersConf) SetUserAgent(value string) *HeadersConf {
	s.simpleReq.headers.Set(hdrUserAgentKey, value)
	return s
}

//SetConentType
func (s *HeadersConf) SetConentType(value string) *HeadersConf {
	s.simpleReq.headers.Set(hdrContentTypeKey, value)
	return s
}

func (s *HeadersConf) ConentType_json() *HeadersConf {
	jsonData, err := json.Marshal(s.simpleReq.tempBody)
	if err == nil {
		s.simpleReq.body = bytes.NewReader(jsonData)
	}
	s.simpleReq.body = bytes.NewReader(jsonData)
	s.simpleReq.headers.Set(hdrContentTypeKey, jsonContentType)
	return s
}

func (s *HeadersConf) ConentType_formData() *HeadersConf {
	//tmp := url.Values{}

	//for k, v := range s.simpleReq.tempBody {
	//	tmp.Add(k, fmt.Sprintf("%v", v))
	//}
	s.simpleReq.headers.Set(hdrContentTypeKey, formDataType)
	return s
}
func (s *HeadersConf) ConentType_formUrlencoded() *HeadersConf {
	s.simpleReq.headers.Set(hdrContentTypeKey, formUrlencodedType)
	return s
}
func (s *HeadersConf) ConentType_textPlain() *HeadersConf {
	s.simpleReq.headers.Set(hdrContentTypeKey, plainTextType)
	return s
}

//
func (s *HeadersConf) SetConentLength(value string) *HeadersConf {
	s.simpleReq.headers.Set(hdrContentLengthKey, value)
	return s
}
func (s *HeadersConf) SetConentEncoding(value string) *HeadersConf {
	s.simpleReq.headers.Set(hdrContentEncodingKey, value)
	return s
}
func (s *HeadersConf) SetConentLocation(value string) *HeadersConf {
	s.simpleReq.headers.Set(hdrLocationKey, value)
	return s
}

//-------------------------------------------------------------
// Extended settings
//随机请求头的User-Agent参数
func (s *HeadersConf) getRandomUerAgent() string {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(userAgentList))
	return userAgentList[index]
}

//设置为随机 User-Agent
func (s *HeadersConf) SetRandomUerAgent() *HeadersConf {
	s.simpleReq.headers.Set(hdrUserAgentKey, s.getRandomUerAgent())
	return s
}

//set Authorization
func (s *HeadersConf) SetAuthorization(value string) *HeadersConf {
	s.simpleReq.headers.Set("Authorization", value)
	return s
}
