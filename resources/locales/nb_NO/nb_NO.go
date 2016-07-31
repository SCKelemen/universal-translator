package nb_NO

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type nb_NO struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'nb_NO' locale
func New() locales.Translator {
	return &nb_NO{
		locale:  "nb_NO",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *nb_NO) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'nb_NO'
func (t *nb_NO) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'nb_NO'
func (t *nb_NO) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
