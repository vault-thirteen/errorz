package errorz

import (
	"errors"
)

const ErrorsSeparator = "; "

func Combine(
	error1 error,
	error2 error,
) error {

	if error1 == nil {
		if error2 == nil {
			return nil
		} else {
			return error2
		}
	} else {
		if error2 == nil {
			return error1
		} else {
			return errors.New(error1.Error() + ErrorsSeparator + error2.Error())
		}
	}
}
