package ksb

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}, AD: ut.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "", Long: "", Medium: "", Short: ""}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{1: "Jan", 2: "Feb", 7: "Jul", 9: "Sep", 8: "Ago", 10: "Okt", 11: "Nov", 12: "Des", 3: "Mac", 4: "Apr", 5: "Mei", 6: "Jun"}, Narrow: ut.CalendarMonthFormatNameValue{4: "A", 5: "M", 9: "S", 10: "O", 12: "D", 1: "J", 3: "M", 6: "J", 7: "J", 8: "A", 11: "N", 2: "F"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{12: "Desemba", 4: "Aplili", 6: "Juni", 8: "Agosti", 10: "Oktoba", 11: "Novemba", 9: "Septemba", 1: "Januali", 2: "Febluali", 3: "Machi", 5: "Mei", 7: "Julai"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{6: "Jmo", 0: "Jpi", 1: "Jtt", 2: "Jmn", 3: "Jtn", 4: "Alh", 5: "Iju"}, Narrow: ut.CalendarDayFormatNameValue{4: "A", 5: "I", 6: "1", 0: "2", 1: "3", 2: "4", 3: "5"}, Short: ut.CalendarDayFormatNameValue(nil), Wide: ut.CalendarDayFormatNameValue{0: "Jumaapii", 1: "Jumaatatu", 2: "Jumaane", 3: "Jumaatano", 4: "Alhamisi", 5: "Ijumaa", 6: "Jumaamosi"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"am": "makeo", "pm": "nyiaghuo"}, Narrow: ut.CalendarPeriodFormatNameValue(nil), Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"am": "makeo", "pm": "nyiaghuo"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "Kabla ya Klisto", Abbrev: "KK", Narrow: ""}}}}
