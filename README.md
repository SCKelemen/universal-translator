## universal-translator
<img align="right" src="https://raw.githubusercontent.com/go-playground/universal-translator/master/logo.png">
![Project status](https://img.shields.io/badge/version-0.9-green.svg)
[![Build Status](https://semaphoreci.com/api/v1/joeybloggs/universal-translator/branches/master/badge.svg)](https://semaphoreci.com/joeybloggs/universal-translator)
[![Coverage Status](https://coveralls.io/repos/github/go-playground/universal-translator/badge.svg?branch=master)](https://coveralls.io/github/go-playground/universal-translator?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-playground/universal-translator)](https://goreportcard.com/report/github.com/go-playground/universal-translator)
[![GoDoc](https://godoc.org/github.com/go-playground/universal-translator?status.svg)](https://godoc.org/github.com/go-playground/universal-translator)
![License](https://img.shields.io/dub/l/vibe-d.svg)

Universal Translator is an i18n Translator for Go/Golang using CLDR data + pluralization rules

Why another i18n library?
--------------------------
Because none of the plural rules seem to be correct out there, including the previous implimentation of this package,
so I took it upon myself to create [locales](https://github.com/go-playground/locales) for everyone to use; this package 
is a thin wrapper around [locales](https://github.com/go-playground/locales) in order to store and translate text for 
use in your applications.

Features
--------
- [x] Rules generated from the latest [CLDR](http://cldr.unicode.org/index/downloads) data, v29
- [x] Contains Cardinal, Ordinal and Range Plural Rules
- [x] Contains Month, Weekday and Timezone translations built in
- [x] Contains Date & Time formatting functions
- [x] Contains Number, Currency, Accounting and Percent formatting functions
- [x] Supports the "Gregorian" calendar only ( my time isn't unlimited, had to draw the line somewhere )
- [ ] Support loading translations from files
- [ ] Exporting translations to file, mainly for getting them professionally translated
- [ ] Code Generation for translation files -> Go code.. i.e. after it has been professionally translated
- [ ] Tests for all languages, I need help with this, please see [here](https://github.com/go-playground/locales/issues/1)

Installation
-----------

Use go get 

```go
go get github.com/go-playground/universal-translator
```

Usage
-------
```go
package main

import (
	"fmt"

	"github.com/go-playground/locales"
	"github.com/go-playground/universal-translator"
)

// only one instance as translators within are shared + goroutine safe
var universalTraslator *ut.UniversalTranslator

func main() {

	// NOTE: this example is omitting allot of error checking for brevity

	universalTraslator, _ = ut.New("en", "en", "en_CA", "nl", "fr")

	en := universalTraslator.GetTranslator("en")

	// generally used after parsing an http 'Accept-Language' header
	// and this will try to find a matching locale you support or
	// fallback locale.
	// en, _ := ut.FindTranslator([]string{"en", "en_CA", "nl"})

	// this will help
	fmt.Println("Cardinal Plural Rules:", en.PluralsCardinal())
	fmt.Println("Ordinal Plural Rules:", en.PluralsOrdinal())
	fmt.Println("Range Plural Rules:", en.PluralsRange())

	// add basic language only translations
	en.Add("welcome", "Welcome {0} to our test")

	// add language translations dependant on cardinal plural rules
	en.AddCardinal("days", "You have {0} day left to register", locales.PluralRuleOne)
	en.AddCardinal("days", "You have {0} days left to register", locales.PluralRuleOther)

	// add language translations dependant on ordinal plural rules
	en.AddOrdinal("day-of-month", "{0}st", locales.PluralRuleOne)
	en.AddOrdinal("day-of-month", "{0}nd", locales.PluralRuleTwo)
	en.AddOrdinal("day-of-month", "{0}rd", locales.PluralRuleFew)
	en.AddOrdinal("day-of-month", "{0}th", locales.PluralRuleOther)

	// add language translations dependant on range plural rules
	// NOTE: only one plural rule for range in 'en' locale
	en.AddRange("between", "It's {0}-{1} days away", locales.PluralRuleOther)

	// now lets use the translations we just added, in the same order we added them

	fmt.Println(en.T("welcome", "Joeybloggs"))

	fmt.Println(en.C("days", 1, 0, string(en.FmtNumber(1, 0)))) // you'd normally have variables defined for 1 and 0
	fmt.Println(en.C("days", 2, 0, string(en.FmtNumber(2, 0))))
	fmt.Println(en.C("days", 10456.25, 2, string(en.FmtNumber(10456.25, 2))))

	fmt.Println(en.O("day-of-month", 1, 0, string(en.FmtNumber(1, 0))))
	fmt.Println(en.O("day-of-month", 2, 0, string(en.FmtNumber(2, 0))))
	fmt.Println(en.O("day-of-month", 3, 0, string(en.FmtNumber(3, 0))))
	fmt.Println(en.O("day-of-month", 4, 0, string(en.FmtNumber(4, 0))))
	fmt.Println(en.O("day-of-month", 10456.25, 0, string(en.FmtNumber(10456.25, 0))))

	fmt.Println(en.R("between", 0, 0, 1, 0, string(en.FmtNumber(0, 0)), string(en.FmtNumber(1, 0))))
	fmt.Println(en.R("between", 1, 0, 2, 0, string(en.FmtNumber(1, 0)), string(en.FmtNumber(2, 0))))
	fmt.Println(en.R("between", 1, 0, 100, 0, string(en.FmtNumber(1, 0)), string(en.FmtNumber(100, 0))))
}
```

Help With Tests
---------------
To anyone interesting in helping or contributing, I sure could use some help creating tests for each language.
Please see issue [here](https://github.com/go-playground/locales/issues/1) for details.
