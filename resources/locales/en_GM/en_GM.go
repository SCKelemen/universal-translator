package en_GM

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type en_GM struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'en_GM' locale
func New() locales.Translator {
	return &en_GM{
		locale:  "en_GM",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *en_GM) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'en_GM'
func (t *en_GM) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'en_GM'
func (t *en_GM) CardinalPluralRule(num string) (locales.PluralRule, error) {

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
