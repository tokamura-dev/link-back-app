package stringutil

import "strings"

/**
 * 文字列が空文字かどうかを判定する
 **/
func IsEmpty(value string) bool {
	return value == "" || strings.TrimSpace(value) == "" || len(value) == 0
}

func ValidPassParam(values []string, primaryKey int) bool {
	if IsArrayEmpty(values) {
		return false
	}
	if len(values) != primaryKey {
		return false
	}
	return true
}

/**
 * 配列が空文字かどうかを判定する
 **/
func IsArrayEmpty(values []string) bool {
	if len(values) == 0 {
		return true
	}
	var isEmpty = false
	for _, v := range values {
		if IsEmpty(v) {
			isEmpty = true
			break
		}
	}
	return isEmpty
}
