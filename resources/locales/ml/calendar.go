package ml

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "y, MMMM d, EEEE", Long: "y, MMMM d", Medium: "y, MMM d", Short: "d/M/yy"}, Time: ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{8: "ഓഗ", 1: "ജനു", 2: "ഫെബ്രു", 5: "മേയ്", 7: "ജൂലൈ", 9: "സെപ്റ്റം", 10: "ഒക്ടോ", 11: "നവം", 12: "ഡിസം", 3: "മാർ", 4: "ഏപ്രി", 6: "ജൂൺ"}, Narrow: ut.CalendarMonthFormatNameValue{4: "ഏ", 7: "ജൂ", 8: "ഓ", 9: "സെ", 10: "ഒ", 11: "ന", 12: "ഡി", 3: "മാ", 2: "ഫെ", 5: "മെ", 6: "ജൂ", 1: "ജ"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{9: "സെപ്റ്റംബർ", 11: "നവംബർ", 1: "ജനുവരി", 2: "ഫെബ്രുവരി", 7: "ജൂലൈ", 8: "ഓഗസ്റ്റ്", 10: "ഒക്\u200cടോബർ", 12: "ഡിസംബർ", 3: "മാർച്ച്", 4: "ഏപ്രിൽ", 5: "മേയ്", 6: "ജൂൺ"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{5: "വെള്ളി", 6: "ശനി", 0: "ഞായർ", 1: "തിങ്കൾ", 2: "ചൊവ്വ", 3: "ബുധൻ", 4: "വ്യാഴം"}, Narrow: ut.CalendarDayFormatNameValue{2: "ചൊ", 3: "ബു", 4: "വ്യാ", 5: "വെ", 6: "ശ", 0: "ഞാ", 1: "തി"}, Short: ut.CalendarDayFormatNameValue{4: "വ്യാ", 5: "വെ", 6: "ശ", 0: "ഞാ", 1: "തി", 2: "ചൊ", 3: "ബു"}, Wide: ut.CalendarDayFormatNameValue{0: "ഞായറാഴ്\u200cച", 1: "തിങ്കളാഴ്\u200cച", 2: "ചൊവ്വാഴ്\u200cച", 3: "ബുധനാഴ്\u200cച", 4: "വ്യാഴാഴ്\u200cച", 5: "വെള്ളിയാഴ്\u200cച", 6: "ശനിയാഴ്\u200cച"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"afternoon2": "ഉച്ചതിരിഞ്ഞ്", "midnight": "അർദ്ധരാത്രി", "am": "AM", "noon": "ഉച്ച", "pm": "PM", "morning1": "പുലർച്ചെ", "morning2": "രാവിലെ", "afternoon1": "ഉച്ചയ്ക്ക്", "evening1": "വൈകുന്നേരം", "evening2": "സന്ധ്യ", "night1": "രാത്രി"}, Narrow: ut.CalendarPeriodFormatNameValue{"morning1": "പുലർച്ചെ", "afternoon2": "ഉച്ചതിരിഞ്ഞ്", "evening1": "വൈകുന്നേരം", "evening2": "സന്ധ്യ", "midnight": "അർദ്ധരാത്രി", "am": "AM", "noon": "ഉച്ച", "pm": "PM", "morning2": "രാവിലെ", "afternoon1": "ഉച്ചയ്ക്ക്", "night1": "രാത്രി"}, Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"am": "AM", "noon": "ഉച്ച", "morning1": "പുലർച്ചെ", "afternoon2": "ഉച്ചതിരിഞ്ഞ്", "evening2": "സന്ധ്യ", "night1": "രാത്രി", "midnight": "അർദ്ധരാത്രി", "pm": "PM", "morning2": "രാവിലെ", "afternoon1": "ഉച്ചയ്ക്ക്", "evening1": "വൈകുന്നേരം"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "ആന്നോ ഡൊമിനി", Abbrev: "എഡി", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "ക്രിസ്\u200cതുവിന് മുമ്പ്", Abbrev: "ക്രി.മു.", Narrow: ""}}}}
