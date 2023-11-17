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
	"encoding/xml"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

// 通用类型解析器
var bodyEntryParsers = map[string]IBodyEntryParser{
	jsonContentType: new(JsonParser),
	formDataType:    new(FormDataParser),
	xmlDataType:     new(XmlParser),
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
	formWriter := multipart.NewWriter(body)
	for k, sv := range BodyEntry {
		if strings.Contains(k, FormFilePathKey.string()) {
			fieldName := k[len(FormFilePathKey):]
			fp := sv.(string)
			filename := filepath.Base(fp)
			//way1
			filePart, _ := formWriter.CreateFormFile(fieldName, filename)
			content, err := os.ReadFile(fp)
			if err != nil {
				panic(err)
			}
			_, _ = filePart.Write(content)
		} else {
			switch multValue := sv.(type) {
			case string:
				_ = formWriter.WriteField(k, multValue)
			case []string:
				sss, _ := sv.([]string)
				for _, v := range sss {
					_ = formWriter.WriteField(k, v)
				}
			case *multipart.FileHeader:
				filePart, _ := formWriter.CreateFormFile(k, multValue.Filename)
				src, err := multValue.Open()
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
			case []byte:
				formWriter.WriteField(k, string(multValue))
			case int:
				formWriter.WriteField(k, fmt.Sprintf("%v", multValue))
			}
		}

	}
	err := formWriter.Close()
	if err != nil {
		panic(err)
	}
	return body, formWriter.FormDataContentType()
}

type XmlParser struct{}

func (f XmlParser) Unmarshal(bodyType EntryMark, BodyEntry map[string]any) (body io.Reader) {
	switch bodyType {
	case MapEntryType:
		xmlData, err := xml.Marshal(BodyEntry[bodyType.string()])
		if err == nil {
			return bytes.NewReader(xmlData)
		} else {
			return strings.NewReader("")
		}
	case ModelEntryType:
		xmlData, err := xml.Marshal(BodyEntry[bodyType.string()])
		if err == nil {
			return bytes.NewReader(xmlData)
		} else {
			return strings.NewReader("")
		}
	case StringEntryType:
		return strings.NewReader(BodyEntry[StringEntryType.string()].(string))
	case BytesEntryType:
		return bytes.NewReader(BodyEntry[BytesEntryType.string()].([]byte))
	default:
		return strings.NewReader("")
	}
}
