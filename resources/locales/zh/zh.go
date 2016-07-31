package zh

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type zh struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'zh' locale
func New() locales.Translator {
	return &zh{
		locale:  "zh",
		plurals: []locales.PluralRule{6},
	}
}

// Locale returns the current translators string locale
func (t *zh) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'zh'
func (t *zh) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'zh'
func (t *zh) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
