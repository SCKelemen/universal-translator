package kok

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "dd-MM-y", Short: "d-M-yy"}, Time: ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: ut.CalendarDateFormat{Full: "", Long: "", Medium: "", Short: ""}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue(nil), Narrow: ut.CalendarMonthFormatNameValue(nil), Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{3: "मार्च", 8: "ओगस्ट", 9: "सेप्टेंबर", 11: "नोव्हेंबर", 6: "जून", 7: "जुलै", 10: "ओक्टोबर", 12: "डिसेंबर", 1: "जानेवारी", 2: "फेब्रुवारी", 4: "एप्रिल", 5: "मे"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{2: "मंगळ", 3: "बुध", 4: "गुरु", 5: "शुक्र", 6: "शनि", 0: "रवि", 1: "सोम"}, Narrow: ut.CalendarDayFormatNameValue(nil), Short: ut.CalendarDayFormatNameValue(nil), Wide: ut.CalendarDayFormatNameValue{4: "गुरुवार", 5: "शुक्रवार", 6: "शनिवार", 0: "आदित्यवार", 1: "सोमवार", 2: "मंगळार", 3: "बुधवार"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"am": "म.पू.", "pm": "म.नं."}, Narrow: ut.CalendarPeriodFormatNameValue(nil), Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"am": "म.पू.", "pm": "म.नं."}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "", Abbrev: "क्रिस्तपूर्व", Narrow: ""}}}}
