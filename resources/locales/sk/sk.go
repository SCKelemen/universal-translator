package sk

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type sk struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'sk' locale
func New() locales.Translator {
	return &sk{
		locale:  "sk",
		plurals: []locales.PluralRule{2, 4, 5, 6},
	}
}

// Locale returns the current translators string locale
func (t *sk) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'sk'
func (t *sk) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'sk'
func (t *sk) CardinalPluralRule(num string) (locales.PluralRule, error) {

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	v := locales.V(num)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne, nil
	} else if i >= 2 && i <= 4 && v == 0 {
		return locales.PluralRuleFew, nil
	} else if v != 0 {
		return locales.PluralRuleMany, nil
	}

	return locales.PluralRuleOther, nil
}
