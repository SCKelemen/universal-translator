package ms

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ms struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ms' locale
func New() locales.Translator {
	return &ms{
		locale:  "ms",
		plurals: []locales.PluralRule{6},
	}
}

// Locale returns the current translators string locale
func (t *ms) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ms'
func (t *ms) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'ms'
func (t *ms) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
