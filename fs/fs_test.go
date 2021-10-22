package fs

import "testing"

// This is how Go tests are structured
func TestSomethingMeaningless(t *testing.T) {
	for _, tc := range []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{
			name:     "simple sum",
			a:        5,
			b:        3,
			expected: 8,
		},
		{
			name:     "sum with 0",
			a:        5,
			b:        0,
			expected: 5,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.a + tc.b
			if got != tc.expected {
				t.Errorf("Unexpected value, Got = %v, Want = %v", got, tc.expected)
			}
		})
	}
}
