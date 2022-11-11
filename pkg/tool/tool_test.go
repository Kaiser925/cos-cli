package tool

import "testing"

func TestGetByteSize(t *testing.T) {
	testcases := []struct {
		i float64
		s string
	}{
		{
			1023,
			"1023B",
		},
		{
			0,
			"0B",
		},
		{
			1024,
			"1.0K",
		},
		{
			1.1 * KB,
			"1.1K",
		},
		{
			1.1 * MB,
			"1.1M",
		},
	}

	for i, tc := range testcases {
		if got := GetByteSize(int64(tc.i)); got != tc.s {
			t.Errorf("test case %d GetByteSize(%v)=%v, want %v", i+1, tc.i, got, tc.s)
		}
	}
}
