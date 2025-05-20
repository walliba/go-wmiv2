package types

import (
	"fmt"
	"unsafe"
)

// TODO: translate to a time.Time

type Timestamp struct {
	Year         uint32
	Month        uint32
	Day          uint32
	Hour         uint32
	Minute       uint32
	Second       uint32
	Microseconds uint32
	Utc          int32
}

func (ts Timestamp) String() string {
	return fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d",
		ts.Year, ts.Month, ts.Day, ts.Hour, ts.Minute, ts.Second)
}

type Interval struct {
	Days         uint32
	Hours        uint32
	Minutes      uint32
	Seconds      uint32
	Microseconds uint32
	__padding1   uint32
	__padding2   uint32
	__padding3   uint32
}

func (iv Interval) String() string {
	return fmt.Sprintf("%d days, %02d:%02d:%02d.%06d", iv.Days, iv.Hours, iv.Minutes, iv.Seconds, iv.Microseconds)
}

type DateTime struct {
	IsTimestamp uint32
	raw         [32]byte
}

func (dt DateTime) String() string {
	if dt.IsTimestamp == 1 {
		ts := (*Timestamp)(unsafe.Pointer(&dt.raw[0]))
		return ts.String()
	} else {
		iv := (*Interval)(unsafe.Pointer(&dt.raw[0]))
		return iv.String()
	}
}
