package de_DE

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type de_DE struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'de_DE' locale
func New() locales.Translator {
	return &de_DE{
		locale:  "de_DE",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *de_DE) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'de_DE'
func (t *de_DE) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'de_DE'
func (t *de_DE) CardinalPluralRule(num string) (locales.PluralRule, error) {

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
