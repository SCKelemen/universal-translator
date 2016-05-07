package en_DM

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE, MMMM d, y GGGG", Long: "MMMM d, y GG", Medium: "MMM d, y", Short: "M/d/yy"}, AD: ut.CalendarDateFormat{Full: "EEEE, MMMM d, y", Long: "MMMM d, y", Medium: "MMM d, y", Short: "M/d/yy"}}, Time: ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: ut.CalendarDateFormat{Full: "{1} 'at' {0}", Long: "{1} 'at' {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{1: "Jan", 2: "Feb", 5: "May", 9: "Sep", 8: "Aug", 11: "Nov", 12: "Dec", 6: "Jun", 7: "Jul", 10: "Oct", 3: "Mar", 4: "Apr"}, Narrow: ut.CalendarMonthFormatNameValue{3: "M", 8: "A", 12: "D", 9: "S", 10: "O", 1: "J", 4: "A", 5: "M", 7: "J", 11: "N", 2: "F", 6: "J"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{8: "August", 10: "October", 1: "January", 2: "February", 6: "June", 11: "November", 5: "May", 7: "July", 9: "September", 3: "March", 12: "December", 4: "April"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{3: "Wed", 4: "Thu", 5: "Fri", 6: "Sat", 0: "Sun", 1: "Mon", 2: "Tue"}, Narrow: ut.CalendarDayFormatNameValue{6: "S", 0: "S", 1: "M", 2: "T", 3: "W", 4: "T", 5: "F"}, Short: ut.CalendarDayFormatNameValue{3: "We", 4: "Th", 5: "Fr", 6: "Sa", 0: "Su", 1: "Mo", 2: "Tu"}, Wide: ut.CalendarDayFormatNameValue{0: "Sunday", 1: "Monday", 2: "Tuesday", 3: "Wednesday", 4: "Thursday", 5: "Friday", 6: "Saturday"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"am": "AM", "noon": "noon", "pm": "PM", "morning1": "in the morning", "afternoon1": "in the afternoon", "evening1": "in the evening", "night1": "at night", "midnight": "midnight"}, Narrow: ut.CalendarPeriodFormatNameValue{"morning1": "in the morning", "afternoon1": "in the afternoon", "evening1": "in the evening", "night1": "at night", "midnight": "mi", "am": "a", "noon": "n", "pm": "p"}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"pm": "PM", "morning1": "morning", "afternoon1": "afternoon", "evening1": "evening", "night1": "night", "midnight": "midnight", "am": "AM", "noon": "noon"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}}}}
