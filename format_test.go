package commas

import (
	"testing"
)

var (
	n        int
	n64      int64
	expected string
	actual   string
)

// Tests for int32
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
	expected = "-1,000"
	actual = Format(n)
	if expected != actual {
		t.Error("expected=", expected, ", actual=", actual)
	}
}

// Tests for int64
func Test1234567890(t *testing.T) {
	n64 = 1234567890
	expected = "1,234,567,890"
	actual = Format64(n64)
	if expected != actual {
		t.Error()
	}
}

func Test064(t *testing.T) {
	n64 = 0
	expected = "0"
	actual = Format64(n64)
	if expected != actual {
		t.Error()
	}
}
