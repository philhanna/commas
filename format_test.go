// test cases for Format - all supported types
package commas

import (
	"testing"
)

func TestFormat(t *testing.T) {
	testCases := []struct {
		number string
		want   string
	}{
		{Format(0), "0"},
		{Format(int(1000)), "1,000"},
		{Format(int(123456)), "123,456"},
		{Format(int(-1234567)), "-1,234,567"},
		{Format(int8(-40)), "-40"},
		{Format(int8(127)), "127"},
		{Format(int16(22222)), "22,222"},
		{Format(int16(-32767)), "-32,767"},
		{Format(int32(1000000)), "1,000,000"},
		{Format(int32(-123456)), "-123,456"},
		{Format(int64(1000000)), "1,000,000"},
		{Format(int64(-123456)), "-123,456"},
		{Format(uint(123456)), "123,456"},
		{Format(uint8(1<<8 - 1)), "255"},
		{Format(uint16(1<<16 - 1)), "65,535"},
		{Format(uint32(1<<32 - 1)), "4,294,967,295"},
		{Format(uint64(1<<64 - 1)), "18,446,744,073,709,551,615"},
	}
	for i, tc := range testCases {
		have := tc.number
		want := tc.want
		if have != want {
			t.Errorf("test %d: have=%s, want=%s", i, have, want)
		}
	}
}
