package jtime

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

type Time struct {
	time.Time
}

func New(t time.Time) Time {
	return Time{
		Time: t,
	}
}

func Now() Time {
	return Time{
		Time: time.Now(),
	}
}

func (j Time) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", j.Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}

func (j *Time) UnmarshalJSON(b []byte) error {
	s := strings.ReplaceAll(string(b), "\"", "")
	t, err := time.Parse("2006-01-02 15:04:05", s)
	if err != nil {
		return err
	}
	*j = New(t)
	return nil
}

func (j Time) Value() (driver.Value, error) {
	var zeroTime time.Time
	if j.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return j.Time, nil
}

func (j *Time) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*j = Time{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
