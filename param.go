/*
 * @FileName:   param.go
 * @Author:		JJXu
 * @CreateTime:	2022/3/1 下午9:07
 * @Description:
 */

package simpleRequest

import (
	"fmt"
	"net/url"
)

type QueryParams struct {
	simpleReq *SimpleRequest
}

//batch settings
func (s *QueryParams) Sets(data map[string]any) *QueryParams {
	for k, v := range data {
		s.simpleReq.queryParams.Set(k, fmt.Sprintf("%v", v))
	}
	return s
}

//single settings
func (s *QueryParams) Set(key string, value any) *QueryParams {
	s.simpleReq.queryParams.Set(key, fmt.Sprintf("%v", value))
	return s
}

//get all queryParams
func (s *QueryParams) Gets() *url.Values {
	return &s.simpleReq.queryParams
}
