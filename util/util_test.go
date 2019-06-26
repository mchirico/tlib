package util

import (
	"testing"
)

func TestConstructDirDefault(t *testing.T) {

	defer NewTlib().ConstructDir()()

	pwd := PWD()

	WriteString("testFile", "sample", 0644)

	if !FindFile("default/testFile", pwd) {
		t.FailNow()
	}

}

func TestConstructDirSubdir(t *testing.T) {

	defer NewTlib(&Tlib{subdir: "subdir"}).ConstructDir()()

	pwd := PWD()

	WriteString("testFile", "sample", 0644)

	if !FindFile("subdir/testFile", pwd) {
		t.FailNow()
	}

}
