package db

import (
	"fmt"
	"time"
)

type Duration time.Duration

func NewDuration(from time.Duration) Duration {
	return Duration(from)
}

func (d Duration) ToDB() (value any, err error) {
	value = int64(time.Duration(d) / time.Millisecond)

	return
}

func (d *Duration) FromDB(bytes []byte) (err error) {
	var ms int64
	if len(bytes) == 0 {
		*d = 0
	} else if _, se := fmt.Sscan(string(bytes), &ms); se != nil {
		err = fmt.Errorf("DurationMs.FromDB: cannot parse %q: %w", bytes, se)
	} else {
		*d = Duration(time.Duration(ms) * time.Millisecond)
	}

	return
}

func (d Duration) Duration() time.Duration {
	return time.Duration(d)
}

func (d Duration) Milliseconds() int64 {
	return int64(time.Duration(d) / time.Millisecond)
}
