package en_SI

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE, MMMM d, y GGGG", Long: "MMMM d, y GG", Medium: "MMM d, y", Short: "M/d/yy"}, AD: ut.CalendarDateFormat{Full: "EEEE, MMMM d, y", Long: "MMMM d, y", Medium: "MMM d, y", Short: "M/d/yy"}}, Time: ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: ut.CalendarDateFormat{Full: "{1} 'at' {0}", Long: "{1} 'at' {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{1: "Jan", 3: "Mar", 8: "Aug", 6: "Jun", 7: "Jul", 9: "Sep", 10: "Oct", 11: "Nov", 12: "Dec", 2: "Feb", 4: "Apr", 5: "May"}, Narrow: ut.CalendarMonthFormatNameValue{10: "O", 8: "A", 11: "N", 12: "D", 1: "J", 7: "J", 2: "F", 6: "J", 9: "S", 3: "M", 4: "A", 5: "M"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{6: "June", 12: "December", 1: "January", 3: "March", 10: "October", 4: "April", 7: "July", 11: "November", 2: "February", 8: "August", 9: "September", 5: "May"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{5: "Fri", 6: "Sat", 0: "Sun", 1: "Mon", 2: "Tue", 3: "Wed", 4: "Thu"}, Narrow: ut.CalendarDayFormatNameValue{3: "W", 4: "T", 5: "F", 6: "S", 0: "S", 1: "M", 2: "T"}, Short: ut.CalendarDayFormatNameValue{2: "Tu", 3: "We", 4: "Th", 5: "Fr", 6: "Sa", 0: "Su", 1: "Mo"}, Wide: ut.CalendarDayFormatNameValue{2: "Tuesday", 3: "Wednesday", 4: "Thursday", 5: "Friday", 6: "Saturday", 0: "Sunday", 1: "Monday"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"evening1": "in the evening", "night1": "at night", "midnight": "midnight", "am": "AM", "noon": "noon", "pm": "PM", "morning1": "in the morning", "afternoon1": "in the afternoon"}, Narrow: ut.CalendarPeriodFormatNameValue{"pm": "p", "morning1": "in the morning", "afternoon1": "in the afternoon", "evening1": "in the evening", "night1": "at night", "midnight": "mi", "am": "a", "noon": "n"}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"afternoon1": "afternoon", "evening1": "evening", "night1": "night", "midnight": "midnight", "am": "AM", "noon": "noon", "pm": "PM", "morning1": "morning"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}}}}
