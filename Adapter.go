package validator

import (
	localsEn "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
)

// Adapter is used to validate structures and variables
// against a set of validation rules.
type Adapter struct {
	validate *validator.Validate
	uni      *ut.UniversalTranslator
}

// NewAdapter creates a new validator adapter instance.
func NewAdapter() (AdapterInterface, error) {

	a := &Adapter{}

	a.validate = validator.New()

	en := localsEn.New()
	a.uni = ut.New(en, en)

	return a, nil
}

// Validate validates bound values of an unpacker struct against
// validation rules defined in that unpacker struct.
func (a *Adapter) Validate(data interface{}) map[string]string {

	// returns nil or ValidationErrors ( []FieldError )
	err := a.validate.Struct(data)
	if err == nil {
		return nil
	}

	// from here you can create your own error messages in whatever language you wish
	errs := err.(validator.ValidationErrors)

	return errs.Translate(a.getTranslator("en"))
}

// ValidateField validates a single variable.
func (a *Adapter) ValidateField(field interface{}, rules string) map[string]string {

	// returns nil or ValidationErrors ( []FieldError )
	err := a.validate.Var(field, rules)
	if err == nil {
		return nil
	}

	// from here you can create your own error messages in whatever language you wish
	errs := err.(validator.ValidationErrors)

	return errs.Translate(a.getTranslator("en"))
}

// getTranslator returns a translator for the given locale.
func (a *Adapter) getTranslator(locale string) ut.Translator {

	// this is usually know or extracted from http 'Accept-Language' header
	// also see uni.FindTranslator(...)
	trans, _ := a.uni.GetTranslator(locale)

	enTranslations.RegisterDefaultTranslations(a.validate, trans)

	return trans
}
