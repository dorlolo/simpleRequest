/*
 * @FileName:   simpleRequest.go
 * @Author:		JuneXu
 * @CreateTime:	2022/3/2 上午12:33
 * @Description:
 */

package simpleRequest

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func NewRequest() *SimpleRequest {
	var (
		hd = http.Header{}
		qp = url.Values{}
	)

	return &SimpleRequest{
		//headers: make(map[string]string),
		//cookies: make(map[string]string),
		timeout:     time.Second * 7,
		queryParams: qp,
		headers:     hd,
		tempBody:    make(map[string]interface{}),
	}
}

type SimpleRequest struct {
	url         string
	queryParams url.Values
	body        io.Reader
	headers     http.Header
	transport   *http.Transport

	tempBody map[string]interface{}
	timeout  time.Duration

	Response http.Response //用于获取完整的返回内容。请注意要在请求之后才能获取
	Request  http.Request  //用于获取完整的请求内容。请注意要在请求之后才能获取
	//cookies           map[string]string
	//data              interface{}
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

func (s *SimpleRequest) NewRequest() *SimpleRequest {
	var qp = url.Values{}
	return &SimpleRequest{
		//headers: make(map[string]string),
		//cookies: make(map[string]string),
		timeout:     time.Second * 7,
		queryParams: qp,
		tempBody:    make(map[string]interface{}),
	}
}

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
	}

	//v0.0.2更新，将request和response内容返回，便于用户进行分析 JuneXu 03-11-2022
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

func (s *SimpleRequest) Post(urls string) (body []byte, err error) {
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

func (s *SimpleRequest) Get(urls string) (body []byte, err error) {
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

// Put method does PUT HTTP request. It's defined in section 4.3.4 of RFC7231.
//func (s *SimpleRequest) Put(url string) (*Response, error) {
//	return r.Execute(MethodPut, url)
//}

// Delete method does DELETE HTTP request. It's defined in section 4.3.5 of RFC7231.
//func (s *SimpleRequest) Delete(url string) (*Response, error) {
//	return r.Execute(MethodDelete, url)
//}

// Options method does OPTIONS HTTP request. It's defined in section 4.3.7 of RFC7231.
//func (s *SimpleRequest) Options(url string) (*Response, error) {
//	return r.Execute(MethodOptions, url)
//}

// Patch method does PATCH HTTP request. It's defined in section 2 of RFC5789.
//func (s *SimpleRequest) Patch(url string) (*Response, error) {
//	return r.Execute(MethodPatch, url)
//}
//------------------------------------------------------
//
//						这里数据
//
func (s *SimpleRequest) initBody() {
	contentTypeData := s.headers.Get(hdrContentTypeKey)
	switch {
	case IsJSONType(contentTypeData):
		jsonData, err := json.Marshal(s.tempBody)
		if err == nil {
			s.body = bytes.NewReader(jsonData)
		} else {
			s.body = bytes.NewReader([]byte("{}"))
		}
	case strings.Contains(contentTypeData, "multipart/form-data"):
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		//data := url.Values{}
		for k, sv := range s.tempBody {
			switch sv.(type) {
			case string:
				strSv, _ := sv.(string)
				_ = writer.WriteField(k, strSv)
			case []string:
				sss, _ := sv.([]string)
				for _, v := range sss {
					_ = writer.WriteField(k, v)
				}
			}
		}
		err := writer.Close()
		if err != nil {
			panic(err)
		}
		s.headers.Set("Content-Type", writer.FormDataContentType())
		s.body = body
	case IsXMLType(contentTypeData):
		//application/soap+xml ,application/xml
		data, _ := s.tempBody[stringBodyType].(string)
		s.body = strings.NewReader(data)
	case strings.Contains(contentTypeData, "text") || strings.Contains(contentTypeData, javaScriptType):
		data, _ := s.tempBody[stringBodyType].(string)
		s.body = strings.NewReader(data)
	case contentTypeData == "" || strings.Contains(contentTypeData, "form-urlencoded"):
		//默认为x-www-form-urlencoded格式
		tmpData := url.Values{}
		for k, v := range tmpData {
			tmpData.Set(k, fmt.Sprintf("%v", v))
		}
		s.body = strings.NewReader(tmpData.Encode())
		s.Headers().ConentType_formUrlencoded()
	default:
		//todo 自动判断数据类型
		tmpData := url.Values{}
		for k, v := range tmpData {
			tmpData.Set(k, fmt.Sprintf("%v", v))
		}
		s.body = strings.NewReader(tmpData.Encode())
	}
}
