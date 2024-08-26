package utils

import "time"

var ChineseLoc = time.FixedZone("CST", 8*3600)

func FormatDate(t time.Time) string {
	return t.In(ChineseLoc).Format("2006-01-02")
}
