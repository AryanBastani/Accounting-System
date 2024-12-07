package strValidate

import (
	"Final/internal/accountSys/consts"
	"errors"
)

func ValidateStr(str *string) error {
	if len(*str) == 0 {
		return errors.New(consts.EMPTY_STR)
	}
	if len(*str) > 64 {
		return errors.New(consts.LONG_STR)
	}

	return nil
}
