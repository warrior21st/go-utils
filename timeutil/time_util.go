package timeutil

import (
	"strings"
	"time"
)

const (
	//标准时间模板字符串
	TimeLayoutStr = "2006-01-02 15:04:05"

	//标准日期模板字符串
	DateLayoutStr = "2006-01-02"
)

//是否为有效的时间
func IsValidTime(t *time.Time) bool {
	return t != nil && t.Unix() > 0
}

//unix时间戳转换为时间
func UnixSecsToTime(seconds int64) *time.Time {
	t := time.Unix(seconds, 0)

	return &t
}

//时间转为unix时间戳
func TimeToUnixSecs(t *time.Time) int64 {
	return t.Unix()
}

//utc时间转北京时间
func UTCTimeToBeijingTime(t *time.Time) *time.Time {
	t1 := t.Add(time.Hour * 8)
	return &t1
}

//转为中文字符串
func TimeToChineseString(t *time.Time) string {
	return t.Format("2006年01月02日 15点04分05秒")
}

//获取指定时间当天的开始时间
func GetTimeDayStart(t *time.Time) *time.Time {
	return StringToTime(TimeToDateString(t) + " 00:00:00")
}

//时间转字符串
func TimeToString(t *time.Time) string {
	return t.Format(TimeLayoutStr)
}

//时间转字符串
func TimeToDateString(t *time.Time) string {
	return t.Format(DateLayoutStr)
}

//字符串转时间
func StringToTime(s string) *time.Time {
	t, err := time.Parse(TimeLayoutStr, StandardTimeString(s))
	if err != nil {
		panic(err)
	}

	return &t
}

//字符串转时间（只保留日期）
func StringToDate(s string) *time.Time {
	t, err := time.Parse(TimeLayoutStr, strings.Split(StandardTimeString(s), " ")[0]+" 00:00:00")
	if err != nil {
		panic(err)
	}

	return &t
}

//将时间字符串标准化
func StandardTimeString(s string) string {
	arr := strings.Split(strings.ReplaceAll(s, "/", "-"), " ")
	dateArr := strings.Split(arr[0], "-")
	for len(dateArr[0]) < 4 {
		dateArr[0] = "0" + dateArr[0]
	}
	if len(dateArr[1]) == 1 {
		dateArr[1] = "0" + dateArr[1]
	}
	if len(dateArr[2]) == 1 {
		dateArr[2] = "0" + dateArr[2]
	}
	result := strings.Join(dateArr, "-")

	if len(arr) > 1 {
		timeArr := strings.Split(arr[1], ":")
		if len(timeArr[0]) == 1 {
			timeArr[0] = "0" + timeArr[0]
		}
		if len(timeArr[1]) == 1 {
			timeArr[1] = "0" + timeArr[1]
		}
		if len(timeArr[2]) == 1 {
			timeArr[2] = "0" + timeArr[2]
		}
		result += " " + strings.Join(timeArr, ":")
	}

	return result
}
