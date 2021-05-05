package fileutil

import (
	"io/ioutil"
	"os"
)

func ReadFile(path string) string {
	return string(ReadFileBytes(path))
}

func ReadFileBytes(path string) []byte {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	// 要记得关闭
	defer f.Close()
	byteValue, _ := ioutil.ReadAll(f)

	return byteValue
}
