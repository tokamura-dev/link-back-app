package usersutil

import (
	"fmt"
	"strconv"
)

/**
 * 社員IDを発番する
 **/
func GetNextEmployeeId(preEmployeeId string) string {
	strEmployeeId, err := strconv.Atoi(preEmployeeId)
	if err != nil {
		panic("文字列から数値に変換できません")
	}
	return fmt.Sprintf("%06d", strEmployeeId+1)
}
