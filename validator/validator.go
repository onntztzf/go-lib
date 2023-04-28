package validator

import (
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
	trans    ut.Translator
)

func Init() error {
	zhTranslator := zh.New()
	uni = ut.New(zhTranslator, zhTranslator)
	trans, _ = uni.GetTranslator("zh")
	validate = binding.Validator.Engine().(*validator.Validate)
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		s := fld.Name
		if f := fld.Tag.Get("form"); len(f) > 0 {
			s = f
		} else if j := fld.Tag.Get("json"); len(j) > 0 {
			s = j
		} else if u := fld.Tag.Get("uri"); len(u) > 0 {
			s = u
		}
		return s
	})
	err := zhTranslations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		return err
	}
	return nil
}

func TranslateValidatorError(err error) string {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		result := validationErrors.Translate(trans)
		errsMsgs := make([]string, 0, len(result))
		for _, v := range result {
			errsMsgs = append(errsMsgs, v)
		}
		return strings.Join(errsMsgs, ",")
	}
	return err.Error()
}
