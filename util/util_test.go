package util

import "testing"

func TestConstructDir(t *testing.T) {
	defer ConstructDir()()
}
