package sw_CD

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type sw_CD struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'sw_CD' locale
func New() locales.Translator {
	return &sw_CD{
		locale:  "sw_CD",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *sw_CD) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'sw_CD'
func (t *sw_CD) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'sw_CD'
func (t *sw_CD) CardinalPluralRule(num string) (locales.PluralRule, error) {

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
