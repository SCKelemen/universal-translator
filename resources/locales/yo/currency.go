package yo

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"SOS": ut.Currency{Currency: "SOS", DisplayName: "Sile ti Orílẹ́ède Somali", Symbol: ""}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "Sile ti Orílẹ́ède Tansania", Symbol: ""}, "INR": ut.Currency{Currency: "INR", DisplayName: "Rupi ti Orílẹ́ède Indina", Symbol: ""}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "Pọọun ti Orílẹ́ède ̣Elena", Symbol: ""}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "Dina ti Orílẹ́ède Libiya", Symbol: ""}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "Lilangeni", Symbol: ""}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "Reminibi ti Orílẹ́ède ṣáínà", Symbol: ""}, "KES": ut.Currency{Currency: "KES", DisplayName: "ṣiili ti Orílẹ́ède Kenya", Symbol: ""}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "Lioni", Symbol: ""}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "Dina ti Orílẹ́ède Báránì", Symbol: ""}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "Dọla ti Orílẹ́ède Kánádà", Symbol: ""}, "ZMK": ut.Currency{Currency: "ZMK", DisplayName: "Kawaṣa ti Orílẹ́ède Saabia (1968–2012)", Symbol: ""}, "MZM": ut.Currency{Currency: "MZM", DisplayName: "Metika ti Orílẹ́ède Mosamibiki", Symbol: ""}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "Faransi ti Orílẹ́ède Bùùrúndì", Symbol: ""}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "Ouguiya ti Orílẹ́ède Maritania", Symbol: ""}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "Faransi ti Orílẹ́ède Ruwanda", Symbol: ""}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "Dina ti Orílẹ́ède Sudani", Symbol: ""}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "Siile ti Orílẹ́ède Uganda", Symbol: ""}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "Faransi ti Orílẹ́ède BEKA", Symbol: ""}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "Wansa ti Orílẹ́ède Àngólà", Symbol: ""}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "Biri ti Orílẹ́ède Eutopia", Symbol: ""}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "Dọla ti Orílẹ́ède Liberia", Symbol: ""}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "Naira ti Orílẹ́ède Nàìjíríà", Symbol: "₦"}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "Uro", Symbol: ""}, "SDP": ut.Currency{Currency: "SDP", DisplayName: "Pọọun ti Orílẹ́ède Sudani", Symbol: ""}, "STD": ut.Currency{Currency: "STD", DisplayName: "Dobira ti Orílẹ́ède Sao tome Ati Pirisipe", Symbol: ""}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "Pọọn ti Orílẹ́ède Bírítísì", Symbol: ""}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "Rupi ti Orílẹ́ède Maritiusi", Symbol: ""}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "Faransi ti Orílẹ́ède Kóngò", Symbol: ""}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "Faransi ti Orílẹ́ède ṣomoriani", Symbol: ""}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "Yeni ti Orílẹ́ède Japani", Symbol: ""}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "Dirami ti Orílẹ́ède Moroko", Symbol: ""}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "Dọla ti Orílẹ́ède Namibia", Symbol: ""}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "Riya ti Orílẹ́ède Saudi", Symbol: ""}, "AED": ut.Currency{Currency: "AED", DisplayName: "Diami ti Awon Orílẹ́ède Arabu", Symbol: ""}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "Pula ti Orílẹ́ède Bọ̀tìsúwánà", Symbol: ""}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "Faransi ti Orílẹ́ède Dibouti", Symbol: ""}, "GNS": ut.Currency{Currency: "GNS", DisplayName: "Faransi ti Orílẹ́ède Gini", Symbol: ""}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "Rupi ti Orílẹ́ède Sayiselesi", Symbol: ""}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "Randi ti Orílẹ́ède Ariwa Afirika", Symbol: ""}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "Faransi ti Orílẹ́ède Siwisi", Symbol: ""}, "TND": ut.Currency{Currency: "TND", DisplayName: "Dina ti Orílẹ́ède Tunisia", Symbol: ""}, "ZWD": ut.Currency{Currency: "ZWD", DisplayName: "Dọla ti Orílẹ́ède Siibabuwe", Symbol: ""}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "Dina ti Orílẹ́ède Àlùgèríánì", Symbol: ""}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "pọọn ti Orílẹ́ède Egipiti", Symbol: ""}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "Kawaṣa ti Orílẹ́ède Saabia", Symbol: ""}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "Faransi ti Orílẹ́ède BIKEAO", Symbol: ""}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "Nakifa ti Orílẹ́ède Eriteriani", Symbol: ""}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "Dalasi ti Orílẹ́ède Gamibia", Symbol: ""}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "Kaṣa ti Orílẹ́ède Malawi", Symbol: ""}, "USD": ut.Currency{Currency: "USD", DisplayName: "Dọla ti Orílẹ́ède Amerika", Symbol: ""}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "Dọla ti Orílẹ́ède Ástràlìá", Symbol: ""}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "Kabofediano ti Orílẹ́ède Esuodo", Symbol: ""}, "LSL": ut.Currency{Currency: "LSL", DisplayName: "Loti ti Orílẹ́ède Lesoto", Symbol: ""}, "GHC": ut.Currency{Currency: "GHC", DisplayName: "ṣidi ti Orílẹ́ède Gana", Symbol: ""}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "Faransi ti Orílẹ́ède Malagasi", Symbol: ""}}
