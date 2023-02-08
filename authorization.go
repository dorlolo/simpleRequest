/*
 *FileName:   auth.go
 *Author:		JJXu
 *CreateTime:	2022/3/24 上午12:09
 *Description:
 */

package simpleRequest

import (
	"encoding/base64"
	"fmt"
)

type Authorization struct {
	simpleReq *SimpleRequest
}

// Basic
// Description: 身份验证，使用bearer 令牌bearer 令牌
// receiver s
// param username
// param password
func (s *Authorization) Bearer(token string) {
	s.simpleReq.headers.Set("Authorization", fmt.Sprintf("Bearer %v", token))
}

// Basic
// Description: 身份验证的基本验证方案
// receiver s
// param username
// param password
func (s *Authorization) Basic(username, password string) {
	authStr := fmt.Sprintf("%v:%v", username, password)
	data := base64.StdEncoding.EncodeToString([]byte(authStr))
	s.simpleReq.headers.Set("Authorization", fmt.Sprintf("Basic %v", data))
}
