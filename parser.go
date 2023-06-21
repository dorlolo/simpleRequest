// Package sRequest -----------------------------
// file      : parser.go
// author    : JJXu
// contact   : wavingBear@163.com
// time      : 2022/12/10 00:48:45
// -------------------------------------------
package simpleRequest

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
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
		return strings.NewReader(BodyEntry[StringEntryType.string()].(string))
	case BytesEntryType:
		return bytes.NewReader(BodyEntry[BytesEntryType.string()].([]byte))
	default:
		body, f.ContentType = multipartCommonParse(BodyEntry)
	}
	return
}
func multipartCommonParse(BodyEntry map[string]any) (reader io.Reader, contentType string) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	for k, sv := range BodyEntry {
		if strings.Contains(k, FormFilePathKey.string()) {
			key := k[len(FormFilePathKey):]
			path := sv.(string)
			filename := filepath.Base(path)
			filePart, _ := writer.CreateFormFile(key, filename)
			content, err := os.ReadFile(path)
			if err != nil {
				panic(err)
			}
			_, _ = filePart.Write(content)
		} else {
			switch sv.(type) {
			case string:
				strSv, _ := sv.(string)
				_ = writer.WriteField(k, strSv)
			case []string:
				sss, _ := sv.([]string)
				for _, v := range sss {
					_ = writer.WriteField(k, v)
				}
			case *multipart.FileHeader:
				file, _ := sv.(*multipart.FileHeader)
				filePart, _ := writer.CreateFormFile(k, file.Filename)
				src, err := file.Open()
				if err != nil {
					panic(err)
					return
				}
				defer src.Close()
				_, err = io.Copy(filePart, src)
				if err != nil {
					panic(err)
					return
				}
			}
		}

	}
	err := writer.Close()
	if err != nil {
		panic(err)
	}
	return body, writer.FormDataContentType()
}
