package my

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"KES": ut.Currency{Currency: "KES", DisplayName: "ကင်ညာသျှီလင်", Symbol: ""}, "KGS": ut.Currency{Currency: "KGS", DisplayName: "ခရူဂစ်စတန်ဆော်မ်", Symbol: ""}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "မြန်မာကျပ်", Symbol: "K"}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "မွန်ဂိုးလီးယားထူးဂရခ်", Symbol: ""}, "UAH": ut.Currency{Currency: "UAH", DisplayName: "ယူကရိန်း", Symbol: ""}, "XDR": ut.Currency{Currency: "XDR", DisplayName: "အထူးထုတ်ယူခွင့်", Symbol: ""}, "AWG": ut.Currency{Currency: "AWG", DisplayName: "အရူဘန် ဂင်းဒါး", Symbol: ""}, "BEF": ut.Currency{Currency: "BEF", DisplayName: "ဘယ်လ်ဂျီယမ် ဖရန့်", Symbol: ""}, "WST": ut.Currency{Currency: "WST", DisplayName: "စမိုအထားလာ", Symbol: ""}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "ခေ့ပ်ဗာဒူ အက်စ်ခူဒို", Symbol: "CVE"}, "MZN": ut.Currency{Currency: "MZN", DisplayName: "မိုဇန်ဘစ်မက်တီခယ်လ်", Symbol: ""}, "CYP": ut.Currency{Currency: "CYP", DisplayName: "ဆိုက်ပရက်စ် ပေါင်", Symbol: ""}, "CZK": ut.Currency{Currency: "CZK", DisplayName: "ချက်ခိုရိုနာ", Symbol: ""}, "KHR": ut.Currency{Currency: "KHR", DisplayName: "ကမ္ဘောဒီးယား ရီးယဲ", Symbol: ""}, "TND": ut.Currency{Currency: "TND", DisplayName: "တူနီရှားဒီနာ", Symbol: ""}, "AMD": ut.Currency{Currency: "AMD", DisplayName: "အာမေးနီးယားဒရမ်း", Symbol: ""}, "BSD": ut.Currency{Currency: "BSD", DisplayName: "ဘဟားမား ဒေါ်လာ", Symbol: ""}, "ESP": ut.Currency{Currency: "ESP", DisplayName: "စပိန် ပယ်စေးတာ", Symbol: ""}, "SDP": ut.Currency{Currency: "SDP", DisplayName: "ဆူဒန် ပေါင်အဟောင်း", Symbol: ""}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "စင်္ကာပူ ဒေါ်လာ", Symbol: ""}, "TJS": ut.Currency{Currency: "TJS", DisplayName: "တာဂျီကစ္စတန်ဆိုမိုနီ", Symbol: ""}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "ကွန်ဂို ဖရန့်", Symbol: ""}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "ဆွစ် ဖရန့်", Symbol: ""}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "ဆေးရှလ်ရူးပီး", Symbol: ""}, "SSP": ut.Currency{Currency: "SSP", DisplayName: "တောင်ဆူဒန်ပေါင်", Symbol: ""}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "ဘော့စနီးယား နှင့် ဟာဇီဂိုဘီးနားမတ်က်", Symbol: ""}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "အီသီယိုးပီးယားဘီးယာ", Symbol: ""}, "JMD": ut.Currency{Currency: "JMD", DisplayName: "ဂျမေကာ ဒေါ်လာ", Symbol: ""}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "ဂျပန်ယန်း", Symbol: "JP¥"}, "KRW": ut.Currency{Currency: "KRW", DisplayName: "တောင်ကိုးရီးယား ဝမ်", Symbol: "₩"}, "QAR": ut.Currency{Currency: "QAR", DisplayName: "ကာတာရီအော်လ်", Symbol: ""}, "AFN": ut.Currency{Currency: "AFN", DisplayName: "အာဖဂန်အာဖဂါနီ", Symbol: ""}, "ALL": ut.Currency{Currency: "ALL", DisplayName: "အယ်လ်ဘီးနီးယားလီခ်", Symbol: ""}, "MVR": ut.Currency{Currency: "MVR", DisplayName: "မော်လ်ဒိုက်ရူးဖီရာ", Symbol: ""}, "VND": ut.Currency{Currency: "VND", DisplayName: "ဗီယက်နမ် ဒေါင်", Symbol: "₫"}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "မသိသို့မဟုတ်မရှိသောငွေကြေး", Symbol: ""}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "လီဗျာ ဒီနာ", Symbol: ""}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "မလာဂစ်စီ အရီရရီ ငွေကြေး", Symbol: ""}, "ISK": ut.Currency{Currency: "ISK", DisplayName: "အိုက်စလန် ခရိုဏာ", Symbol: ""}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "တန်ဇန်းနီးယားသျှီလင်", Symbol: ""}, "XAG": ut.Currency{Currency: "XAG", DisplayName: "ငွေ", Symbol: ""}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "ဘူရွန်ဒီ ဖရန့်", Symbol: ""}, "DEM": ut.Currency{Currency: "DEM", DisplayName: "ဂျာမဏီ မတ်", Symbol: ""}, "UYU": ut.Currency{Currency: "UYU", DisplayName: "ဥရုဂွေးပီဆို", Symbol: ""}, "JOD": ut.Currency{Currency: "JOD", DisplayName: "ဂျော်ဒန်ဒီနား", Symbol: ""}, "RON": ut.Currency{Currency: "RON", DisplayName: "ရိုမေးနီယားလယ်အို", Symbol: ""}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "ဘင်္ဂလားဒေ့ရှ် တာကာ", Symbol: ""}, "CRC": ut.Currency{Currency: "CRC", DisplayName: "ကော့စ်တာရီကာ ခိုလုံး", Symbol: "စီအာစီ"}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "အဲလ်ဂျီရီယန် ဒီနာ", Symbol: ""}, "KYD": ut.Currency{Currency: "KYD", DisplayName: "ကေမန် ကျွန်းစု ဒေါ်လာ", Symbol: ""}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "လာအိုခိပ်", Symbol: ""}, "TMT": ut.Currency{Currency: "TMT", DisplayName: "တာခ်မီန့စ်တန်မာနတ်", Symbol: ""}, "AED": ut.Currency{Currency: "AED", DisplayName: "အာရပ်ဒူဟမ်း", Symbol: ""}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "အင်ဂိုလာ ကန်ဇာ", Symbol: ""}, "TWD": ut.Currency{Currency: "TWD", DisplayName: "ထိုင်ဝမ် ဒေါ်လာအသစ်", Symbol: "NT$"}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "သီရိလင်္ကာ ရူပီး", Symbol: ""}, "MKD": ut.Currency{Currency: "MKD", DisplayName: "မာစီဒိုးနီးယားဒီနာ", Symbol: ""}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "ဆိုမာလီသျှီလင်", Symbol: ""}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "တွန်ဂါဗန်ဂါ", Symbol: ""}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "အီရီတရီအာနာ့ခ်ဖာ", Symbol: ""}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "ဂီးနီ ဖရန့်", Symbol: "GNF"}, "ILP": ut.Currency{Currency: "ILP", DisplayName: "အစ္စရေး ပေါင်", Symbol: ""}, "IQD": ut.Currency{Currency: "IQD", DisplayName: "အီရပ်ဒီနား", Symbol: ""}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "ကိုမိုရိုစ် ဖရန့်", Symbol: ""}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "မောရီတာနီအာအူဂီးယာ", Symbol: "MRO"}, "RSD": ut.Currency{Currency: "RSD", DisplayName: "ဆားဘီးယားဒယ်နား", Symbol: ""}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "ကင်မရွန်းဖရန့်", Symbol: "FCFA"}, "GEL": ut.Currency{Currency: "GEL", DisplayName: "ဂျော်ဂျီယာလားရီ", Symbol: ""}, "GTQ": ut.Currency{Currency: "GTQ", DisplayName: "ဂွာတီမာလာ ခက်ဇော်လ်", Symbol: "GTQ"}, "LVL": ut.Currency{Currency: "LVL", DisplayName: "လတ်ဗီးယားလတ်", Symbol: ""}, "OMR": ut.Currency{Currency: "OMR", DisplayName: "အိုမန်ရီအော်လ်", Symbol: ""}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "ဒိန်းမတ်ခရိုဏာ", Symbol: ""}, "GYD": ut.Currency{Currency: "GYD", DisplayName: "ဂူရာနာ ဒေါ်လာ", Symbol: ""}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "ဂျီဘူတီ ဖရန့်", Symbol: ""}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "တောင်အဖရိက ရန်း", Symbol: ""}, "ZWD": ut.Currency{Currency: "ZWD", DisplayName: "ဇင်ဘာဘွေ ဒေါ်လာ", Symbol: ""}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "ဘာရိန်းဒီနား", Symbol: ""}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "ဘော့စ်ဝါနာ ပုလ", Symbol: ""}, "HUF": ut.Currency{Currency: "HUF", DisplayName: "ဟန်ဂေရီယံဖော်ရင့်တ်", Symbol: ""}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "ခရူဂစ်စတန်ထိန်ဂျီ", Symbol: ""}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "ဆွီဒင် ခရိုဏာ", Symbol: ""}, "YER": ut.Currency{Currency: "YER", DisplayName: "ယီမင်ရီအော်လ်", Symbol: ""}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "ဗြိတိသျှ ပေါင်", Symbol: "£"}, "HTG": ut.Currency{Currency: "HTG", DisplayName: "ဟေတီဂူးအော်ဒ်", Symbol: ""}, "PKR": ut.Currency{Currency: "PKR", DisplayName: "ပါကစ္စတန် ရူပီး", Symbol: ""}, "USD": ut.Currency{Currency: "USD", DisplayName: "အမေရိကန် ဒေါ်လာ", Symbol: "US$"}, "LBP": ut.Currency{Currency: "LBP", DisplayName: "လက်ဘနွန် ပေါင်", Symbol: ""}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "နော်ဝေ ခရိုဏာ", Symbol: ""}, "MOP": ut.Currency{Currency: "MOP", DisplayName: "မကာအိုပါတားကား", Symbol: ""}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "မော်ရေရှားစ် ရူပီ", Symbol: ""}, "TRY": ut.Currency{Currency: "TRY", DisplayName: "တူရကီ လိုင်ရာ", Symbol: ""}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "ဂန်ဘီယာ ဒါလာစီ", Symbol: "GMD"}, "MDL": ut.Currency{Currency: "MDL", DisplayName: "မောလ်ဒိုဗာလယ်အို", Symbol: ""}, "ILS": ut.Currency{Currency: "ILS", DisplayName: "အစ္စရေးရှဲကလ်အသစ်", Symbol: "₪"}, "PYG": ut.Currency{Currency: "PYG", DisplayName: "ပါရာဂွေးဂွါးအ်နီး", Symbol: ""}, "TRL": ut.Currency{Currency: "TRL", DisplayName: "ရှေးဟောင်းတူရကီ လိုင်ရာ", Symbol: ""}, "BND": ut.Currency{Currency: "BND", DisplayName: "ဘရူနိုင်း ဒေါ်လာ", Symbol: "BND"}, "BZD": ut.Currency{Currency: "BZD", DisplayName: "ဘေလီဇ် ဒေါ်လာ", Symbol: "ဒေါ်လာ"}, "FJD": ut.Currency{Currency: "FJD", DisplayName: "ဖီဂျီ ဒေါ်လာ", Symbol: ""}, "GIP": ut.Currency{Currency: "GIP", DisplayName: "ဂျီဘရော်လ်တာ ပေါင်", Symbol: ""}, "INR": ut.Currency{Currency: "INR", DisplayName: "အိန္ဒိယ ရူပီး", Symbol: "₹"}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "မိုရိုကို ဒရမ်", Symbol: ""}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "နမ်မီးဘီးယား ဒေါ်လာ", Symbol: ""}, "NIO": ut.Currency{Currency: "NIO", DisplayName: "နီကာရာဂွာ ခိုးဒိုဘာ", Symbol: ""}, "BGN": ut.Currency{Currency: "BGN", DisplayName: "ဘူဂေးရီးယားလက်ဖ်", Symbol: ""}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "ကနေဒါ ဒေါ်လာ", Symbol: "CA$"}, "PEN": ut.Currency{Currency: "PEN", DisplayName: "ပီရူး နူအီဗိုဆိုးလ်", Symbol: ""}, "RUR": ut.Currency{Currency: "RUR", DisplayName: "ရုရှ ရူဘယ် (၁၉၉၁–၁၉၉၈)", Symbol: ""}, "LTL": ut.Currency{Currency: "LTL", DisplayName: "လစ်သူယေးနီးယားလီတားစ်", Symbol: ""}, "BMD": ut.Currency{Currency: "BMD", DisplayName: "ဘာမူဒါ ဒေါ်လာ", Symbol: "BMD"}, "COP": ut.Currency{Currency: "COP", DisplayName: "ကိုလံဘီယာ ပီဆို", Symbol: ""}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "အိုင်ဗရီးကိုးစ်ဖရန့်", Symbol: "CFA"}, "XTS": ut.Currency{Currency: "XTS", DisplayName: "စမ်းသပ် ငွေကြေး ကုဒ်", Symbol: ""}, "SUR": ut.Currency{Currency: "SUR", DisplayName: "ဆိုဗီယက် ရူဗယ်", Symbol: ""}, "THB": ut.Currency{Currency: "THB", DisplayName: "ထိုင်းဘတ်", Symbol: "฿"}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "မာလာဝီခွါးချာ", Symbol: ""}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "စိန့်ဟဲလီနာ ပေါင်", Symbol: "SHP"}, "USN": ut.Currency{Currency: "USN", DisplayName: "အမေရိကန် ဒေါ်လာ (နောက်နေ့)", Symbol: ""}, "VEF": ut.Currency{Currency: "VEF", DisplayName: "ဗင်နီဇွဲလား ဘိုလီဗာ", Symbol: ""}, "BYB": ut.Currency{Currency: "BYB", DisplayName: "ဘီလာရုစ် ရူဘယ်အသစ် (၁၉၉၄–၁၉၉၉)", Symbol: ""}, "DOP": ut.Currency{Currency: "DOP", DisplayName: "ဒိုမီနီကန် ပီဆို", Symbol: ""}, "KPW": ut.Currency{Currency: "KPW", DisplayName: "မြောက်ကိုးရီးယား ဝမ်", Symbol: ""}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "ရုရှ ရူဘယ်", Symbol: ""}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "ဆီအဲရာ လီအိုနီယန် လီအိုနီ", Symbol: "SLL"}, "XAU": ut.Currency{Currency: "XAU", DisplayName: "ရွှေ", Symbol: ""}, "XPT": ut.Currency{Currency: "XPT", DisplayName: "ပလက်တီနမ်", Symbol: ""}, "AZN": ut.Currency{Currency: "AZN", DisplayName: "အဇာဘိုင်ဂျန်မာနတ်", Symbol: ""}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "တရုတ် ယွမ်", Symbol: "CN¥"}, "TTD": ut.Currency{Currency: "TTD", DisplayName: "ထရိုင်နီဒတ်နှင့်တိုဘာဂိုဒေါ်လာ", Symbol: ""}, "FRF": ut.Currency{Currency: "FRF", DisplayName: "ပြင်သစ် ဖရန့်", Symbol: ""}, "HRK": ut.Currency{Currency: "HRK", DisplayName: "ခရိုအေးရှားခူးနာ", Symbol: ""}, "BUK": ut.Currency{Currency: "BUK", DisplayName: "ဗမာ ကျပ်", Symbol: ""}, "FKP": ut.Currency{Currency: "FKP", DisplayName: "ဖောက်ကလန် ကျွန်းစု ပေါင်", Symbol: ""}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "နိုင်ဂျီးရီးယားနိုင်းရာ", Symbol: "NGN"}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "အရှေ့ကာရီဘီယံဒေါ်လာ", Symbol: "EC$"}, "BOB": ut.Currency{Currency: "BOB", DisplayName: "ဘိုလီဘီယံ ဘိုလီဘီအားနို", Symbol: ""}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "ဘရာဇီး ရီးယဲ", Symbol: "R$"}, "KWD": ut.Currency{Currency: "KWD", DisplayName: "ကူဝိတ်ဒီနာ", Symbol: ""}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "ပိုလန် ဇ\u200cလော့တီ", Symbol: ""}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "ဇင်ဘာဘွေခွါးချာ", Symbol: ""}, "ARS": ut.Currency{Currency: "ARS", DisplayName: "အာဂျင်တီးနား ပီဆို", Symbol: ""}, "IRR": ut.Currency{Currency: "IRR", DisplayName: "အီရန်ရီအော်လ်", Symbol: ""}, "CUP": ut.Currency{Currency: "CUP", DisplayName: "ကျူးဘား ပီဆို", Symbol: ""}, "HKD": ut.Currency{Currency: "HKD", DisplayName: "ဟောင်ကောင် ဒေါ်လာ", Symbol: "HK$"}, "PGK": ut.Currency{Currency: "PGK", DisplayName: "ပါပူရာနယူးဂီနီခီးနာ", Symbol: ""}, "SRD": ut.Currency{Currency: "SRD", DisplayName: "ဆူရီနိမ်း ဒေါ်လာ", Symbol: ""}, "UZS": ut.Currency{Currency: "UZS", DisplayName: "ဥဘက်ကစ္စတန်ဆော်မ်", Symbol: ""}, "XPF": ut.Currency{Currency: "XPF", DisplayName: "CFP ဖရန့်", Symbol: "CFPF"}, "BOP": ut.Currency{Currency: "BOP", DisplayName: "ဘိုလီးဘီးယား ပီဆို", Symbol: ""}, "BYR": ut.Currency{Currency: "BYR", DisplayName: "ဘီလာရုစ် ရူဘယ်", Symbol: ""}, "VUV": ut.Currency{Currency: "VUV", DisplayName: "ဗာနုအာတူဗားထူ", Symbol: ""}, "GHS": ut.Currency{Currency: "GHS", DisplayName: "ဂါနာ ဆဲဒီ", Symbol: "GHS"}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "လိုင်ဘေးရီးယား ဒေါ်လာ", Symbol: "LRD"}, "SBD": ut.Currency{Currency: "SBD", DisplayName: "ဆော်လမွန်ကျွန်းစု ဒေါ်လာ", Symbol: ""}, "NZD": ut.Currency{Currency: "NZD", DisplayName: "နယူးဇီလန် ဒေါ်လာ", Symbol: "NZ$"}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "ရဝန်ဒါ ဖရန့်", Symbol: ""}, "BBD": ut.Currency{Currency: "BBD", DisplayName: "ဘာဘဒီယံဒေါ်လာ", Symbol: ""}, "CLP": ut.Currency{Currency: "CLP", DisplayName: "ချီလီ ပီဆို", Symbol: ""}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "အီဂျစ် ပေါင်", Symbol: ""}, "HNL": ut.Currency{Currency: "HNL", DisplayName: "ဟွန်ဒူးရပ်စ် လန်းပီးရာ", Symbol: ""}, "MXN": ut.Currency{Currency: "MXN", DisplayName: "မက္ကဆီကို ပီဆို", Symbol: "MX$"}, "MYR": ut.Currency{Currency: "MYR", DisplayName: "မလေးရှား ရင်းဂစ်", Symbol: ""}, "ANG": ut.Currency{Currency: "ANG", DisplayName: "နယ်သာလန် အန်တီလန် ဂင်းဒါး", Symbol: ""}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "ဩစတြေးလျ ဒေါ်လာ", Symbol: "A$"}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "ဆူဒန် ပေါင်", Symbol: ""}, "STD": ut.Currency{Currency: "STD", DisplayName: "စိန့်တိုမီနှင့်ပရင်စီပ့် ဒိုဘရာ", Symbol: ""}, "SYP": ut.Currency{Currency: "SYP", DisplayName: "ဆီးရီးယား ပေါင်", Symbol: ""}, "XBB": ut.Currency{Currency: "XBB", DisplayName: "ဥရောပငွေကြေးစံနစ်", Symbol: ""}, "NPR": ut.Currency{Currency: "NPR", DisplayName: "နီပေါ ရူပီး", Symbol: ""}, "PHP": ut.Currency{Currency: "PHP", DisplayName: "ဖိလစ်ပိုင် ပီဆို", Symbol: ""}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "ယူရို", Symbol: "€"}, "PAB": ut.Currency{Currency: "PAB", DisplayName: "ပနားမား ဘလ်ဘိုးအာ", Symbol: ""}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "စွာဇီလန်လီလန်းဂီနီ", Symbol: ""}, "USS": ut.Currency{Currency: "USS", DisplayName: "အမေရိကန် ဒေါ်လာ (တနေ့တည်း)", Symbol: ""}, "BTN": ut.Currency{Currency: "BTN", DisplayName: "ဘူတန်အံဂါလ်ထရန်", Symbol: ""}, "CUC": ut.Currency{Currency: "CUC", DisplayName: "နိုင်ငံခြားငွေလဲလှယ်နိုင်သော ကျူးဘားပီဆိုငွေ", Symbol: ""}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "ဆော်ဒီအာရေးဗီးယားရီယော်လ်", Symbol: ""}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "ယူဂန္ဓာသျှီလင်", Symbol: ""}, "ARP": ut.Currency{Currency: "ARP", DisplayName: "အာဂျင်တီးနား ပီဆို (၁၉၈၃–၁၉၈၅)", Symbol: ""}, "IDR": ut.Currency{Currency: "IDR", DisplayName: "အင်ဒိုနီးရှား ရူပီးယား", Symbol: ""}}
