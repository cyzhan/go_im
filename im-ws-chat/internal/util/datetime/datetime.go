package datetime

import "time"

var (
	eaZone = time.FixedZone("UTC-4", -4*60*60)
)

func EestAmerica() string {
	now := time.Now().In(eaZone)
	return now.Format("2006-01-02 15:04:05")
}
