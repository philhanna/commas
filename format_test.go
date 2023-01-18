package commas

import (
	"testing"
)

var (
	n        int
	expected string
	actual   string
)

func Test123456(t *testing.T) {
	n = 123456
	expected = "123,456"
	actual = Format(n)
	if expected != actual {
		t.Error()
	}
}

func Test0(t *testing.T) {
	n = 0
	expected = "0"
	actual = Format(n)
	if expected != actual {
		t.Error()
	}
}

func Test1000000(t *testing.T) {
	n = 1000000
	expected = "1,000,000"
	actual = Format(n)
	if expected != actual {
		t.Error()
	}
}

func TestNegative1000(t *testing.T) {
	n = -1000
	expected = "-1000"
	actual = Format(n)
	if expected != actual {
		t.Error("expected=", expected, ", actual=", actual)
	}
}
