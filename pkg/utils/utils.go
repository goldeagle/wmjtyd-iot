package utils

import (
	"bytes"
	"math"
	"reflect"
	"strconv"
	"time"
)

// StringBytesBufferJoin : 拼接字符串 bytes.Buffer模式
func StringBytesBufferJoin(con ...string) string {
	stringBytesBuffer := bytes.Buffer{}
	for _, s := range con {
		stringBytesBuffer.WriteString(s)
	}
	return stringBytesBuffer.String()
}

// 时间转13位时间戳 时间戳（毫秒）
func TimeToUnix(e time.Time) int64 {
	return e.UnixNano() / 1e6
}

// 13位时间戳转时间格式
func UnixToTime(e string) (datatime time.Time, err error) {
	data, err := strconv.ParseInt(e, 10, 64)
	datatime = time.Unix(data/1000, 0)
	return
}

// 时间转时间戳
func Strtime2Int(datetime string) int64 {
	//日期转化为时间戳
	timeLayout := "2006-01-02 15:04:05" //转化所需模板
	tmp, _ := time.ParseInLocation(timeLayout, datetime, time.Local)
	timestamp := tmp.Unix() //转化为时间戳 类型是int64
	return timestamp
}

// 核心方法利用reflect.Typeof(mm).Kind()
// 利用反射先遍历slice的值，再进行类型转换
func Map2Array(m map[string]interface{}) []map[string]interface{} {
	var list []map[string]interface{}
	if reflect.TypeOf(m).Kind() == reflect.Slice {
		s := reflect.ValueOf(m)
		for i := 0; i < s.Len(); i++ {
			ele := s.Index(i)
			list = append(list, ele.Interface().(map[string]interface{}))
		}
	}
	return list
}

// 将大切片按指定长度切割成小切片
func SpiltIntList(list []int, size int) [][]int {
	lens := len(list)
	mod := math.Ceil(float64(lens) / float64(size))
	spliltList := make([][]int, 0)
	for i := 0; i < int(mod); i++ {
		tmpList := make([]int, 0, size)
		if i == int(mod)-1 {
			tmpList = list[i*size:]
		} else {
			tmpList = list[i*size : i*size+size]
		}
		spliltList = append(spliltList, tmpList)
	}
	return spliltList
}

// 将大切片按指定长度切割成小切片
func SpiltStringList(list []string, size int) [][]string {
	lens := len(list)
	mod := math.Ceil(float64(lens) / float64(size))
	spliltList := make([][]string, 0)
	for i := 0; i < int(mod); i++ {
		tmpList := make([]string, 0, size)
		if i == int(mod)-1 {
			tmpList = list[i*size:]
		} else {
			tmpList = list[i*size : i*size+size]
		}
		spliltList = append(spliltList, tmpList)
	}
	return spliltList
}
