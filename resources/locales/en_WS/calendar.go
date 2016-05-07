package en_WS

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE, MMMM d, y GGGG", Long: "MMMM d, y GG", Medium: "MMM d, y", Short: "M/d/yy"}, AD: ut.CalendarDateFormat{Full: "EEEE, MMMM d, y", Long: "MMMM d, y", Medium: "MMM d, y", Short: "M/d/yy"}}, Time: ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: ut.CalendarDateFormat{Full: "{1} 'at' {0}", Long: "{1} 'at' {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{8: "Aug", 10: "Oct", 6: "Jun", 7: "Jul", 9: "Sep", 5: "May", 2: "Feb", 3: "Mar", 4: "Apr", 11: "Nov", 12: "Dec", 1: "Jan"}, Narrow: ut.CalendarMonthFormatNameValue{7: "J", 4: "A", 1: "J", 2: "F", 6: "J", 9: "S", 10: "O", 8: "A", 11: "N", 12: "D", 3: "M", 5: "M"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{5: "May", 9: "September", 10: "October", 3: "March", 6: "June", 12: "December", 4: "April", 7: "July", 8: "August", 1: "January", 2: "February", 11: "November"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{1: "Mon", 2: "Tue", 3: "Wed", 4: "Thu", 5: "Fri", 6: "Sat", 0: "Sun"}, Narrow: ut.CalendarDayFormatNameValue{5: "F", 6: "S", 0: "S", 1: "M", 2: "T", 3: "W", 4: "T"}, Short: ut.CalendarDayFormatNameValue{5: "Fr", 6: "Sa", 0: "Su", 1: "Mo", 2: "Tu", 3: "We", 4: "Th"}, Wide: ut.CalendarDayFormatNameValue{2: "Tuesday", 3: "Wednesday", 4: "Thursday", 5: "Friday", 6: "Saturday", 0: "Sunday", 1: "Monday"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"morning1": "in the morning", "afternoon1": "in the afternoon", "evening1": "in the evening", "night1": "at night", "midnight": "midnight", "am": "AM", "noon": "noon", "pm": "PM"}, Narrow: ut.CalendarPeriodFormatNameValue{"midnight": "mi", "am": "a", "noon": "n", "pm": "p", "morning1": "in the morning", "afternoon1": "in the afternoon", "evening1": "in the evening", "night1": "at night"}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"night1": "night", "midnight": "midnight", "am": "AM", "noon": "noon", "pm": "PM", "morning1": "morning", "afternoon1": "afternoon", "evening1": "evening"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}}}}
