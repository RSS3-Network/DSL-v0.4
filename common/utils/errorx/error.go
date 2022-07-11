package errorx

import "errors"

func IsExpectedError(err error, expectedErrors ...error) bool {
	return !IsUnexpectedError(err, expectedErrors...)
}

func IsUnexpectedError(err error, expectedErrors ...error) bool {
	for _, expectedError := range expectedErrors {
		if errors.Is(err, expectedError) {
			return false
		}
	}

	return true
}
