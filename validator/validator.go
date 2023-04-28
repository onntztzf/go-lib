package validator

import (
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
