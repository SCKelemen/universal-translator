package ii_CN

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ii_CN struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ii_CN' locale
func New() locales.Translator {
	return &ii_CN{
		locale:  "ii_CN",
		plurals: []locales.PluralRule{6},
	}
}

// Locale returns the current translators string locale
func (t *ii_CN) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ii_CN'
func (t *ii_CN) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'ii_CN'
func (t *ii_CN) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
