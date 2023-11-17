// Package simpleRequest -----------------------------
// @file      : utils_test.go
// @author    : JJXu
// @contact   : wavingbear@163.com
// @time      : 2023/11/17 16:37
// -------------------------------------------
package simpleRequest

import (
	"fmt"
	"testing"
)

func Test_mapToXML(t *testing.T) {
	type args struct {
		m map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "多级xml测试",
			args: args{
				m: map[string]interface{}{
					"UserInfo": map[string]any{
						"Name":      "JJXu",
						"Age":       18,
						"isTrueMan": true,
						"assets": map[string]any{
							"car":   "BMW",
							"house": "shanghai",
						},
					},
				},
			},
			want:    "<UserInfo>\n<Age>18</Age>\n<isTrueMan>true</isTrueMan>\n<assets></assets>\n<Name>JJXu</Name>\n</UserInfo>",
			wantErr: false,
		},
		{
			name:    "错误格式测试",
			args:    args{},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := mapToXML(tt.args.m)
			if (err != nil) != tt.wantErr {
				t.Errorf("mapToXML() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if string(got) != tt.want {
				t.Errorf("mapToXML() got = %v, want %v", string(got), tt.want)
			}
		})
	}
}
func Test_mapToXML2(t *testing.T) {
	person := map[string]interface{}{
		"userInfo": map[string]interface{}{
			"name": "John",
			"age":  30,
			"address": map[string]interface{}{
				"street": "123 Main St",
				"city":   "New York",
			},
		},
	}

	xmlBytes, err := mapToXML(person)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	xmlString := string(xmlBytes)
	fmt.Println(xmlString)
}
