package fileutil

import (
	"io/fs"
	"io/ioutil"
	"os"
)

func ReadFile(path string) string {
	return string(ReadFileBytes(path))
}

func ReadFileBytes(path string) []byte {
	b, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return b
}

func WriteFile(path string, content string) {
	err := ioutil.WriteFile(path, []byte(content), fs.ModePerm)
	if err != nil {
		panic(err)
	}
}

func AppendToFile(path string, content string) {
	err := ioutil.WriteFile(path, []byte(content), fs.ModeAppend)
	if err != nil {
		panic(err)
	}
}
