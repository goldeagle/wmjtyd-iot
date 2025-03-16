package utils

import (
	"math/rand"
	"strings"
	"time"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var digits = "0123456789"
var all = "ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
	"abcdefghijklmnopqrstuvwxyz" +
	digits

// 随机长度的字符串
func RandString(length int) string {
	rand.Seed(time.Now().UnixNano())

	buf := make([]byte, length)
	buf[0] = digits[rand.Intn(len(digits))]
	for i := 1; i < length; i++ {
		buf[i] = all[rand.Intn(len(all))]
	}
	rand.Shuffle(len(buf), func(i, j int) {
		buf[i], buf[j] = buf[j], buf[i]
	})
	return string(buf) // E.g. "3i[g0|)z"
}

// SnakeToCamel 将蛇形命名转换为驼峰命名
func SnakeToCamel(s string) string {
	var result string
	words := strings.Split(s, "_")
	for i, word := range words {
		if i == 0 {
			result += word
		} else {
			result += cases.Title(language.English).String(word)
		}
	}
	return result
}

// CamelToSnake 将驼峰命名转换为蛇形命名
func CamelToSnake(s string) string {
	var result string
	for i, char := range s {
		if i > 0 && unicode.IsUpper(char) {
			result += "_"
		}
		result += string(unicode.ToLower(char))
	}
	return result
}
