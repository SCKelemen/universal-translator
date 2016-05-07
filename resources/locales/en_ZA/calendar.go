package en_ZA

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE, dd MMMM y", Long: "dd MMMM y", Medium: "dd MMM y", Short: "y/MM/dd"}, AD: ut.CalendarDateFormat{Full: "EEEE, dd MMMM y", Long: "dd MMMM y", Medium: "dd MMM y", Short: "y/MM/dd"}}, Time: ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: ut.CalendarDateFormat{Full: "{1} 'at' {0}", Long: "{1} 'at' {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{9: "Sep", 8: "Aug", 11: "Nov", 12: "Dec", 2: "Feb", 4: "Apr", 5: "May", 10: "Oct", 1: "Jan", 3: "Mar", 6: "Jun", 7: "Jul"}, Narrow: ut.CalendarMonthFormatNameValue{5: "M", 8: "A", 6: "J", 1: "J", 3: "M", 4: "A", 7: "J", 11: "N", 2: "F", 9: "S", 10: "O", 12: "D"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{11: "November", 5: "May", 8: "August", 10: "October", 9: "September", 1: "January", 2: "February", 3: "March", 6: "June", 12: "December", 4: "April", 7: "July"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{1: "Mon", 2: "Tue", 3: "Wed", 4: "Thu", 5: "Fri", 6: "Sat", 0: "Sun"}, Narrow: ut.CalendarDayFormatNameValue{3: "W", 4: "T", 5: "F", 6: "S", 0: "S", 1: "M", 2: "T"}, Short: ut.CalendarDayFormatNameValue{6: "Sa", 0: "Su", 1: "Mo", 2: "Tu", 3: "We", 4: "Th", 5: "Fr"}, Wide: ut.CalendarDayFormatNameValue{4: "Thursday", 5: "Friday", 6: "Saturday", 0: "Sunday", 1: "Monday", 2: "Tuesday", 3: "Wednesday"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"afternoon1": "in the afternoon", "evening1": "in the evening", "night1": "at night", "midnight": "midnight", "am": "AM", "noon": "noon", "pm": "PM", "morning1": "in the morning"}, Narrow: ut.CalendarPeriodFormatNameValue{"am": "a", "noon": "n", "pm": "p", "morning1": "in the morning", "afternoon1": "in the afternoon", "evening1": "in the evening", "night1": "at night", "midnight": "mi"}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"afternoon1": "afternoon", "evening1": "evening", "night1": "night", "midnight": "midnight", "am": "AM", "noon": "noon", "pm": "PM", "morning1": "morning"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}}}}
