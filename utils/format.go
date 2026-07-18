package utils

import "time"

const responseTimeLayout = "2006-01-02 15:04:05"

func FormatTimestamp(value time.Time) string {
	if value.IsZero() {
		return ""
	}
	return value.Format(responseTimeLayout)
}
