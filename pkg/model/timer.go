package model

import (
	"strconv"
	"time"
)

// 时间格式
type TimeFormat string

const (
	DATETIMEMS       TimeFormat = "2006-01-02 15:04:05.000"
	DATETIMEMS_SLASH TimeFormat = "2006/01/02 15:04:05.000"
	DATETIME         TimeFormat = "2006-01-02 15:04:05"
	DATETIME_SLASH   TimeFormat = "2006/01/02 15:04:05"
	DATE             TimeFormat = "2006-01-02"
	DATE_SLASH       TimeFormat = "2006/01/02"
	DATE_NONE        TimeFormat = "20060102"
	TIME             TimeFormat = "15:04:05"
	UNIX_S           TimeFormat = "1136185445"
	UNIX_MS          TimeFormat = "1136185445000"
)

// 对时间操作进行扩展处理
type Time time.Time

// json解析和反向解析的操作,相关的时间信息转换
func (t *Time) UnmarshalJSON(data []byte) error {
	dataStrTemp := string(data)
	if len(dataStrTemp) <= 2 {
		return nil
	}
	dataStr := dataStrTemp[1:(len(dataStrTemp) - 1)]
	lenStr := len(dataStr)
	var err error = nil
	if lenStr == len(DATETIMEMS) {
		now, err := time.ParseInLocation(string(DATETIMEMS), dataStr, time.Local)
		if err == nil {
			*t = Time(now)
			return nil
		}
		now, err = time.ParseInLocation(string(DATETIMEMS_SLASH), dataStr, time.Local)
		if err == nil {
			*t = Time(now)
			return nil
		}
	}
	if lenStr == len(DATETIME) {
		now, err := time.ParseInLocation(string(DATETIME), dataStr, time.Local)
		if err == nil {
			*t = Time(now)
			return nil
		}
		now, err = time.ParseInLocation(string(DATETIME_SLASH), dataStr, time.Local)
		if err == nil {
			*t = Time(now)
			return nil
		}
	}
	if lenStr == len(DATE) {
		now, err := time.ParseInLocation(string(DATE), dataStr, time.Local)
		if err == nil {
			*t = Time(now)
			return nil
		}
		now, err = time.ParseInLocation(string(DATE_SLASH), dataStr, time.Local)
		if err == nil {
			*t = Time(now)
			return nil
		}
	}

	if lenStr == len(DATE_NONE) {
		now, err := time.ParseInLocation(string(DATE_NONE), dataStr, time.Local)
		if err == nil {
			*t = Time(now)
			return nil
		}
		now, err = time.ParseInLocation(string(DATE_SLASH), dataStr, time.Local)
		if err == nil {
			*t = Time(now)
			return nil
		}
	}
	if lenStr == len(UNIX_S) {
		b, err := strconv.ParseInt(dataStr, 10, 64)
		if nil == err {
			now := time.Unix(b, 0)
			*t = Time(now)
			return nil
		}
	}
	if lenStr == len(UNIX_MS) {
		b, err := strconv.ParseInt(dataStr, 10, 64)
		if nil == err {
			now := time.UnixMilli(b)
			*t = Time(now)
			return nil
		}
	}
	return err
}

func (t *Time) MarshalJSON() ([]byte, error) {
	b := time.Time(*t).UnixMilli()
	s := strconv.FormatInt(b, 10)
	return []byte(s), nil
}

func (t Time) String() string {
	return time.Time(t).Format(string(DATETIME))
}

func (t Time) DateTime() string {
	return time.Time(t).Format(string(DATETIME))
}

func (t Time) Date() string {
	return time.Time(t).Format(string(DATE))
}

func (t Time) Time() string {
	return time.Time(t).Format(string(TIME))
}

func (t Time) Format(format string) string {
	return time.Time(t).Format(format)
}

func TransTime2Time(time time.Time) *Time {
	timer := Time(time)
	return &timer
}
