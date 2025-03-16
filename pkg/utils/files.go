package utils

import (
	"io/ioutil"
	"os"
)

// FileGetContents 把整个文件读入一个字符串中
func FileGetContents(filename string) (string, error) {
	data, err := ioutil.ReadFile(filename)
	return string(data), err
}

// FilePutContents 把一个字符串写入文件中
func FilePutContents(filename string, data string, mode os.FileMode) error {
	return ioutil.WriteFile(filename, []byte(data), mode)
}
