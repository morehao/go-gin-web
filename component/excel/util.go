package excel

import (
	"regexp"
	"strings"
)

func Trim(str string) string {
	if len(str) == 0 {
		return ""
	}
	s := strings.Replace(str, " ", "", -1)
	// 替换所有空白字符（包括空格、制表符、换行符等）
	return regexp.MustCompile(`\s`).ReplaceAllString(s, "")
}
