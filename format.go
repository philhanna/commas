package commas

import "fmt"

type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

// Format return the thousands-separated value of any of the integer numeric types
func Format[V Number](n V) string {
	// Convert the number to a string
	s := fmt.Sprint(n)

	// If it is negative, remove the leading "-" and remember that you did so
	negative := string(s[0]) == "-"
	if negative {
		s = s[1:]
	}

	// Now build up the comma-separated version
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

	// If we stripped the sign off of a negative number, add it back now
	if negative {
		sb = "-" + sb
	}

	return sb
}
