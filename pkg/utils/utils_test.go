package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//TestGetHourRangeFromUnixTimestamp ...
func TestGetHourRangeFromUnixTimestamp(t *testing.T)  {
	testUnix := 1603114500
	expectedStart := "2020-10-19 15:00:00"
	expectedEnd := "2020-10-19 15:59:59"
	
	start,end := GetDateRangeFromTimestamp(int64(testUnix))
	assert.Equal(t, expectedStart, start, "Start datetime")
	assert.Equal(t, expectedEnd, end, "End datetime")
}




