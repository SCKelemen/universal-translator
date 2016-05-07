package pl

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "dd.MM.y", Short: "dd.MM.y"}, AD: ut.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "dd.MM.y", Short: "dd.MM.y"}}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{7: "lip", 8: "sie", 10: "paź", 11: "lis", 3: "mar", 4: "kwi", 5: "maj", 6: "cze", 9: "wrz", 12: "gru", 1: "sty", 2: "lut"}, Narrow: ut.CalendarMonthFormatNameValue{6: "c", 7: "l", 8: "s", 10: "p", 1: "s", 3: "m", 4: "k", 5: "m", 11: "l", 2: "l", 9: "w", 12: "g"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{2: "luty", 3: "marzec", 4: "kwiecień", 9: "wrzesień", 10: "październik", 11: "listopad", 1: "styczeń", 5: "maj", 6: "czerwiec", 7: "lipiec", 8: "sierpień", 12: "grudzień"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{0: "niedz.", 1: "pon.", 2: "wt.", 3: "śr.", 4: "czw.", 5: "pt.", 6: "sob."}, Narrow: ut.CalendarDayFormatNameValue{1: "P", 2: "W", 3: "Ś", 4: "C", 5: "P", 6: "S", 0: "N"}, Short: ut.CalendarDayFormatNameValue{5: "pią", 6: "sob", 0: "nie", 1: "pon", 2: "wto", 3: "śro", 4: "czw"}, Wide: ut.CalendarDayFormatNameValue{5: "piątek", 6: "sobota", 0: "niedziela", 1: "poniedziałek", 2: "wtorek", 3: "środa", 4: "czwartek"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"am": "AM", "pm": "PM"}, Narrow: ut.CalendarPeriodFormatNameValue{"am": "a", "pm": "p"}, Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"morning2": "przedpołudnie", "afternoon1": "popołudnie", "evening1": "wieczór", "midnight": "północ", "noon": "południe", "pm": "PM", "morning1": "rano", "am": "AM", "night1": "noc"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "n.e.", Abbrev: "n.e.", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "p.n.e.", Abbrev: "p.n.e.", Narrow: ""}}}}
