// Package simpleRequest -----------------------------
// file      : simpleRequest_test.go
// author    : JJXu
// contact   : wavingBear@163.com
// time      : 2022/12/9 20:34:52
// -------------------------------------------
package simpleRequest

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"
)

type api struct {
	Name string `json:"name" form:"name" query:"name"`
}

func httpserver() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			var data api
			jsonDecoder := json.NewDecoder(r.Body)
			jsonDecoder.Decode(&data)
			if data.Name == "JJXu" {
				io.WriteString(w, "ok")
			} else {
				io.WriteString(w, "false")
			}

		default:
			io.WriteString(w, "false")
		}
	})
	//fmt.Println("http服务启动了")
	http.ListenAndServe(":8989", nil)
}

func TestPost_withSet(t *testing.T) {
	go httpserver()
	var r = NewRequest()
	r.Headers().ConentType_json()
	r.Body().Set("name", "JJXu")
	result, err := r.POST("http://localhost:8989/")
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log(string(result))
	}
}

func TestPost_withSets(t *testing.T) {
	go httpserver()

	var r = NewRequest()
	r.Headers().ConentType_json()
	r.Body().Sets(map[string]any{
		"name": "JJXu",
	})
	result, err := r.POST("http://localhost:8989/")
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log(string(result))
	}
}

func TestPost_withSetModel(t *testing.T) {
	go httpserver()

	var r = NewRequest()
	r.Headers().ConentType_json()
	var entry = api{
		Name: "JJXu",
	}
	r.Body().SetModel(&entry)
	result, err := r.POST("http://localhost:8989/")
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log(string(result))
	}
}

// url中的query param字符串参数会被r.QueryParams()中的key值覆盖
func TestQueryUrl2(t *testing.T) {
	go httpserver()

	var r = NewRequest()
	r.Headers().ConentType_formUrlencoded()
	r.QueryParams().Set("a", "123")
	r.QueryParams().Set("b", "456")
	_, err := r.POST("http://localhost:8989?a=1&b=2&c=3")
	if err != nil {
		t.Error(err.Error())
	} else {
		if r.Request.URL.RawQuery != "a=123&b=456&c=3" {
			t.Errorf("query params want '%s' but get '%s'", "a=123&b=456&c=3", r.Request.URL.RawQuery)
		}
	}

}

// 请求后，r.Request.Body中的内容仍旧可读
func TestQueryUseStringBody(t *testing.T) {
	go httpserver()
	var r = NewRequest()
	r.Headers().ConentType_json()
	bodyData := "{'a'=123,'b'=56}"
	r.Body().SetString(bodyData)
	_, err := r.POST("http://localhost:8989")
	if err != nil {
		t.Error(err)
		return
	}
	body, err := io.ReadAll(r.Request.Body)
	if err != nil {
		t.Error(err)
		return
	}
	if string(body) != bodyData {
		t.Errorf("request body want '%s' but get '%s'", bodyData, body)
	}
}

func TestEmptyBody(t *testing.T) {
	go httpserver()
	var r = NewRequest()
	r.Headers().ConentType_json()
	_, err := r.POST("http://localhost:8989")
	if err != nil {
		t.Error(err)
		return
	}
	body, err := io.ReadAll(r.Request.Body)
	if err != nil {
		t.Error(err)
		return
	}
	if string(body) != "{}" {
		t.Errorf("request body want '%s' but get '%s'", "{}", body)
	}

	r.Headers().ConentType_formUrlencoded()
	_, err = r.POST("http://localhost:8989")
	if err != nil {
		t.Error(err)
		return
	}
	body, err = io.ReadAll(r.Request.Body)
	if err != nil {
		t.Error(err)
		return
	}
	if string(body) != "" {
		t.Errorf("request body want '%s' but get '%s'", "", body)
	}
}
