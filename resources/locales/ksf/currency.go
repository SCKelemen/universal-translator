package ksf

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"AUD": ut.Currency{Currency: "AUD", DisplayName: "mɔni mǝ á ɔstralí", Symbol: ""}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "mɔni mǝ á arabí saodí", Symbol: ""}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "mɔni mǝ á dyibutí", Symbol: ""}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "mɔni mǝ á pɛrɛsǝ́", Symbol: ""}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "mɔni mǝ á gambí", Symbol: ""}, "TND": ut.Currency{Currency: "TND", DisplayName: "mɔni mǝ á tunɛsí", Symbol: ""}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "mɔni mǝ á kanada", Symbol: ""}, "MZM": ut.Currency{Currency: "MZM", DisplayName: "mɔni mǝ á mosambík", Symbol: ""}, "AED": ut.Currency{Currency: "AED", DisplayName: "mɔni mǝ á bǝlɔŋ bǝ kaksa bɛ táatáaŋzǝn", Symbol: ""}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "mɔni mǝ á kɔngó", Symbol: ""}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "mɔni mǝ á marɔk", Symbol: ""}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "mɔni mǝ á namibí", Symbol: ""}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "mɔni mǝ á sudan", Symbol: ""}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "mɔni mǝ á sɛntɛ́len", Symbol: ""}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "mɔni mǝ á zambí", Symbol: ""}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "mɔni mǝ á ɛritrɛ́", Symbol: ""}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "mɔni mǝ á ɛtyɔpí", Symbol: ""}, "USD": ut.Currency{Currency: "USD", DisplayName: "mɔni mǝ á amɛrika", Symbol: ""}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "mɔni mǝ á afríka aná wɛs", Symbol: ""}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "mɔni mǝ á ingɛrís", Symbol: ""}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "mɔni mǝ á libɛrya", Symbol: ""}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "mɔni mǝ á mwarís", Symbol: ""}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "mɔni mǝ á ɛjípt", Symbol: ""}, "KES": ut.Currency{Currency: "KES", DisplayName: "mɔni mǝ á kɛnya", Symbol: ""}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "mɔni mǝ á barǝ́n", Symbol: ""}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "mɔni mǝ á burundí", Symbol: ""}, "GHC": ut.Currency{Currency: "GHC", DisplayName: "mɔni mǝ á gána", Symbol: ""}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "mɔni mǝ á libí", Symbol: ""}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "mɔni mǝ á swazilan", Symbol: ""}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "mɔni mǝ á mwaritaní", Symbol: ""}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "mɔni mǝ á nijɛ́rya", Symbol: ""}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "mɔni mǝ á rwanda", Symbol: ""}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "mɔni mǝ á syɛraleon", Symbol: ""}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "mɔni mǝ á uganda", Symbol: ""}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "mɔni mǝ á swís", Symbol: ""}, "LSL": ut.Currency{Currency: "LSL", DisplayName: "mɔni mǝ á lǝsóto", Symbol: ""}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "fráŋ", Symbol: ""}, "GNS": ut.Currency{Currency: "GNS", DisplayName: "mɔni mǝ á ginɛ́", Symbol: ""}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "mɔni mǝ á madagaska", Symbol: ""}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "mɔni mǝ á somalí", Symbol: ""}, "ZMK": ut.Currency{Currency: "ZMK", DisplayName: "mɔni mǝ á zambí (1968–2012)", Symbol: ""}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "mɔni mǝ á botswana", Symbol: ""}, "INR": ut.Currency{Currency: "INR", DisplayName: "mɔni mǝ á indí", Symbol: ""}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "mɔni mǝ á komɔr", Symbol: ""}, "STD": ut.Currency{Currency: "STD", DisplayName: "mɔni mǝ á saotomɛ́ ri priŋsib", Symbol: ""}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "mɔni mǝ á tanzaní", Symbol: ""}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "mɔni mǝ á angóla", Symbol: ""}, "ZWD": ut.Currency{Currency: "ZWD", DisplayName: "mɔni mǝ á zimbabwɛ́", Symbol: ""}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "mɔni mǝ á cín", Symbol: ""}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "mɔni mǝ á kapvɛr", Symbol: ""}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "mɔni mǝ á aljɛrí", Symbol: ""}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "mɔni mǝ á japɔ́ŋ", Symbol: ""}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "mɔni mǝ á malawi", Symbol: ""}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "mɔni mǝ á sɛcɛl", Symbol: ""}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "mɔni mǝ á afrik anǝ a sud", Symbol: ""}}
