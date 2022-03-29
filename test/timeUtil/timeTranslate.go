/**
 * @Author Puzzle
 * @Date 2021/11/18 1:36 下午
 **/

package timeUtil

import (
	"fmt"
	"time"
)

func GetTimestampMillisecond() int64 {
	now := time.Now()
	return now.UnixNano() / 1e6
}

func StringToTime(strTime string) (*time.Time, error) {
	const TIME_LAYOUT = "2006-01-02 15:04:05" //此时间不可更改
	timeobj, err := time.ParseInLocation(TIME_LAYOUT, strTime, Loc)
	return &timeobj, err
}

func StringToTimeWithFormat(strTime string, timeFormat string) (*time.Time, error) {
	timeobj, err := time.ParseInLocation(timeFormat, strTime, Loc)
	return &timeobj, err
}

//去除精确时间后面的小数点
func NowTimeToTime(layout string) *time.Time {
	otime := time.Now().Format(layout) //"2006-01-02 15:04:05" and so on
	tt, _ := StringToTime(otime)
	return tt
}

// timeToString
func TimeToString(timer time.Time) string {
	//t1 := time.Now()
	t2 := timer.Format("2006-01-02 15:04:05")
	fmt.Println(t2)
	return t2
}

// todo
//func commonParse_stringToTime(timeStr string) *time.Time {
//const spaceList =[4,2,2,2,2,2]
//var timeMap struct {
//	year   string
//	month  string
//	day    string
//	hour   string
//	minute string
//	second string
//}
//
//for k, v := range timeStr {
//	fmt.Println()
//}

//测试能否被int64化，如果能够转化说明全是数字
// 替换-为""
// 替换/为""
// 替换:为""

//}
