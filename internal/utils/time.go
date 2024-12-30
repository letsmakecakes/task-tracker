package utils

import "time"

// TimeFormat is the standard format used for parsing and formatting time.
const TimeFormat = "2006-01-02 15:04:05"

// FormatTime formats a given time.Time object into a string using the standard TimeFormat.
func FormatTime(t time.Time) string {
	return t.Format(TimeFormat)
}
