/*
 * @FileName:   getTime_test.go
 * @Author:		JuneXu
 * @CreateTime:	2022/3/21 下午5:58
 * @Description:
 */

package timeUtil

import (
	"testing"
)

func TestTimett(t *testing.T) {
	t.Log(ThisMorming(TimeFormat.Normal_YMDhms))
	t.Log(Currentime())
	t.Log(Currentime2(TimeFormat.Normal_YMDhms))
	t.Log(HoursAgo(5, TimeFormat.Normal_YMDhms))
	t.Log(HoursAgo(5, TimeFormat.Normal_YMDhms))
}
