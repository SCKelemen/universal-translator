package xog

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}, AD: ut.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "", Long: "", Medium: "", Short: ""}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{9: "Seb", 1: "Jan", 3: "Mar", 5: "Maa", 6: "Juu", 7: "Jul", 12: "Des", 2: "Feb", 4: "Apu", 8: "Agu", 10: "Oki", 11: "Nov"}, Narrow: ut.CalendarMonthFormatNameValue{3: "M", 4: "A", 9: "S", 10: "O", 12: "D", 8: "A", 11: "N", 1: "J", 2: "F", 5: "M", 6: "J", 7: "J"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{12: "Desemba", 2: "Febwaliyo", 7: "Julaayi", 8: "Agusito", 5: "Maayi", 6: "Juuni", 9: "Sebuttemba", 10: "Okitobba", 11: "Novemba", 1: "Janwaliyo", 3: "Marisi", 4: "Apuli"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{4: "Kuna", 5: "Kuta", 6: "Muka", 0: "Sabi", 1: "Bala", 2: "Kubi", 3: "Kusa"}, Narrow: ut.CalendarDayFormatNameValue{0: "S", 1: "B", 2: "B", 3: "S", 4: "K", 5: "K", 6: "M"}, Short: ut.CalendarDayFormatNameValue(nil), Wide: ut.CalendarDayFormatNameValue{1: "Balaza", 2: "Owokubili", 3: "Owokusatu", 4: "Olokuna", 5: "Olokutaanu", 6: "Olomukaaga", 0: "Sabiiti"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"am": "Munkyo", "pm": "Eigulo"}, Narrow: ut.CalendarPeriodFormatNameValue(nil), Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"am": "Munkyo", "pm": "Eigulo"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "Kulisto nga azilawo", Abbrev: "AZ", Narrow: ""}}}}
