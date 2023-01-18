package commas

import (
	"testing"
)

func Test123456(t *testing.T) {
	n := 123456
	expected := "123,456"
	actual := Format(n)
	if expected != actual {
		t.Error()
	}
}
