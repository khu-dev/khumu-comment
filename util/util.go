package util

import "errors"

// expectedErrs 중에 err이 해당되는 에러가 있다면 true
// 없다면 false를 리턴
func ErrorIn(err error, expectedErrs []error) bool {
	for _, expectedErr := range expectedErrs {
		if errors.Is(err, expectedErr) {
			return true
		}
	}
	return false
}
