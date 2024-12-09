# GOLANG CLI: CURRENCY CONVERTER

A simple CLI tool to convert currencies using real-time exchange rates from the RapidApi service.

## Installation
1. Clone the repo

2. Install dependencies.
```
go mod init github.com/eliuttth-dev/golang-cli-currency-converter
go mod tidy
```

3. Set up API key: Add .env file with the field API_KEY=<YOUR_OWN_RAPID_API_KEY>

4. Run the program
`go run main.go`

## Currencies Table

# List of Currency Symbols

| Currency Name               | Symbol | Code |
|-----------------------------|--------|------|
| Afghan Afghani              | ؋      | AFN  |
| Albanian Lek                | L      | ALL  |
| Algerian Dinar              | دج     | DZD  |
| Andorran Euro               | €      | EUR  |
| Angolan Kwanza              | Kz     | AOA  |
| Argentine Peso              | $      | ARS  |
| Armenian Dram               | ֏      | AMD  |
| Australian Dollar           | $      | AUD  |
| Azerbaijani Manat           | ₼      | AZN  |
| Bahraini Dinar              | .د.ب   | BHD  |
| Bangladeshi Taka            | ৳      | BDT  |
| Belarusian Ruble            | Br     | BYN  |
| Belize Dollar               | $      | BZD  |
| Bermudian Dollar            | $      | BMD  |
| Bhutanese Ngultrum          | Nu.    | BTN  |
| Bolivian Boliviano          | Bs.    | BOB  |
| Bosnian Convertible Mark    | KM     | BAM  |
| Botswanan Pula              | P      | BWP  |
| Brazilian Real              | R$     | BRL  |
| British Pound Sterling      | £      | GBP  |
| Brunei Dollar               | $      | BND  |
| Bulgarian Lev               | лв     | BGN  |
| Burundian Franc             | FBu    | BIF  |
| Cambodian Riel              | ៛      | KHR  |
| Canadian Dollar             | $      | CAD  |
| Cape Verdean Escudo         | $      | CVE  |
| Cayman Islands Dollar       | $      | KYD  |
| Central African CFA Franc   | Fr     | XAF  |
| Chilean Peso                | $      | CLP  |
| Chinese Yuan Renminbi       | ¥      | CNY  |
| Colombian Peso              | $      | COP  |
| Comorian Franc              | CF     | KMF  |
| Congolese Franc             | FC     | CDF  |
| Costa Rican Colón           | ₡      | CRC  |
| Croatian Kuna               | kn     | HRK  |
| Cuban Peso                  | $      | CUP  |
| Czech Koruna                | Kč     | CZK  |
| Danish Krone                | kr     | DKK  |
| Djiboutian Franc            | Fdj    | DJF  |
| Dominican Peso              | $      | DOP  |
| East Caribbean Dollar       | $      | XCD  |
| Egyptian Pound              | £      | EGP  |
| Eritrean Nakfa              | Nfk    | ERN  |
| Estonian Euro               | €      | EUR  |
| Ethiopian Birr              | Br     | ETB  |
| Euro                        | €      | EUR  |
| Fijian Dollar               | $      | FJD  |
| Gambian Dalasi              | D      | GMD  |
| Georgian Lari               | ₾      | GEL  |
| Ghanaian Cedi               | ₵      | GHS  |
| Guatemalan Quetzal          | Q      | GTQ  |
| Guinean Franc               | FG     | GNF  |
| Guyanese Dollar             | $      | GYD  |
| Haitian Gourde              | G      | HTG  |
| Honduran Lempira            | L      | HNL  |
| Hong Kong Dollar            | $      | HKD  |
| Hungarian Forint            | Ft     | HUF  |
| Icelandic Króna             | kr     | ISK  |
| Indian Rupee                | ₹      | INR  |
| Indonesian Rupiah           | Rp     | IDR  |
| Iranian Rial                | ﷼      | IRR  |
| Iraqi Dinar                 | ع.د    | IQD  |
| Israeli New Shekel          | ₪      | ILS  |
| Jamaican Dollar             | $      | JMD  |
| Japanese Yen                | ¥      | JPY  |
| Jordanian Dinar             | د.ا    | JOD  |
| Kazakhstani Tenge           | ₸      | KZT  |
| Kenyan Shilling             | Sh     | KES  |
| Kuwaiti Dinar               | د.ك    | KWD  |
| Kyrgyzstani Som             | с      | KGS  |
| Lao Kip                     | ₭      | LAK  |
| Lebanese Pound              | ل.ل    | LBP  |
| Lesotho Loti                | L      | LSL  |
| Liberian Dollar             | $      | LRD  |
| Libyan Dinar                | ل.د    | LYD  |
| Macanese Pataca             | MOP$   | MOP  |
| Macedonian Denar            | ден    | MKD  |
| Malagasy Ariary             | Ar     | MGA  |
| Malawian Kwacha             | MK     | MWK  |
| Malaysian Ringgit           | RM     | MYR  |
| Maldivian Rufiyaa           | ރ.     | MVR  |
| Mauritanian Ouguiya         | UM     | MRU  |
| Mauritian Rupee             | ₨      | MUR  |
| Mexican Peso                | $      | MXN  |
| Moldovan Leu                | L      | MDL  |
| Mongolian Tögrög            | ₮      | MNT  |
| Moroccan Dirham             | د.م.   | MAD  |
| Mozambican Metical          | MT     | MZN  |
| Myanmar Kyat                | K      | MMK  |
| Namibian Dollar             | $      | NAD  |
| Nepalese Rupee              | ₨      | NPR  |
| Netherlands Antillean Guilder | ƒ    | ANG  |
| New Taiwan Dollar           | NT$    | TWD  |
| New Zealand Dollar          | $      | NZD  |
| Nicaraguan Córdoba          | C$     | NIO  |
| Nigerian Naira              | ₦      | NGN  |
| North Korean Won            | ₩      | KPW  |
| Norwegian Krone             | kr     | NOK  |
| Omani Rial                  | ﷼      | OMR  |
| Pakistani Rupee             | ₨      | PKR  |
| Panamanian Balboa           | B/.    | PAB  |
| Papua New Guinean Kina      | K      | PGK  |
| Paraguayan Guarani          | ₲      | PYG  |
| Peruvian Sol                | S/.    | PEN  |
| Philippine Peso             | ₱      | PHP  |
| Polish Zloty                | zł     | PLN  |
| Qatari Riyal                | ﷼      | QAR  |
| Romanian Leu                | lei    | RON  |
| Russian Ruble               | ₽      | RUB  |
| Rwandan Franc               | RF     | RWF  |
| Salvadoran Colón            | $      | SVC  |
| Samoan Tala                 | T      | WST  |
| Saudi Riyal                 | ﷼      | SAR  |
| Serbian Dinar               | дин.   | RSD  |
| Seychellois Rupee           | ₨      | SCR  |
| Singapore Dollar            | $      | SGD  |
| Solomon Islands Dollar      | $      | SBD  |
| Somali Shilling             | Sh     | SOS  |
| South African Rand          | R      | ZAR  |
| South Korean Won            | ₩      | KRW  |
| South Sudanese Pound        | £      | SSP  |
| Sri Lankan Rupee            | ₨      | LKR  |
| Sudanese Pound              | £      | SDG  |
| Surinamese Dollar           | $      | SRD  |
| Swedish Krona               | kr     | SEK  |
| Swiss Franc                 | CHF    | CHF  |
| Syrian Pound                | £S     | SYP  |
| São Tomé and Príncipe Dobra | Db     | STN  |
| Tajikistani Somoni          | ЅM     | TJS  |
| Tanzanian Shilling          | Sh     | TZS  |
| Thai Baht                   | ฿      | THB  |
| Timorese Dollar             | $      | USD  |
| Tongan Paʻanga              | T$     | TOP  |
| Trinidad and Tobago Dollar  | $      | TTD  |
| Tunisian Dinar              | د.ت    | TND  |
| Turkish Lira                | ₺      | TRY  |
| Turkmenistani Manat         | T      | TMT  |
| Ugandan Shilling            | Sh     | UGX  |
| Ukrainian Hryvnia           | ₴      | UAH  |
| United Arab Emirates Dirham | د.إ    | AED  |
| United States Dollar        | $      | USD  |
| Uruguayan Peso              | $U     | UYU  |
| Uzbekistani Som             | сўм    | UZS  |
| Vanuatu Vatu                | VT     | VUV  |
| Venezuelan Bolívar Soberano | Bs.S.  | VES  |
| Vietnamese Dong             | ₫      | VND  |
| Yemeni Rial                 | ﷼      | YER  |
| Zambian Kwacha              | ZK     | ZMW  |
| Zimbabwean Dollar           | $      | ZWL  |

