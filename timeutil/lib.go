package timeutil

import (
	"strconv"
	"time"
)

func timeStamp(level int64) string {
	microTimeStamp := time.Now().UnixNano() / level
	ts := strconv.Itoa(int(microTimeStamp))
	return ts
}

// Get the millisecond string timestamp in the format: 1578185349287
func GetMsTS() string {
	return timeStamp(1e6)
}

// Get second string timestamp in the format: 1582688231
func GetTS() string {
	return timeStamp(1e9)
}

// 获取当前日期 格式如: 20200104
func getDate() string {
	timeStr := time.Now().Format("20060102")
	return timeStr
}

// 获取当前时间 格式如: 20200104234744
func getTime() string {
	timeStr := time.Now().Format("20060102150405")
	return timeStr
}

// 获取几天前的日期
func daysAfter(days int) string {
	return ""
}

// 获取今天日期的字符串
func getTodayStr() string {
	timeStr := time.Now().Format("20060102")
	return timeStr
}
