package fs

import (
	"path"
	"testing"
)

func TestMem(t *testing.T) {
	for _, tc := range []struct {
		desc  string
		user  string
		mount string
	}{
		{
			desc:  "create",
			user:  "nico",
			mount: "/Users/anicolaspp/vfs/",
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			fs, err := NewMemFileSystem(tc.mount, tc.user)
			if err != nil {
				t.Error(err)
			}

			want := path.Join(tc.mount, tc.user)
			if fs.Root() != want {
				t.Errorf("Wrong fs rool, got = %v, want = %v", fs.Root(), want)
			}

			want = tc.mount
			if fs.MountPoint() != want {
				t.Errorf("Wrong fs mountpoint, got = %v, want = %v", fs.MountPoint(), want)
			}
		})
	}
}

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
