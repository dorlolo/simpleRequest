/*
 * @FileName:   simpleRequest_test.go
 * @Author:		JuneXu
 * @CreateTime:	2022/3/3 下午11:34
 * @Description:
 */

package test

import (
	"fmt"
	"github.com/dorlolo/simpleRequest"

	"testing"
	"time"
)

func TestRequest(t *testing.T) {
	var r = simpleRequest.NewRequest()
	//---设置请求头
	r.Headers().Set("token", "d+jfdji*D%1=")
	//串联使用示例：设置Conent-Type为applicaiton/json 并且 随机user-agent
	r.Headers().ConentType_json().SetRandomUerAgent()

	//设置params
	r.QueryParams().Set("user", "dorlolo")
	//支持一次性添加,不会覆盖上面user
	pamarsBulid := make(map[string]interface{})
	pamarsBulid["passwd"] = "123456"
	pamarsBulid["action"] = "login"
	r.QueryParams().Sets(pamarsBulid)

	//--添加body
	r.Body().Set("beginDate", "2022-03-01").Set("endDate", "2022-03-03")

	//--其它请求参数
	r.TimeOut(time.Second * 30) //请求超时,默认7秒
	r.SkipCertVerify()          //跳过证书验证

	//--发送请求,这里返回的直接是body中的数据，等后续增加功能
	res, err := r.Get("www.webSite.com/end/point")
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(res)
	}

}
