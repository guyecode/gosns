package validators

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

// 验证国内手机号
var MobileNumber validator.Func = func(fl validator.FieldLevel) bool {
	number, ok := fl.Field().Interface().(string)
	if ok {
		if match, _ := regexp.MatchString(`^1[3|5|6|7|8|9]\d{9}$`, number); match {
			return match
		}
	}
	return false
}