package fa

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "y/M/d"}, Time: ut.CalendarDateFormat{Full: "H:mm:ss (zzzz)", Long: "H:mm:ss (z)", Medium: "H:mm:ss", Short: "H:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1}، ساعت {0}", Long: "{1}، ساعت {0}", Medium: "{1}،\u200f {0}", Short: "{1}،\u200f {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{8: "اوت", 10: "اکتبر", 11: "نوامبر", 2: "فوریه", 3: "مارس", 5: "مه", 7: "ژوئیه", 12: "دسامبر", 1: "ژانویه", 4: "آوریل", 6: "ژوئن", 9: "سپتامبر"}, Narrow: ut.CalendarMonthFormatNameValue{3: "م", 9: "س", 10: "ا", 11: "ن", 12: "د", 1: "ژ", 2: "ف", 4: "آ", 5: "م", 6: "ژ", 7: "ژ", 8: "ا"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{3: "مارس", 4: "آوریل", 5: "مه", 6: "ژوئن", 7: "ژوئیه", 11: "نوامبر", 1: "ژانویه", 2: "فوریه", 8: "اوت", 9: "سپتامبر", 10: "اکتبر", 12: "دسامبر"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{6: "شنبه", 0: "یکشنبه", 1: "دوشنبه", 2: "سه\u200cشنبه", 3: "چهارشنبه", 4: "پنجشنبه", 5: "جمعه"}, Narrow: ut.CalendarDayFormatNameValue{4: "پ", 5: "ج", 6: "ش", 0: "ی", 1: "د", 2: "س", 3: "چ"}, Short: ut.CalendarDayFormatNameValue{0: "۱ش", 1: "۲ش", 2: "۳ش", 3: "۴ش", 4: "۵ش", 5: "ج", 6: "ش"}, Wide: ut.CalendarDayFormatNameValue{0: "یکشنبه", 1: "دوشنبه", 2: "سه\u200cشنبه", 3: "چهارشنبه", 4: "پنجشنبه", 5: "جمعه", 6: "شنبه"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"evening1": "سر شب", "night1": "شب", "midnight": "نیمه\u200cشب", "am": "ق.ظ.", "noon": "ظ", "pm": "ب.ظ.", "morning1": "صبح", "afternoon1": "عصر"}, Narrow: ut.CalendarPeriodFormatNameValue{"pm": "ب.ظ.", "morning1": "صبح", "afternoon1": "عصر", "evening1": "سر شب", "night1": "شب", "midnight": "نیمه\u200cشب", "am": "ق.ظ.", "noon": "ظهر"}, Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"am": "قبل\u200cازظهر", "noon": "ظهر", "pm": "بعدازظهر", "morning1": "صبح", "afternoon1": "عصر", "evening1": "سر شب", "night1": "شب", "midnight": "نیمه\u200cشب"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "میلادی", Abbrev: "م.", Narrow: "م"}, BC: ut.CalendarEraFormatNames{Full: "قبل از میلاد", Abbrev: "ق.م.", Narrow: "ق"}}}}
