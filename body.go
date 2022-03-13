/*
 * @FileName:   body.go
 * @Author:		JuneXu
 * @CreateTime:	2022/3/2 上午1:23
 * @Description:
 */

package simpleRequest

var (
	stringBodyType = "__STRING_BODY__"
)

type BodyConf struct {
	simpleReq *SimpleRequest
}

func (s *BodyConf) Set(key string, value interface{}) *BodyConf {
	s.simpleReq.tempBody[key] = value
	return s
}
func (s *BodyConf) Sets(data map[string]interface{}) *BodyConf {
	for k, v := range data {
		s.simpleReq.tempBody[k] = v
	}
	return s
}
func (s *BodyConf) SetString(strData string) *BodyConf {
	s.simpleReq.tempBody[stringBodyType] = strData
	return s
}
