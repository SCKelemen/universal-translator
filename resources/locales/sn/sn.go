package sn

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type sn struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'sn' locale
func New() locales.Translator {
	return &sn{
		locale:  "sn",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *sn) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'sn'
func (t *sn) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'sn'
func (t *sn) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
