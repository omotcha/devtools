package time

import (
	"fmt"
	"testing"
)

func TestCurrentUnixTime(t *testing.T) {
	fmt.Println("Current Unix Time: ", currentUnixTime())
}

func TestParseUnixTime(t *testing.T) {
	fmt.Println("Target Unix Time: ", YMD2UnixTime(2077, 7, 7, 7, 7, 7))
	y, mon, d, h, min, s := UnixTime2YMD(3392838427)
	fmt.Println("Target time params: ", y, mon, d, h, min, s)
}

func TestDelayedUnixTime(t *testing.T) {
	fmt.Println("Delayed Unix Time: ", delayedUnixTime("3m"))
}
