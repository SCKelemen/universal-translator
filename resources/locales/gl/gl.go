package gl

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type gl struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'gl' locale
func New() locales.Translator {
	return &gl{
		locale:  "gl",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *gl) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'gl'
func (t *gl) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'gl'
func (t *gl) CardinalPluralRule(num string) (locales.PluralRule, error) {

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	v := locales.V(num)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
