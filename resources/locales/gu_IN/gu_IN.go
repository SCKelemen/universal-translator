package gu_IN

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type gu_IN struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'gu_IN' locale
func New() locales.Translator {
	return &gu_IN{
		locale:  "gu_IN",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *gu_IN) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'gu_IN'
func (t *gu_IN) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'gu_IN'
func (t *gu_IN) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if (i == 0) || (n == 1) {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
