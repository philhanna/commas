package commas

import "fmt"

func Format(d int) string {
	s := fmt.Sprint(d)
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
