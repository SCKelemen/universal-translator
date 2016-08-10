package dsb

import (
	"math"

	"github.com/go-playground/universal-translator/resources/locales"
)

type dsb struct {
	locale   string
	plurals  []locales.PluralRule
	decimal  []byte
	group    []byte
	minus    []byte
	percent  []byte
	perMille []byte
	symbol   []byte
}

// New returns a new instance of translator for the 'dsb' locale
func New() locales.Translator {
	return &dsb{
		locale:   "dsb",
		plurals:  []locales.PluralRule{2, 3, 4, 6},
		decimal:  []byte{0x2c},
		group:    []byte{0x2e},
		minus:    []byte{0x2d},
		percent:  []byte{0x25},
		perMille: []byte{0xe2, 0x80, 0xb0},
		symbol:   []byte{},
	}
}

// Locale returns the current translators string locale
func (t *dsb) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'dsb'
func (t *dsb) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'dsb'
func (t *dsb) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)
	f := locales.F(n, v)

	if (v == 0 && i%100 == 1) || (f%100 == 1) {
		return locales.PluralRuleOne
	} else if (v == 0 && i%100 == 2) || (f%100 == 2) {
		return locales.PluralRuleTwo
	} else if (v == 0 && i%100 >= 3 && i%100 <= 4) || (f%100 >= 3 && f%100 <= 4) {
		return locales.PluralRuleFew
	}

	return locales.PluralRuleOther
}
