package util

import "strings"

// withMessage, withStack 같은 타입들에서 가장 바깥 에러메시지만 얻음.
func GetSimpleErrorMessage(err error) string {
	if err == nil {
		return ""
	}

	splitted := strings.Split(err.Error(), ": ")
	return splitted[0]
}
