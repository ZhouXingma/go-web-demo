package model

import (
	"database/sql"
	"time"
)

func NewSqlNullString(s *string) sql.NullString {
	var value = ""
	var valid = false
	if nil != s {
		value = *s
		valid = true
	}
	return sql.NullString{
		String: value,
		Valid:  valid,
	}
}

func NewSqlNullByte(s *byte) sql.NullByte {
	var value byte = 0
	var valid = false
	if nil != s {
		value = *s
		valid = true
	}
	return sql.NullByte{
		Byte:  value,
		Valid: valid,
	}
}

func NewSqlNullInt16(s *int16) sql.NullInt16 {
	var value int16 = 0
	var valid = false
	if nil != s {
		value = *s
		valid = true
	}
	return sql.NullInt16{
		Int16: value,
		Valid: valid,
	}
}

func NewSqlNullInt32(s *int32) sql.NullInt32 {
	var value int32 = 0
	var valid = false
	if nil != s {
		value = *s
		valid = true
	}
	return sql.NullInt32{
		Int32: value,
		Valid: valid,
	}
}

func NewSqlNullInt64(s *int64) sql.NullInt64 {
	var value int64 = 0
	var valid = false
	if nil != s {
		value = *s
		valid = true
	}
	return sql.NullInt64{
		Int64: value,
		Valid: valid,
	}
}

func NewSqlNullFloat64(s *float64) sql.NullFloat64 {
	var value float64 = 0.0
	var valid = false
	if nil != s {
		value = *s
		valid = true
	}
	return sql.NullFloat64{
		Float64: value,
		Valid:   valid,
	}
}

func NewSqlNullBool(s *bool) sql.NullBool {
	var value = false
	var valid = false
	if nil != s {
		value = *s
		valid = true
	}
	return sql.NullBool{
		Bool:  value,
		Valid: valid,
	}
}

func NewSqlNullTime(s *time.Time) sql.NullTime {
	var value = time.Now()
	var valid = false
	if nil != s {
		value = *s
		valid = true
	}
	return sql.NullTime{
		Time:  value,
		Valid: valid,
	}
}

func NewSqlNullTimeOfTime(s *Time) sql.NullTime {
	var value = time.Now()
	var valid = false
	if nil != s {
		value = time.Time(*s)
		valid = true
	}
	return sql.NullTime{
		Time:  value,
		Valid: valid,
	}
}
