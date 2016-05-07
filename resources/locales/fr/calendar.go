package fr

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}, AD: ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} 'à' {0}", Long: "{1} 'à' {0}", Medium: "{1} 'à' {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{2: "févr.", 5: "mai", 9: "sept.", 10: "oct.", 11: "nov.", 12: "déc.", 1: "janv.", 4: "avr.", 6: "juin", 7: "juil.", 8: "août", 3: "mars"}, Narrow: ut.CalendarMonthFormatNameValue{10: "O", 11: "N", 1: "J", 3: "M", 4: "A", 5: "M", 8: "A", 2: "F", 6: "J", 7: "J", 9: "S", 12: "D"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{11: "novembre", 12: "décembre", 3: "mars", 5: "mai", 7: "juillet", 10: "octobre", 8: "août", 9: "septembre", 1: "janvier", 2: "février", 4: "avril", 6: "juin"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{5: "ven.", 6: "sam.", 0: "dim.", 1: "lun.", 2: "mar.", 3: "mer.", 4: "jeu."}, Narrow: ut.CalendarDayFormatNameValue{2: "M", 3: "M", 4: "J", 5: "V", 6: "S", 0: "D", 1: "L"}, Short: ut.CalendarDayFormatNameValue{4: "je", 5: "ve", 6: "sa", 0: "di", 1: "lu", 2: "ma", 3: "me"}, Wide: ut.CalendarDayFormatNameValue{0: "dimanche", 1: "lundi", 2: "mardi", 3: "mercredi", 4: "jeudi", 5: "vendredi", 6: "samedi"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"am": "AM", "noon": "midi", "pm": "PM", "morning1": "mat.", "afternoon1": "ap.m.", "evening1": "soir", "night1": "nuit", "midnight": "min."}, Narrow: ut.CalendarPeriodFormatNameValue{"noon": "midi", "pm": "PM", "morning1": "mat.", "afternoon1": "ap.m.", "evening1": "soir", "night1": "nuit", "midnight": "min.", "am": "AM"}, Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"morning1": "matin", "afternoon1": "après-midi", "evening1": "soir", "night1": "nuit", "midnight": "minuit", "am": "AM", "noon": "midi", "pm": "PM"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "après Jésus-Christ", Abbrev: "ap. J.-C.", Narrow: "ap. J.-C."}, BC: ut.CalendarEraFormatNames{Full: "avant Jésus-Christ", Abbrev: "av. J.-C.", Narrow: "av. J.-C."}}}}
