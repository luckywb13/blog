package util

import (
	"github.com/gin-gonic/gin/binding"
	lZh "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	tZh "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strings"
)

var Validate *validator.Validate
var trans ut.Translator

func init() {
	zh := lZh.New()
	uni := ut.New(zh, zh)
	trans, _ = uni.GetTranslator("zh")

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			label := field.Tag.Get("label")
			if label == "" {
				return field.Name
			}
			return label
		})
		if err := tZh.RegisterDefaultTranslations(v, trans); err != nil {
			panic(err)
		}
	}
}

func Translate(errs validator.ValidationErrors) string {
	var errList []string
	for _, e := range errs {
		errList = append(errList, e.Translate(trans))
	}
	return strings.Join(errList, "|")
}
