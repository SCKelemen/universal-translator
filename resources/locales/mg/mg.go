package mg

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type mg struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'mg' locale
func New() locales.Translator {
	return &mg{
		locale:  "mg",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *mg) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'mg'
func (t *mg) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'mg'
func (t *mg) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n >= 0 && n <= 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
