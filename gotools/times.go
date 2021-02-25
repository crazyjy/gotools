/**********

#### times 常用时间方法

```
* NowDataTime //获取东八区当前时间 Y-m-d H:i:s
* NowDateTimeMap //获取东八区当前时间map Y m d H i s
* NowDateTimeArr //获取东八区当前时间 时，分，秒 ["23","05","40"]
* MonthRange //获取当月日期范围 ["2021-02-01", "2021-02-28"]
* LostMonth //获取前几个月日期
* FutureMonth //获取后几个月日期
* LostDay //获取前几天日期
* FutureDay //获取后几天日期
* LostTimeS //获取前几秒时间
* FutureTimeS //获取后几秒时间
```

*****/

package gotools

import (
	"fmt"
	"strconv"
	"time"
)

/**
 * 格式化为东八区时间
 */
func initTime() time.Time {
	cstZone := time.FixedZone("CST", 8*3600)  // 东八
	t := time.Now().In(cstZone)
	return t
}

/**
 * 获取东八区当前时间 Y-m-d H:i:s
 */
func NowDataTime() string {
	return initTime().Format("2006-01-02 15:04:05")
}

/**
 * 获取东八区当前时间map Y m d H i s
 */
func NowDateTimeMap() map[string]string {
	t := initTime()
	data := make(map[string]string, 6)
	data["Y"] = fmt.Sprintf("%v", t.Year())
	data["m"] = fmt.Sprintf("%v", t.Format("01"))
	data["d"] = fmt.Sprintf("%v", t.Day())
	data["H"] = fmt.Sprintf("%v", t.Format("15"))
	data["i"] = fmt.Sprintf("%v", t.Format("04"))
	data["s"] = fmt.Sprintf("%v", t.Format("05"))
	return data
}

/**
 * 获取东八区当前时间 时，分，秒 ["23","05","40"]
 */
func NowDateTimeArr() [6]string {
	data := [6]string{"0000","00","00","00","00","00"}
	t := initTime()
	data[0] = fmt.Sprintf("%v", t.Year())
	data[1] = fmt.Sprintf("%v", t.Format("01"))
	data[2] = fmt.Sprintf("%v", t.Day())
	data[3] = fmt.Sprintf("%v", t.Hour())
	data[4] = fmt.Sprintf("%v", t.Minute())
	data[5] = fmt.Sprintf("%v", t.Second())
	if len(data[0]) == 1 {
		data[0] = "0"+data[0]
	}
	if len(data[1]) == 1 {
		data[1] = "0"+data[1]
	}
	if len(data[2]) == 1 {
		data[2] = "0"+data[2]
	}
	return data
}

/**
 * 获取当月日期范围
 */
func MonthRange() [2]string {
	now := initTime()
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()

	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)

	beginDate := fmt.Sprintf("%v", firstOfMonth)
	endData := fmt.Sprintf("%v", lastOfMonth)

	return [2]string{beginDate[:10], endData[:10]}
}

/**
 * 获取前几个月日期
 */
func LostMonth(i int) string {
	curTime := initTime()
	lostTime := fmt.Sprintf("%v", curTime.AddDate(0, -i, 0))
	return lostTime[:10]
}

/**
 * 获取后几个月日期
 */
func FutureMonth(i int) string  {
	curTime := initTime()
	futureTime := fmt.Sprintf("%v", curTime.AddDate(0, +i, 0))
	return futureTime[:10]
}

/**
 * 获取前几天日期
 */
func LostDay(i int) string {
	curTime := initTime()
	lostTime := fmt.Sprintf("%v", curTime.AddDate(0, 0, -i))
	return lostTime[:10]
}

/**
 * 获取后几天日期
 */
func FutureDay(i int) string  {
	curTime := initTime()
	futureTime := fmt.Sprintf("%v", curTime.AddDate(0, 0, +i))
	return futureTime[:10]
}

/**
 * 获取前几秒时间
 */
func LostTimeS(i int) string {
	st, _ := time.ParseDuration("-" + strconv.Itoa(i) + "s")
	lostTime := fmt.Sprintf("%v", initTime().Add(st))
	return lostTime[:19]
}

/**
 * 获取后几秒时间
 */
func FutureTimeS(i int) string  {
	st, _ := time.ParseDuration(strconv.Itoa(i) + "s")
	futureTime:= fmt.Sprintf("%v", initTime().Add(st))
	return futureTime[:19]
}