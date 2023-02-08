// Package simpleRequest -----------------------------
// file      : options.go
// author    : JJXu
// contact   : wavingBear@163.com
// time      : 2022/12/10 01:45:37
// -------------------------------------------
package simpleRequest

type OPTION func(r *SimpleRequest) *SimpleRequest

// OptionNewBodyEntryParser 新增或覆盖BodyEntryParser
func OptionNewBodyEntryParser(contentType string, parser IBodyEntryParser) OPTION {
	return func(r *SimpleRequest) *SimpleRequest {
		r.bodyEntryParsers[contentType] = parser
		return r
	}
}
