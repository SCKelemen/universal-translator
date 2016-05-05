package en_NG

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE, MMMM d, y", Long: "MMMM d, y", Medium: "MMM d, y", Short: "M/d/yy"}, Time: ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: ut.CalendarDateFormat{Full: "{1} 'at' {0}", Long: "{1} 'at' {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{5: "May", 6: "Jun", 11: "Nov", 2: "Feb", 9: "Sep", 10: "Oct", 12: "Dec", 1: "Jan", 3: "Mar", 7: "Jul", 8: "Aug", 4: "Apr"}, Narrow: ut.CalendarMonthFormatNameValue{9: "S", 12: "D", 2: "F", 4: "A", 7: "J", 5: "M", 6: "J", 3: "M", 11: "N", 10: "O", 1: "J", 8: "A"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{12: "December", 2: "February", 8: "August", 9: "September", 1: "January", 10: "October", 11: "November", 3: "March", 6: "June", 7: "July", 4: "April", 5: "May"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{5: "Fri", 6: "Sat", 0: "Sun", 1: "Mon", 2: "Tue", 3: "Wed", 4: "Thu"}, Narrow: ut.CalendarDayFormatNameValue{2: "T", 3: "W", 4: "T", 5: "F", 6: "S", 0: "S", 1: "M"}, Short: ut.CalendarDayFormatNameValue{3: "We", 4: "Th", 5: "Fr", 6: "Sa", 0: "Su", 1: "Mo", 2: "Tu"}, Wide: ut.CalendarDayFormatNameValue{5: "Friday", 6: "Saturday", 0: "Sunday", 1: "Monday", 2: "Tuesday", 3: "Wednesday", 4: "Thursday"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"afternoon1": "in the afternoon", "evening1": "in the evening", "night1": "at night", "midnight": "midnight", "am": "AM", "noon": "noon", "pm": "PM", "morning1": "in the morning"}, Narrow: ut.CalendarPeriodFormatNameValue{"morning1": "in the morning", "afternoon1": "in the afternoon", "evening1": "in the evening", "night1": "at night", "midnight": "mi", "am": "a", "noon": "n", "pm": "p"}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"midnight": "midnight", "am": "AM", "noon": "noon", "pm": "PM", "morning1": "morning", "afternoon1": "afternoon", "evening1": "evening", "night1": "night"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}}}}
