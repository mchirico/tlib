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
	FindFunc func(substr string, pwd string) bool
	MockDir  string `default:"../test-fixtures"`
	SubDir   string
}

func NewTlib(t ...*Tlib) *Tlib {

	if t == nil {
		tlib := &Tlib{FindFunc: FindFile, MockDir: "../test-fixtures", SubDir: "default"}
		return tlib

	}

	if t[0].SubDir == "" {
		t[0].SubDir = "default"
	}

	if t[0].FindFunc == nil {
		t[0].FindFunc = FindFile
	}

	if t[0].MockDir == "" {
		t[0].MockDir = "../test-fixtures"
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

func AppendString(file string, string string) {
	f, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(string); err != nil {
		panic(err)
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

func ReadFile(file string) string {
	dat, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	return string(dat)
}

func FileContents(file string) (map[string]string, bool) {

	m := map[string]string{}
	found := false

	files := ListFiles(PWD())
	for _, v := range files {
		if strings.Contains(v, file) {
			contents := ReadFile(v)
			m[v] = contents
			found = true
		}
	}

	return m, found
}

func PWD() string {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return pwd
}

func (t *Tlib) ConstructDir() func() {

	old, err := os.Getwd()
	if err != nil {
		log.Fatalf("can't get current dir: %s\n", err)
	}

	mockdir := filepath.Join(t.MockDir, t.SubDir)
	err = Mkdir(mockdir)
	if err != nil {
		log.Printf("ConstructDir Failed: %s\n", err)
	}
	err = os.Chdir(mockdir)
	if err != nil {
		log.Printf("os.Chdir(%s) failed\n", mockdir)
		log.Printf("MockDir: %s\n", mockdir)
		log.Fatalf("can't Chdir. Error: %s\n", err)

	}

	return func() {
		os.Chdir(old)
		c, _ := os.Getwd()
		fmt.Printf("current: %s\n", c)

		err := os.Chdir(t.MockDir)
		if err != nil {
			log.Fatalf("can't cd")
		}

		if t.SubDir != "" {
			Rmdir(t.SubDir)
		}

		os.Chdir(old)

	}
}
