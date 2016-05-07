package sw

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"MKD": ut.Currency{Currency: "MKD", DisplayName: "Denar ya Masedonia", Symbol: "MKD"}, "MOP": ut.Currency{Currency: "MOP", DisplayName: "Pataca ya Macau", Symbol: "MOP"}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "Leone", Symbol: "SLL"}, "TJS": ut.Currency{Currency: "TJS", DisplayName: "Somoni ya Tajikistani", Symbol: "TJS"}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "CFA faranga za BCEAO", Symbol: "CFA"}, "GYD": ut.Currency{Currency: "GYD", DisplayName: "Dola ya Guyana", Symbol: "GYD"}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "Yen ya Japani", Symbol: "JP¥"}, "KPW": ut.Currency{Currency: "KPW", DisplayName: "Won ya Korea Kaskazini", Symbol: "KPW"}, "PGK": ut.Currency{Currency: "PGK", DisplayName: "Kina ya Papua New Guinea", Symbol: "PGK"}, "FJD": ut.Currency{Currency: "FJD", DisplayName: "Dola ya Fiji", Symbol: "FJD"}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "Kip ya Laosi", Symbol: "LAK"}, "STD": ut.Currency{Currency: "STD", DisplayName: "Dobra ya Sao Tome na Principe", Symbol: "STD"}, "AMD": ut.Currency{Currency: "AMD", DisplayName: "Dram ya Armenia", Symbol: "AMD"}, "ARS": ut.Currency{Currency: "ARS", DisplayName: "Peso ya Ajentina", Symbol: "ARS"}, "BOB": ut.Currency{Currency: "BOB", DisplayName: "Boliviano ya Bolivia", Symbol: "BOB"}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "Dola ya Singapore", Symbol: "SGD"}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "Shilingi ya Somalia", Symbol: "SOS"}, "SRD": ut.Currency{Currency: "SRD", DisplayName: "Dola ya Suriname", Symbol: "SRD"}, "SSP": ut.Currency{Currency: "SSP", DisplayName: "Pauni ya Sudani Kusini", Symbol: "SSP"}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "Faranga ya Uswisi", Symbol: "CHF"}, "HRK": ut.Currency{Currency: "HRK", DisplayName: "Kuna ya Kroeshia", Symbol: "HRK"}, "INR": ut.Currency{Currency: "INR", DisplayName: "Rupia ya India", Symbol: "₹"}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "Krone ya Norwe", Symbol: "NOK"}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "Kwanza ya Angola", Symbol: "AOA"}, "OMR": ut.Currency{Currency: "OMR", DisplayName: "Rial ya Omani", Symbol: "OMR"}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "Krona ya Uswidi", Symbol: "SEK"}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "Dola ya Karibea ya Mashariki", Symbol: "EC$"}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "Kwacha ya Zambia", Symbol: "ZMW"}, "BTN": ut.Currency{Currency: "BTN", DisplayName: "Ngultrum ya Bhutan", Symbol: "BTN"}, "IQD": ut.Currency{Currency: "IQD", DisplayName: "Dinar ya Iraki", Symbol: "IQD"}, "KES": ut.Currency{Currency: "KES", DisplayName: "Shilingi ya Kenya", Symbol: "Ksh"}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "Riyal ya Saudia", Symbol: "SAR"}, "BGN": ut.Currency{Currency: "BGN", DisplayName: "Lev ya Bulgaria", Symbol: "BGN"}, "CZK": ut.Currency{Currency: "CZK", DisplayName: "Koruna ya Jamhuri ya Cheki", Symbol: "CZK"}, "HUF": ut.Currency{Currency: "HUF", DisplayName: "Forint ya Hungaria", Symbol: "HUF"}, "JMD": ut.Currency{Currency: "JMD", DisplayName: "Dola ya Jamaica", Symbol: "JMD"}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "Tugrik ya Mongolia", Symbol: "MNT"}, "PHP": ut.Currency{Currency: "PHP", DisplayName: "Peso ya Ufilipino", Symbol: "PHP"}, "RON": ut.Currency{Currency: "RON", DisplayName: "Leu ya Romania", Symbol: "RON"}, "ALL": ut.Currency{Currency: "ALL", DisplayName: "Lek ya Albania", Symbol: "ALL"}, "BBD": ut.Currency{Currency: "BBD", DisplayName: "Dola ya Barbados", Symbol: "BBD"}, "BMD": ut.Currency{Currency: "BMD", DisplayName: "Dola ya Bermuda", Symbol: "BMD"}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "Real ya Brazil", Symbol: "R$"}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "Nakfa ya Eritrea", Symbol: "ERN"}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "Ouguiya ya Moritania", Symbol: "MRO"}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "Rupia ya Morisi", Symbol: "MUR"}, "AED": ut.Currency{Currency: "AED", DisplayName: "Dirham ya Falme za Kiarabu", Symbol: "AED"}, "AZN": ut.Currency{Currency: "AZN", DisplayName: "Manat ya Azebaijani", Symbol: "AZN"}, "HKD": ut.Currency{Currency: "HKD", DisplayName: "Dola ya Hong Kong", Symbol: "HK$"}, "HTG": ut.Currency{Currency: "HTG", DisplayName: "Gourde ya Haiti", Symbol: "HTG"}, "NIO": ut.Currency{Currency: "NIO", DisplayName: "Cordoba ya Nikaragua", Symbol: "NIO"}, "TWD": ut.Currency{Currency: "TWD", DisplayName: "Dola ya Taiwan", Symbol: "NT$"}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "Shilingi ya Uganda", Symbol: "UGX"}, "AWG": ut.Currency{Currency: "AWG", DisplayName: "Florin ya Aruba", Symbol: "AWG"}, "GNS": ut.Currency{Currency: "GNS", DisplayName: "Faranga ya Gine", Symbol: ""}, "MZM": ut.Currency{Currency: "MZM", DisplayName: "Metikali ya Msumbiji (1980–2006)", Symbol: ""}, "NPR": ut.Currency{Currency: "NPR", DisplayName: "Rupia ya Nepali", Symbol: "NPR"}, "WST": ut.Currency{Currency: "WST", DisplayName: "Tala ya Samoa", Symbol: "WST"}, "ZWD": ut.Currency{Currency: "ZWD", DisplayName: "Dola ya Zimbabwe", Symbol: ""}, "CRC": ut.Currency{Currency: "CRC", DisplayName: "Colon ya Kostarika", Symbol: "CRC"}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "Birr ya Uhabeshi", Symbol: "ETB"}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "Kyat ya Myama", Symbol: "MMK"}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "Pula ya Botswana", Symbol: "BWP"}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "Pauni ya Uingereza", Symbol: "£"}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "Faranga ya Komoro", Symbol: "KMF"}, "LSL": ut.Currency{Currency: "LSL", DisplayName: "Loti ya Lesoto", Symbol: ""}, "PYG": ut.Currency{Currency: "PYG", DisplayName: "Guarani ya Paragwai", Symbol: "PYG"}, "TRY": ut.Currency{Currency: "TRY", DisplayName: "Lira ya Uturuki", Symbol: "TRY"}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "Paʻanga ya Tonga", Symbol: "TOP"}, "CUP": ut.Currency{Currency: "CUP", DisplayName: "Peso ya Cuba", Symbol: "CUP"}, "DOP": ut.Currency{Currency: "DOP", DisplayName: "Peso ya Dominika", Symbol: "DOP"}, "GHC": ut.Currency{Currency: "GHC", DisplayName: "Sedi ya Ghana", Symbol: ""}, "GIP": ut.Currency{Currency: "GIP", DisplayName: "Pauni ya Gibraltar", Symbol: "GIP"}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "Dirham ya Moroko", Symbol: "MAD"}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "Naira ya Nijeria", Symbol: "NGN"}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "Pauni ya Santahelena", Symbol: "SHP"}, "VUV": ut.Currency{Currency: "VUV", DisplayName: "Vatu ya Vanuatu", Symbol: "VUV"}, "ZMK": ut.Currency{Currency: "ZMK", DisplayName: "Kwacha ya Zambia (1968–2012)", Symbol: ""}, "IDR": ut.Currency{Currency: "IDR", DisplayName: "Rupiah ya Indonesia", Symbol: "IDR"}, "LBP": ut.Currency{Currency: "LBP", DisplayName: "Pauni ya Lebanon", Symbol: "LBP"}, "TMT": ut.Currency{Currency: "TMT", DisplayName: "Manat ya Turukimenistani", Symbol: "TMT"}, "UYU": ut.Currency{Currency: "UYU", DisplayName: "Peso ya Urugwai", Symbol: "UYU"}, "AFN": ut.Currency{Currency: "AFN", DisplayName: "Afghani ya Afghanistan", Symbol: "AFN"}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "Dola ya Australia", Symbol: "A$"}, "BND": ut.Currency{Currency: "BND", DisplayName: "Dola ya Brunei", Symbol: "BND"}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "Lilangeni", Symbol: "SZL"}, "UAH": ut.Currency{Currency: "UAH", DisplayName: "Hryvnia ya Ukrania", Symbol: "UAH"}, "USD": ut.Currency{Currency: "USD", DisplayName: "Dola ya Marekani", Symbol: "US$"}, "YER": ut.Currency{Currency: "YER", DisplayName: "Rial ya Yemeni", Symbol: "YER"}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "Faranga ya Kongo", Symbol: "CDF"}, "COP": ut.Currency{Currency: "COP", DisplayName: "Peso ya Kolombia", Symbol: "COP"}, "FKP": ut.Currency{Currency: "FKP", DisplayName: "Pauni ya Visiwa vya Falkland", Symbol: "FKP"}, "GTQ": ut.Currency{Currency: "GTQ", DisplayName: "Quetzal ya Guatemala", Symbol: "GTQ"}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "Tenge ya Kazakistani", Symbol: "KZT"}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "Ruble ya Urusi", Symbol: "RUB"}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "CFA faranga ya BEAC", Symbol: "FCFA"}, "ISK": ut.Currency{Currency: "ISK", DisplayName: "Krona ya Aisilandi", Symbol: "ISK"}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "Dola ya Namibia", Symbol: "NAD"}, "PEN": ut.Currency{Currency: "PEN", DisplayName: "Nuevo Sol ya Peru", Symbol: "PEN"}, "SYP": ut.Currency{Currency: "SYP", DisplayName: "Pauni ya Syria", Symbol: "SYP"}, "PAB": ut.Currency{Currency: "PAB", DisplayName: "Balboa ya Panama", Symbol: "PAB"}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "Convertible Mark ya Bosnia na Hezegovina", Symbol: "BAM"}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "Dinar ya Bahareni", Symbol: "BHD"}, "JOD": ut.Currency{Currency: "JOD", DisplayName: "Dinar ya Yordani", Symbol: "JOD"}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "Ariari ya Madagaska", Symbol: "MGA"}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "Rupia ya Ushelisheli", Symbol: "SCR"}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "Dola ya Kanada", Symbol: "CA$"}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "Yuro", Symbol: "€"}, "KRW": ut.Currency{Currency: "KRW", DisplayName: "Won ya Korea Kusini", Symbol: "₩"}, "KWD": ut.Currency{Currency: "KWD", DisplayName: "Dinar ya Kuwaiti", Symbol: "KWD"}, "BZD": ut.Currency{Currency: "BZD", DisplayName: "Dola ya Belize", Symbol: "BZD"}, "GHS": ut.Currency{Currency: "GHS", DisplayName: "Cedi ya Ghana", Symbol: "GHS"}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "Faranga ya Guinea", Symbol: "GNF"}, "KHR": ut.Currency{Currency: "KHR", DisplayName: "Riel ya Kambodia", Symbol: "KHR"}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "Dinari ya Libya", Symbol: "LYD"}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "Kwacha ya Malawi", Symbol: "MWK"}, "ILS": ut.Currency{Currency: "ILS", DisplayName: "Sheqel Mpya ya Israeli", Symbol: "₪"}, "MVR": ut.Currency{Currency: "MVR", DisplayName: "Rufiyaa ya Maldivi", Symbol: "MVR"}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "Faranga ya Rwanda", Symbol: "RWF"}, "VND": ut.Currency{Currency: "VND", DisplayName: "Dong ya Vietnam", Symbol: "₫"}, "TTD": ut.Currency{Currency: "TTD", DisplayName: "Dola ya Trinidad na Tobago", Symbol: "TTD"}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "Faranga ya Burundi", Symbol: "BIF"}, "BYR": ut.Currency{Currency: "BYR", DisplayName: "Ruble ya Belarusi", Symbol: "BYR"}, "CLP": ut.Currency{Currency: "CLP", DisplayName: "Peso ya Chile", Symbol: "CLP"}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "Rupia ya Sri Lanka", Symbol: "LKR"}, "MYR": ut.Currency{Currency: "MYR", DisplayName: "Ringgit ya Malaysia", Symbol: "MYR"}, "NZD": ut.Currency{Currency: "NZD", DisplayName: "Dola ya Nyuzilandi", Symbol: "NZ$"}, "TND": ut.Currency{Currency: "TND", DisplayName: "Dinari ya Tunisia", Symbol: "TND"}, "XPF": ut.Currency{Currency: "XPF", DisplayName: "Faranga ya CFP", Symbol: "CFPF"}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "Faranga ya Jibuti", Symbol: "DJF"}, "LVL": ut.Currency{Currency: "LVL", DisplayName: "Lats ya Lativia", Symbol: "LVL"}, "MDL": ut.Currency{Currency: "MDL", DisplayName: "Leu ya Moldova", Symbol: "MDL"}, "MXN": ut.Currency{Currency: "MXN", DisplayName: "Peso ya Meksiko", Symbol: "MX$"}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "Zloty ya Polandi", Symbol: "PLN"}, "SDP": ut.Currency{Currency: "SDP", DisplayName: "Pauni ya Sudani (1957–1998)", Symbol: ""}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "Shilingi ya Tanzania", Symbol: "TSh"}, "ANG": ut.Currency{Currency: "ANG", DisplayName: "Guilder ya Antili za Kiholanzi", Symbol: "ANG"}, "BSD": ut.Currency{Currency: "BSD", DisplayName: "Dola ya Bahamas", Symbol: "BSD"}, "GEL": ut.Currency{Currency: "GEL", DisplayName: "Lari ya Georgia", Symbol: "GEL"}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "Pauni ya Misri", Symbol: "EGP"}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "Dola ya Liberia", Symbol: "LRD"}, "QAR": ut.Currency{Currency: "QAR", DisplayName: "Rial ya Qatari", Symbol: "QAR"}, "UZS": ut.Currency{Currency: "UZS", DisplayName: "Som ya Uzibekistani", Symbol: "UZS"}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "Randi ya Afrika Kusini", Symbol: "ZAR"}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "Krone ya Denmaki", Symbol: "DKK"}, "HNL": ut.Currency{Currency: "HNL", DisplayName: "Lempira ya Hondurasi", Symbol: "HNL"}, "IRR": ut.Currency{Currency: "IRR", DisplayName: "Rial ya Iran", Symbol: "IRR"}, "PKR": ut.Currency{Currency: "PKR", DisplayName: "Rupia ya Pakistani", Symbol: "PKR"}, "RSD": ut.Currency{Currency: "RSD", DisplayName: "Dinar ya Serbia", Symbol: "RSD"}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "Sarafu isiyojulikana", Symbol: ""}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "Dalasi ya Gambia", Symbol: "GMD"}, "LTL": ut.Currency{Currency: "LTL", DisplayName: "Litas ya Lithuania", Symbol: "LTL"}, "SBD": ut.Currency{Currency: "SBD", DisplayName: "Dola ya Visiwa vya Solomon", Symbol: "SBD"}, "MZN": ut.Currency{Currency: "MZN", DisplayName: "Metikali ya Msumbiji", Symbol: "MZN"}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "Taka ya Bangladeshi", Symbol: "BDT"}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "Yuan ya Uchina", Symbol: "CN¥"}, "CUC": ut.Currency{Currency: "CUC", DisplayName: "Peso ya Cuba Inayoweza Kubadilishwa", Symbol: "CUC"}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "Eskudo ya Cape Verde", Symbol: "CVE"}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "Dinari ya Aljeria", Symbol: "DZD"}, "KGS": ut.Currency{Currency: "KGS", DisplayName: "Som ya Kyrgystani", Symbol: "KGS"}, "KYD": ut.Currency{Currency: "KYD", DisplayName: "Dola ya Visiwa vya Cayman", Symbol: "KYD"}, "THB": ut.Currency{Currency: "THB", DisplayName: "Baht ya Tailandi", Symbol: "฿"}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "Pauni ya Sudani", Symbol: "SDG"}, "VEF": ut.Currency{Currency: "VEF", DisplayName: "Bolivar ya Venezuela", Symbol: "VEF"}}
