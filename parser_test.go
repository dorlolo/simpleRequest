// Package simpleRequest -----------------------------
// @file      : parser_test.go
// @author    : JJXu
// @contact   : wavingbear@163.com
// @time      : 2024/1/11 18:35
// -------------------------------------------
package simpleRequest

import (
	"io"
	"reflect"
	"testing"
)

func TestFormDataParser_Unmarshal(t *testing.T) {
	type fields struct {
		ContentType string
	}
	type args struct {
		bodyType  EntryMark
		BodyEntry map[string]any
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantBody io.Reader
	}{
		{
			name:   "StringEntryType",
			fields: fields{},
			args: args{
				bodyType:  StringEntryType,
				BodyEntry: map[string]any{StringEntryType.string(): "test"},
			},
			wantBody: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &FormDataParser{
				ContentType: tt.fields.ContentType,
			}
			if gotBody := f.Unmarshal(tt.args.bodyType, tt.args.BodyEntry); !reflect.DeepEqual(gotBody, tt.wantBody) {
				t.Errorf("Unmarshal() = %v, want %v", gotBody, tt.wantBody)
			}
		})
	}
}

func TestJsonParser_Unmarshal(t *testing.T) {
	type args struct {
		bodyType  EntryMark
		BodyEntry map[string]any
	}
	tests := []struct {
		name string
		args args
		want io.Reader
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			js := JsonParser{}
			if got := js.Unmarshal(tt.args.bodyType, tt.args.BodyEntry); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Unmarshal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXmlParser_Unmarshal(t *testing.T) {
	type args struct {
		bodyType  EntryMark
		BodyEntry map[string]any
	}
	tests := []struct {
		name     string
		args     args
		wantBody io.Reader
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := XmlParser{}
			if gotBody := f.Unmarshal(tt.args.bodyType, tt.args.BodyEntry); !reflect.DeepEqual(gotBody, tt.wantBody) {
				t.Errorf("Unmarshal() = %v, want %v", gotBody, tt.wantBody)
			}
		})
	}
}

func Test_multipartCommonParse(t *testing.T) {
	type args struct {
		BodyEntry map[string]any
	}
	tests := []struct {
		name            string
		args            args
		wantReader      io.Reader
		wantContentType string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotReader, gotContentType := multipartCommonParse(tt.args.BodyEntry)
			if !reflect.DeepEqual(gotReader, tt.wantReader) {
				t.Errorf("multipartCommonParse() gotReader = %v, want %v", gotReader, tt.wantReader)
			}
			if gotContentType != tt.wantContentType {
				t.Errorf("multipartCommonParse() gotContentType = %v, want %v", gotContentType, tt.wantContentType)
			}
		})
	}
}
