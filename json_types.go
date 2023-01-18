package vismanet

import (
	"time"
)

type Date struct {
	time.Time
}

func (d Date) MarshalSchema() string {
	return d.Time.Format("2006-01-02")
}

func (d *Date) UnmarshalCSV(s string) error {
	var err error
	if s == "" {
		return nil
	}

	d.Time, err = time.Parse("2006-01-02", s)
	return err
}

// func (d Date) MarshalCSV() (string, error) {
// 	return strconv.Itoa(v.A), nil
// }

type DateTime struct {
	time.Time
}

func (d DateTime) MarshalSchema() string {
	return d.Time.Format(time.RFC3339)
}
