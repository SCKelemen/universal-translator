package dje

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM, y", Short: "d/M/y"}, AD: ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM, y", Short: "d/M/y"}}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "", Long: "", Medium: "", Short: ""}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{9: "Sek", 11: "Noo", 12: "Dee", 1: "Žan", 4: "Awi", 7: "Žuy", 8: "Ut", 10: "Okt", 2: "Fee", 3: "Mar", 5: "Me", 6: "Žuw"}, Narrow: ut.CalendarMonthFormatNameValue{6: "Ž", 9: "S", 10: "O", 11: "N", 12: "D", 1: "Ž", 3: "M", 4: "A", 5: "M", 7: "Ž", 8: "U", 2: "F"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{10: "Oktoobur", 2: "Feewiriye", 3: "Marsi", 5: "Me", 7: "Žuyye", 8: "Ut", 12: "Deesanbur", 1: "Žanwiye", 4: "Awiril", 6: "Žuweŋ", 9: "Sektanbur", 11: "Noowanbur"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{6: "Asi", 0: "Alh", 1: "Ati", 2: "Ata", 3: "Ala", 4: "Alm", 5: "Alz"}, Narrow: ut.CalendarDayFormatNameValue{3: "L", 4: "M", 5: "Z", 6: "S", 0: "H", 1: "T", 2: "T"}, Short: ut.CalendarDayFormatNameValue(nil), Wide: ut.CalendarDayFormatNameValue{6: "Asibti", 0: "Alhadi", 1: "Atinni", 2: "Atalaata", 3: "Alarba", 4: "Alhamisi", 5: "Alzuma"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"am": "Subbaahi", "pm": "Zaarikay b"}, Narrow: ut.CalendarPeriodFormatNameValue(nil), Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"am": "Subbaahi", "pm": "Zaarikay b"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "Isaa jine", Abbrev: "IJ", Narrow: ""}}}}
