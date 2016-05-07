package af

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE, dd MMMM y", Long: "dd MMMM y", Medium: "dd MMM y", Short: "y-MM-dd"}, AD: ut.CalendarDateFormat{Full: "EEEE, dd MMMM y", Long: "dd MMMM y", Medium: "dd MMM y", Short: "y-MM-dd"}}, Time: ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{7: "Jul.", 10: "Okt.", 11: "Nov.", 3: "Mrt.", 6: "Jun.", 4: "Apr.", 5: "Mei", 8: "Aug.", 9: "Sep.", 12: "Des.", 1: "Jan.", 2: "Feb."}, Narrow: ut.CalendarMonthFormatNameValue{9: "S", 10: "O", 12: "D", 4: "A", 5: "M", 7: "J", 8: "A", 11: "N", 1: "J", 2: "F", 3: "M", 6: "J"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{2: "Februarie", 6: "Junie", 7: "Julie", 10: "Oktober", 12: "Desember", 1: "Januarie", 3: "Maart", 4: "April", 5: "Mei", 8: "Augustus", 9: "September", 11: "November"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{6: "Sa.", 0: "So.", 1: "Ma.", 2: "Di.", 3: "Wo.", 4: "Do.", 5: "Vr."}, Narrow: ut.CalendarDayFormatNameValue{3: "W", 4: "D", 5: "V", 6: "S", 0: "S", 1: "M", 2: "D"}, Short: ut.CalendarDayFormatNameValue{1: "Ma.", 2: "Di.", 3: "Wo.", 4: "Do.", 5: "Vr.", 6: "Sa.", 0: "So."}, Wide: ut.CalendarDayFormatNameValue{5: "Vrydag", 6: "Saterdag", 0: "Sondag", 1: "Maandag", 2: "Dinsdag", 3: "Woensdag", 4: "Donderdag"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"evening1": "aand", "night1": "nag", "midnight": "middernag", "am": "vm.", "pm": "nm.", "morning1": "oggend", "afternoon1": "middag"}, Narrow: ut.CalendarPeriodFormatNameValue{"pm": "n", "morning1": "o", "afternoon1": "m", "evening1": "a", "night1": "n", "midnight": "mn", "am": "v"}, Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"morning1": "oggend", "afternoon1": "middag", "evening1": "aand", "night1": "nag", "midnight": "middernag", "am": "vm.", "pm": "nm."}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "na Christus", Abbrev: "n.C.", Narrow: "n.C."}, BC: ut.CalendarEraFormatNames{Full: "voor Christus", Abbrev: "v.C.", Narrow: "v.C."}}}}
