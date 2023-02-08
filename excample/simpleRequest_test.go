/*
 *FileName:   simpleRequest_test.go
 *Author:		JJXu
 *CreateTime:	2022/3/3 下午11:34
 *Description:
 */

package excample

import (
	"fmt"
	"github.com/dorlolo/simpleRequest"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestRequest(t *testing.T) {
	var r = simpleRequest.NewRequest()
	//---设置请求头
	r.Headers().Set("token", "d+jfdji*D%1=")
	//串联使用示例：设置Conent-Type为applicaiton/json 并且 随机user-agent
	r.Headers().ConentType_json().SetRandomUerAgent()

	//设置params
	r.QueryParams().Set("user", "dorlolo")
	//批量添加,不会覆盖上面user
	pamarsBulid := map[string]any{
		"passwd": "123456",
		"action": "login",
	}
	r.QueryParams().Sets(pamarsBulid)

	//--添加body
	r.Body().Set("beginDate", "2022-03-01").Set("endDate", "2022-03-03")

	//--其它请求参数
	r.TimeOut(time.Second * 30) //请求超时,默认7秒
	r.SkipCertVerify()          //跳过证书验证

	//--发送请求,这里返回的直接是body中的数据，等后续增加功能
	res, err := r.GET("http://www.webSite.com/end/point")
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(res)
	}

}

// 测试content-type 为 multipart/form-data格式的数据请求
func TestAuth_fotmData(t *testing.T) {
	req := simpleRequest.NewRequest()
	req.Headers().ConentType_formData()
	req.Headers().SetRandomUerAgent()
	req.Body().Set("grant_type", "password").
		Set("client_id", "smz").
		Set("client_secret", "smz").
		Set("scope", "getdata").
		Set("username", "shiming_zyf").
		Set("password", "zyf499bbcb9")

	var URL = ""

	data, _ := req.POST(URL)
	t.Log(string(data))
}

// 测试令牌验证
func TestAuthorization(t *testing.T) {
	req := simpleRequest.NewRequest()
	req.Authorization().Bearer("19f0591e-fab1-4447-90c3-1c60aef78fbd")
	req.Body().
		Set("prjnumber", "3205072020100901A01000").
		Set("date", "20220324")
	data, err := req.PUT("")
	t.Log(string(data))
	t.Log(err)

}

func TestXml(t *testing.T) {
	idcard := "320324196705101880"
	pass := "96778"
	urlAddr := "http://218.4.84.171:5445/AppWebService/GHBackBone_SAMWS.asmx?Content-Type=application/soap+xml;charset=utf-8"
	body := fmt.Sprintf(`<?xml version="1.0" encoding="utf-8"?>
<soap12:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap12="http://www.w3.org/2003/05/soap-envelope">
    <soap12:Body>
        <GetWorkerSAM xmlns="http://tempuri.org/">
            <zjhm>%v</zjhm>
            <pass>%v</pass>
        </GetWorkerSAM>
    </soap12:Body>
</soap12:Envelope>`, idcard, pass)
	req := simpleRequest.NewRequest()
	req.Headers().
		Set("Content-Type", "application/soap+xml;charset=utf-8").
		SetRandomUerAgent()
	req.Body().SetString(body)
	data, err := req.POST(urlAddr)
	t.Log(string(data))
	t.Log(err)
	return
}

func TestIsJsonType(t *testing.T) {
	var headers = http.Header{}
	headers.Set("Content-Type", "application/json")
	headers.Add("Content-Type", "charset=UTF-8")
	RES := simpleRequest.IsJSONType(headers.Get("Content-Type"))
	t.Log(RES)

}
func TestIsXmlType(t *testing.T) {
	var headers = http.Header{}
	headers.Add("Content-Type", "application/soap+xml;charset=utf-8")
	RES := simpleRequest.IsXMLType(headers.Get("Content-Type"))
	t.Log(RES)
}

func TestTextPlain(t *testing.T) {

	var headers = http.Header{}
	headers.Add("Content-Type", "text/plain;charset=utf-8")

	res := strings.Contains(headers.Get("Content-Type"), "text")
	if res {
		t.Log(res)
	} else {
		t.Log(false)
	}

}
