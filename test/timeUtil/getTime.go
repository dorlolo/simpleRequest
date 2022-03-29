/*
 * @FileName:   getTime.go
 * @Author:		JuneXu
 * @CreateTime:	2022/3/1 下午6:35
 * @Description:
 */

package timeUtil

import (
	"fmt"
	"time"
)

//ThisMormingTime 今天凌晨
func ThisMorming(format string) (strTime string) {
	thisTime := time.Now()
	year := thisTime.Year()
	month := MonthMap[thisTime.Month().String()]
	day := fmt.Sprintf("%02d", thisTime.Day())
	strTime = fmt.Sprintf("%v-%v-%v 00:00:00", year, month, day)
	if format != TimeFormat.Normal_YMDhms {
		t1, _ := time.ParseInLocation(TimeFormat.Normal_YMDhms, strTime, Loc)
		strTime = t1.Format(format)
	}
	return strTime
}

//ThisMorningUnix 获取当日凌晨的时间戳
func ThisMorningToUnix() int64 {
	thist := time.Now()
	zero_tm := time.Date(thist.Year(), thist.Month(), thist.Day(), 0, 0, 0, 0, thist.Location()).Unix()
	return zero_tm
}

//ThisTimeUnix 获取当前时间的时间戳
func CurrentimeToUnix() int64 {
	return time.Now().Unix()
}

//CurrenStrtime 获取当前时间字符串
func CurrenStrtime(format string) (strTime string) {
	strTime = time.Now().Format(format)
	return
}

//Currentime 获取当前时间
func Currentime() (thisTime time.Time) {
	thisTime = time.Now().In(Loc)
	return
}

//Currentime 获取当前时间
func Currentime2(format string) (strTime string) {
	strTime = time.Now().Format(format)
	return
}

//HoursAgo 若干小时之前的时间
func HoursAgo(hours time.Duration, format string) (lastTimeStr string) {
	lastStamp := time.Now().Unix() - int64((time.Hour * hours).Seconds())
	lastTime := time.Unix(lastStamp, 0).In(Loc)
	lastTimeStr = lastTime.Format(format)
	return
}
