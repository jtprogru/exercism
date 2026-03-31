package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	const myFormat = "1/02/2006 15:04:05"
	t, _ := time.Parse(myFormat, date)
	return t
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	const myFormat = "January 2, 2006 15:04:05"
	d, _ := time.Parse(myFormat, date)
	return d.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	const myFormat = "Monday, January 2, 2006 15:04:05"
	d, _ := time.Parse(myFormat, date)
	if d.Hour() == 12 {
		return true
	} else {
		return false
	}
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	const myFormat = "1/2/2006 15:04:05"
	t, _ := time.Parse(myFormat, date)

	return fmt.Sprintf(
		"You have an appointment on %s, %s %v, %v, at %v:%v.",
		t.Weekday(), t.Month(), t.Day(),
		t.Year(), t.Hour(), t.Minute())
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}
