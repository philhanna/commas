package commas

import "fmt"

type Number interface {
	int | int64 | uint64
}

// Format return the comma-separated value of a int or int64
func Format[V Number](n V) string {
	s := fmt.Sprint(n)
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
