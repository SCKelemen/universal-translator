package sr

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"MUR": ut.Currency{Currency: "MUR", DisplayName: "Маурицијска рупија", Symbol: "MUR"}, "PLZ": ut.Currency{Currency: "PLZ", DisplayName: "Пољски злоти (1950–1995)", Symbol: ""}, "QAR": ut.Currency{Currency: "QAR", DisplayName: "Катарски ријал", Symbol: "QAR"}, "RSD": ut.Currency{Currency: "RSD", DisplayName: "Српски динар", Symbol: "RSD"}, "RUR": ut.Currency{Currency: "RUR", DisplayName: "Руска рубља (1991–1998)", Symbol: ""}, "TJS": ut.Currency{Currency: "TJS", DisplayName: "Tаџихистански сомон", Symbol: "TJS"}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "Швајцарски франак", Symbol: "CHF"}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "Британска фунта", Symbol: "£"}, "UYP": ut.Currency{Currency: "UYP", DisplayName: "Уругвајски пезо (1975–1993)", Symbol: ""}, "ZWD": ut.Currency{Currency: "ZWD", DisplayName: "Зимбабвеански долар (1980–2008)", Symbol: ""}, "CUC": ut.Currency{Currency: "CUC", DisplayName: "Кубански конвертибилни пезос", Symbol: "CUC"}, "OMR": ut.Currency{Currency: "OMR", DisplayName: "Омански ријал", Symbol: "OMR"}, "EEK": ut.Currency{Currency: "EEK", DisplayName: "Естонска кроон", Symbol: ""}, "PGK": ut.Currency{Currency: "PGK", DisplayName: "Папуанска кина", Symbol: "PGK"}, "UAH": ut.Currency{Currency: "UAH", DisplayName: "Украјинска хривња", Symbol: "UAH"}, "AMD": ut.Currency{Currency: "AMD", DisplayName: "Јерменски драм", Symbol: "AMD"}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "Џибутански франак", Symbol: "DJF"}, "MKD": ut.Currency{Currency: "MKD", DisplayName: "Македонски денар", Symbol: "MKD"}, "YDD": ut.Currency{Currency: "YDD", DisplayName: "Јеменски динар", Symbol: ""}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "Кинески јуан", Symbol: "CN¥"}, "HKD": ut.Currency{Currency: "HKD", DisplayName: "Хонгконшки долар", Symbol: "HK$"}, "BMD": ut.Currency{Currency: "BMD", DisplayName: "Бермудски долар", Symbol: "BMD"}, "VEB": ut.Currency{Currency: "VEB", DisplayName: "Венецуелански боливар (1871–2008)", Symbol: ""}, "MXP": ut.Currency{Currency: "MXP", DisplayName: "Мексички сребрни пезо (1861–1992)", Symbol: ""}, "WST": ut.Currency{Currency: "WST", DisplayName: "Самоанска тала", Symbol: "WST"}, "XPT": ut.Currency{Currency: "XPT", DisplayName: "Платина", Symbol: ""}, "AOK": ut.Currency{Currency: "AOK", DisplayName: "Анголијска кванза (1977–1990)", Symbol: ""}, "FIM": ut.Currency{Currency: "FIM", DisplayName: "Финска марка", Symbol: ""}, "BRR": ut.Currency{Currency: "BRR", DisplayName: "Бразилски крузеиро", Symbol: ""}, "BUK": ut.Currency{Currency: "BUK", DisplayName: "Бурмански кјат", Symbol: ""}, "GIP": ut.Currency{Currency: "GIP", DisplayName: "Гибралтарска фунта", Symbol: "GIP"}, "ALL": ut.Currency{Currency: "ALL", DisplayName: "Албански лек", Symbol: "ALL"}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "Анголска кванза", Symbol: "AOA"}, "NZD": ut.Currency{Currency: "NZD", DisplayName: "Новозеландски долар", Symbol: "NZD"}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "Либеријски долар", Symbol: "LRD"}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "Норвешка круна", Symbol: "NOK"}, "GRD": ut.Currency{Currency: "GRD", DisplayName: "Грчка драхма", Symbol: ""}, "GYD": ut.Currency{Currency: "GYD", DisplayName: "Гвајански долар", Symbol: "GYD"}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "Лаошки кип", Symbol: "LAK"}, "VND": ut.Currency{Currency: "VND", DisplayName: "Вијетнамски донг", Symbol: "VND"}, "ARS": ut.Currency{Currency: "ARS", DisplayName: "Аргентински пезос", Symbol: "ARS"}, "ESP": ut.Currency{Currency: "ESP", DisplayName: "Шпанска пезета", Symbol: ""}, "PES": ut.Currency{Currency: "PES", DisplayName: "Перуански сол", Symbol: ""}, "SVC": ut.Currency{Currency: "SVC", DisplayName: "Салвадорски колон", Symbol: ""}, "UGS": ut.Currency{Currency: "UGS", DisplayName: "Угандски шилинг (1966–1987)", Symbol: ""}, "XFO": ut.Currency{Currency: "XFO", DisplayName: "Француски златни франак", Symbol: ""}, "LTL": ut.Currency{Currency: "LTL", DisplayName: "Литвански литас", Symbol: "LTL"}, "MCF": ut.Currency{Currency: "MCF", DisplayName: "Монегаскански франак", Symbol: ""}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "Сомалијски шилинг", Symbol: "SOS"}, "CNX": ut.Currency{Currency: "CNX", DisplayName: "Долар кинеске народне банке", Symbol: ""}, "PKR": ut.Currency{Currency: "PKR", DisplayName: "Пакистанскa рупиja", Symbol: "PKR"}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "Пољски злот", Symbol: "PLN"}, "YUR": ut.Currency{Currency: "YUR", DisplayName: "Југословенски реформирани динар", Symbol: ""}, "BRZ": ut.Currency{Currency: "BRZ", DisplayName: "Стари бразилски крузеиро", Symbol: ""}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "Малавијска квача", Symbol: "MWK"}, "LUC": ut.Currency{Currency: "LUC", DisplayName: "Луксембуршки конвертибилни франак", Symbol: ""}, "MAF": ut.Currency{Currency: "MAF", DisplayName: "Марокански франак", Symbol: ""}, "NIO": ut.Currency{Currency: "NIO", DisplayName: "Никарагванска златна кордоба", Symbol: "NIO"}, "SKK": ut.Currency{Currency: "SKK", DisplayName: "Словачка круна", Symbol: ""}, "UYU": ut.Currency{Currency: "UYU", DisplayName: "Уругвајски пезос", Symbol: "UYU"}, "YUD": ut.Currency{Currency: "YUD", DisplayName: "Југословенски тврди динар", Symbol: ""}, "FRF": ut.Currency{Currency: "FRF", DisplayName: "Француски франак", Symbol: ""}, "JOD": ut.Currency{Currency: "JOD", DisplayName: "Јордански динар", Symbol: "JOD"}, "IDR": ut.Currency{Currency: "IDR", DisplayName: "Индонежанска рупија", Symbol: "IDR"}, "USD": ut.Currency{Currency: "USD", DisplayName: "Амерички долар", Symbol: "US$"}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "Конгоански франак", Symbol: "CDF"}, "XPD": ut.Currency{Currency: "XPD", DisplayName: "Паладијум", Symbol: ""}, "BEF": ut.Currency{Currency: "BEF", DisplayName: "Белгијски франак", Symbol: ""}, "BOP": ut.Currency{Currency: "BOP", DisplayName: "Боливијски пезо", Symbol: ""}, "CSK": ut.Currency{Currency: "CSK", DisplayName: "Чехословачка тврда круна", Symbol: ""}, "KPW": ut.Currency{Currency: "KPW", DisplayName: "Севернокорејски вон", Symbol: "KPW"}, "SRD": ut.Currency{Currency: "SRD", DisplayName: "Суринамски долар", Symbol: "SRD"}, "XBB": ut.Currency{Currency: "XBB", DisplayName: "Европска новчана јединица", Symbol: ""}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "Eритрејска накфa", Symbol: "ERN"}, "ESB": ut.Currency{Currency: "ESB", DisplayName: "Шпанска пезета (конвертибилнирачун)", Symbol: ""}, "MXN": ut.Currency{Currency: "MXN", DisplayName: "Мексички пезос", Symbol: "MX$"}, "XFU": ut.Currency{Currency: "XFU", DisplayName: "Француски UIC-франак", Symbol: ""}, "ADP": ut.Currency{Currency: "ADP", DisplayName: "Андорска пезета", Symbol: ""}, "ECV": ut.Currency{Currency: "ECV", DisplayName: "Еквадорски унидад де валор константе", Symbol: ""}, "BGN": ut.Currency{Currency: "BGN", DisplayName: "Бугарски лев", Symbol: "[BGN]"}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "Алжирски динар", Symbol: "DZD"}, "KWD": ut.Currency{Currency: "KWD", DisplayName: "Кувајтски динар", Symbol: "KWD"}, "LUF": ut.Currency{Currency: "LUF", DisplayName: "Луксембуршки франак", Symbol: ""}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "Мауританијска oгија", Symbol: "MRO"}, "SUR": ut.Currency{Currency: "SUR", DisplayName: "Совјетска рубља", Symbol: ""}, "AFN": ut.Currency{Currency: "AFN", DisplayName: "Авганистански авгани", Symbol: "AFN"}, "BGL": ut.Currency{Currency: "BGL", DisplayName: "Бугарски тврди лев", Symbol: ""}, "TMT": ut.Currency{Currency: "TMT", DisplayName: "Туркменистански манат", Symbol: "TMT"}, "TRY": ut.Currency{Currency: "TRY", DisplayName: "Турска лира", Symbol: "TRY"}, "ILP": ut.Currency{Currency: "ILP", DisplayName: "Израелска фунта", Symbol: ""}, "KES": ut.Currency{Currency: "KES", DisplayName: "Кенијски шилинг", Symbol: "KES"}, "KYD": ut.Currency{Currency: "KYD", DisplayName: "Кајмански долар", Symbol: "KYD"}, "LVL": ut.Currency{Currency: "LVL", DisplayName: "Латвијски лати", Symbol: "LVL"}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "Марокански дирхам", Symbol: "MAD"}, "SDD": ut.Currency{Currency: "SDD", DisplayName: "Стари судански динар", Symbol: ""}, "BTN": ut.Currency{Currency: "BTN", DisplayName: "Бутански нгултрум", Symbol: "BTN"}, "DEM": ut.Currency{Currency: "DEM", DisplayName: "Немачка марка", Symbol: ""}, "ZWR": ut.Currency{Currency: "ZWR", DisplayName: "Зимбабвеански долар (2008)", Symbol: ""}, "LVR": ut.Currency{Currency: "LVR", DisplayName: "Латвијска рубља", Symbol: ""}, "XEU": ut.Currency{Currency: "XEU", DisplayName: "Европска валутна јединица", Symbol: ""}, "YER": ut.Currency{Currency: "YER", DisplayName: "Јеменски риjал", Symbol: "YER"}, "COU": ut.Currency{Currency: "COU", DisplayName: "Унидад де валоршки реал", Symbol: ""}, "HRD": ut.Currency{Currency: "HRD", DisplayName: "Хрватски динар", Symbol: ""}, "MGF": ut.Currency{Currency: "MGF", DisplayName: "Малагасијски франак", Symbol: ""}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "Сијералеонски леоне", Symbol: "SLL"}, "CHE": ut.Currency{Currency: "CHE", DisplayName: "WIR евро", Symbol: ""}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "Гамбијски даласи", Symbol: "GMD"}, "UYI": ut.Currency{Currency: "UYI", DisplayName: "Уругвајски пезо ен унидадес индексадас", Symbol: ""}, "KRW": ut.Currency{Currency: "KRW", DisplayName: "Јужнокорејски Вон", Symbol: "KRW"}, "SDP": ut.Currency{Currency: "SDP", DisplayName: "Стара суданска фунта", Symbol: ""}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "Сејшелска рупија", Symbol: "SCR"}, "GNS": ut.Currency{Currency: "GNS", DisplayName: "Гвинејски сили", Symbol: ""}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "Коморски франак", Symbol: "KMF"}, "GQE": ut.Currency{Currency: "GQE", DisplayName: "Екваторијално-гвинејски еквеле", Symbol: ""}, "MVR": ut.Currency{Currency: "MVR", DisplayName: "Малдивска руфија", Symbol: "MVR"}, "TMM": ut.Currency{Currency: "TMM", DisplayName: "Туркменистански манат (1993–2009)", Symbol: ""}, "UZS": ut.Currency{Currency: "UZS", DisplayName: "Узбекистански сом", Symbol: "UZS"}, "CHW": ut.Currency{Currency: "CHW", DisplayName: "WIR франак", Symbol: ""}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "Намибијски долар", Symbol: "NAD"}, "XDR": ut.Currency{Currency: "XDR", DisplayName: "Посебна цртаћа права", Symbol: ""}, "DOP": ut.Currency{Currency: "DOP", DisplayName: "Доминикански пезос", Symbol: "DOP"}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "Свазилендски лилангени", Symbol: "SZL"}, "VUV": ut.Currency{Currency: "VUV", DisplayName: "Вануатски вату", Symbol: "VUV"}, "XPF": ut.Currency{Currency: "XPF", DisplayName: "CFP франак", Symbol: "CFPF"}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "Јужноафрички ранд", Symbol: "ZAR"}, "BEL": ut.Currency{Currency: "BEL", DisplayName: "Белгијски франак (финансијски)", Symbol: ""}, "KGS": ut.Currency{Currency: "KGS", DisplayName: "Киргистански сом", Symbol: "KGS"}, "PTE": ut.Currency{Currency: "PTE", DisplayName: "Португалски ескудо", Symbol: ""}, "SIT": ut.Currency{Currency: "SIT", DisplayName: "Словеначки толар", Symbol: ""}, "VNN": ut.Currency{Currency: "VNN", DisplayName: "Вијетнамски донг (1978–1985)", Symbol: ""}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "ЦФА франак БЦЕАО", Symbol: "CFA"}, "BOL": ut.Currency{Currency: "BOL", DisplayName: "Стари боливијски боливијано", Symbol: ""}, "BYB": ut.Currency{Currency: "BYB", DisplayName: "Белоруска нова рубља (1994–1999)", Symbol: ""}, "BRB": ut.Currency{Currency: "BRB", DisplayName: "Бразилски нови крузеиро (1967–1986)", Symbol: ""}, "CUP": ut.Currency{Currency: "CUP", DisplayName: "Кубански пезос", Symbol: "CUP"}, "IQD": ut.Currency{Currency: "IQD", DisplayName: "Ирачки динар", Symbol: "IQD"}, "AOR": ut.Currency{Currency: "AOR", DisplayName: "Анголијска кванза реађустадо (1995–1999)", Symbol: ""}, "ARP": ut.Currency{Currency: "ARP", DisplayName: "Аргентински пезо (1983–1985)", Symbol: ""}, "RON": ut.Currency{Currency: "RON", DisplayName: "Румунски леј (1952–2006)", Symbol: "RON"}, "AON": ut.Currency{Currency: "AON", DisplayName: "Анголијска нова кванза (1990–2000)", Symbol: ""}, "PEN": ut.Currency{Currency: "PEN", DisplayName: "Перуански нови сол", Symbol: "PEN"}, "YUN": ut.Currency{Currency: "YUN", DisplayName: "Југословенски конвертибилни динар", Symbol: ""}, "GHS": ut.Currency{Currency: "GHS", DisplayName: "Гански седи", Symbol: "GHS"}, "INR": ut.Currency{Currency: "INR", DisplayName: "Индијска рупија", Symbol: "₹"}, "KHR": ut.Currency{Currency: "KHR", DisplayName: "Kамбоџански ријел", Symbol: "KHR"}, "NIC": ut.Currency{Currency: "NIC", DisplayName: "Никарагванска кордоба", Symbol: ""}, "ZRN": ut.Currency{Currency: "ZRN", DisplayName: "Заирски нови заир", Symbol: ""}, "BYR": ut.Currency{Currency: "BYR", DisplayName: "Белоруска рубља", Symbol: "[BYR]"}, "DDM": ut.Currency{Currency: "DDM", DisplayName: "Источно-немачка марка", Symbol: ""}, "MYR": ut.Currency{Currency: "MYR", DisplayName: "Малезијски рингит", Symbol: "MYR"}, "XAU": ut.Currency{Currency: "XAU", DisplayName: "Злато", Symbol: ""}, "BGM": ut.Currency{Currency: "BGM", DisplayName: "Бугарски социјалистички лев", Symbol: ""}, "IRR": ut.Currency{Currency: "IRR", DisplayName: "Ирански риjал", Symbol: "IRR"}, "XBA": ut.Currency{Currency: "XBA", DisplayName: "Европска композитна јединица", Symbol: ""}, "GWE": ut.Currency{Currency: "GWE", DisplayName: "Португалска гвинеја ескудо", Symbol: ""}, "TJR": ut.Currency{Currency: "TJR", DisplayName: "Таџихистанска рубља", Symbol: ""}, "XRE": ut.Currency{Currency: "XRE", DisplayName: "RINET фонд", Symbol: ""}, "ARL": ut.Currency{Currency: "ARL", DisplayName: "Аргентински пезос леј", Symbol: ""}, "ISK": ut.Currency{Currency: "ISK", DisplayName: "Исландска круна", Symbol: "ISK"}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "Шведска круна", Symbol: "SEK"}, "UAK": ut.Currency{Currency: "UAK", DisplayName: "Украјински карбованети", Symbol: ""}, "XBC": ut.Currency{Currency: "XBC", DisplayName: "Европска јединица рачуна (XBC)", Symbol: ""}, "CLE": ut.Currency{Currency: "CLE", DisplayName: "Чилеански ескудо", Symbol: ""}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "Суданска фунта", Symbol: "SDG"}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "Мадагаскарски ариари", Symbol: "MGA"}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "Монголски тугрик", Symbol: "MNT"}, "MZM": ut.Currency{Currency: "MZM", DisplayName: "Стари мозамбијски метикал", Symbol: ""}, "RHD": ut.Currency{Currency: "RHD", DisplayName: "Родејскидолар", Symbol: ""}, "SBD": ut.Currency{Currency: "SBD", DisplayName: "Соломонски долар", Symbol: "SBD"}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "Сингапурски долар", Symbol: "SGD"}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "Бразилски реал", Symbol: "R$"}, "ILS": ut.Currency{Currency: "ILS", DisplayName: "Израелски нови шекел", Symbol: "₪"}, "SRG": ut.Currency{Currency: "SRG", DisplayName: "Суринамски гилдер", Symbol: ""}, "YUM": ut.Currency{Currency: "YUM", DisplayName: "Југословенски нови динар", Symbol: ""}, "ISJ": ut.Currency{Currency: "ISJ", DisplayName: "Стара исландска круна", Symbol: ""}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "Канадски долар", Symbol: "CA$"}, "FKP": ut.Currency{Currency: "FKP", DisplayName: "Фокландска фунта", Symbol: "FKP"}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "Шриланканскa рупиja", Symbol: "LKR"}, "PYG": ut.Currency{Currency: "PYG", DisplayName: "Парагвајски гварани", Symbol: "PYG"}, "TND": ut.Currency{Currency: "TND", DisplayName: "Туниски динар", Symbol: "TND"}, "VEF": ut.Currency{Currency: "VEF", DisplayName: "Венецуелански боливар", Symbol: "VEF"}, "AZN": ut.Currency{Currency: "AZN", DisplayName: "Азербејџански манат", Symbol: "AZN"}, "BOB": ut.Currency{Currency: "BOB", DisplayName: "Боливијски боливијано", Symbol: "BOB"}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "Саудијски ријал", Symbol: "SAR"}, "GHC": ut.Currency{Currency: "GHC", DisplayName: "Гански цеди (1979–2007)", Symbol: ""}, "ITL": ut.Currency{Currency: "ITL", DisplayName: "Италијанска лира", Symbol: ""}, "MTL": ut.Currency{Currency: "MTL", DisplayName: "Малтешка лира", Symbol: ""}, "MZN": ut.Currency{Currency: "MZN", DisplayName: "Мозамбички метикал", Symbol: "MZN"}, "SYP": ut.Currency{Currency: "SYP", DisplayName: "Сиријска фунта", Symbol: "SYP"}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "Замбијска квача", Symbol: "ZMW"}, "BAN": ut.Currency{Currency: "BAN", DisplayName: "Босанско-херцеговачки нови динар", Symbol: ""}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "Зеленортски ескудо", Symbol: "CVE"}, "KRH": ut.Currency{Currency: "KRH", DisplayName: "Јужнокорејски хван", Symbol: ""}, "KRO": ut.Currency{Currency: "KRO", DisplayName: "Стари јужнокорејски вон", Symbol: ""}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "Боцванска пула", Symbol: "BWP"}, "ESA": ut.Currency{Currency: "ESA", DisplayName: "Шпанска пезета (рачун)", Symbol: ""}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "Данска круна", Symbol: "DKK"}, "MTP": ut.Currency{Currency: "MTP", DisplayName: "Малтешка фунта", Symbol: ""}, "AED": ut.Currency{Currency: "AED", DisplayName: "УАЕ дирхам", Symbol: "AED"}, "CZK": ut.Currency{Currency: "CZK", DisplayName: "Чешка круна", Symbol: "CZK"}, "LTT": ut.Currency{Currency: "LTT", DisplayName: "Литвански талонас", Symbol: ""}, "NLG": ut.Currency{Currency: "NLG", DisplayName: "Холандски гулден", Symbol: ""}, "THB": ut.Currency{Currency: "THB", DisplayName: "Тајландски бат", Symbol: "THB"}, "USS": ut.Currency{Currency: "USS", DisplayName: "САД долар (исти дан)", Symbol: ""}, "XAG": ut.Currency{Currency: "XAG", DisplayName: "Сребро", Symbol: ""}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "Источнокарипски долар", Symbol: "EC$"}, "ARA": ut.Currency{Currency: "ARA", DisplayName: "Аргентински аустрал", Symbol: ""}, "AWG": ut.Currency{Currency: "AWG", DisplayName: "Арубански флорин", Symbol: "AWG"}, "ZRZ": ut.Currency{Currency: "ZRZ", DisplayName: "Заирски заир", Symbol: ""}, "HUF": ut.Currency{Currency: "HUF", DisplayName: "Мађарска форинта", Symbol: "HUF"}, "LBP": ut.Currency{Currency: "LBP", DisplayName: "Либанска фунта", Symbol: "LBP"}, "LUL": ut.Currency{Currency: "LUL", DisplayName: "Луксембуршки финансијски франак", Symbol: ""}, "TRL": ut.Currency{Currency: "TRL", DisplayName: "Турска лира (1922–2005)", Symbol: ""}, "TWD": ut.Currency{Currency: "TWD", DisplayName: "Нови тајвански долар", Symbol: "NT$"}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "Бангладешка така", Symbol: "BDT"}, "BZD": ut.Currency{Currency: "BZD", DisplayName: "Белиски долар", Symbol: "BZD"}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "Евро", Symbol: "€"}, "ZMK": ut.Currency{Currency: "ZMK", DisplayName: "Замбијска квача (1968–2012)", Symbol: ""}, "BBD": ut.Currency{Currency: "BBD", DisplayName: "Барбадошки долар", Symbol: "BBD"}, "COP": ut.Currency{Currency: "COP", DisplayName: "Колумбијски пезос", Symbol: "COP"}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "CFA франак BEAC", Symbol: "FCFA"}, "AFA": ut.Currency{Currency: "AFA", DisplayName: "Авганистански авгани (1927–2002)", Symbol: ""}, "TPE": ut.Currency{Currency: "TPE", DisplayName: "Тиморшки ескудо", Symbol: ""}, "HNL": ut.Currency{Currency: "HNL", DisplayName: "Хондурашка лемпира", Symbol: "HNL"}, "STD": ut.Currency{Currency: "STD", DisplayName: "Саотомска добра", Symbol: "STD"}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "Угандски шилинг", Symbol: "UGX"}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "Бахреински динар", Symbol: "BHD"}, "GEL": ut.Currency{Currency: "GEL", DisplayName: "Грузијски лари", Symbol: "GEL"}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "Египатска фунта", Symbol: "EGP"}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "Босанско-херцеговачка конвертибилна марка", Symbol: "КМ"}, "CLP": ut.Currency{Currency: "CLP", DisplayName: "Чилеански пезос", Symbol: "CLP"}, "HRK": ut.Currency{Currency: "HRK", DisplayName: "Хрватска куна", Symbol: "HRK"}, "ILR": ut.Currency{Currency: "ILR", DisplayName: "Стари израелски шекели", Symbol: ""}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "Казахстански тенге", Symbol: "KZT"}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "Нигеријска наира", Symbol: "NGN"}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "Непозната валута", Symbol: ""}, "BRC": ut.Currency{Currency: "BRC", DisplayName: "Бразилијски крузадо", Symbol: ""}, "CRC": ut.Currency{Currency: "CRC", DisplayName: "Костарикански колон", Symbol: "CRC"}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "Св. јеленска фунта", Symbol: "SHP"}, "SSP": ut.Currency{Currency: "SSP", DisplayName: "Јужносуданска фунта", Symbol: "SSP"}, "MKN": ut.Currency{Currency: "MKN", DisplayName: "Стари македонски денар", Symbol: ""}, "MXV": ut.Currency{Currency: "MXV", DisplayName: "Мексички унидад де инверсион (UDI)", Symbol: ""}, "GWP": ut.Currency{Currency: "GWP", DisplayName: "Гвинеја Бисао Пезо", Symbol: ""}, "PAB": ut.Currency{Currency: "PAB", DisplayName: "Панамска балбоа", Symbol: "PAB"}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "Руска рубља", Symbol: "RUB"}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "Тонганска панга", Symbol: "TOP"}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "Танзанијски шилинг", Symbol: "TZS"}, "USN": ut.Currency{Currency: "USN", DisplayName: "САД долар (следећи дан)", Symbol: ""}, "BAD": ut.Currency{Currency: "BAD", DisplayName: "Босанско-Херцеговачки динар", Symbol: ""}, "BOV": ut.Currency{Currency: "BOV", DisplayName: "Боливијски мвдол", Symbol: ""}, "NPR": ut.Currency{Currency: "NPR", DisplayName: "Непалскa рупиja", Symbol: "NPR"}, "BRN": ut.Currency{Currency: "BRN", DisplayName: "Бразилијски нови крузадо", Symbol: ""}, "MDC": ut.Currency{Currency: "MDC", DisplayName: "Молдовански купон", Symbol: ""}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "Јапански јен", Symbol: "¥"}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "Мјанмарски кјат", Symbol: "MMK"}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "Руандски франак", Symbol: "RWF"}, "ZWL": ut.Currency{Currency: "ZWL", DisplayName: "Зимбабвеански долар (2009)", Symbol: ""}, "ATS": ut.Currency{Currency: "ATS", DisplayName: "Аустријски шилинг", Symbol: ""}, "BEC": ut.Currency{Currency: "BEC", DisplayName: "Белгијски франак (конвертибилни)", Symbol: ""}, "BRE": ut.Currency{Currency: "BRE", DisplayName: "Бразилски крузеиро (1990–1993)", Symbol: ""}, "BSD": ut.Currency{Currency: "BSD", DisplayName: "Бахамски долар", Symbol: "BSD"}, "BGO": ut.Currency{Currency: "BGO", DisplayName: "Стари бугарски лев", Symbol: ""}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "Бурундски франак", Symbol: "BIF"}, "CYP": ut.Currency{Currency: "CYP", DisplayName: "Кипарска фунта", Symbol: ""}, "MOP": ut.Currency{Currency: "MOP", DisplayName: "Макаоска патака", Symbol: "MOP"}, "XBD": ut.Currency{Currency: "XBD", DisplayName: "Европска јединица рачуна (XBD)", Symbol: ""}, "ANG": ut.Currency{Currency: "ANG", DisplayName: "Холандскоантилски гулден", Symbol: "ANG"}, "AZM": ut.Currency{Currency: "AZM", DisplayName: "Азербејџански манат (1993–2006)", Symbol: ""}, "HTG": ut.Currency{Currency: "HTG", DisplayName: "Хаићански гурд", Symbol: "HTG"}, "LSL": ut.Currency{Currency: "LSL", DisplayName: "Лесото лоти", Symbol: ""}, "PHP": ut.Currency{Currency: "PHP", DisplayName: "Филипински пезос", Symbol: "PHP"}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "Етиопски бир", Symbol: "ETB"}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "Гвинејски франак", Symbol: "GNF"}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "Либијски динар", Symbol: "LYD"}, "MZE": ut.Currency{Currency: "MZE", DisplayName: "Мозамбијски ескудо", Symbol: ""}, "ROL": ut.Currency{Currency: "ROL", DisplayName: "Румунски леј", Symbol: ""}, "ALK": ut.Currency{Currency: "ALK", DisplayName: "стари албански лек", Symbol: ""}, "JMD": ut.Currency{Currency: "JMD", DisplayName: "Јамајчански долар", Symbol: "JMD"}, "PEI": ut.Currency{Currency: "PEI", DisplayName: "Перуански инти", Symbol: ""}, "XTS": ut.Currency{Currency: "XTS", DisplayName: "Код тестиране валуте", Symbol: ""}, "CLF": ut.Currency{Currency: "CLF", DisplayName: "Чилеовски унидадес се фоменто", Symbol: ""}, "GTQ": ut.Currency{Currency: "GTQ", DisplayName: "Гватемалски кецал", Symbol: "GTQ"}, "FJD": ut.Currency{Currency: "FJD", DisplayName: "Фиџијски долар", Symbol: "FJD"}, "ZAL": ut.Currency{Currency: "ZAL", DisplayName: "Јужно-афрички ранд (финансијски)", Symbol: ""}, "CSD": ut.Currency{Currency: "CSD", DisplayName: "Стари српски динар", Symbol: ""}, "ECS": ut.Currency{Currency: "ECS", DisplayName: "Еквадорски сакр", Symbol: ""}, "BND": ut.Currency{Currency: "BND", DisplayName: "Брунејски долар", Symbol: "BND"}, "GEK": ut.Currency{Currency: "GEK", DisplayName: "Грузијски купон ларит", Symbol: ""}, "IEP": ut.Currency{Currency: "IEP", DisplayName: "Ирска фунта", Symbol: ""}, "MDL": ut.Currency{Currency: "MDL", DisplayName: "Молдавски леј", Symbol: "MDL"}, "MLF": ut.Currency{Currency: "MLF", DisplayName: "Малијански франак", Symbol: ""}, "TTD": ut.Currency{Currency: "TTD", DisplayName: "Тринидад-тобагошки долар", Symbol: "TTD"}, "ARM": ut.Currency{Currency: "ARM", DisplayName: "Аргентински пезос монедо национал", Symbol: ""}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "Аустралијски долар", Symbol: "AUD"}}
