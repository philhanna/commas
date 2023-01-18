package commas

import "fmt"

// Format returns the comma-separated value of an int
func Format(d int) string {
	s := fmt.Sprint(d)
	return format(s)
}

// Format64 returns the comma-separated value of an int64
func Format64(d int64) string {
	s := fmt.Sprint(d)
	return format(s)
}

// format is an internal function used by Format and Format64
func format(s string) string {
	sb := ""
	m := 0
	for {
		if len(s) == 0 {
			break
		}
		lastChar := string(s[len(s)-1])
		s = s[0 : len(s)-1]
		if m%3 == 0 && sb != "" {
			sb = "," + sb
		}
		m += 1
		sb = lastChar + sb
	}
	return sb
}
