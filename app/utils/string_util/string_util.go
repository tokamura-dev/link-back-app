package stringutil

import "strings"

/**
 * 文字列が空文字かどうかを判定する
 **/
func IsEmpty(value string) bool {
	return value == "" || strings.TrimSpace(value) == "" || len(value) == 0
}
