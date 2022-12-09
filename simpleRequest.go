/*
 * @FileName:   simpleRequest.go
 * @Author:		JJXu
 * @CreateTime:	2022/3/2 上午12:33
 * @Description:
 */

package simpleRequest

import (
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func NewRequest(opts ...OPTION) *SimpleRequest {
	var (
		hd = http.Header{}
		qp = url.Values{}
	)

	var r = &SimpleRequest{
		//headers: make(map[string]string),
		//cookies: make(map[string]string),
		timeout:          time.Second * 7,
		queryParams:      qp,
		headers:          hd,
		BodyEntries:      make(map[string]any),
		bodyEntryParsers: bodyEntryParsers,
	}
	if len(opts) > 0 {
		for _, o := range opts {
			r = o(r)
		}
	}
	return r
}

type SimpleRequest struct {
	url         string
	queryParams url.Values
	body        io.Reader
	headers     http.Header
	transport   *http.Transport

	BodyEntryMark    EntryMark
	BodyEntries      map[string]any
	bodyEntryParsers map[string]IBodyEntryParser

	timeout time.Duration

	Response http.Response //用于获取完整的返回内容。请注意要在请求之后才能获取
	Request  http.Request  //用于获取完整的请求内容。请注意要在请求之后才能获取
	//cookies           map[string]string
	//data              any
	//cli               *http.Client
	//debug             bool
	//method            string
	//time              int64
	//disableKeepAlives bool
	//tlsClientConfig   *tls.Config
	//jar               http.CookieJar
	//proxy             func(*http.Request) (*url.URL, error)
	//checkRedirect     func(req *http.Request, via []*http.Request) error
}

//func (s *SimpleRequest) NewRequest() *SimpleRequest {
//	var qp = url.Values{}
//	return &SimpleRequest{
//		//headers: make(map[string]string),
//		//cookies: make(map[string]string),
//		timeout:     time.Second * 7,
//		queryParams: qp,
//		BodyEntries:   make(map[string]any),
//	}
//}

//------------------------------------------------------
//
//						数据准备

//Authorization 添加令牌的方法集合
func (s *SimpleRequest) Authorization() *Authorization {
	return &Authorization{
		simpleReq: s,
	}
}

//Headers 添加请求头
func (s *SimpleRequest) Headers() *HeadersConf {
	return &HeadersConf{
		simpleReq: s,
	}
}

//Body 添加请求体
func (s *SimpleRequest) Body() *BodyConf {
	return &BodyConf{
		simpleReq: s,
	}
}

//QueryParams 添加url后面的参数
func (s *SimpleRequest) QueryParams() *QueryParams {
	return &QueryParams{
		simpleReq: s,
	}
}

//跳过证书验证
func (s *SimpleRequest) SkipCertVerify() *SimpleRequest {

	s.transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	return s
}

//设置超时时间
func (s *SimpleRequest) TimeOut(t time.Duration) *SimpleRequest {
	s.timeout = t
	return s
}

//------------------------------------------------------
//
//						发送请求
//
//发送postt请求
func (s *SimpleRequest) do(request *http.Request) (body []byte, err error) {
	//3. 建立http客户端
	client := &http.Client{
		Timeout: s.timeout,
	}
	if s.transport != nil {
		client.Transport = s.transport
	}
	//3.1 发送数据
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println("【Request Error】:", err.Error())
		return
	}

	//v0.0.2更新，将request和response内容返回，便于用户进行分析 JJXu 03-11-2022
	if resp != nil {
		s.Response = *resp
	}
	if request != nil {
		s.Request = *request
	}

	defer resp.Body.Close()
	//3.2 获取数据
	body, err = ioutil.ReadAll(resp.Body)
	return
}

// POST method does POST HTTP request. It's defined in section 2 of RFC5789.
func (s *SimpleRequest) POST(urls string) (body []byte, err error) {
	s.initBody()
	r, err := http.NewRequest(http.MethodPost, urls, s.body)
	if err != nil {
		return nil, err
	}
	//headers
	for k := range s.headers {
		r.Header[k] = append(r.Header[k], s.headers[k]...)
		s.headers.Del(k)
	}

	//queryParams
	r.URL.RawQuery = s.queryParams.Encode()

	body, err = s.do(r)

	return
}

// GET method does GET HTTP request. It's defined in section 2 of RFC5789.
func (s *SimpleRequest) GET(urls string) (body []byte, err error) {
	// body
	s.initBody()
	r, err := http.NewRequest(http.MethodGet, urls, s.body)
	if err != nil {
		return nil, err
	}
	//headers
	for k := range s.headers {
		r.Header[k] = append(r.Header[k], s.headers[k]...)
		s.headers.Del(k)
	}
	//queryParams
	r.URL.RawQuery = s.queryParams.Encode()

	body, err = s.do(r)
	return
}

