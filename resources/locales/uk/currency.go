package uk

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"MYR": ut.Currency{Currency: "MYR", DisplayName: "малайзійський рингіт", Symbol: "MYR"}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "нігерійська найра", Symbol: "NGN"}, "XFU": ut.Currency{Currency: "XFU", DisplayName: "французький франк UIC", Symbol: ""}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "китайський юань", Symbol: "CNY"}, "ECS": ut.Currency{Currency: "ECS", DisplayName: "еквадорський сукре", Symbol: ""}, "MOP": ut.Currency{Currency: "MOP", DisplayName: "патака Макао", Symbol: "MOP"}, "BAD": ut.Currency{Currency: "BAD", DisplayName: "динар (Боснія і Герцеговина)", Symbol: ""}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "малагасійський аріарі", Symbol: "MGA"}, "NPR": ut.Currency{Currency: "NPR", DisplayName: "непальська рупія", Symbol: "NPR"}, "USD": ut.Currency{Currency: "USD", DisplayName: "долар США", Symbol: "USD"}, "BZD": ut.Currency{Currency: "BZD", DisplayName: "белізький долар", Symbol: "BZD"}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "ескудо Кабо-Верде", Symbol: "CVE"}, "KPW": ut.Currency{Currency: "KPW", DisplayName: "північнокорейський вон", Symbol: "KPW"}, "TJR": ut.Currency{Currency: "TJR", DisplayName: "таджицький рубль", Symbol: ""}, "YUD": ut.Currency{Currency: "YUD", DisplayName: "югославський твердий динар", Symbol: ""}, "AFA": ut.Currency{Currency: "AFA", DisplayName: "афгані (1927–2002)", Symbol: ""}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "руандійський франк", Symbol: "RWF"}, "WST": ut.Currency{Currency: "WST", DisplayName: "самоанська тала", Symbol: "WST"}, "COU": ut.Currency{Currency: "COU", DisplayName: "одиниця реальної вартості", Symbol: ""}, "GRD": ut.Currency{Currency: "GRD", DisplayName: "грецька драхма", Symbol: ""}, "JMD": ut.Currency{Currency: "JMD", DisplayName: "ямайський долар", Symbol: "JMD"}, "GHC": ut.Currency{Currency: "GHC", DisplayName: "ганський седі (1979–2007)", Symbol: ""}, "BGN": ut.Currency{Currency: "BGN", DisplayName: "болгарський лев", Symbol: "BGN"}, "DOP": ut.Currency{Currency: "DOP", DisplayName: "домініканський песо", Symbol: "DOP"}, "FKP": ut.Currency{Currency: "FKP", DisplayName: "фунт Фолклендських островів", Symbol: "FKP"}, "PEI": ut.Currency{Currency: "PEI", DisplayName: "перуанський інті", Symbol: ""}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "тонганська паанга", Symbol: "TOP"}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "свазілендський лілангені", Symbol: "SZL"}, "XBA": ut.Currency{Currency: "XBA", DisplayName: "європейська складена валютна одиниця", Symbol: ""}, "HKD": ut.Currency{Currency: "HKD", DisplayName: "гонконгський долар", Symbol: "HKD"}, "MXV": ut.Currency{Currency: "MXV", DisplayName: "мексиканський юнідад де інверсіон", Symbol: ""}, "ROL": ut.Currency{Currency: "ROL", DisplayName: "старий румунський лей", Symbol: ""}, "MZN": ut.Currency{Currency: "MZN", DisplayName: "мозамбіцький метикал", Symbol: "MZN"}, "OMR": ut.Currency{Currency: "OMR", DisplayName: "оманський ріал", Symbol: "OMR"}, "USN": ut.Currency{Currency: "USN", DisplayName: "долар США (наступного дня)", Symbol: ""}, "ZWR": ut.Currency{Currency: "ZWR", DisplayName: "зімбабвійський долар (2008)", Symbol: ""}, "ANG": ut.Currency{Currency: "ANG", DisplayName: "нідерландський антильський гульден", Symbol: "ANG"}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "алжирський динар", Symbol: "DZD"}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "маврикійська рупія", Symbol: "MUR"}, "MXN": ut.Currency{Currency: "MXN", DisplayName: "мексиканський песо", Symbol: "MXN"}, "NIC": ut.Currency{Currency: "NIC", DisplayName: "нікарагуанська кордоба (1988–1991)", Symbol: ""}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "австралійський долар", Symbol: "AUD"}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "конвертована марка Боснії і Герцеговини", Symbol: "BAM"}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "швейцарський франк", Symbol: "CHF"}, "XAG": ut.Currency{Currency: "XAG", DisplayName: "срібло", Symbol: ""}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "ангольська кванза", Symbol: "AOA"}, "CLP": ut.Currency{Currency: "CLP", DisplayName: "чилійський песо", Symbol: "CLP"}, "GNS": ut.Currency{Currency: "GNS", DisplayName: "гвінейське сілі", Symbol: ""}, "AOR": ut.Currency{Currency: "AOR", DisplayName: "ангольська кванза реаджастадо (1995–1999)", Symbol: ""}, "HUF": ut.Currency{Currency: "HUF", DisplayName: "угорський форинт", Symbol: "HUF"}, "TND": ut.Currency{Currency: "TND", DisplayName: "туніський динар", Symbol: "TND"}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "угандійський шилінг", Symbol: "UGX"}, "ECV": ut.Currency{Currency: "ECV", DisplayName: "еквадорський юнідад де валор константе", Symbol: ""}, "GQE": ut.Currency{Currency: "GQE", DisplayName: "еквеле (Екваторіальна Ґвінея)", Symbol: ""}, "KHR": ut.Currency{Currency: "KHR", DisplayName: "камбоджійський рієль", Symbol: "KHR"}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "гвінейський франк", Symbol: "GNF"}, "SSP": ut.Currency{Currency: "SSP", DisplayName: "південносуданський фунт", Symbol: "SSP"}, "XEU": ut.Currency{Currency: "XEU", DisplayName: "європейська валютна одиниця", Symbol: ""}, "BEF": ut.Currency{Currency: "BEF", DisplayName: "бельгійський франк", Symbol: ""}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "бурундійський франк", Symbol: "BIF"}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "канадський долар", Symbol: "CAD"}, "UAH": ut.Currency{Currency: "UAH", DisplayName: "українська гривня", Symbol: "₴"}, "XAU": ut.Currency{Currency: "XAU", DisplayName: "золото", Symbol: ""}, "FIM": ut.Currency{Currency: "FIM", DisplayName: "фінляндська марка", Symbol: ""}, "NIO": ut.Currency{Currency: "NIO", DisplayName: "нікарагуанська кордоба", Symbol: "NIO"}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "саудівський ріал", Symbol: "SAR"}, "NZD": ut.Currency{Currency: "NZD", DisplayName: "новозеландський долар", Symbol: "NZD"}, "XBB": ut.Currency{Currency: "XBB", DisplayName: "одиниця європейського валютного фонду", Symbol: ""}, "BYR": ut.Currency{Currency: "BYR", DisplayName: "білоруський рубль", Symbol: "BYR"}, "LBP": ut.Currency{Currency: "LBP", DisplayName: "ліванський фунт", Symbol: "LBP"}, "MXP": ut.Currency{Currency: "MXP", DisplayName: "мексиканське срібне песо (1861–1992)", Symbol: ""}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "гамбійський даласі", Symbol: "GMD"}, "PES": ut.Currency{Currency: "PES", DisplayName: "перуанський сол", Symbol: ""}, "QAR": ut.Currency{Currency: "QAR", DisplayName: "катарський ріал", Symbol: "QAR"}, "AMD": ut.Currency{Currency: "AMD", DisplayName: "вірменський драм", Symbol: "AMD"}, "BOB": ut.Currency{Currency: "BOB", DisplayName: "болівійський болівіано", Symbol: "BOB"}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "євро", Symbol: "EUR"}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "танзанійський шилінг", Symbol: "TZS"}, "UYU": ut.Currency{Currency: "UYU", DisplayName: "уругвайський песо", Symbol: "UYU"}, "BOP": ut.Currency{Currency: "BOP", DisplayName: "болівійське песо", Symbol: ""}, "MZE": ut.Currency{Currency: "MZE", DisplayName: "мозамбіцький ескудо", Symbol: ""}, "SYP": ut.Currency{Currency: "SYP", DisplayName: "сирійський фунт", Symbol: "SYP"}, "MLF": ut.Currency{Currency: "MLF", DisplayName: "малійський франк", Symbol: ""}, "SKK": ut.Currency{Currency: "SKK", DisplayName: "словацька крона", Symbol: ""}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "південноафриканський ранд", Symbol: "ZAR"}, "ARS": ut.Currency{Currency: "ARS", DisplayName: "аргентинський песо", Symbol: "ARS"}, "BND": ut.Currency{Currency: "BND", DisplayName: "брунейський долар", Symbol: "BND"}, "FJD": ut.Currency{Currency: "FJD", DisplayName: "фіджійський долар", Symbol: "FJD"}, "CUP": ut.Currency{Currency: "CUP", DisplayName: "кубинський песо", Symbol: "CUP"}, "MKD": ut.Currency{Currency: "MKD", DisplayName: "македонський денар", Symbol: "MKD"}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "невідома грошова одиниця", Symbol: ""}, "LUF": ut.Currency{Currency: "LUF", DisplayName: "люксембурзький франк", Symbol: ""}, "RHD": ut.Currency{Currency: "RHD", DisplayName: "родезійський долар", Symbol: ""}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "західноафриканський франк", Symbol: "CFA"}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "бразильський реал", Symbol: "BRL"}, "BSD": ut.Currency{Currency: "BSD", DisplayName: "багамський долар", Symbol: "BSD"}, "LTL": ut.Currency{Currency: "LTL", DisplayName: "литовський літ", Symbol: "LTL"}, "BRE": ut.Currency{Currency: "BRE", DisplayName: "бразильське крузейро (1990–1993)", Symbol: ""}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "конголезький франк", Symbol: "CDF"}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "леоне Сьєрра-Леоне", Symbol: "SLL"}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "мавританська угія", Symbol: "MRO"}, "UAK": ut.Currency{Currency: "UAK", DisplayName: "український карбованець", Symbol: "крб."}, "YER": ut.Currency{Currency: "YER", DisplayName: "єменський ріал", Symbol: "YER"}, "CHW": ut.Currency{Currency: "CHW", DisplayName: "франк WIR", Symbol: ""}, "COP": ut.Currency{Currency: "COP", DisplayName: "колумбійський песо", Symbol: "COP"}, "ITL": ut.Currency{Currency: "ITL", DisplayName: "італійська ліра", Symbol: ""}, "SRG": ut.Currency{Currency: "SRG", DisplayName: "суринамський гульден", Symbol: ""}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "бангладеська така", Symbol: "BDT"}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "єгипетський фунт", Symbol: "EGP"}, "GHS": ut.Currency{Currency: "GHS", DisplayName: "ганський седі", Symbol: "GHS"}, "FRF": ut.Currency{Currency: "FRF", DisplayName: "французький франк", Symbol: ""}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "кʼят Мʼянми", Symbol: "MMK"}, "PKR": ut.Currency{Currency: "PKR", DisplayName: "пакистанська рупія", Symbol: "PKR"}, "SBD": ut.Currency{Currency: "SBD", DisplayName: "долар Соломонових Островів", Symbol: "SBD"}, "XTS": ut.Currency{Currency: "XTS", DisplayName: "код тестування валюти", Symbol: ""}, "ALL": ut.Currency{Currency: "ALL", DisplayName: "албанський лек", Symbol: "ALL"}, "CHE": ut.Currency{Currency: "CHE", DisplayName: "євро WIR", Symbol: ""}, "CLF": ut.Currency{Currency: "CLF", DisplayName: "чилійський юнідадес де фоменто", Symbol: ""}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "замбійська квача", Symbol: "ZMW"}, "LUC": ut.Currency{Currency: "LUC", DisplayName: "люксембурґський франк (конвертований)", Symbol: ""}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "лівійський динар", Symbol: "LYD"}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "намібійський долар", Symbol: "NAD"}, "ZMK": ut.Currency{Currency: "ZMK", DisplayName: "замбійська квача (1968–2012)", Symbol: ""}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "бахрейнський динар", Symbol: "BHD"}, "EEK": ut.Currency{Currency: "EEK", DisplayName: "естонська крона", Symbol: ""}, "INR": ut.Currency{Currency: "INR", DisplayName: "індійська рупія", Symbol: "INR"}, "CZK": ut.Currency{Currency: "CZK", DisplayName: "чеська крона", Symbol: "CZK"}, "KRW": ut.Currency{Currency: "KRW", DisplayName: "південнокорейський вон", Symbol: "KRW"}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "фунт острова Святої Єлени", Symbol: "SHP"}, "AFN": ut.Currency{Currency: "AFN", DisplayName: "афганський афгані", Symbol: "AFN"}, "ESA": ut.Currency{Currency: "ESA", DisplayName: "іспанська песета (\"А\" рахунок)", Symbol: ""}, "IDR": ut.Currency{Currency: "IDR", DisplayName: "індонезійська рупія", Symbol: "IDR"}, "UYI": ut.Currency{Currency: "UYI", DisplayName: "уругвайський песо в індексованих одиницях", Symbol: ""}, "XPT": ut.Currency{Currency: "XPT", DisplayName: "платина", Symbol: ""}, "DEM": ut.Currency{Currency: "DEM", DisplayName: "німецька марка", Symbol: ""}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "казахстанський тенге", Symbol: "KZT"}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "марокканський дирхам", Symbol: "MAD"}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "норвезька крона", Symbol: "NOK"}, "PGK": ut.Currency{Currency: "PGK", DisplayName: "кіна Папуа Нової Гвінеї", Symbol: "PGK"}, "XBD": ut.Currency{Currency: "XBD", DisplayName: "європейська розрахункова одиниця XBD", Symbol: ""}, "GWE": ut.Currency{Currency: "GWE", DisplayName: "ескудо португальської гвінеї", Symbol: ""}, "HRK": ut.Currency{Currency: "HRK", DisplayName: "хорватська куна", Symbol: "HRK"}, "JOD": ut.Currency{Currency: "JOD", DisplayName: "йорданський динар", Symbol: "JOD"}, "IRR": ut.Currency{Currency: "IRR", DisplayName: "іранський ріал", Symbol: "IRR"}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "шрі-ланкійська рупія", Symbol: "LKR"}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "польський злотий", Symbol: "PLN"}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "російський рубль", Symbol: "RUB"}, "TMT": ut.Currency{Currency: "TMT", DisplayName: "туркменський манат", Symbol: "TMT"}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "ботсванська пула", Symbol: "BWP"}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "джибутійський франк", Symbol: "DJF"}, "ILS": ut.Currency{Currency: "ILS", DisplayName: "ізраїльський новий шекель", Symbol: "ILS"}, "PHP": ut.Currency{Currency: "PHP", DisplayName: "філіппінський песо", Symbol: "PHP"}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "шведська крона", Symbol: "SEK"}, "ZAL": ut.Currency{Currency: "ZAL", DisplayName: "південноафриканський фінансовий ранд", Symbol: ""}, "CSK": ut.Currency{Currency: "CSK", DisplayName: "чехословацька тверда крона", Symbol: ""}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "японська єна", Symbol: "¥"}, "KYD": ut.Currency{Currency: "KYD", DisplayName: "долар Кайманових островів", Symbol: "KYD"}, "PYG": ut.Currency{Currency: "PYG", DisplayName: "парагвайський гуарані", Symbol: "PYG"}, "TJS": ut.Currency{Currency: "TJS", DisplayName: "таджицький сомоні", Symbol: "TJS"}, "BEC": ut.Currency{Currency: "BEC", DisplayName: "бельгійський франк (конвертований)", Symbol: ""}, "MAF": ut.Currency{Currency: "MAF", DisplayName: "марокканський франк", Symbol: ""}, "MDL": ut.Currency{Currency: "MDL", DisplayName: "молдовський лей", Symbol: "MDL"}, "PLZ": ut.Currency{Currency: "PLZ", DisplayName: "польський злотий (1950–1995)", Symbol: ""}, "GEL": ut.Currency{Currency: "GEL", DisplayName: "грузинський ларі", Symbol: "GEL"}, "ILP": ut.Currency{Currency: "ILP", DisplayName: "ізраїльський фунт", Symbol: ""}, "LUL": ut.Currency{Currency: "LUL", DisplayName: "люксембурґський франк (фінансовий)", Symbol: ""}, "RUR": ut.Currency{Currency: "RUR", DisplayName: "російський рубль (1991–1998)", Symbol: ""}, "SUR": ut.Currency{Currency: "SUR", DisplayName: "радянський рубль", Symbol: ""}, "TPE": ut.Currency{Currency: "TPE", DisplayName: "тіморський ескудо", Symbol: ""}, "UYP": ut.Currency{Currency: "UYP", DisplayName: "уругвайське песо (1975–1993)", Symbol: ""}, "AON": ut.Currency{Currency: "AON", DisplayName: "ангольська нова кванза (1990–2000)", Symbol: ""}, "BBD": ut.Currency{Currency: "BBD", DisplayName: "барбадоський долар", Symbol: "BBD"}, "MVR": ut.Currency{Currency: "MVR", DisplayName: "мальдівська руфія", Symbol: "MVR"}, "CUC": ut.Currency{Currency: "CUC", DisplayName: "кубинський конвертований песо", Symbol: "CUC"}, "YDD": ut.Currency{Currency: "YDD", DisplayName: "єменський динар", Symbol: ""}, "SVC": ut.Currency{Currency: "SVC", DisplayName: "сальвадорський колон", Symbol: ""}, "VEF": ut.Currency{Currency: "VEF", DisplayName: "венесуельський болівар", Symbol: "VEF"}, "XFO": ut.Currency{Currency: "XFO", DisplayName: "французький золотий франк", Symbol: ""}, "BOV": ut.Currency{Currency: "BOV", DisplayName: "болівійський мвдол", Symbol: ""}, "LVR": ut.Currency{Currency: "LVR", DisplayName: "латвійський рубль", Symbol: ""}, "SIT": ut.Currency{Currency: "SIT", DisplayName: "словенський толар", Symbol: ""}, "AZM": ut.Currency{Currency: "AZM", DisplayName: "азербайджанський манат (1993–2006)", Symbol: ""}, "PTE": ut.Currency{Currency: "PTE", DisplayName: "португальський ескудо", Symbol: ""}, "STD": ut.Currency{Currency: "STD", DisplayName: "добра Сан-Томе і Принсіпі", Symbol: "STD"}, "IQD": ut.Currency{Currency: "IQD", DisplayName: "іракський динар", Symbol: "IQD"}, "LSL": ut.Currency{Currency: "LSL", DisplayName: "лесотський лоті", Symbol: ""}, "RSD": ut.Currency{Currency: "RSD", DisplayName: "сербський динар", Symbol: "RSD"}, "BEL": ut.Currency{Currency: "BEL", DisplayName: "бельгійський франк (фінансовий)", Symbol: ""}, "BRN": ut.Currency{Currency: "BRN", DisplayName: "бразильське нове крузадо", Symbol: ""}, "BUK": ut.Currency{Currency: "BUK", DisplayName: "бірманський кіат", Symbol: ""}, "TMM": ut.Currency{Currency: "TMM", DisplayName: "туркменський манат (1993–2009)", Symbol: ""}, "USS": ut.Currency{Currency: "USS", DisplayName: "долар США (цього дня)", Symbol: ""}, "ZWD": ut.Currency{Currency: "ZWD", DisplayName: "зімбабвійський долар", Symbol: ""}, "CRC": ut.Currency{Currency: "CRC", DisplayName: "костариканський колон", Symbol: "CRC"}, "SDD": ut.Currency{Currency: "SDD", DisplayName: "суданський динар", Symbol: ""}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "сінгапурський долар", Symbol: "SGD"}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "монгольський тугрик", Symbol: "MNT"}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "сейшельська рупія", Symbol: "SCR"}, "VUV": ut.Currency{Currency: "VUV", DisplayName: "вануатський вату", Symbol: "VUV"}, "XRE": ut.Currency{Currency: "XRE", DisplayName: "фонди RINET", Symbol: ""}, "ARP": ut.Currency{Currency: "ARP", DisplayName: "аргентинський песо (1983–1985)", Symbol: ""}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "англійський фунт", Symbol: "GBP"}, "HRD": ut.Currency{Currency: "HRD", DisplayName: "хорватський динар", Symbol: ""}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "еритрейська накфа", Symbol: "ERN"}, "GEK": ut.Currency{Currency: "GEK", DisplayName: "грузинський купон", Symbol: ""}, "GIP": ut.Currency{Currency: "GIP", DisplayName: "гібралтарський фунт", Symbol: "GIP"}, "MGF": ut.Currency{Currency: "MGF", DisplayName: "мадагаскарський франк", Symbol: ""}, "ARA": ut.Currency{Currency: "ARA", DisplayName: "аргентинський австрал", Symbol: ""}, "BGL": ut.Currency{Currency: "BGL", DisplayName: "болгарський твердий лев", Symbol: ""}, "BRC": ut.Currency{Currency: "BRC", DisplayName: "бразильське крузадо", Symbol: ""}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "центральноафриканський франк", Symbol: "FCFA"}, "ZRZ": ut.Currency{Currency: "ZRZ", DisplayName: "заїрський заїр", Symbol: ""}, "ATS": ut.Currency{Currency: "ATS", DisplayName: "австрійський шилінг", Symbol: ""}, "LVL": ut.Currency{Currency: "LVL", DisplayName: "латвійський лат", Symbol: "LVL"}, "MTP": ut.Currency{Currency: "MTP", DisplayName: "мальтійський фунт", Symbol: ""}, "KGS": ut.Currency{Currency: "KGS", DisplayName: "киргизький сом", Symbol: "KGS"}, "SRD": ut.Currency{Currency: "SRD", DisplayName: "суринамський долар", Symbol: "SRD"}, "BMD": ut.Currency{Currency: "BMD", DisplayName: "бермудський долар", Symbol: "BMD"}, "BRR": ut.Currency{Currency: "BRR", DisplayName: "бразильське крузейро", Symbol: ""}, "HNL": ut.Currency{Currency: "HNL", DisplayName: "гондураська лемпіра", Symbol: "HNL"}, "TRY": ut.Currency{Currency: "TRY", DisplayName: "турецька ліра", Symbol: "TRY"}, "ADP": ut.Currency{Currency: "ADP", DisplayName: "андоррська песета", Symbol: ""}, "BTN": ut.Currency{Currency: "BTN", DisplayName: "бутанський нгултрум", Symbol: "BTN"}, "MZM": ut.Currency{Currency: "MZM", DisplayName: "старий мозамбіцький метикал", Symbol: ""}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "данська крона", Symbol: "DKK"}, "RON": ut.Currency{Currency: "RON", DisplayName: "румунський лей", Symbol: "RON"}, "UGS": ut.Currency{Currency: "UGS", DisplayName: "угандійський шилінг (1966–1987)", Symbol: ""}, "ZRN": ut.Currency{Currency: "ZRN", DisplayName: "заїрський новий заїр", Symbol: ""}, "ZWL": ut.Currency{Currency: "ZWL", DisplayName: "зімбабвійський долар (2009)", Symbol: ""}, "IEP": ut.Currency{Currency: "IEP", DisplayName: "ірландський фунт", Symbol: ""}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "коморський франк", Symbol: "KMF"}, "SDP": ut.Currency{Currency: "SDP", DisplayName: "старий суданський фунт", Symbol: ""}, "YUN": ut.Currency{Currency: "YUN", DisplayName: "югославський конвертований динар", Symbol: ""}, "ESP": ut.Currency{Currency: "ESP", DisplayName: "іспанська песета", Symbol: ""}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "східнокарибський долар", Symbol: "XCD"}, "BYB": ut.Currency{Currency: "BYB", DisplayName: "білоруський новий рубль (1994–1999)", Symbol: ""}, "TWD": ut.Currency{Currency: "TWD", DisplayName: "новий тайванський долар", Symbol: "TWD"}, "VND": ut.Currency{Currency: "VND", DisplayName: "вʼєтнамський донг", Symbol: "VND"}, "LTT": ut.Currency{Currency: "LTT", DisplayName: "литовський талон", Symbol: ""}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "малавійська квача", Symbol: "MWK"}, "NLG": ut.Currency{Currency: "NLG", DisplayName: "нідерландський гульден", Symbol: ""}, "AZN": ut.Currency{Currency: "AZN", DisplayName: "азербайджанський манат", Symbol: "AZN"}, "BRB": ut.Currency{Currency: "BRB", DisplayName: "бразильське нове крузейро (1967–1986)", Symbol: ""}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "ефіопський бир", Symbol: "ETB"}, "XPD": ut.Currency{Currency: "XPD", DisplayName: "паладій", Symbol: ""}, "YUM": ut.Currency{Currency: "YUM", DisplayName: "югославський новий динар", Symbol: ""}, "AED": ut.Currency{Currency: "AED", DisplayName: "дирхам ОАЕ", Symbol: "AED"}, "GTQ": ut.Currency{Currency: "GTQ", DisplayName: "гватемальський кетсаль", Symbol: "GTQ"}, "KES": ut.Currency{Currency: "KES", DisplayName: "кенійський шилінг", Symbol: "KES"}, "MTL": ut.Currency{Currency: "MTL", DisplayName: "мальтійська ліра", Symbol: ""}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "сомалійський шилінг", Symbol: "SOS"}, "TTD": ut.Currency{Currency: "TTD", DisplayName: "долар Тринідаду і Тобаго", Symbol: "TTD"}, "UZS": ut.Currency{Currency: "UZS", DisplayName: "узбецький сум", Symbol: "UZS"}, "GWP": ut.Currency{Currency: "GWP", DisplayName: "песо Гвінеї-Бісау", Symbol: ""}, "ISK": ut.Currency{Currency: "ISK", DisplayName: "ісландська крона", Symbol: "ISK"}, "KWD": ut.Currency{Currency: "KWD", DisplayName: "кувейтський динар", Symbol: "KWD"}, "PAB": ut.Currency{Currency: "PAB", DisplayName: "панамська бальбоа", Symbol: "PAB"}, "XBC": ut.Currency{Currency: "XBC", DisplayName: "європейська розрахункова одиниця XBC", Symbol: ""}, "DDM": ut.Currency{Currency: "DDM", DisplayName: "марка НДР", Symbol: ""}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "ліберійський долар", Symbol: "LRD"}, "TRL": ut.Currency{Currency: "TRL", DisplayName: "турецька ліра (1922–2005)", Symbol: ""}, "THB": ut.Currency{Currency: "THB", DisplayName: "таїландський бат", Symbol: "THB"}, "XDR": ut.Currency{Currency: "XDR", DisplayName: "спеціальні права запозичення", Symbol: ""}, "XPF": ut.Currency{Currency: "XPF", DisplayName: "французький тихоокеанський франк", Symbol: "CFPF"}, "GYD": ut.Currency{Currency: "GYD", DisplayName: "гаянський долар", Symbol: "GYD"}, "HTG": ut.Currency{Currency: "HTG", DisplayName: "гаїтянський гурд", Symbol: "HTG"}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "суданський фунт", Symbol: "SDG"}, "AWG": ut.Currency{Currency: "AWG", DisplayName: "арубський флорин", Symbol: "AWG"}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "лаоський кіп", Symbol: "LAK"}, "VEB": ut.Currency{Currency: "VEB", DisplayName: "венесуельський болівар (1871–2008)", Symbol: ""}, "AOK": ut.Currency{Currency: "AOK", DisplayName: "ангольська кванза (1977–1990)", Symbol: ""}, "CSD": ut.Currency{Currency: "CSD", DisplayName: "старий сербський динар", Symbol: ""}, "ESB": ut.Currency{Currency: "ESB", DisplayName: "іспанська песета (конвертовані рахунки)", Symbol: ""}, "CYP": ut.Currency{Currency: "CYP", DisplayName: "кіпрський фунт", Symbol: ""}, "PEN": ut.Currency{Currency: "PEN", DisplayName: "перуанський новий сол", Symbol: "PEN"}}
