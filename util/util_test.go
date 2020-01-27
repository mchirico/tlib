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

	defer NewTlib(&Tlib{SubDir: "SubDir"}).ConstructDir()()

	pwd := PWD()

	WriteString("testFile", "sample", 0644)

	if !FindFile("SubDir/testFile", pwd) {
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

func TestAppend(t *testing.T) {

	defer NewTlib().ConstructDir()()

	pwd := PWD()

	WriteString("testFile", "sample", 0644)
	AppendString("testFile", " modified")

	if !FindFile("default/testFile", pwd) {
		t.FailNow()
	}

	if m, found := FileContents("testFile"); found {
		for _, v := range m {
			if v != "sample modified" {
				t.Fatalf("Internal contents not found.")
			}
		}
	} else {
		t.FailNow()
	}

}
