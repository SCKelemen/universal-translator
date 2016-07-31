package sv

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type sv struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'sv' locale
func New() locales.Translator {
	return &sv{
		locale:  "sv",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *sv) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'sv'
func (t *sv) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'sv'
func (t *sv) CardinalPluralRule(num string) (locales.PluralRule, error) {

	v := locales.V(num)

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if i == 1 && v == 0 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
