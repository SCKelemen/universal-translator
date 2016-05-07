package es_CR

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE, d 'de' MMMM 'de' y", Long: "d 'de' MMMM 'de' y", Medium: "d MMM y", Short: "d/M/yy"}, AD: ut.CalendarDateFormat{Full: "EEEE, d 'de' MMMM 'de' y", Long: "d 'de' MMMM 'de' y", Medium: "d MMM y", Short: "d/M/yy"}}, Time: ut.CalendarDateFormat{Full: "H:mm:ss (zzzz)", Long: "H:mm:ss z", Medium: "H:mm:ss", Short: "H:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1}, {0}", Long: "{1}, {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{5: "may.", 6: "jun.", 8: "ago.", 4: "abr.", 7: "jul.", 10: "oct.", 12: "dic.", 11: "nov.", 2: "feb.", 3: "mar.", 9: "sept.", 1: "ene."}, Narrow: ut.CalendarMonthFormatNameValue{8: "A", 9: "S", 12: "D", 2: "F", 10: "O", 3: "M", 6: "J", 5: "M", 11: "N", 1: "E", 7: "J", 4: "A"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{5: "mayo", 3: "marzo", 4: "abril", 8: "agosto", 10: "octubre", 11: "noviembre", 6: "junio", 9: "septiembre", 1: "enero", 12: "diciembre", 7: "julio", 2: "febrero"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{2: "mar.", 3: "mié.", 4: "jue.", 5: "vie.", 6: "sáb.", 0: "dom.", 1: "lun."}, Narrow: ut.CalendarDayFormatNameValue{5: "V", 6: "S", 0: "D", 1: "L", 2: "M", 3: "X", 4: "J"}, Short: ut.CalendarDayFormatNameValue{2: "MA", 3: "MI", 4: "JU", 5: "VI", 6: "SA", 0: "DO", 1: "LU"}, Wide: ut.CalendarDayFormatNameValue{4: "jueves", 5: "viernes", 6: "sábado", 0: "domingo", 1: "lunes", 2: "martes", 3: "miércoles"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"noon": "mediodía", "pm": "p. m.", "morning1": "madrugada", "morning2": "mañana", "evening1": "tarde", "night1": "noche", "am": "a. m."}, Narrow: ut.CalendarPeriodFormatNameValue{"morning1": "madrugada", "morning2": "mañana", "evening1": "tarde", "night1": "noche", "am": "a. m.", "noon": "mediodía", "pm": "p. m."}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"noon": "mediodía", "pm": "p. m.", "morning1": "madrugada", "morning2": "mañana", "evening1": "tarde", "night1": "noche", "am": "a. m."}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}}}}
