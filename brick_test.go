package main

import (
	"testing"
)

// TestPackageNameFromDirectory
func TestPackageNameFromDirectory(t *testing.T) {
	var tests = []struct {
		dir          string
		expectedName string
	}{
		{"test", "test"},
		{"/home/sph/foo", "foo"},
		{"/home/sph/foo/", "foo"},
		{"some-package", "some-package"},
		{"/dir/some-package", "some-package"},
	}

	for _, test := range tests {
		name, err := packageNameFromDirectory(test.dir)
		if err != nil {
			t.Error(err)
			continue
		}

		if name != test.expectedName {
			t.Errorf("packageNameFromDirectory: expected name %q, got %q\n", test.expectedName, name)
		}
	}
}
