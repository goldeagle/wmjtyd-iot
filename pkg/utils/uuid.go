package utils

import (
	"strings"

	uuid "github.com/satori/go.uuid"
)

/**
 * 生成uuid
 */
func GetUuid() string {
	return uuid.NewV4().String()
}

/**
 * 将一个uuid字符串中添加连字符
 * @param string uuid
 * @return string
 */
func IncompactUuid(uuid string) string {
	r := []string{
		uuid[0:8],
		uuid[8:12],
		uuid[12:16],
		uuid[16:20],
		uuid[20:],
	}
	return strings.Join(r, "-")
}

/**
 * 去除 uuid 字符串中的连字符
 * @param string uuid
 * @return string
 */
func CompactUuid(uuid string) string {
	r := strings.Split(uuid, "-")

	return strings.Join(r, "")
}
