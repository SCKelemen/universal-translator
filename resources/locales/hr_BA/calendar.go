package hr_BA

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE, d. MMMM y.", Long: "d. MMMM y.", Medium: "d. MMM y.", Short: "dd.MM.y."}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} 'u' {0}", Long: "{1} 'u' {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{2: "velj", 9: "ruj", 12: "pro", 4: "tra", 7: "srp", 8: "kol", 5: "svi", 6: "lip", 11: "stu", 1: "sij", 3: "ožu", 10: "lis"}, Narrow: ut.CalendarMonthFormatNameValue{5: "5.", 11: "11.", 6: "6.", 9: "9.", 7: "7.", 8: "8.", 2: "2.", 4: "4.", 10: "10.", 12: "12.", 1: "1.", 3: "3."}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{1: "siječanj", 4: "travanj", 10: "listopad", 11: "studeni", 2: "veljača", 6: "lipanj", 12: "prosinac", 5: "svibanj", 3: "ožujak", 7: "srpanj", 8: "kolovoz", 9: "rujan"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{4: "čet", 5: "pet", 6: "sub", 0: "ned", 1: "pon", 2: "uto", 3: "sri"}, Narrow: ut.CalendarDayFormatNameValue{0: "n", 1: "p", 2: "u", 3: "s", 4: "č", 5: "p", 6: "s"}, Short: ut.CalendarDayFormatNameValue{6: "sub", 0: "ned", 1: "pon", 2: "uto", 3: "sri", 4: "čet", 5: "pet"}, Wide: ut.CalendarDayFormatNameValue{4: "četvrtak", 5: "petak", 6: "subota", 0: "nedjelja", 1: "ponedjeljak", 2: "utorak", 3: "srijeda"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"am": "AM", "pm": "PM"}, Narrow: ut.CalendarPeriodFormatNameValue{"am": "AM", "pm": "PM"}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"pm": "PM", "morning1": "jutro", "afternoon1": "popodne", "evening1": "večer", "night1": "noć", "midnight": "ponoć", "am": "AM", "noon": "podne"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}}}}
