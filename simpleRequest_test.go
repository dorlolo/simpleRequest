// Package simpleRequest -----------------------------
// @file      : simpleRequest_test.go
// @author    : JJXu
// @contact   : wavingBear@163.com
// @time      : 2022/12/9 20:34:52
// -------------------------------------------
package simpleRequest

import (
	"encoding/json"
	"fmt"
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
	fmt.Println("http服务启动了")
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
