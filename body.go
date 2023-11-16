/*
 *FileName:   body.go
 *Author:		JJXu
 *CreateTime:	2022/3/2 上午1:23
 *Description:
 */

package simpleRequest

import (
	"mime/multipart"
)

// EntryMark 请求体条目标记，用于标记输入的body内容格式
type EntryMark string

func (b EntryMark) string() string {
	return string(b)
}

const (
	StringEntryType    EntryMark = "__STRING_ENTRY__"
	BytesEntryType     EntryMark = "__BYTES_ENTRY__"
	ModelEntryType     EntryMark = "__MODEL_ENTRY__"
	MapEntryType       EntryMark = "__MAP_ENTRY__"
	MultipartEntryType EntryMark = "__MULTIPART_ENTRY__"
	FormFilePathKey    EntryMark = "__FORM_FILE_PATH_KEY__"
)

type BodyConf struct {
	simpleReq *SimpleRequest
}

func (s *BodyConf) Set(key string, value any) *BodyConf {
	s.simpleReq.BodyEntries[key] = value
	return s
}
func (s *BodyConf) Sets(data map[string]any) *BodyConf {
	s.simpleReq.BodyEntryMark = MapEntryType
	for k, v := range data {
		s.simpleReq.BodyEntries[k] = v
	}
	return s
}
func (s *BodyConf) SetString(strData string) *BodyConf {
	s.simpleReq.BodyEntryMark = StringEntryType
	s.simpleReq.BodyEntries[StringEntryType.string()] = strData
	return s
}

func (s *BodyConf) SetBytes(byteData []byte) *BodyConf {
	s.simpleReq.BodyEntryMark = BytesEntryType
	s.simpleReq.BodyEntries[BytesEntryType.string()] = byteData
	return s
}

func (s *BodyConf) SetModel(model any) *BodyConf {
	s.simpleReq.BodyEntryMark = ModelEntryType
	s.simpleReq.BodyEntries[ModelEntryType.string()] = model
	return s
}
func (s *BodyConf) SetFromDataFile(key, filePath string) *BodyConf {
	s.simpleReq.BodyEntryMark = MultipartEntryType
	s.simpleReq.BodyEntries[FormFilePathKey.string()+key] = filePath
	if s.simpleReq.headers.Get(hdrContentTypeKey) == "" {
		s.simpleReq.headers.Set(hdrContentTypeKey, formDataType)
	}
	return s
}
func (s *BodyConf) SetFromDataMultipartFile(key string, multFile *multipart.FileHeader) *BodyConf {
	s.simpleReq.BodyEntryMark = MultipartEntryType
	s.simpleReq.BodyEntries[key] = multFile
	if s.simpleReq.headers.Get(hdrContentTypeKey) == "" {
		s.simpleReq.headers.Set(hdrContentTypeKey, formDataType)
	}
	return s
}
