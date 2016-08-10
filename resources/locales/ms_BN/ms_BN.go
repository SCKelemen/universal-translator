package ms_BN

import "github.com/go-playground/universal-translator/resources/locales"

type ms_BN struct {
	locale   string
	plurals  []locales.PluralRule
	decimal  []byte
	group    []byte
	minus    []byte
	percent  []byte
	perMille []byte
	symbol   []byte
}

// New returns a new instance of translator for the 'ms_BN' locale
func New() locales.Translator {
	return &ms_BN{
		locale:   "ms_BN",
		plurals:  []locales.PluralRule{6},
		decimal:  []byte{0x2c},
		group:    []byte{0x2e},
		minus:    []byte{0x5b, 0x5d, 0x62, 0x79, 0x74, 0x65, 0x7b, 0x30, 0x78, 0x32, 0x64, 0x7d},
		percent:  []byte{0x5b, 0x5d, 0x62, 0x79, 0x74, 0x65, 0x7b, 0x30, 0x78, 0x32, 0x35, 0x7d},
		perMille: []byte{0x5b, 0x5d, 0x62, 0x79, 0x74, 0x65, 0x7b, 0x30, 0x78, 0x65, 0x32, 0x2c, 0x20, 0x30, 0x78, 0x38, 0x30, 0x2c, 0x20, 0x30, 0x78, 0x62, 0x30, 0x7d},
		symbol:   []byte{0x5b, 0x5d, 0x62, 0x79, 0x74, 0x65, 0x7b, 0x7d},
	}
}

// Locale returns the current translators string locale
func (t *ms_BN) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ms_BN'
func (t *ms_BN) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'ms_BN'
func (t *ms_BN) cardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}
