// Package leap implements a simple utility func
// to check leap years in Gregorian calendar.
package leap

// IsLeapYear returns true if given year was a leap year
func IsLeapYear(y int) bool {
	return y%4 == 0 && (y%100 != 0 || y%400 == 0)
}
