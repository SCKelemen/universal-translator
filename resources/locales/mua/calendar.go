package mua

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "d/M/y"}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "", Long: "", Medium: "", Short: ""}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{2: "CLA", 7: "MLI", 8: "MAM", 10: "FMU", 11: "FGW", 12: "FYU", 1: "FLO", 3: "CKI", 4: "FMF", 5: "MAD", 6: "MBI", 9: "FDE"}, Narrow: ut.CalendarMonthFormatNameValue{4: "F", 6: "B", 8: "M", 10: "U", 11: "W", 12: "Y", 2: "A", 3: "I", 5: "D", 7: "L", 9: "E", 1: "O"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{8: "Madǝmbii", 11: "Fĩi Gwahlle", 2: "Cokcwaklaŋne", 4: "Fĩi Marfoo", 7: "Mamǝŋgwãalii", 6: "Mamǝŋgwãafahbii", 9: "Fĩi Dǝɓlii", 10: "Fĩi Mundaŋ", 12: "Fĩi Yuru", 1: "Fĩi Loo", 3: "Cokcwaklii", 5: "Madǝǝuutǝbijaŋ"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{4: "Cka", 5: "Cga", 6: "Cze", 0: "Cya", 1: "Cla", 2: "Czi", 3: "Cko"}, Narrow: ut.CalendarDayFormatNameValue{0: "Y", 1: "L", 2: "Z", 3: "O", 4: "A", 5: "G", 6: "E"}, Short: ut.CalendarDayFormatNameValue(nil), Wide: ut.CalendarDayFormatNameValue{0: "Com’yakke", 1: "Comlaaɗii", 2: "Comzyiiɗii", 3: "Comkolle", 4: "Comkaldǝɓlii", 5: "Comgaisuu", 6: "Comzyeɓsuu"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"am": "comme", "pm": "lilli"}, Narrow: ut.CalendarPeriodFormatNameValue(nil), Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"am": "comme", "pm": "lilli"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "KǝPel Kristu", Abbrev: "KK", Narrow: ""}}}}
