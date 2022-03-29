/*
 * @FileName:   format.go
 * @Author:		JuneXu
 * @CreateTime:	2022/2/25 下午2:30
 * @Description:
 */

package timeUtil

import "time"

var MonthMap = make(map[string]string)

type DefineTimeFormat struct {
	//常规时间格式(日期带横杠)
	Normal_YMDhms string
	Normal_YMD    string
	Normal_hms    string
	//带斜杠的时间格式
	Slash_YMDhms string
	Slash_YMD    string
	//无间隔符
	NoSpacer_YMDhms string
	NoSpacer_YMD    string
}

var TimeFormat DefineTimeFormat
var Loc *time.Location

func init() {
	MonthMap[""] = "00"
	MonthMap["January"] = "01"
	MonthMap["February"] = "02"
	MonthMap["March"] = "03"
	MonthMap["April"] = "04"
	MonthMap["May"] = "05"
	MonthMap["June"] = "06"
	MonthMap["July"] = "07"
	MonthMap["August"] = "08"
	MonthMap["September"] = "09"
	MonthMap["October"] = "10"
	MonthMap["November"] = "11"
	MonthMap["December"] = "12"

	TimeFormat = DefineTimeFormat{
		Normal_YMDhms:   "2006-01-02 15:04:05",
		Normal_YMD:      "2006-01-02",
		Normal_hms:      "15:04:05",
		Slash_YMDhms:    "2006/01/02 15:04:05",
		Slash_YMD:       "2006/01/02",
		NoSpacer_YMDhms: "20060102150405",
		NoSpacer_YMD:    "20060102",
	}
	Loc, _ = time.LoadLocation("Asia/Shanghai")
}
