/*
 * @FileName:   base64_test.go
 * @Author:		JuneXu
 * @CreateTime:	2022/3/24 上午12:03
 * @Description:
 */

package test

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestBs(t *testing.T) {
	authStr := fmt.Sprintf("%v:%v", "aaa", "bbb")
	data := base64.StdEncoding.EncodeToString([]byte(authStr))
	t.Log(data)

} //YWFhOmJiYg==
