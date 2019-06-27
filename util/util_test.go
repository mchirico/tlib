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

	if m, found := FileContents("testFile"); found {
		for _, v := range m {
			if v != "sample" {
				t.Fatalf("Internal contents not found.")
			}
		}
	} else {
		t.FailNow()
	}
}
