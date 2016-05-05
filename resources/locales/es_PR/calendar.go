package es_PR

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE, d 'de' MMMM 'de' y", Long: "d 'de' MMMM 'de' y", Medium: "MM/dd/y", Short: "MM/dd/yy"}, Time: ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: ut.CalendarDateFormat{Full: "{1}, {0}", Long: "{1}, {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{8: "ago.", 10: "oct.", 6: "jun.", 9: "sept.", 1: "ene.", 2: "feb.", 4: "abr.", 3: "mar.", 5: "may.", 7: "jul.", 12: "dic.", 11: "nov."}, Narrow: ut.CalendarMonthFormatNameValue{12: "D", 1: "E", 4: "A", 9: "S", 5: "M", 7: "J", 8: "A", 2: "F", 3: "M", 6: "J", 10: "O", 11: "N"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{7: "julio", 4: "abril", 5: "mayo", 6: "junio", 11: "noviembre", 1: "enero", 10: "octubre", 8: "agosto", 12: "diciembre", 3: "marzo", 2: "febrero", 9: "septiembre"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{5: "vie.", 6: "sáb.", 0: "dom.", 1: "lun.", 2: "mar.", 3: "mié.", 4: "jue."}, Narrow: ut.CalendarDayFormatNameValue{3: "X", 4: "J", 5: "V", 6: "S", 0: "D", 1: "L", 2: "M"}, Short: ut.CalendarDayFormatNameValue{6: "SA", 0: "DO", 1: "LU", 2: "MA", 3: "MI", 4: "JU", 5: "VI"}, Wide: ut.CalendarDayFormatNameValue{3: "miércoles", 4: "jueves", 5: "viernes", 6: "sábado", 0: "domingo", 1: "lunes", 2: "martes"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"noon": "mediodía", "pm": "p. m.", "morning1": "madrugada", "morning2": "mañana", "evening1": "tarde", "night1": "noche", "am": "a. m."}, Narrow: ut.CalendarPeriodFormatNameValue{"evening1": "tarde", "night1": "noche", "am": "a. m.", "noon": "mediodía", "pm": "p. m.", "morning1": "madrugada", "morning2": "mañana"}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"evening1": "tarde", "night1": "noche", "am": "a. m.", "noon": "mediodía", "pm": "p. m.", "morning1": "madrugada", "morning2": "mañana"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}}}}
