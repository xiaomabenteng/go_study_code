package src

import (
	"gopkg.in/go-playground/validator.v9"
	"regexp"
)
//这是v8验证的方法，当前版本gin无法使用这种方式自定义验证方法
//func TopicUrl(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
//	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
//
//	fmt.Println(topStruct)
//		return false
//	}
//v9 自定义验证方法
func TopicUrl(fl validator.FieldLevel) bool {
	if _,ok:=fl.Top().Interface().(*TopicModel);ok {
		if matched,_:=regexp.MatchString(`^\w{4,10}$`,fl.Field().String());matched{
			return true
		}
	}
	return false
}
