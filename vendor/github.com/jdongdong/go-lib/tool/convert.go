package tool

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func ToString(s interface{}) string {
	switch vlu := s.(type) {
	case int:
		return strconv.Itoa(vlu)
	case int64:
		return strconv.FormatInt(vlu, 10)
	case float64:
		return fmt.Sprintf("%g", s)
	case string:
		return vlu
	}
	return ""
}

func StrToTime(s string, format string) time.Time {
	tempFormat := strings.Replace(format, "yyyy", "2006", -1)
	tempFormat = strings.Replace(format, "yy", "06", -1)
	tempFormat = strings.Replace(tempFormat, "MM", "01", -1)
	tempFormat = strings.Replace(tempFormat, "dd", "02", -1)
	tempFormat = strings.Replace(tempFormat, "HH", "15", -1)
	tempFormat = strings.Replace(tempFormat, "mm", "04", -1)
	tempFormat = strings.Replace(tempFormat, "ss", "05", -1)
	loc, _ := time.LoadLocation("Local")
	formatTime, _ := time.ParseInLocation(tempFormat, s, loc)
	return formatTime
}

func TimeToStr(t time.Time, format string) string {
	tempFormat := strings.Replace(format, "yyyy", "2006", -1)
	tempFormat = strings.Replace(tempFormat, "yy", "06", -1)
	tempFormat = strings.Replace(tempFormat, "MM", "01", -1)
	tempFormat = strings.Replace(tempFormat, "dd", "02", -1)
	tempFormat = strings.Replace(tempFormat, "HH", "15", -1)
	tempFormat = strings.Replace(tempFormat, "mm", "04", -1)
	tempFormat = strings.Replace(tempFormat, "ss", "05", -1)
	formatTime := t.Format(tempFormat)
	return formatTime
}
