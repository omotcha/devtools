package time

import "time"

func currentUnixTime() int64 {
	now := time.Now().Unix()
	return now
}

func YMD2UnixTime(year int, month int, day int, hour int, minute int, second int) int64 {
	ret := time.Date(year, time.Month(month), day, hour, minute, second, 0, time.Local).Unix()
	return ret
}

func delayedUnixTime(delay string) int64 {
	now := time.Now()
	dur, _ := time.ParseDuration(delay)
	ret := now.Add(dur)
	return ret.Unix()
}

func UnixTime2YMD(ut int64) (year int, month int, day int, hour int, minute int, second int) {
	t := time.Unix(ut, 0)
	return t.Year(), int(t.Month()), t.Day(), t.Hour(), t.Minute(), t.Second()
}
