package date_utils

import "time"

const timeLayout = "2006-01-02 15:04:05"

func GetNowString() string {
	return time.Now().Format(timeLayout)
}

func GetNowSEpoch() int64 {
	return time.Now().Unix()
}
