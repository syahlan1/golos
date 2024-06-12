package formatTime

import (
	"database/sql"
	"database/sql/driver"
	"time"
)

const (
	YYYY_MM_DD_HH_MM_SS = "2006-01-02 15:04:05"
)

type WrapDateTime sql.NullTime

func (t *WrapDateTime) UnmarshalJSON(data []byte) (err error) {
	if string(data) == "null" {
		t.Valid = false
		return nil
	}

	var now time.Time
	now, err = time.ParseInLocation(`"`+YYYY_MM_DD_HH_MM_SS+`"`, string(data), time.Local)
	t.Valid = true
	t.Time = now
	return
}

func (t WrapDateTime) MarshalJSON() ([]byte, error) {
	if !t.Valid {
		return []byte("null"), nil
	}
	b := make([]byte, 0, len(YYYY_MM_DD_HH_MM_SS)+2)
	b = append(b, '"')
	b = t.Time.AppendFormat(b, YYYY_MM_DD_HH_MM_SS)
	b = append(b, '"')
	return b, nil
}

func (t WrapDateTime) String() string {
	if !t.Valid {
		return "null"
	}
	return t.Time.Format(YYYY_MM_DD_HH_MM_SS)
}

// Value insert timestamp into mysql need this function.
func (t WrapDateTime) Value() (driver.Value, error) {
	if !t.Valid {
		return nil, nil
	}
	return t.Time, nil
}

// Scan value time.Time
func (t *WrapDateTime) Scan(v interface{}) error {
	return (*sql.NullTime)(t).Scan(v)
}

func NewWrapDateTime(t time.Time) WrapDateTime {
	if t.IsZero() {
		return WrapDateTime{Valid: false}
	}
	return WrapDateTime{Valid: true, Time: t}
}
