package pt_PT

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"BSD": ut.Currency{Currency: "BSD", DisplayName: "Dólar das Bahamas", Symbol: "BSD"}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "Tenge do Cazaquistão", Symbol: "KZT"}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "Kip de Laos", Symbol: "LAK"}, "UAH": ut.Currency{Currency: "UAH", DisplayName: "Hryvnia da Ucrânia", Symbol: "UAH"}, "VUV": ut.Currency{Currency: "VUV", DisplayName: "Vatu de Vanuatu", Symbol: "VUV"}, "MXV": ut.Currency{Currency: "MXV", DisplayName: "Unidad de Inversion (UDI) mexicana", Symbol: ""}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "Dólar de Singapura", Symbol: "SGD"}, "TMT": ut.Currency{Currency: "TMT", DisplayName: "Manat do Turquemenistão", Symbol: "TMT"}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "Dólar das Caraíbas Orientais", Symbol: "EC$"}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "Franco jibutiano", Symbol: "DJF"}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "Franco guineense", Symbol: "GNF"}, "ILS": ut.Currency{Currency: "ILS", DisplayName: "Sheqel novo israelita", Symbol: "₪"}, "SRD": ut.Currency{Currency: "SRD", DisplayName: "Dólar do Suriname", Symbol: "SRD"}, "TJS": ut.Currency{Currency: "TJS", DisplayName: "Somoni do Tajaquistão", Symbol: "TJS"}, "YUN": ut.Currency{Currency: "YUN", DisplayName: "Dinar conversível jugoslavo", Symbol: ""}, "ZMK": ut.Currency{Currency: "ZMK", DisplayName: "Kwacha zambiano (1968–2012)", Symbol: "ZMK"}, "AED": ut.Currency{Currency: "AED", DisplayName: "Dirham dos Emirados Árabes Unidos", Symbol: "AED"}, "AWG": ut.Currency{Currency: "AWG", DisplayName: "Florim de Aruba", Symbol: "AWG"}, "BTN": ut.Currency{Currency: "BTN", DisplayName: "Ngultrum do Butão", Symbol: "BTN"}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "Dirham marroquino", Symbol: "MAD"}, "PGK": ut.Currency{Currency: "PGK", DisplayName: "Kina da Papua-Nova Guiné", Symbol: "PGK"}, "RON": ut.Currency{Currency: "RON", DisplayName: "Leu romeno", Symbol: "RON"}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "Xelim somali", Symbol: "SOS"}, "ZWD": ut.Currency{Currency: "ZWD", DisplayName: "Dólar do Zimbabwe", Symbol: ""}, "NIC": ut.Currency{Currency: "NIC", DisplayName: "Córdoba nicaraguano (1988–1991)", Symbol: ""}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "Taka de Bangladesh", Symbol: "BDT"}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "Libra esterlina britânica", Symbol: "£"}, "GHC": ut.Currency{Currency: "GHC", DisplayName: "Cedi do Gana", Symbol: ""}, "MLF": ut.Currency{Currency: "MLF", DisplayName: "Franco do Mali", Symbol: ""}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "Ouguiya da Mauritânia", Symbol: "MRO"}, "MVR": ut.Currency{Currency: "MVR", DisplayName: "Rupia das Ilhas Maldivas", Symbol: "MVR"}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "Dólar da Namíbia", Symbol: "NAD"}, "PAB": ut.Currency{Currency: "PAB", DisplayName: "Balboa do Panamá", Symbol: "PAB"}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "Lilangeni da Suazilândia", Symbol: "SZL"}, "AFA": ut.Currency{Currency: "AFA", DisplayName: "Afeghani (1927–2002)", Symbol: ""}, "AFN": ut.Currency{Currency: "AFN", DisplayName: "Afegani do Afeganistão", Symbol: "AFN"}, "ECV": ut.Currency{Currency: "ECV", DisplayName: "Unidad de Valor Constante (UVC) do Equador", Symbol: ""}, "FJD": ut.Currency{Currency: "FJD", DisplayName: "Dólar de Fiji", Symbol: "FJD"}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "Dalasi da Gâmbia", Symbol: "GMD"}, "KYD": ut.Currency{Currency: "KYD", DisplayName: "Dólar das Ilhas Caimão", Symbol: "KYD"}, "NIO": ut.Currency{Currency: "NIO", DisplayName: "Córdoba nicaraguano", Symbol: "NIO"}, "YUM": ut.Currency{Currency: "YUM", DisplayName: "Super Dinar jugoslavo", Symbol: ""}, "GYD": ut.Currency{Currency: "GYD", DisplayName: "Dólar da Guiana", Symbol: "GYD"}, "PEN": ut.Currency{Currency: "PEN", DisplayName: "Novo sol peruano", Symbol: "PEN"}, "YER": ut.Currency{Currency: "YER", DisplayName: "Rial iemenita", Symbol: "YER"}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "Marco bósnio-herzegóvino conversível", Symbol: "BAM"}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "Zloti polaco", Symbol: "PLN"}, "PLZ": ut.Currency{Currency: "PLZ", DisplayName: "Zloti polaco (1950–1995)", Symbol: ""}, "YUD": ut.Currency{Currency: "YUD", DisplayName: "Dinar forte jugoslavo", Symbol: ""}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "Dinar baremita", Symbol: "BHD"}, "BYB": ut.Currency{Currency: "BYB", DisplayName: "Rublo novo bielorusso (1994–1999)", Symbol: ""}, "MZN": ut.Currency{Currency: "MZN", DisplayName: "Metical de Moçambique", Symbol: "MZN"}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "Rial saudita", Symbol: "SAR"}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "Dólar canadiano", Symbol: "CA$"}, "FKP": ut.Currency{Currency: "FKP", DisplayName: "Libra das Ilhas Falkland", Symbol: "FKP"}, "OMR": ut.Currency{Currency: "OMR", DisplayName: "Rial de Omã", Symbol: "OMR"}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "Paʻanga de Tonga", Symbol: "TOP"}, "TTD": ut.Currency{Currency: "TTD", DisplayName: "Dólar de Trindade e Tobago", Symbol: "TTD"}, "CRC": ut.Currency{Currency: "CRC", DisplayName: "Colon costa-riquenho", Symbol: "CRC"}, "CZK": ut.Currency{Currency: "CZK", DisplayName: "Coroa checa", Symbol: "CZK"}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "Franco comoriano", Symbol: "KMF"}, "LTL": ut.Currency{Currency: "LTL", DisplayName: "Litas da Lituânia", Symbol: "LTL"}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "Kwacha do Malawi", Symbol: "MWK"}, "QAR": ut.Currency{Currency: "QAR", DisplayName: "Rial do Catar", Symbol: "QAR"}, "THB": ut.Currency{Currency: "THB", DisplayName: "Baht da Tailândia", Symbol: "฿"}, "XEU": ut.Currency{Currency: "XEU", DisplayName: "Unidade da Moeda Europeia", Symbol: ""}, "AZN": ut.Currency{Currency: "AZN", DisplayName: "Manat do Azerbaijão", Symbol: "AZN"}, "BAD": ut.Currency{Currency: "BAD", DisplayName: "Dinar da Bósnia-Herzegóvina", Symbol: ""}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "Pula de Botswana", Symbol: "BWP"}, "HNL": ut.Currency{Currency: "HNL", DisplayName: "Lempira das Honduras", Symbol: "HNL"}, "IRR": ut.Currency{Currency: "IRR", DisplayName: "Rial iraniano", Symbol: "IRR"}, "MXP": ut.Currency{Currency: "MXP", DisplayName: "Peso Plata mexicano (1861–1992)", Symbol: ""}, "TND": ut.Currency{Currency: "TND", DisplayName: "Dinar tunisino", Symbol: "TND"}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "Franco CFA (BCEAO)", Symbol: "CFA"}, "PTE": ut.Currency{Currency: "PTE", DisplayName: "Escudo português", Symbol: "\u200b"}, "AMD": ut.Currency{Currency: "AMD", DisplayName: "Dram arménio", Symbol: "AMD"}, "BZD": ut.Currency{Currency: "BZD", DisplayName: "Dólar belizense", Symbol: "BZD"}, "GHS": ut.Currency{Currency: "GHS", DisplayName: "Cedi de Gana", Symbol: "GHS"}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "Rupia do Sri Lanka", Symbol: "LKR"}, "LVL": ut.Currency{Currency: "LVL", DisplayName: "Lats da Letónia", Symbol: "LVL"}, "MKD": ut.Currency{Currency: "MKD", DisplayName: "Dinar macedónio", Symbol: "MKD"}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "Tugrik da Mongólia", Symbol: "MNT"}, "XDR": ut.Currency{Currency: "XDR", DisplayName: "direito especial de saque", Symbol: ""}, "MDL": ut.Currency{Currency: "MDL", DisplayName: "Leu moldavo", Symbol: "MDL"}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "Franco CFA (BEAC)", Symbol: "FCFA"}, "BEC": ut.Currency{Currency: "BEC", DisplayName: "Franco belga (convertível)", Symbol: ""}, "CYP": ut.Currency{Currency: "CYP", DisplayName: "Libra de Chipre", Symbol: ""}, "GTQ": ut.Currency{Currency: "GTQ", DisplayName: "Quetzal da Guatemala", Symbol: "GTQ"}, "KGS": ut.Currency{Currency: "KGS", DisplayName: "Som do Quirguistão", Symbol: "KGS"}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "Kyat de Mianmar", Symbol: "MMK"}, "UZS": ut.Currency{Currency: "UZS", DisplayName: "Som do Uzbequistão", Symbol: "UZS"}, "BND": ut.Currency{Currency: "BND", DisplayName: "Dólar bruneíno", Symbol: "BND"}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "Ariari de Madagáscar", Symbol: "MGA"}, "MOP": ut.Currency{Currency: "MOP", DisplayName: "Pataca de Macau", Symbol: "MOP"}, "USD": ut.Currency{Currency: "USD", DisplayName: "Dólar dos Estados Unidos", Symbol: "US$"}}
