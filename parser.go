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
	"fmt"
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
	formWriter := multipart.NewWriter(body)
	for k, sv := range BodyEntry {
		if strings.Contains(k, FormFilePathKey.string()) {
			key := k[len(FormFilePathKey):]
			fp := sv.(string)
			filename := filepath.Base(fp)
			filePart, _ := formWriter.CreateFormFile(key, filename)
			content, err := os.ReadFile(fp)
			if err != nil {
				panic(err)
			}
			_, _ = filePart.Write(content)

			// way 2
			file, err := os.Open(fp)
			if err != nil {
				panic(err)
			}
			defer file.Close()
			fieldName := k[len(FormFilePathKey.string()):]
			formPart, err := formWriter.CreateFormFile(fieldName, filepath.Base(fp))
			if err != nil {
				panic(err)
			}
			if _, err = io.Copy(formPart, file); err != nil {
				return
			}
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