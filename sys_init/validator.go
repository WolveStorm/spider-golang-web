package sys_init

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"spider-golang-web/global"
)

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

func ReplaceGinBinding(language string) {
	RegisterTranslator(language)
}

func RegisterTranslator(language string) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		zh := zh.New()
		uni := ut.New(zh, en, zh)
		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			jTag := field.Tag.Get("json")
			return jTag
		})
		// this is usually know or extracted from http 'Accept-Language' header
		// also see uni.FindTranslator(...)
		global.Trans, _ = uni.GetTranslator("zh")
		switch language {
		case "en":
			en_translations.RegisterDefaultTranslations(v, global.Trans)
		case "zh":
			zh_translations.RegisterDefaultTranslations(v, global.Trans)
		default:
			zh_translations.RegisterDefaultTranslations(v, global.Trans)
		}
	}
}
