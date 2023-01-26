package commas

import (
	"testing"
)

func TestFormat(t *testing.T) {
	type testCase struct {
		number uint64
		want   string
	}
	testCases := []testCase{
		{uint64(1<<64 - 1), "18,446,744,073,709,551,615"},
	}
	for i, tc := range testCases {
		have := Format(tc.number)
		want := tc.want
		if have != want {
			t.Errorf("test %d: have=%s, want=%s", i, have, want)
		}
	}
}

func Test123456(t *testing.T) {
	n := 123456
	want := "123,456"
	have := Format(n)
	if want != have {
		t.Errorf("have=%s, want=%s", have, want)
	}
}

func Test0(t *testing.T) {
	n := 0
	want := "0"
	have := Format(n)
	if want != have {
		t.Errorf("have=%s, want=%s", have, want)
	}
}

func Test1000000(t *testing.T) {
	n := 1000000
	want := "1,000,000"
	have := Format(n)
	if want != have {
		t.Errorf("have=%s, want=%s", have, want)
	}
}

func TestNegative1000(t *testing.T) {
	n := -1000
	want := "-1,000"
	have := Format(n)
	if want != have {
		t.Errorf("have=%s, want=%s", have, want)
	}
}

func Test1234567890(t *testing.T) {
	n64 := 1234567890
	want := "1,234,567,890"
	have := Format(n64)
	if want != have {
		t.Errorf("have=%s, want=%s", have, want)
	}
}

func Test064(t *testing.T) {
	n64 := 0
	want := "0"
	have := Format(n64)
	if want != have {
		t.Errorf("have=%s, want=%s", have, want)
	}
}
