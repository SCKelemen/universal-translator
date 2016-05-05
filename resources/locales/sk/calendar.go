package sk

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE, d. MMMM y", Long: "d. MMMM y", Medium: "d. M. y", Short: "d.M.yy"}, Time: ut.CalendarDateFormat{Full: "H:mm:ss zzzz", Long: "H:mm:ss z", Medium: "H:mm:ss", Short: "H:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1}, {0}", Long: "{1}, {0}", Medium: "{1}, {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{1: "jan", 2: "feb", 3: "mar", 4: "apr", 6: "jún", 5: "máj", 7: "júl", 8: "aug", 9: "sep", 10: "okt", 11: "nov", 12: "dec"}, Narrow: ut.CalendarMonthFormatNameValue{2: "f", 4: "a", 5: "m", 6: "j", 7: "j", 8: "a", 10: "o", 1: "j", 3: "m", 9: "s", 11: "n", 12: "d"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{1: "január", 3: "marec", 4: "apríl", 6: "jún", 8: "august", 11: "november", 2: "február", 5: "máj", 7: "júl", 9: "september", 10: "október", 12: "december"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{6: "so", 0: "ne", 1: "po", 2: "ut", 3: "st", 4: "št", 5: "pi"}, Narrow: ut.CalendarDayFormatNameValue{4: "š", 5: "p", 6: "s", 0: "n", 1: "p", 2: "u", 3: "s"}, Short: ut.CalendarDayFormatNameValue{0: "ne", 1: "po", 2: "ut", 3: "st", 4: "št", 5: "pi", 6: "so"}, Wide: ut.CalendarDayFormatNameValue{0: "nedeľa", 1: "pondelok", 2: "utorok", 3: "streda", 4: "štvrtok", 5: "piatok", 6: "sobota"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"am": "AM", "morning1": "ráno", "afternoon1": "popol.", "evening1": "večer", "midnight": "poln.", "noon": "pol.", "pm": "PM", "morning2": "dopol.", "night1": "noc"}, Narrow: ut.CalendarPeriodFormatNameValue{"pm": "PM", "morning1": "ráno", "morning2": "dop.", "evening1": "več.", "midnight": "poln.", "am": "AM", "noon": "pol.", "afternoon1": "pop.", "night1": "noc"}, Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"midnight": "polnoc", "afternoon1": "popoludnie", "evening1": "večer", "night1": "noc", "am": "AM", "noon": "poludnie", "pm": "PM", "morning1": "ráno", "morning2": "dopoludnie"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "po Kristovi", Abbrev: "po Kr.", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "pred Kristom", Abbrev: "pred Kr.", Narrow: ""}}}}
