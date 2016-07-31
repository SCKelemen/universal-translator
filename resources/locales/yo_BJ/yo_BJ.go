package yo_BJ

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type yo_BJ struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'yo_BJ' locale
func New() locales.Translator {
	return &yo_BJ{
		locale:  "yo_BJ",
		plurals: []locales.PluralRule{6},
	}
}

// Locale returns the current translators string locale
func (t *yo_BJ) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'yo_BJ'
func (t *yo_BJ) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'yo_BJ'
func (t *yo_BJ) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
