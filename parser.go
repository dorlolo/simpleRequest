// Package simpleRequest -----------------------------
// @file      : parser.go
// @author    : JJXu
// @contact   : wavingBear@163.com
// @time      : 2022/12/10 00:48:45
// -------------------------------------------
package simpleRequest

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"strings"
)

var bodyEntryParsers = map[string]IBodyEntryParser{
	jsonContentType: new(JsonParser),
	formDataType:    new(FormDataParser),
}

type IBodyEntryParser interface {
	Unmarshal(bodyType EntryMark, BodyEntry map[string]any) io.Reader
}

type JsonParser struct{}

func (JsonParser) Unmarshal(bodyType EntryMark, BodyEntry map[string]any) io.Reader {
	switch bodyType {
	case StringEntryType:
		return strings.NewReader(BodyEntry[StringEntryType.string()].(string))
	case BytesEntryType:
		return bytes.NewReader(BodyEntry[BytesEntryType.string()].([]byte))
	case ModelEntryType:
		jsonData, err := json.Marshal(BodyEntry[ModelEntryType.string()])
		if err == nil {
			return bytes.NewReader(jsonData)
		}
		return strings.NewReader("{}")
	case MapEntryType:
		jsonData, err := json.Marshal(BodyEntry)
		if err == nil {
			return bytes.NewReader(jsonData)
		} else {
			return strings.NewReader("{}")
		}
	default:
		if len(BodyEntry) > 0 {
			jsonData, err := json.Marshal(BodyEntry)
			if err == nil {
				return bytes.NewReader(jsonData)
			}
		}
		return strings.NewReader("{}")
	}
}

type FormDataParser struct {
	ContentType string
}

func (f *FormDataParser) Unmarshal(bodyType EntryMark, BodyEntry map[string]any) (body io.Reader) {
	switch bodyType {
	case MapEntryType:
		body, f.ContentType = multipartCommonParse(BodyEntry)
	case ModelEntryType:
		tb := BodyEntry[ModelEntryType.string()]
		buffer, err := json.Marshal(tb)
		if err != nil {
			panic(err.Error())
		}
		var mapper map[string]any
		err = json.Unmarshal(buffer, &mapper)
		if err != nil {
			panic(err.Error())
		}
		body, f.ContentType = multipartCommonParse(mapper)
	case StringEntryType:
		f.ContentType = formDataType
		return strings.NewReader(BodyEntry[StringEntryType.string()].(string))
	case BytesEntryType:
		f.ContentType = formDataType
		return bytes.NewReader(BodyEntry[BytesEntryType.string()].([]byte))
	default:
		body, f.ContentType = multipartCommonParse(BodyEntry)
	}
	f.ContentType = formDataType
	return nil
}
func multipartCommonParse(BodyEntry map[string]any) (reader io.Reader, contentType string) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	for k, sv := range BodyEntry {
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
	return body, writer.FormDataContentType()
}
