package stringUtil

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func StringToInt(num string) (int, error) {
	n, err := strconv.Atoi(num)
	if err != nil {
		return n, err
	}
	return n, nil
}

func ParseUint(num string) int {
	parsed, err := strconv.ParseUint(num, 10, 32)
	if nil != err {
		fmt.Println("Parsing uint error: ", err)
		return 0
	}
	return int(parsed)
}

func Split(str string, delimeter string) []string {
	return strings.Split(str, delimeter)
}

func Typeof(v interface{}) string {
	switch v.(type) {
	case int:
		return "int"
	case float64:
		return "float64"
	case uint64:
		return "uint64"
	case string:
		return "string"
	case uint32:
		return "uint32"
	case uint16:
		return "uint16"
	default:
		return "unknown"
	}
}

func IsLetter(s string) bool {
	for _, r := range s {
		if !(r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
			return true
		}
	}
	return false
}

func ContainsChar(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return true
		}
	}
	return false
}
