package ca

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ca struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ca' locale
func New() locales.Translator {
	return &ca{
		locale:  "ca",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *ca) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ca'
func (t *ca) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'ca'
func (t *ca) CardinalPluralRule(num string) (locales.PluralRule, error) {

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
