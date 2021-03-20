package utils

import (
	"math"
	"time"
)

//GetDateRangeFromTimestamp ...
func GetDateRangeFromTimestamp(ts int64) (start, end string) {
	timunx := time.Unix(ts, 0)
	start = timunx.Format("2006-01-02 15") + ":00:00"
	end = timunx.Format("2006-01-02 15") + ":59:59"
	return
}

//WeiToEther ...
func WeiToEther(wei float64) (eth float32) {
	return float32(wei / math.Pow10(18))
}

//TimestampToDateTime ...
func TimestampToDateTime(ts int64) string {
	timunx := time.Unix(ts, 0)
	return timunx.Format("2006-01-02 15:04:05")
}
