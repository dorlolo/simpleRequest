/*
 * @FileName:   dataTest.go
 * @Author:		JuneXu
 * @CreateTime:	2022/3/24 下午9:18
 * @Description:
 */

package test

import (
	"encoding/json"
	"testing"
)

func TestData(t *testing.T) {
	type dt map[string][]string
	var newData = dt{
		"aa": []string{"aa", "cc"},
	}
	res, err := json.Marshal(&newData)
	t.Log(err)
	t.Log(res)
}
