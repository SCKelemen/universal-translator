package fi

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "cccc d. MMMM y", Long: "d. MMMM y", Medium: "d.M.y", Short: "d.M.y"}, AD: ut.CalendarDateFormat{Full: "cccc d. MMMM y", Long: "d. MMMM y", Medium: "d.M.y", Short: "d.M.y"}}, Time: ut.CalendarDateFormat{Full: "H.mm.ss zzzz", Long: "H.mm.ss z", Medium: "H.mm.ss", Short: "H.mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} 'klo' {0}", Long: "{1} 'klo' {0}", Medium: "{1} 'klo' {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{3: "maalis", 4: "huhti", 5: "touko", 6: "kesä", 7: "heinä", 10: "loka", 12: "joulu", 1: "tammi", 2: "helmi", 8: "elo", 9: "syys", 11: "marras"}, Narrow: ut.CalendarMonthFormatNameValue{2: "H", 3: "M", 7: "H", 8: "E", 12: "J", 1: "T", 4: "H", 5: "T", 6: "K", 9: "S", 10: "L", 11: "M"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{2: "helmikuu", 4: "huhtikuu", 7: "heinäkuu", 11: "marraskuu", 10: "lokakuu", 12: "joulukuu", 1: "tammikuu", 3: "maaliskuu", 5: "toukokuu", 6: "kesäkuu", 8: "elokuu", 9: "syyskuu"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{0: "su", 1: "ma", 2: "ti", 3: "ke", 4: "to", 5: "pe", 6: "la"}, Narrow: ut.CalendarDayFormatNameValue{0: "S", 1: "M", 2: "T", 3: "K", 4: "T", 5: "P", 6: "L"}, Short: ut.CalendarDayFormatNameValue{1: "ma", 2: "ti", 3: "ke", 4: "to", 5: "pe", 6: "la", 0: "su"}, Wide: ut.CalendarDayFormatNameValue{5: "perjantai", 6: "lauantai", 0: "sunnuntai", 1: "maanantai", 2: "tiistai", 3: "keskiviikko", 4: "torstai"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"afternoon1": "iltap.", "midnight": "keskiyö", "am": "ap.", "noon": "keskip.", "pm": "ip.", "morning1": "aamu", "morning2": "aamup.", "evening1": "ilta", "night1": "yö"}, Narrow: ut.CalendarPeriodFormatNameValue{"midnight": "ky.", "am": "ap.", "noon": "kp.", "night1": "yö", "pm": "ip.", "morning1": "aamu", "morning2": "ap.", "afternoon1": "ip.", "evening1": "ilta"}, Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"afternoon1": "iltapäivä", "evening1": "ilta", "night1": "yö", "midnight": "keskiyö", "am": "ap.", "noon": "keskipäivä", "morning2": "aamupäivä", "pm": "ip.", "morning1": "aamu"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "jälkeen Kristuksen syntymän", Abbrev: "jKr.", Narrow: "jK"}, BC: ut.CalendarEraFormatNames{Full: "ennen Kristuksen syntymää", Abbrev: "eKr.", Narrow: "eK"}}}}
