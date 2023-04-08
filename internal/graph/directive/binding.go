package directive

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	validate *validator.Validate
	trans    ut.Translator
)

func init() {
	validate = validator.New()
	en := en.New()
	uni := ut.New(en, en)
	trans, _ = uni.GetTranslator("en")
	_ = en_translations.RegisterDefaultTranslations(validate, trans)
}

func Binding(ctx context.Context, obj interface{}, next graphql.Resolver, constraint string, field string) (interface{}, error) {
	val, err := next(ctx)
	if err != nil {
		panic(err)
	}
	if field == "" {
		field = *graphql.GetPathContext(ctx).Field
	}

	err = validate.Var(val, constraint)
	if err != nil {
		validationErrors, ok := err.(validator.ValidationErrors)
		if !ok {
			return val, nil
		}
		transErr := fmt.Errorf("%s%+v", field, validationErrors[0].Translate(trans))
		return val, transErr
	}

	return val, nil
}

func ValidateAddTranslation(tag string, message string) {
	_ = validate.RegisterTranslation(tag, trans, func(ut ut.Translator) error {
		return ut.Add(tag, message, true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(tag, fe.Field())

		return t
	})
}
