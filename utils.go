/*
 *FileName:   utils.go
 *Author:		JJXu
 *CreateTime:	2022/3/29 上午11:16
 *Description:
 */

package simpleRequest

import (
	"encoding/xml"
	"fmt"
)

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

type xmlMapEntry struct {
	XMLName xml.Name
	Value   interface{} `xml:",chardata"`
}

//	func MapToXml(data map[string]any) ([]byte, error) {
//		xmlData, err := mapToXML(data)
//		if err != nil {
//			return nil, err
//		}
//		return xml.MarshalIndent(xmlData, "", "  ")
//	}
//
//	func mapToXML(m map[string]interface{}) (xmlMap map[string]xmlMapEntry, err error) {
//		if len(m) > 1 {
//			return nil, errors.New("xml format must have a root name,the map value must like this: map[string]interface{}{\"rootName\":map[string]interface{}{}}")
//		}
//		xmlMap = make(map[string]xmlMapEntry)
//		var rootName string
//		for root, data := range m {
//			rootName = root
//			for k, v := range data.(map[string]interface{}) {
//				switch typeV := v.(type) {
//				case map[string]interface{}:
//					subXmlMap, err := mapToXML(typeV)
//					if err != nil {
//						return
//					}
//
//				default:
//					entry := xmlMapEntry{XMLName: xml.Name{Local: k}, Value: v}
//					xmlMap[k] = entry
//				}
//			}
//		}
//
//		xmlData := struct {
//			XMLName xml.Name
//			Data    []xmlMapEntry `xml:",any"`
//		}{
//			XMLName: xml.Name{Local: rootName},
//			Data:    make([]xmlMapEntry, 0, len(xmlMap)),
//		}
//
//		for _, v := range xmlMap {
//			xmlData.Data = append(xmlData.Data, v)
//		}
//
//		return xml.MarshalIndent(xmlData, "", "  ")
//	}

func mapToXML(m map[string]interface{}) ([]byte, error) {
	xmlData := make([]xmlNode, 0)

	for k, v := range m {
		node := xmlNode{
			XMLName: xml.Name{Local: k},
		}

		switch value := v.(type) {
		case map[string]interface{}:
			childXML, err := mapToXML(value)
			if err != nil {
				return nil, err
			}
			node.Data = childXML
		default:
			node.Data = []byte(fmt.Sprintf("%v", v))
		}

		xmlData = append(xmlData, node)
	}

	return xml.MarshalIndent(xmlData, "", "  ")
}

type xmlNode struct {
	XMLName xml.Name
	Data    []byte `xml:",innerxml"`
}
