package en_TZ

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE, MMMM d, y", Long: "MMMM d, y", Medium: "MMM d, y", Short: "M/d/yy"}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} 'at' {0}", Long: "{1} 'at' {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{1: "Jan", 5: "May", 9: "Sep", 8: "Aug", 2: "Feb", 6: "Jun", 7: "Jul", 10: "Oct", 12: "Dec", 3: "Mar", 11: "Nov", 4: "Apr"}, Narrow: ut.CalendarMonthFormatNameValue{1: "J", 5: "M", 8: "A", 9: "S", 10: "O", 3: "M", 4: "A", 11: "N", 6: "J", 2: "F", 7: "J", 12: "D"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{9: "September", 1: "January", 4: "April", 12: "December", 3: "March", 11: "November", 6: "June", 7: "July", 8: "August", 2: "February", 5: "May", 10: "October"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{4: "Thu", 5: "Fri", 6: "Sat", 0: "Sun", 1: "Mon", 2: "Tue", 3: "Wed"}, Narrow: ut.CalendarDayFormatNameValue{5: "F", 6: "S", 0: "S", 1: "M", 2: "T", 3: "W", 4: "T"}, Short: ut.CalendarDayFormatNameValue{3: "We", 4: "Th", 5: "Fr", 6: "Sa", 0: "Su", 1: "Mo", 2: "Tu"}, Wide: ut.CalendarDayFormatNameValue{4: "Thursday", 5: "Friday", 6: "Saturday", 0: "Sunday", 1: "Monday", 2: "Tuesday", 3: "Wednesday"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"evening1": "in the evening", "night1": "at night", "midnight": "midnight", "am": "AM", "noon": "noon", "pm": "PM", "morning1": "in the morning", "afternoon1": "in the afternoon"}, Narrow: ut.CalendarPeriodFormatNameValue{"pm": "p", "morning1": "in the morning", "afternoon1": "in the afternoon", "evening1": "in the evening", "night1": "at night", "midnight": "mi", "am": "a", "noon": "n"}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"evening1": "evening", "night1": "night", "midnight": "midnight", "am": "AM", "noon": "noon", "pm": "PM", "morning1": "morning", "afternoon1": "afternoon"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}}}}