//通用的请求方法
func (s *SimpleRequest) LaunchTo(urls, method string) (body []byte, err error) {
	// body
	s.initBody()
	r, err := http.NewRequest(method, urls, s.body)
	if err != nil {
		return nil, err
	}
	//headers
	for k := range s.headers {
		r.Header[k] = append(r.Header[k], s.headers[k]...)
		s.headers.Del(k)
	}
	//queryParams
	r.URL.RawQuery = s.queryParams.Encode()

	body, err = s.do(r)
	return
}

// PUT method does PUT HTTP request. It's defined in section 4.3.4 of RFC7231.
func (s *SimpleRequest) PUT(url string) (body []byte, err error) {
	return s.LaunchTo(url, http.MethodPut)
}

// DELETE method does DELETE HTTP request. It's defined in section 2 of RFC5789.
func (s *SimpleRequest) DELETE(url string) (body []byte, err error) {
	return s.LaunchTo(url, http.MethodDelete)
}

// Patch method does Patch HTTP request. It's defined in section 2 of RFC5789.
func (s *SimpleRequest) PATCH(url string) (body []byte, err error) {
	return s.LaunchTo(url, http.MethodPatch)
}

// HEAD method does HEAD HTTP request. It's defined in section 2 of RFC5789.
func (s *SimpleRequest) HEAD(url string) (body []byte, err error) {
	return s.LaunchTo(url, http.MethodHead)
}

// CONNECT method does CONNECT HTTP request. It's defined in section 2 of RFC5789.
func (s *SimpleRequest) CONNECT(url string) (body []byte, err error) {
	return s.LaunchTo(url, http.MethodConnect)
}

// OPTIONS method does OPTIONS HTTP request. It's defined in section 2 of RFC5789.
func (s *SimpleRequest) OPTIONS(url string) (body []byte, err error) {
	return s.LaunchTo(url, http.MethodOptions)
}

// TRACE method does TRACE HTTP request. It's defined in section 2 of RFC5789.
func (s *SimpleRequest) TRACE(url string) (body []byte, err error) {
	return s.LaunchTo(url, http.MethodTrace)
}

//------------------------------------------------------
//
//						这里数据
//
func (s *SimpleRequest) initBody() {
	contentTypeData := s.headers.Get(hdrContentTypeKey)
	switch {
	case IsJSONType(contentTypeData):
		var parser, ok = s.bodyEntryParsers[jsonContentType]
		if !ok {
			panic(fmt.Sprintf("cannot find %s type parser", contentTypeData))
		}
		s.body = parser.Unmarshal(s.BodyEntryMark, s.BodyEntries)

	case strings.Contains(contentTypeData, formDataType):
		var parser, ok = s.bodyEntryParsers[formDataType]
		if !ok {
			panic(fmt.Sprintf("cannot find %s type parser", contentTypeData))
		}
		s.body = parser.Unmarshal(s.BodyEntryMark, s.BodyEntries)
		fdParser := parser.(*FormDataParser)
		s.headers.Set("Content-Type", fdParser.ContentType)

	case IsXMLType(contentTypeData):
		//application/soap+xml ,application/xml
		var parser, ok = s.bodyEntryParsers[xmlDataType]
		if !ok {
			data, _ := s.BodyEntries[StringEntryType.string()].(string)
			s.body = strings.NewReader(data)
		}
		s.body = parser.Unmarshal(s.BodyEntryMark, s.BodyEntries)

	case strings.Contains(contentTypeData, "text") || strings.Contains(contentTypeData, javaScriptType):
		var parser, ok = s.bodyEntryParsers[textPlainType]
		if !ok {
			data, _ := s.BodyEntries[StringEntryType.string()].(string)
			s.body = strings.NewReader(data)
		}
		s.body = parser.Unmarshal(s.BodyEntryMark, s.BodyEntries)

	case contentTypeData == "" || strings.Contains(contentTypeData, "form-urlencoded"):
		//默认为x-www-form-urlencoded格式
		var parser, ok = s.bodyEntryParsers["form-urlencoded"]
		if !ok {
			tmpData := url.Values{}
			for k, v := range s.BodyEntries {
				tmpData.Set(k, fmt.Sprintf("%v", v))
			}
			s.body = strings.NewReader(tmpData.Encode())
			s.Headers().ConentType_formUrlencoded()
		}
		s.body = parser.Unmarshal(s.BodyEntryMark, s.BodyEntries)
	default:
		//todo 自动判断数据类型
		tmpData := url.Values{}
		for k, v := range tmpData {
			tmpData.Set(k, fmt.Sprintf("%v", v))
		}
		s.body = strings.NewReader(tmpData.Encode())
	}
}
