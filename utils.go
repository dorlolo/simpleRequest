/*
 *FileName:   utils.go
 *Author:		JJXu
 *CreateTime:	2022/3/29 上午11:16
 *Description:
 */

package simpleRequest

func IsJSONType(ct string) bool {
	return jsonCheck.MatchString(ct)
}

// IsXMLType method is to check XML content type or not
func IsXMLType(ct string) bool {
	return xmlCheck.MatchString(ct)
}

func IsInArray(arr []string, str string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}
