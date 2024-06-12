package formatTime

import (
	"database/sql"
	"database/sql/driver"
	"time"
)

const (
	YYYY_MM_DD = "2006-01-02"
)

type WrapDate sql.NullTime

func (t *WrapDate) UnmarshalJSON(data []byte) (err error) {
	if string(data) == "null" {
		t.Valid = false
		return nil
	}

	var now time.Time
	now, err = time.ParseInLocation(`"`+YYYY_MM_DD+`"`, string(data), time.Local)
	t.Valid = true
	t.Time = now
	return
}

func (t WrapDate) MarshalJSON() ([]byte, error) {
	if !t.Valid {
		return []byte("null"), nil
	}
	b := make([]byte, 0, len(YYYY_MM_DD)+2)
	b = append(b, '"')
	b = t.Time.AppendFormat(b, YYYY_MM_DD)
	b = append(b, '"')
	return b, nil
}

func (t WrapDate) String() string {
	if !t.Valid {
		return "null"
	}
	return t.Time.Format(YYYY_MM_DD)
}

// Value insert timestamp into mysql need this function.
func (t WrapDate) Value() (driver.Value, error) {
	if !t.Valid {
		return nil, nil
	}
	return t.Time, nil
}

// Scan value time.Time
func (t *WrapDate) Scan(v interface{}) error {
	return (*sql.NullTime)(t).Scan(v)
}

func NewWrapDate(t time.Time) WrapDate {
	if t.IsZero() {
		return WrapDate{Valid: false}
	}
	return WrapDate{Valid: true, Time: t}
}
