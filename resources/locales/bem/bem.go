package bem

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type bem struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'bem' locale
func New() locales.Translator {
	return &bem{
		locale:  "bem",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *bem) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'bem'
func (t *bem) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'bem'
func (t *bem) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
