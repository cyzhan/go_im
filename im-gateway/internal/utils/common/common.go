package common

import "time"

var (
	timeZone *time.Location
)

func init() {
	timeZone, _ = time.LoadLocation("America/Grenada")
}

func UtcMinus4() string {
	return time.Now().In(timeZone).Format("2006-01-02 15:04:05")
}

func ParseBoolToInt32(b bool) int32 {
	if b {
		return 1
	}

	return 0
}

func RemoveDuplicate[T string | int | int32 | int64](sliceList []T) []T {
	allKeys := make(map[T]bool)
	list := make([]T, 0, len(sliceList))
	for _, item := range sliceList {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}
