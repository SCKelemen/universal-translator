package dje_NE

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type dje_NE struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'dje_NE' locale
func New() locales.Translator {
	return &dje_NE{
		locale:  "dje_NE",
		plurals: nil,
	}
}

// Locale returns the current translators string locale
func (t *dje_NE) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'dje_NE'
func (t *dje_NE) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'dje_NE'
func (t *dje_NE) CardinalPluralRule(num string) (locales.PluralRule, error) {
	return locales.PluralRuleUnknown, nil
}
