package util

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Tlib struct {
	findFunc func(substr string, pwd string) bool
	mockdir  string `default:"../test-fixtures"`
	subdir   string
}

func NewTlib(t ...*Tlib) *Tlib {

	if t == nil {
		tlib := &Tlib{findFunc: FindFile, mockdir: "../test-fixtures", subdir: "default"}
		return tlib

	}

	if t[0].subdir == "" {
		t[0].subdir = "default"
	}

	if t[0].findFunc == nil {
		t[0].findFunc = FindFile
	}

	if t[0].mockdir == "" {
		t[0].mockdir = "../test-fixtures"
	}

	if len(t) > 1 {
		log.Fatalf("len(t) > 1")
	}

	return t[0]

}

func WriteString(file string, string string, perm os.FileMode) {
	data := []byte(string)
	err := ioutil.WriteFile(file, data, perm)
	if err != nil {
		log.Fatalf("err: %s\n", err)
	}
}

func Mkdir(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, os.ModePerm)
		return nil
	}
	return fmt.Errorf("Problem in pkg.Mkdir. Could not create: %s\n", path)
}

func Rmdir(path string) {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		os.RemoveAll(path)
	}
}

func ListFiles(pwd string) []string {
	var files []string

	err := filepath.Walk(pwd, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	return files
}

func FindFile(substr string, pwd string) bool {
	files := ListFiles(pwd)
	for _, v := range files {
		if strings.Contains(v, substr) {
			return true
		}
	}
	return false

}

func PWD() string {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}
	return pwd
}

func (t *Tlib) ConstructDir() func() {

	old, err := os.Getwd()
	if err != nil {
		log.Fatalf("can't get current dir: %s\n", err)
	}

	mockdir := filepath.Join(t.mockdir, t.subdir)
	err = Mkdir(mockdir)
	if err != nil {
		log.Fatalf("ConstructDir Failed: %s\n", err)
	}
	os.Chdir(mockdir)

	return func() {
		os.Chdir(old)
		c, _ := os.Getwd()
		fmt.Printf("current: %s\n", c)

		err := os.Chdir(t.mockdir)
		if err != nil {
			log.Fatalf("can't cd")
		}

		if t.subdir != "" {
			Rmdir(t.subdir)
		}

		os.Chdir(old)

	}
}
