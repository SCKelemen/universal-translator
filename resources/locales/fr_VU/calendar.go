package fr_VU

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}, AD: ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}}, Time: ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: ut.CalendarDateFormat{Full: "{1} 'à' {0}", Long: "{1} 'à' {0}", Medium: "{1} 'à' {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{2: "févr.", 5: "mai", 3: "mars", 6: "juin", 7: "juil.", 10: "oct.", 12: "déc.", 8: "août", 4: "avr.", 11: "nov.", 1: "janv.", 9: "sept."}, Narrow: ut.CalendarMonthFormatNameValue{1: "J", 6: "J", 9: "S", 2: "F", 3: "M", 11: "N", 10: "O", 7: "J", 12: "D", 4: "A", 5: "M", 8: "A"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{8: "août", 10: "octobre", 9: "septembre", 1: "janvier", 2: "février", 11: "novembre", 4: "avril", 6: "juin", 12: "décembre", 3: "mars", 5: "mai", 7: "juillet"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{0: "dim.", 1: "lun.", 2: "mar.", 3: "mer.", 4: "jeu.", 5: "ven.", 6: "sam."}, Narrow: ut.CalendarDayFormatNameValue{2: "M", 3: "M", 4: "J", 5: "V", 6: "S", 0: "D", 1: "L"}, Short: ut.CalendarDayFormatNameValue{2: "ma", 3: "me", 4: "je", 5: "ve", 6: "sa", 0: "di", 1: "lu"}, Wide: ut.CalendarDayFormatNameValue{3: "mercredi", 4: "jeudi", 5: "vendredi", 6: "samedi", 0: "dimanche", 1: "lundi", 2: "mardi"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"noon": "midi", "pm": "PM", "morning1": "mat.", "afternoon1": "ap.m.", "evening1": "soir", "night1": "nuit", "midnight": "min.", "am": "AM"}, Narrow: ut.CalendarPeriodFormatNameValue{"pm": "PM", "morning1": "mat.", "afternoon1": "ap.m.", "evening1": "soir", "night1": "nuit", "midnight": "min.", "am": "AM", "noon": "midi"}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"am": "AM", "noon": "midi", "pm": "PM", "morning1": "matin", "afternoon1": "après-midi", "evening1": "soir", "night1": "nuit", "midnight": "minuit"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}}}}
