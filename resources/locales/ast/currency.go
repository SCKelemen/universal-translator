package ast

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"BGM": ut.Currency{Currency: "BGM", DisplayName: "Lev socialista búlgaru", Symbol: "BGM"}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "Dólar canadiense", Symbol: "CA$"}, "DDM": ut.Currency{Currency: "DDM", DisplayName: "Marcu d’Alemaña Oriental", Symbol: "DDM"}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "Zloty polacu", Symbol: "PLN"}, "XPF": ut.Currency{Currency: "XPF", DisplayName: "francu CFP", Symbol: "CFPF"}, "RHD": ut.Currency{Currency: "RHD", DisplayName: "Dólar rodesianu", Symbol: "RHD"}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "francu CFA BCEAO", Symbol: "CFA"}, "XRE": ut.Currency{Currency: "XRE", DisplayName: "Fondos RINET", Symbol: "XRE"}, "PEI": ut.Currency{Currency: "PEI", DisplayName: "Inti peruanu", Symbol: "PEI"}, "ZMK": ut.Currency{Currency: "ZMK", DisplayName: "Kwacha zambianu (1968–2012)", Symbol: "ZMK"}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "Kwacha zambianu", Symbol: "ZMW"}, "NZD": ut.Currency{Currency: "NZD", DisplayName: "dólar neozelandés", Symbol: "NZ$"}, "SKK": ut.Currency{Currency: "SKK", DisplayName: "Corona eslovaca", Symbol: "SKK"}, "BOP": ut.Currency{Currency: "BOP", DisplayName: "Pesu bolivianu", Symbol: "BOP"}, "BRE": ut.Currency{Currency: "BRE", DisplayName: "Cruzeiro brasilanu (1990–1993)", Symbol: "BRE"}, "CNX": ut.Currency{Currency: "CNX", DisplayName: "Dólar del Bancu Popular Chinu", Symbol: ""}, "IDR": ut.Currency{Currency: "IDR", DisplayName: "rupiah indonesia", Symbol: "IDR"}, "KPW": ut.Currency{Currency: "KPW", DisplayName: "Won norcoreanu", Symbol: "KPW"}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "Tugrik mongol", Symbol: "MNT"}, "NIC": ut.Currency{Currency: "NIC", DisplayName: "Córdoba nicaraguanu (1988–1991)", Symbol: "NIC"}, "PTE": ut.Currency{Currency: "PTE", DisplayName: "Escudu portugués", Symbol: "PTE"}, "AZN": ut.Currency{Currency: "AZN", DisplayName: "Manat azerbaixanu", Symbol: "AZN"}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "Taka bangladexí", Symbol: "BDT"}, "CUC": ut.Currency{Currency: "CUC", DisplayName: "Pesu cubanu convertible", Symbol: "CUC"}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "dinar libiu", Symbol: "LYD"}, "MKD": ut.Currency{Currency: "MKD", DisplayName: "Denar macedoniu", Symbol: "MKD"}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "Kwacha malauianu", Symbol: "MWK"}, "BGN": ut.Currency{Currency: "BGN", DisplayName: "Lev búlgaru", Symbol: "BGN"}, "GHC": ut.Currency{Currency: "GHC", DisplayName: "Cedi ghanianu (1979–2007)", Symbol: "GHC"}, "GYD": ut.Currency{Currency: "GYD", DisplayName: "dólar guyanés", Symbol: "GYD"}, "RON": ut.Currency{Currency: "RON", DisplayName: "Leu rumanu", Symbol: "RON"}, "XFO": ut.Currency{Currency: "XFO", DisplayName: "Francu oru francés", Symbol: "XFO"}, "AON": ut.Currency{Currency: "AON", DisplayName: "Kwanza nuevu angolanu (1990–2000)", Symbol: "AON"}, "ECS": ut.Currency{Currency: "ECS", DisplayName: "Sucre ecuatorianu", Symbol: "ECS"}, "ECV": ut.Currency{Currency: "ECV", DisplayName: "Unidá ecuatoriana de valor constante", Symbol: "ECV"}, "ESB": ut.Currency{Currency: "ESB", DisplayName: "Peseta española (cuenta convertible)", Symbol: "ESB"}, "PLZ": ut.Currency{Currency: "PLZ", DisplayName: "Zloty polacu (1950–1995)", Symbol: "PLZ"}, "BUK": ut.Currency{Currency: "BUK", DisplayName: "Kyat birmanu", Symbol: "BUK"}, "ISK": ut.Currency{Currency: "ISK", DisplayName: "corona islandesa", Symbol: "ISK"}, "ARA": ut.Currency{Currency: "ARA", DisplayName: "Austral arxentín", Symbol: "ARA"}, "ESP": ut.Currency{Currency: "ESP", DisplayName: "Peseta española", Symbol: "ESP"}, "MVR": ut.Currency{Currency: "MVR", DisplayName: "Rufiyaa maldiviana", Symbol: "MVR"}, "QAR": ut.Currency{Currency: "QAR", DisplayName: "Rial qatarín", Symbol: "QAR"}, "XTS": ut.Currency{Currency: "XTS", DisplayName: "Códigu monetariu de prueba", Symbol: "XTS"}, "AED": ut.Currency{Currency: "AED", DisplayName: "Dirḥam de los Emiratos Árabes Xuníos", Symbol: "AED"}, "BOV": ut.Currency{Currency: "BOV", DisplayName: "Bolivianos mvdol", Symbol: "BOV"}, "ROL": ut.Currency{Currency: "ROL", DisplayName: "Leu rumanu (1952–2006)", Symbol: "ROL"}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "Francu ruandés", Symbol: "RWF"}, "ADP": ut.Currency{Currency: "ADP", DisplayName: "Peseta andorrana", Symbol: "ADP"}, "AOR": ut.Currency{Currency: "AOR", DisplayName: "Kwanza angolanu reaxustáu (1995–1999)", Symbol: "AOR"}, "CLF": ut.Currency{Currency: "CLF", DisplayName: "Unidá de cuenta chilena (UF)", Symbol: "CLF"}, "GEK": ut.Currency{Currency: "GEK", DisplayName: "Kupon larit xeorxanu", Symbol: "GEK"}, "IRR": ut.Currency{Currency: "IRR", DisplayName: "Rial iranín", Symbol: "IRR"}, "KYD": ut.Currency{Currency: "KYD", DisplayName: "dólar caimanés", Symbol: "KYD"}, "ZWL": ut.Currency{Currency: "ZWL", DisplayName: "Dólar zimbabuanu (2009)", Symbol: "ZWL"}, "HNL": ut.Currency{Currency: "HNL", DisplayName: "Lempira hondurana", Symbol: "HNL"}, "SRG": ut.Currency{Currency: "SRG", DisplayName: "Florín surinamés", Symbol: "SRG"}, "SUR": ut.Currency{Currency: "SUR", DisplayName: "Rublu soviéticu", Symbol: "SUR"}, "UAK": ut.Currency{Currency: "UAK", DisplayName: "Karbovanets ucraína", Symbol: "UAK"}, "UYU": ut.Currency{Currency: "UYU", DisplayName: "pesu uruguayu", Symbol: "UYU"}, "ZWD": ut.Currency{Currency: "ZWD", DisplayName: "Dólar zimbabuanu (1980–2008)", Symbol: "ZWD"}, "AFA": ut.Currency{Currency: "AFA", DisplayName: "Afganí afganistanu (1927–2002)", Symbol: "AFA"}, "COP": ut.Currency{Currency: "COP", DisplayName: "pesu colombianu", Symbol: "COP"}, "BYR": ut.Currency{Currency: "BYR", DisplayName: "Rublu bielorrusu", Symbol: "BYR"}, "TND": ut.Currency{Currency: "TND", DisplayName: "dinar tunecín", Symbol: "TND"}, "USS": ut.Currency{Currency: "USS", DisplayName: "Dólar d’EE.XX. (mesmu día)", Symbol: "USS"}, "XUA": ut.Currency{Currency: "XUA", DisplayName: "unidá de cuenta ADB", Symbol: "XUA"}, "ZRN": ut.Currency{Currency: "ZRN", DisplayName: "Zaire nuevu zairiegu (1993–1998)", Symbol: "ZRN"}, "ARL": ut.Currency{Currency: "ARL", DisplayName: "Pesu Ley arxentín (1970–1983)", Symbol: "ARL"}, "BOL": ut.Currency{Currency: "BOL", DisplayName: "Boliviano de Bolivia (1863–1963)", Symbol: "BOL"}, "KRH": ut.Currency{Currency: "KRH", DisplayName: "Hwan surcoreanu (1953–1962)", Symbol: "KRH"}, "PKR": ut.Currency{Currency: "PKR", DisplayName: "Rupia paquistanina", Symbol: "PKR"}, "USN": ut.Currency{Currency: "USN", DisplayName: "Dólar d’EE.XX. (día siguiente)", Symbol: "USN"}, "BRB": ut.Currency{Currency: "BRB", DisplayName: "Cruzeiro nuevu brasilanu (1967–1986)", Symbol: "BRB"}, "TJR": ut.Currency{Currency: "TJR", DisplayName: "Rublu taxiquistanín", Symbol: "TJR"}, "XSU": ut.Currency{Currency: "XSU", DisplayName: "Sucre", Symbol: "XSU"}, "BRR": ut.Currency{Currency: "BRR", DisplayName: "Cruzeiro brasilanu (1993–1994)", Symbol: "BRR"}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "francu suizu", Symbol: "CHF"}, "ISJ": ut.Currency{Currency: "ISJ", DisplayName: "Corona islandesa (1918–1981)", Symbol: ""}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "Tenge kazaquistanín", Symbol: "KZT"}, "BGO": ut.Currency{Currency: "BGO", DisplayName: "Lev búlgaru (1879–1952)", Symbol: "BGO"}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "Yuan chinu", Symbol: "CN¥"}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "corona danesa", Symbol: "DKK"}, "GNS": ut.Currency{Currency: "GNS", DisplayName: "syli guineanu", Symbol: "GNS"}, "KHR": ut.Currency{Currency: "KHR", DisplayName: "riel camboyanu", Symbol: "KHR"}, "JOD": ut.Currency{Currency: "JOD", DisplayName: "Dinar xordanu", Symbol: "JOD"}, "RUR": ut.Currency{Currency: "RUR", DisplayName: "Rublu rusu (1991–1998)", Symbol: "RUR"}, "ARP": ut.Currency{Currency: "ARP", DisplayName: "Pesu arxentín (1983–1985)", Symbol: "ARP"}, "ARS": ut.Currency{Currency: "ARS", DisplayName: "pesu arxentín", Symbol: "ARS"}, "ATS": ut.Currency{Currency: "ATS", DisplayName: "Chelín austriacu", Symbol: "ATS"}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "Dinar baḥreiní", Symbol: "BHD"}, "BZD": ut.Currency{Currency: "BZD", DisplayName: "Dólar belizianu", Symbol: "BZD"}, "HUF": ut.Currency{Currency: "HUF", DisplayName: "Forint húngaru", Symbol: "HUF"}, "OMR": ut.Currency{Currency: "OMR", DisplayName: "Rial omanianu", Symbol: "OMR"}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "Shilling tanzanianu", Symbol: "TZS"}, "XAG": ut.Currency{Currency: "XAG", DisplayName: "Plata", Symbol: "XAG"}, "YDD": ut.Currency{Currency: "YDD", DisplayName: "Dinar yemenín", Symbol: "YDD"}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "Dólar australianu", Symbol: "A$"}, "RSD": ut.Currency{Currency: "RSD", DisplayName: "dinar serbiu", Symbol: "RSD"}, "VND": ut.Currency{Currency: "VND", DisplayName: "dong vietnamín", Symbol: "₫"}, "XPD": ut.Currency{Currency: "XPD", DisplayName: "Paladiu", Symbol: "XPD"}, "SDP": ut.Currency{Currency: "SDP", DisplayName: "llibra sudanesa (1957–1998)", Symbol: "SDP"}, "YUR": ut.Currency{Currency: "YUR", DisplayName: "Dinar reformáu yugoslavu (1992–1993)", Symbol: "YUR"}, "ANG": ut.Currency{Currency: "ANG", DisplayName: "Guílder de les Antilles Neerlandeses", Symbol: "ANG"}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "Francu comoranu", Symbol: "KMF"}, "LUF": ut.Currency{Currency: "LUF", DisplayName: "Francu luxemburgués", Symbol: "LUF"}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "Ariary malgaxe", Symbol: "MGA"}, "MZN": ut.Currency{Currency: "MZN", DisplayName: "Metical mozambicanu", Symbol: "MZN"}, "PES": ut.Currency{Currency: "PES", DisplayName: "Sol peruanu (1863–1965)", Symbol: "PES"}, "AOK": ut.Currency{Currency: "AOK", DisplayName: "Kwanza angolanu (1977–1991)", Symbol: "AOK"}, "AWG": ut.Currency{Currency: "AWG", DisplayName: "Florín arubanu", Symbol: "AWG"}, "BEC": ut.Currency{Currency: "BEC", DisplayName: "Francu belga (convertible)", Symbol: "BEC"}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "dalasi gambianu", Symbol: "GMD"}, "NIO": ut.Currency{Currency: "NIO", DisplayName: "Córdoba nicaraguanu", Symbol: "NIO"}, "TTD": ut.Currency{Currency: "TTD", DisplayName: "dólar de Trinidá y Tobagu", Symbol: "TTD"}, "HTG": ut.Currency{Currency: "HTG", DisplayName: "Gourde haitianu", Symbol: "HTG"}, "SRD": ut.Currency{Currency: "SRD", DisplayName: "dólar surinamés", Symbol: "SRD"}, "VEB": ut.Currency{Currency: "VEB", DisplayName: "Bolívar venezolanu (1871–2008)", Symbol: "VEB"}, "XBA": ut.Currency{Currency: "XBA", DisplayName: "Unidá Compuesta Europea", Symbol: "XBA"}, "GIP": ut.Currency{Currency: "GIP", DisplayName: "llibra de Xibraltar", Symbol: "GIP"}, "HKD": ut.Currency{Currency: "HKD", DisplayName: "Dólar hongkonés", Symbol: "HK$"}, "ITL": ut.Currency{Currency: "ITL", DisplayName: "Llira italiana", Symbol: "ITL"}, "SVC": ut.Currency{Currency: "SVC", DisplayName: "Colón salvadorianu", Symbol: "SVC"}, "TRY": ut.Currency{Currency: "TRY", DisplayName: "Llira turca", Symbol: "TRY"}, "UYI": ut.Currency{Currency: "UYI", DisplayName: "Pesu uruguayu (Unidaes indexaes)", Symbol: "UYI"}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "Kwanza angolanu", Symbol: "AOA"}, "BND": ut.Currency{Currency: "BND", DisplayName: "dólar bruneyanu", Symbol: "BND"}, "MCF": ut.Currency{Currency: "MCF", DisplayName: "Francu monegascu", Symbol: "MCF"}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "naira nixeriana", Symbol: "NGN"}, "UGS": ut.Currency{Currency: "UGS", DisplayName: "Shilling ugandés (1966–1987)", Symbol: "UGS"}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "leone sierralleonés", Symbol: "SLL"}, "TMM": ut.Currency{Currency: "TMM", DisplayName: "Manat turcomanu (1993–2009)", Symbol: "TMM"}, "BGL": ut.Currency{Currency: "BGL", DisplayName: "Lev fuerte búlgaru", Symbol: "BGL"}, "BOB": ut.Currency{Currency: "BOB", DisplayName: "Bolivianu de Bolivia", Symbol: "BOB"}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "Nakfa eritréu", Symbol: "ERN"}, "KWD": ut.Currency{Currency: "KWD", DisplayName: "Dinar kuwaitianu", Symbol: "KWD"}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "kip laosianu", Symbol: "LAK"}, "MGF": ut.Currency{Currency: "MGF", DisplayName: "Francu malgaxe", Symbol: "MGF"}, "UAH": ut.Currency{Currency: "UAH", DisplayName: "Grivna ucraína", Symbol: "UAH"}, "AMD": ut.Currency{Currency: "AMD", DisplayName: "Dram armeniu", Symbol: "AMD"}, "LUL": ut.Currency{Currency: "LUL", DisplayName: "Francu financieru luxemburgués", Symbol: "LUL"}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "corona noruega", Symbol: "NOK"}, "VNN": ut.Currency{Currency: "VNN", DisplayName: "Dong vietnamín (1978–1985)", Symbol: "VNN"}, "XBB": ut.Currency{Currency: "XBB", DisplayName: "Unidá monetaria europea", Symbol: "XBB"}, "BBD": ut.Currency{Currency: "BBD", DisplayName: "Dólar barbadianu", Symbol: "BBD"}, "GQE": ut.Currency{Currency: "GQE", DisplayName: "Ekwele de Guinea Ecuatorial", Symbol: "GQE"}, "TJS": ut.Currency{Currency: "TJS", DisplayName: "Somoni taxiquistanín", Symbol: "TJS"}, "XDR": ut.Currency{Currency: "XDR", DisplayName: "Drechos especiales de xiru", Symbol: "XDR"}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "Rupia de Sri Lanka", Symbol: "LKR"}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "Lilangeni suazilandés", Symbol: "SZL"}, "FIM": ut.Currency{Currency: "FIM", DisplayName: "Marcu finlandés", Symbol: "FIM"}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "francu guineanu", Symbol: "GNF"}, "LUC": ut.Currency{Currency: "LUC", DisplayName: "Francu convertible luxemburgués", Symbol: "LUC"}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "Rupia mauriciana", Symbol: "MUR"}, "CSK": ut.Currency{Currency: "CSK", DisplayName: "Corona fuerte checoslovaca", Symbol: "CSK"}, "ILP": ut.Currency{Currency: "ILP", DisplayName: "Llibra israelina", Symbol: "ILP"}, "LBP": ut.Currency{Currency: "LBP", DisplayName: "Llibra libanesa", Symbol: "LBP"}, "SBD": ut.Currency{Currency: "SBD", DisplayName: "dólar salomonés", Symbol: "SBD"}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "dólar singapuranu", Symbol: "SGD"}, "VEF": ut.Currency{Currency: "VEF", DisplayName: "bolívar venezolanu", Symbol: "VEF"}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "dólar del Caribe Oriental", Symbol: "EC$"}, "ARM": ut.Currency{Currency: "ARM", DisplayName: "Pesu arxentín (1881–1970)", Symbol: "ARM"}, "BEF": ut.Currency{Currency: "BEF", DisplayName: "Francu belga", Symbol: "BEF"}, "CZK": ut.Currency{Currency: "CZK", DisplayName: "Corona checa", Symbol: "CZK"}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "llibra esterlina", Symbol: "£"}, "KES": ut.Currency{Currency: "KES", DisplayName: "Shilling kenianu", Symbol: "KES"}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "dirḥam marroquín", Symbol: "MAD"}, "HRK": ut.Currency{Currency: "HRK", DisplayName: "Kuna croata", Symbol: "HRK"}, "FKP": ut.Currency{Currency: "FKP", DisplayName: "llibra malviniana", Symbol: "FKP"}, "MKN": ut.Currency{Currency: "MKN", DisplayName: "Denar macedoniu (1992–1993)", Symbol: "MKN"}, "MXN": ut.Currency{Currency: "MXN", DisplayName: "Pesu mexicanu", Symbol: "MX$"}, "MZM": ut.Currency{Currency: "MZM", DisplayName: "Metical mozambicanu (1980–2006)", Symbol: "MZM"}, "CLP": ut.Currency{Currency: "CLP", DisplayName: "pesu chilenu", Symbol: "CLP"}, "CSD": ut.Currency{Currency: "CSD", DisplayName: "Dinar serbiu (2002–2006)", Symbol: "CSD"}, "GEL": ut.Currency{Currency: "GEL", DisplayName: "Lari xeorxanu", Symbol: "GEL"}, "GWE": ut.Currency{Currency: "GWE", DisplayName: "Escudo de Guinea portuguesa", Symbol: "GWE"}, "MDC": ut.Currency{Currency: "MDC", DisplayName: "Cupón moldavu", Symbol: "MDC"}, "MTL": ut.Currency{Currency: "MTL", DisplayName: "Llira maltesa", Symbol: "MTL"}, "SYP": ut.Currency{Currency: "SYP", DisplayName: "Llibra siria", Symbol: "SYP"}, "TPE": ut.Currency{Currency: "TPE", DisplayName: "Escudu timorés", Symbol: "TPE"}, "ZAL": ut.Currency{Currency: "ZAL", DisplayName: "Rand sudafricanu (financieru)", Symbol: "ZAL"}, "SIT": ut.Currency{Currency: "SIT", DisplayName: "Tolar eslovenu", Symbol: "SIT"}, "THB": ut.Currency{Currency: "THB", DisplayName: "baht tailandés", Symbol: "฿"}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "Pula botsuaniana", Symbol: "BWP"}, "BYB": ut.Currency{Currency: "BYB", DisplayName: "Rublu nuevu bielorrusu (1994–1999)", Symbol: "BYB"}, "FRF": ut.Currency{Currency: "FRF", DisplayName: "Francu francés", Symbol: "FRF"}, "PYG": ut.Currency{Currency: "PYG", DisplayName: "guaraní paraguayu", Symbol: "PYG"}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "Rupia seixelesa", Symbol: "SCR"}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "llibra sudanesa", Symbol: "SDG"}, "CHW": ut.Currency{Currency: "CHW", DisplayName: "Francu WIR", Symbol: "CHW"}, "GHS": ut.Currency{Currency: "GHS", DisplayName: "cedi ghanianu", Symbol: "GHS"}, "IQD": ut.Currency{Currency: "IQD", DisplayName: "Dinar iraquín", Symbol: "IQD"}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "Rublu rusu", Symbol: "RUB"}, "MAF": ut.Currency{Currency: "MAF", DisplayName: "francu marroquín", Symbol: "MAF"}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "Dólar namibianu", Symbol: "NAD"}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "Francu CFA centroafricanu", Symbol: "FCFA"}, "CYP": ut.Currency{Currency: "CYP", DisplayName: "Llibra xipriota", Symbol: "CYP"}, "MTP": ut.Currency{Currency: "MTP", DisplayName: "Llibra maltesa", Symbol: "MTP"}, "TWD": ut.Currency{Currency: "TWD", DisplayName: "Dólar nuevu taiwanés", Symbol: "NT$"}, "XBC": ut.Currency{Currency: "XBC", DisplayName: "Unidá de cuenta europea (XBC)", Symbol: "XBC"}, "ILR": ut.Currency{Currency: "ILR", DisplayName: "Xequel israelín (1980–1985)", Symbol: ""}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "corona sueca", Symbol: "SEK"}, "YUD": ut.Currency{Currency: "YUD", DisplayName: "Dinar fuerte yugoslavu (1966–1990)", Symbol: "YUD"}, "MXP": ut.Currency{Currency: "MXP", DisplayName: "Pesu de plata mexicanu (1861–1992)", Symbol: "MXP"}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "Shilling ugandés", Symbol: "UGX"}, "AFN": ut.Currency{Currency: "AFN", DisplayName: "Afganí afganistanu", Symbol: "AFN"}, "BAD": ut.Currency{Currency: "BAD", DisplayName: "Dinar de Bosnia-Herzegovina (1992–1994)", Symbol: "BAD"}, "BRN": ut.Currency{Currency: "BRN", DisplayName: "Cruzado nuevu brasilanu (1989–1990)", Symbol: "BRN"}, "GWP": ut.Currency{Currency: "GWP", DisplayName: "Pesu de Guinea-Bisáu", Symbol: "GWP"}, "IEP": ut.Currency{Currency: "IEP", DisplayName: "Llibra irlandesa", Symbol: "IEP"}, "KGS": ut.Currency{Currency: "KGS", DisplayName: "Som kirguistanín", Symbol: "KGS"}, "VUV": ut.Currency{Currency: "VUV", DisplayName: "vatu vanuatuanu", Symbol: "VUV"}, "DOP": ut.Currency{Currency: "DOP", DisplayName: "Pesu dominicanu", Symbol: "DOP"}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "dinar arxelín", Symbol: "DZD"}, "SSP": ut.Currency{Currency: "SSP", DisplayName: "llibra sursudanesa", Symbol: "SSP"}, "ZWR": ut.Currency{Currency: "ZWR", DisplayName: "Dólar zimbabuanu (2008)", Symbol: "ZWR"}, "CUP": ut.Currency{Currency: "CUP", DisplayName: "Pesu cubanu", Symbol: "CUP"}, "STD": ut.Currency{Currency: "STD", DisplayName: "dobra de Santu Tomé y Príncipe", Symbol: "STD"}, "BTN": ut.Currency{Currency: "BTN", DisplayName: "Ngultrum butanés", Symbol: "BTN"}, "MXV": ut.Currency{Currency: "MXV", DisplayName: "Unidá d’inversión mexicana", Symbol: "MXV"}, "XFU": ut.Currency{Currency: "XFU", DisplayName: "Francu UIC francés", Symbol: "XFU"}, "FJD": ut.Currency{Currency: "FJD", DisplayName: "dólar fixanu", Symbol: "FJD"}, "LVL": ut.Currency{Currency: "LVL", DisplayName: "Lats letón", Symbol: "LVL"}, "MLF": ut.Currency{Currency: "MLF", DisplayName: "Francu malianu", Symbol: "MLF"}, "PGK": ut.Currency{Currency: "PGK", DisplayName: "kina papuana", Symbol: "PGK"}, "TMT": ut.Currency{Currency: "TMT", DisplayName: "Manat turcomanu", Symbol: "TMT"}, "XBD": ut.Currency{Currency: "XBD", DisplayName: "Unidá de cuenta europea (XBD)", Symbol: "XBD"}, "NLG": ut.Currency{Currency: "NLG", DisplayName: "Florín neerlandés", Symbol: "NLG"}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "Riyal saudita", Symbol: "SAR"}, "YUN": ut.Currency{Currency: "YUN", DisplayName: "Dinar convertible yugoslavu (1990–1992)", Symbol: "YUN"}, "AZM": ut.Currency{Currency: "AZM", DisplayName: "Manat azerbaixanu (1993–2006)", Symbol: "AZM"}, "XAU": ut.Currency{Currency: "XAU", DisplayName: "Oru", Symbol: "XAU"}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "dólar liberianu", Symbol: "LRD"}, "LTT": ut.Currency{Currency: "LTT", DisplayName: "Talonas lituanu", Symbol: "LTT"}, "MZE": ut.Currency{Currency: "MZE", DisplayName: "Escudu mozambicanu", Symbol: "MZE"}, "NPR": ut.Currency{Currency: "NPR", DisplayName: "Rupia nepalesa", Symbol: "NPR"}, "XPT": ut.Currency{Currency: "XPT", DisplayName: "Platín", Symbol: "XPT"}, "YUM": ut.Currency{Currency: "YUM", DisplayName: "Dinar nuevu yugoslavu (1994–2002)", Symbol: "YUM"}, "CRC": ut.Currency{Currency: "CRC", DisplayName: "Colón costarricanu", Symbol: "CRC"}, "KRW": ut.Currency{Currency: "KRW", DisplayName: "Won surcoreanu", Symbol: "₩"}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "kyat de Myanmar", Symbol: "MMK"}, "MOP": ut.Currency{Currency: "MOP", DisplayName: "Pataca de Macáu", Symbol: "MOP"}, "PAB": ut.Currency{Currency: "PAB", DisplayName: "Balboa panamiegu", Symbol: "PAB"}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "Shilling somalín", Symbol: "SOS"}, "KRO": ut.Currency{Currency: "KRO", DisplayName: "Won surcoreanu (1945–1953)", Symbol: "KRO"}, "LTL": ut.Currency{Currency: "LTL", DisplayName: "Litas lituanu", Symbol: "LTL"}, "LVR": ut.Currency{Currency: "LVR", DisplayName: "Rublu letón", Symbol: "LVR"}, "BAN": ut.Currency{Currency: "BAN", DisplayName: "Dinar nuevu de Bosnia-Herzegovina (1994–1997)", Symbol: "BAN"}, "PHP": ut.Currency{Currency: "PHP", DisplayName: "pesu filipín", Symbol: "PHP"}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "paʻanga tonganu", Symbol: "TOP"}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "Francu burundianu", Symbol: "BIF"}, "HRD": ut.Currency{Currency: "HRD", DisplayName: "Dinar croata", Symbol: "HRD"}, "JMD": ut.Currency{Currency: "JMD", DisplayName: "Dólar xamaicanu", Symbol: "JMD"}, "ZRZ": ut.Currency{Currency: "ZRZ", DisplayName: "Zaire zairiegu (1971–1993)", Symbol: "ZRZ"}, "ESA": ut.Currency{Currency: "ESA", DisplayName: "Peseta española (cuenta A)", Symbol: "ESA"}, "GTQ": ut.Currency{Currency: "GTQ", DisplayName: "Quetzal guatemalianu", Symbol: "GTQ"}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "Yen xaponés", Symbol: "¥"}, "LSL": ut.Currency{Currency: "LSL", DisplayName: "Loti de Lesothu", Symbol: "LSL"}, "YER": ut.Currency{Currency: "YER", DisplayName: "Rial yemenín", Symbol: "YER"}, "ALL": ut.Currency{Currency: "ALL", DisplayName: "Lek albanés", Symbol: "ALL"}, "MYR": ut.Currency{Currency: "MYR", DisplayName: "ringgit malasiu", Symbol: "MYR"}, "PEN": ut.Currency{Currency: "PEN", DisplayName: "sol nuevu peruanu", Symbol: "PEN"}, "UYP": ut.Currency{Currency: "UYP", DisplayName: "Pesu uruguayu (1975–1993)", Symbol: "UYP"}, "UZS": ut.Currency{Currency: "UZS", DisplayName: "Som uzbequistanín", Symbol: "UZS"}, "BRZ": ut.Currency{Currency: "BRZ", DisplayName: "Cruzeiru brasilanu (1942–1967)", Symbol: "BRZ"}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "Birr etíope", Symbol: "ETB"}, "GRD": ut.Currency{Currency: "GRD", DisplayName: "Dracma griegu", Symbol: "GRD"}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "llibra de Santa Lena", Symbol: "SHP"}, "ILS": ut.Currency{Currency: "ILS", DisplayName: "Xequel nuevu israelín", Symbol: "₪"}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "Rand sudafricanu", Symbol: "ZAR"}, "ALK": ut.Currency{Currency: "ALK", DisplayName: "Lek albanés (1946–1965)", Symbol: ""}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "real brasilanu", Symbol: "R$"}, "CHE": ut.Currency{Currency: "CHE", DisplayName: "Euru WIR", Symbol: "CHE"}, "CLE": ut.Currency{Currency: "CLE", DisplayName: "Escudu chilenu", Symbol: "CLE"}, "SDD": ut.Currency{Currency: "SDD", DisplayName: "dinar sudanés (1992–2007)", Symbol: "SDD"}, "BRC": ut.Currency{Currency: "BRC", DisplayName: "Cruzado brasilanu (1986–1989)", Symbol: "BRC"}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "escudu cabuverdianu", Symbol: "CVE"}, "USD": ut.Currency{Currency: "USD", DisplayName: "Dólar estaunidense", Symbol: "$"}, "BSD": ut.Currency{Currency: "BSD", DisplayName: "Dólar bahamés", Symbol: "BSD"}, "DEM": ut.Currency{Currency: "DEM", DisplayName: "Marcu alemán", Symbol: "DEM"}, "EEK": ut.Currency{Currency: "EEK", DisplayName: "Corona estonia", Symbol: "EEK"}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "llibra exipciana", Symbol: "EGP"}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "ouguiya mauritanu", Symbol: "MRO"}, "BMD": ut.Currency{Currency: "BMD", DisplayName: "Dólar bermudianu", Symbol: "BMD"}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "Francu xibutianu", Symbol: "DJF"}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "euro", Symbol: "€"}, "WST": ut.Currency{Currency: "WST", DisplayName: "tala samoanu", Symbol: "WST"}, "XEU": ut.Currency{Currency: "XEU", DisplayName: "Unidá de divisa europea", Symbol: "XEU"}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "marcu convertible de Bosnia-Herzegovina", Symbol: "BAM"}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "francu congolés", Symbol: "CDF"}, "COU": ut.Currency{Currency: "COU", DisplayName: "Unidá de valor real colombiana", Symbol: "COU"}, "MDL": ut.Currency{Currency: "MDL", DisplayName: "Leu moldavu", Symbol: "MDL"}, "BEL": ut.Currency{Currency: "BEL", DisplayName: "Francu belga (financieru)", Symbol: "BEL"}, "INR": ut.Currency{Currency: "INR", DisplayName: "Rupia india", Symbol: "₹"}, "MVP": ut.Currency{Currency: "MVP", DisplayName: "Rupia maldiviana (1947–1981)", Symbol: ""}, "TRL": ut.Currency{Currency: "TRL", DisplayName: "Llira turca (1922–2005)", Symbol: "TRL"}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "Divisa desconocida", Symbol: "XXX"}}
