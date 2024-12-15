package timeutil

import "time"

func GetTimeNow() time.Time {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return time.Time{}
	}

	// 現在時刻を日本時間に変換
	return time.Now().In(jst)
}
