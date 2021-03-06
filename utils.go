/*
 * @FileName:   utils.go
 * @Author:		JuneXu
 * @CreateTime:	2022/3/29 上午11:16
 * @Description:
 */

package simpleRequest

func IsJSONType(ct string) bool {
	return jsonCheck.MatchString(ct)
}

// IsXMLType method is to check XML content type or not
func IsXMLType(ct string) bool {
	return xmlCheck.MatchString(ct)
}
