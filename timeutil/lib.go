package timeutil

import (
	"strconv"
	"strings"
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

// 获取当前时间 格式如: 20200104234744
func getTime() string {
	return Format("%Y%m%d%X")
}

// 获取几天前的日期
func daysAfter(days int) string {
	return ""
}

// 获取今天日期的字符串
func getTodayStr() string {
	return Format("%Y%m%d")
}

func Format(format string) string {
	format = strings.Replace(format, "%Y", "2006", -1)
	format = strings.Replace(format, "%y", "06", -1)
	format = strings.Replace(format, "%m", "01", -1)
	format = strings.Replace(format, "%d", "02", -1)
	// %w
	format = strings.Replace(format, "%I", "03", -1)
	format = strings.Replace(format, "%p", "PM", -1)
	format = strings.Replace(format, "%f", ".000000", -1)

	format = strings.Replace(format, "%x", "01/02/06", -1)
	format = strings.Replace(format, "%X", "15:04:05", -1)
	format = strings.Replace(format, "%H", "15", -1)
	format = strings.Replace(format, "%M", "04", -1)
	format = strings.Replace(format, "%S", "05", -1)

	// 周几
	format = strings.Replace(format, "%a", "Mon", -1)
	format = strings.Replace(format, "%A", "Monday", -1)

	// 月份
	format = strings.Replace(format, "%b", "Jan", -1)
	format = strings.Replace(format, "%B", "January", -1)

	// %z
	format = strings.Replace(format, "%Z", "MST", -1)
	timeStr := time.Now().Format(format)

	// %j
	// %U
	// %W
	// %c
	// %G
	// %u
	// %V
	return timeStr
}
