// Package simpleRequest -----------------------------
// @file      : options_test.go
// @author    : JJXu
// @contact   : wavingbear@163.com
// @time      : 2024/1/12 11:27
// -------------------------------------------
package simpleRequest

import (
	"io"
	"testing"
)

func TestOptionDisableDefaultContentType(t *testing.T) {
	go httpserver()
	var r = NewRequest(OptionDisableDefaultContentType())
	r.Headers()
	bodyData := "{'a'=123,'b'=56}"
	r.Body().SetString(bodyData)
	_, err := r.POST("http://localhost:8989")
	if err != nil {
		t.Error(err)
		return
	}
	if r.Request.Header.Get(hdrContentTypeKey) != "" {
		t.Errorf("query params want '%s' but get '%s'", "", r.Request.Header.Get(hdrContentTypeKey))
	}
}

func TestOptionOptionDisableCopyRequestBody(t *testing.T) {
	go httpserver()
	var r = NewRequest(OptionDisableCopyRequestBody())
	r.Headers()
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
	if string(body) != "" {
		t.Errorf("query params want '%s' but get '%s'", "", body)
	}
}
