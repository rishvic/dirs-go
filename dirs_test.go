package dirs

import "testing"

func TestCaching(t *testing.T) {
	// Explicitly enable caching
	DisableCache = false

	testcases := []struct{ s1, s2 string }{
		{"1", "2"},
		{"Lorem ipsum", "dolor sit amet"},
		{"Is this the real life?", "Is this just fantasy?"},
	}

	for _, tc := range testcases {
		var c cache
		// Load tc.s1 into cache
		s1, err := c.cur(mockF(tc.s1))
		if err != nil || s1 != tc.s1 {
			t.Errorf("Dir: (%q, %v), want (%q, nil)", s1, err, tc.s1)
		}

		// If tc.s2 is different from tc.s1, check if cache uses previous data
		if tc.s2 == tc.s1 {
			t.Log("testcase's s1 == s2, continuing...")
			continue
		}
		s2, err := c.cur(mockF(tc.s2))
		if err != nil || s2 != tc.s1 {
			t.Errorf("Cached Dir: (%q, %v), want (%q, nil)", s2, err, tc.s1)
		}
	}
}

func TestDisabledCaching(t *testing.T) {
	// Explicitly disable caching
	DisableCache = true

	testcases := []struct{ s1, s2 string }{
		{"1", "2"},
		{"Lorem ipsum", "dolor sit amet"},
		{"Is this the real life?", "Is this just fantasy?"},
	}

	for _, tc := range testcases {
		var c cache
		// Load tc.s1 into cache
		s1, err := c.cur(mockF(tc.s1))
		if err != nil || s1 != tc.s1 {
			t.Errorf("Dir: (%q, %v), want (%q, <nil>)", s1, err, tc.s1)
		}

		// If tc.s2 is different from tc.s1, check if cache uses new data
		if tc.s2 == tc.s1 {
			t.Log("testcase's s1 == s2, continuing...")
			continue
		}
		s2, err := c.cur(mockF(tc.s2))
		if err != nil || s2 != tc.s2 {
			t.Errorf("Cached Dir: (%q, %v), want (%q, <nil>)", s2, err, tc.s2)
		}
	}
}

// Mock function that returns the passed string as return value of function
func mockF(str string) func() (string, error) {
	return func() (string, error) { return str, nil }
}

// Mock function that always returns an empty string and error
func mockErrF() (string, error) { return "", ErrNotFound }

// Mock function that returns string, but an error along with it
func mockWeirdF(str string) func() (string, error) {
	return func() (string, error) { return str, ErrNotFound }
}
