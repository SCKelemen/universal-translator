package smn

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "", Long: "", Medium: "", Short: ""}, Time: ut.CalendarDateFormat{Full: "", Long: "", Medium: "", Short: ""}, DateTime: ut.CalendarDateFormat{Full: "", Long: "", Medium: "", Short: ""}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue(nil), Narrow: ut.CalendarMonthFormatNameValue(nil), Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{2: "kuovâmáánu", 4: "cuáŋuimáánu", 5: "vyesimáánu", 7: "syeinimáánu", 8: "porgemáánu", 9: "čohčâmáánu", 1: "uđđâivemáánu", 3: "njuhčâmáánu", 6: "kesimáánu", 10: "roovvâdmáánu", 11: "skammâmáánu", 12: "juovlâmáánu"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{2: "ma", 3: "ko", 4: "tu", 5: "vá", 6: "lá", 0: "pa", 1: "vu"}, Narrow: ut.CalendarDayFormatNameValue{4: "T", 5: "V", 6: "L", 0: "P", 1: "V", 2: "M", 3: "K"}, Short: ut.CalendarDayFormatNameValue(nil), Wide: ut.CalendarDayFormatNameValue{0: "pasepeivi", 1: "vuossargâ", 2: "majebargâ", 3: "koskokko", 4: "tuorâstâh", 5: "vástuppeivi", 6: "lávurdâh"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue(nil), Narrow: ut.CalendarPeriodFormatNameValue(nil), Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue(nil)}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}}}}
