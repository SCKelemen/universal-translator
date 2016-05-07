package lt

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "y 'm'. MMMM d 'd'., EEEE", Long: "y 'm'. MMMM d 'd'.", Medium: "y-MM-dd", Short: "y-MM-dd"}, AD: ut.CalendarDateFormat{Full: "y 'm'. MMMM d 'd'., EEEE", Long: "y 'm'. MMMM d 'd'.", Medium: "y-MM-dd", Short: "y-MM-dd"}}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{10: "spal.", 11: "lapkr.", 12: "gruod.", 4: "bal.", 5: "geg.", 6: "birž.", 8: "rugp.", 9: "rugs.", 1: "saus.", 2: "vas.", 3: "kov.", 7: "liep."}, Narrow: ut.CalendarMonthFormatNameValue{1: "S", 3: "K", 5: "G", 8: "R", 9: "R", 10: "S", 11: "L", 2: "V", 4: "B", 6: "B", 7: "L", 12: "G"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{3: "kovas", 7: "liepa", 8: "rugpjūtis", 9: "rugsėjis", 10: "spalis", 12: "gruodis", 1: "sausis", 2: "vasaris", 4: "balandis", 5: "gegužė", 6: "birželis", 11: "lapkritis"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{2: "an", 3: "tr", 4: "kt", 5: "pn", 6: "št", 0: "sk", 1: "pr"}, Narrow: ut.CalendarDayFormatNameValue{4: "K", 5: "P", 6: "Š", 0: "S", 1: "P", 2: "A", 3: "T"}, Short: ut.CalendarDayFormatNameValue{0: "Sk", 1: "Pr", 2: "An", 3: "Tr", 4: "Kt", 5: "Pn", 6: "Št"}, Wide: ut.CalendarDayFormatNameValue{6: "šeštadienis", 0: "sekmadienis", 1: "pirmadienis", 2: "antradienis", 3: "trečiadienis", 4: "ketvirtadienis", 5: "penktadienis"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"pm": "popiet", "morning1": "rytas", "afternoon1": "diena", "evening1": "vakaras", "night1": "naktis", "midnight": "vidurnaktis", "am": "priešpiet", "noon": "vidurdienis"}, Narrow: ut.CalendarPeriodFormatNameValue{"am": "pr. p.", "noon": "vidurdienis", "pm": "pop.", "morning1": "rytas", "afternoon1": "diena", "evening1": "vakaras", "night1": "naktis", "midnight": "vidurnaktis"}, Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"afternoon1": "diena", "evening1": "vakaras", "night1": "naktis", "midnight": "vidurnaktis", "am": "priešpiet", "noon": "vidurdienis", "pm": "popiet", "morning1": "rytas"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "po Kristaus", Abbrev: "po Kr.", Narrow: "po Kr."}, BC: ut.CalendarEraFormatNames{Full: "prieš Kristų", Abbrev: "pr. Kr.", Narrow: "pr. Kr."}}}}
