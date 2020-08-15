package iso31661

var (
	names   = []string{"Afghanistan", "Åland Islands", "Albania", "Algeria", "American Samoa", "Andorra", "Angola", "Anguilla", "Antarctica", "Antigua and Barbuda", "Argentina", "Armenia", "Aruba", "Australia", "Austria", "Azerbaijan", "Bahamas", "Bahrain", "Bangladesh", "Barbados", "Belarus", "Belgium", "Belize", "Benin", "Bermuda", "Bhutan", "Bolivia (Plurinational State of)", "Bonaire, Sint Eustatius and Saba", "Bosnia and Herzegovina", "Botswana", "Bouvet Island", "Brazil", "British Indian Ocean Territory", "Brunei Darussalam", "Bulgaria", "Burkina Faso", "Burundi", "Cabo Verde", "Cambodia", "Cameroon", "Canada", "Cayman Islands", "Central African Republic", "Chad", "Chile", "China", "Christmas Island", "Cocos (Keeling) Islands", "Colombia", "Comoros", "Congo", "Congo, Democratic Republic of the", "Cook Islands", "Costa Rica", "Côte d'Ivoire", "Croatia", "Cuba", "Curaçao", "Cyprus", "Czechia", "Denmark", "Djibouti", "Dominica", "Dominican Republic", "Ecuador", "Egypt", "El Salvador", "Equatorial Guinea", "Eritrea", "Estonia", "Eswatini", "Ethiopia", "Falkland Islands (Malvinas)", "Faroe Islands", "Fiji", "Finland", "France", "French Guiana", "French Polynesia", "French Southern Territories", "Gabon", "Gambia", "Georgia", "Germany", "Ghana", "Gibraltar", "Greece", "Greenland", "Grenada", "Guadeloupe", "Guam", "Guatemala", "Guernsey", "Guinea", "Guinea-Bissau", "Guyana", "Haiti", "Heard Island and McDonald Islands", "Holy See", "Honduras", "Hong Kong", "Hungary", "Iceland", "India", "Indonesia", "Iran (Islamic Republic of)", "Iraq", "Ireland", "Isle of Man", "Israel", "Italy", "Jamaica", "Japan", "Jersey", "Jordan", "Kazakhstan", "Kenya", "Kiribati", "Korea (Democratic People's Republic of)", "Korea, Republic of", "Kuwait", "Kyrgyzstan", "Lao People's Democratic Republic", "Latvia", "Lebanon", "Lesotho", "Liberia", "Libya", "Liechtenstein", "Lithuania", "Luxembourg", "Macao", "Madagascar", "Malawi", "Malaysia", "Maldives", "Mali", "Malta", "Marshall Islands", "Martinique", "Mauritania", "Mauritius", "Mayotte", "Mexico", "Micronesia (Federated States of)", "Moldova, Republic of", "Monaco", "Mongolia", "Montenegro", "Montserrat", "Morocco", "Mozambique", "Myanmar", "Namibia", "Nauru", "Nepal", "Netherlands", "New Caledonia", "New Zealand", "Nicaragua", "Niger", "Nigeria", "Niue", "Norfolk Island", "North Macedonia", "Northern Mariana Islands", "Norway", "Oman", "Pakistan", "Palau", "Palestine, State of", "Panama", "Papua New Guinea", "Paraguay", "Peru", "Philippines", "Pitcairn", "Poland", "Portugal", "Puerto Rico", "Qatar", "Réunion", "Romania", "Russian Federation", "Rwanda", "Saint Barthélemy", "Saint Helena, Ascension and Tristan da Cunha", "Saint Kitts and Nevis", "Saint Lucia", "Saint Martin (French part)", "Saint Pierre and Miquelon", "Saint Vincent and the Grenadines", "Samoa", "San Marino", "Sao Tome and Principe", "Saudi Arabia", "Senegal", "Serbia", "Seychelles", "Sierra Leone", "Singapore", "Sint Maarten (Dutch part)", "Slovakia", "Slovenia", "Solomon Islands", "Somalia", "South Africa", "South Georgia and the South Sandwich Islands", "South Sudan", "Spain", "Sri Lanka", "Sudan", "Suriname", "Svalbard and Jan Mayen", "Sweden", "Switzerland", "Syrian Arab Republic", "Taiwan, Province of China", "Tajikistan", "Tanzania, United Republic of", "Thailand", "Timor-Leste", "Togo", "Tokelau", "Tonga", "Trinidad and Tobago", "Tunisia", "Turkey", "Turkmenistan", "Turks and Caicos Islands", "Tuvalu", "Uganda", "Ukraine", "United Arab Emirates", "United Kingdom of Great Britain and Northern Ireland", "United States of America", "United States Minor Outlying Islands", "Uruguay", "Uzbekistan", "Vanuatu", "Venezuela (Bolivarian Republic of)", "Viet Nam", "Virgin Islands (British)", "Virgin Islands (U.S.)", "Wallis and Futuna", "Western Sahara", "Yemen", "Zambia", "Zimbabwe"}
	alpha2  = map[string]int{"AD": 5, "AE": 233, "AF": 0, "AG": 9, "AI": 7, "AL": 2, "AM": 11, "AO": 6, "AQ": 8, "AR": 10, "AS": 4, "AT": 14, "AU": 13, "AW": 12, "AX": 1, "AZ": 15, "BA": 28, "BB": 19, "BD": 18, "BE": 21, "BF": 35, "BG": 34, "BH": 17, "BI": 36, "BJ": 23, "BL": 185, "BM": 24, "BN": 33, "BO": 26, "BQ": 27, "BR": 31, "BS": 16, "BT": 25, "BV": 30, "BW": 29, "BY": 20, "BZ": 22, "CA": 40, "CC": 47, "CD": 51, "CF": 42, "CG": 50, "CH": 215, "CI": 54, "CK": 52, "CL": 44, "CM": 39, "CN": 45, "CO": 48, "CR": 53, "CU": 56, "CV": 37, "CW": 57, "CX": 46, "CY": 58, "CZ": 59, "DE": 83, "DJ": 61, "DK": 60, "DM": 62, "DO": 63, "DZ": 3, "EC": 64, "EE": 69, "EG": 65, "EH": 245, "ER": 68, "ES": 209, "ET": 71, "FI": 75, "FJ": 74, "FK": 72, "FM": 144, "FO": 73, "FR": 76, "GA": 80, "GB": 234, "GD": 88, "GE": 82, "GF": 77, "GG": 92, "GH": 84, "GI": 85, "GL": 87, "GM": 81, "GN": 93, "GP": 89, "GQ": 67, "GR": 86, "GS": 207, "GT": 91, "GU": 90, "GW": 94, "GY": 95, "HK": 100, "HM": 97, "HN": 99, "HR": 55, "HT": 96, "HU": 101, "ID": 104, "IE": 107, "IL": 109, "IM": 108, "IN": 103, "IO": 32, "IQ": 106, "IR": 105, "IS": 102, "IT": 110, "JE": 113, "JM": 111, "JO": 114, "JP": 112, "KE": 116, "KG": 121, "KH": 38, "KI": 117, "KM": 49, "KN": 187, "KP": 118, "KR": 119, "KW": 120, "KY": 41, "KZ": 115, "LA": 122, "LB": 124, "LC": 188, "LI": 128, "LK": 210, "LR": 126, "LS": 125, "LT": 129, "LU": 130, "LV": 123, "LY": 127, "MA": 150, "MC": 146, "MD": 145, "ME": 148, "MF": 189, "MG": 132, "MH": 138, "MK": 164, "ML": 136, "MM": 152, "MN": 147, "MO": 131, "MP": 165, "MQ": 139, "MR": 140, "MS": 149, "MT": 137, "MU": 141, "MV": 135, "MW": 133, "MX": 143, "MY": 134, "MZ": 151, "NA": 153, "NC": 157, "NE": 160, "NF": 163, "NG": 161, "NI": 159, "NL": 156, "NO": 166, "NP": 155, "NR": 154, "NU": 162, "NZ": 158, "OM": 167, "PA": 171, "PE": 174, "PF": 78, "PG": 172, "PH": 175, "PK": 168, "PL": 177, "PM": 190, "PN": 176, "PR": 179, "PS": 170, "PT": 178, "PW": 169, "PY": 173, "QA": 180, "RE": 181, "RO": 182, "RS": 197, "RU": 183, "RW": 184, "SA": 195, "SB": 204, "SC": 198, "SD": 211, "SE": 214, "SG": 200, "SH": 186, "SI": 203, "SJ": 213, "SK": 202, "SL": 199, "SM": 193, "SN": 196, "SO": 205, "SR": 212, "SS": 208, "ST": 194, "SV": 66, "SX": 201, "SY": 216, "SZ": 70, "TC": 229, "TD": 43, "TF": 79, "TG": 222, "TH": 220, "TJ": 218, "TK": 223, "TL": 221, "TM": 228, "TN": 226, "TO": 224, "TR": 227, "TT": 225, "TV": 230, "TW": 217, "TZ": 219, "UA": 232, "UG": 231, "UM": 236, "US": 235, "UY": 237, "UZ": 238, "VA": 98, "VC": 191, "VE": 240, "VG": 242, "VI": 243, "VN": 241, "VU": 239, "WF": 244, "WS": 192, "YE": 246, "YT": 142, "ZA": 206, "ZM": 247, "ZW": 248}
	alpha3  = map[string]int{"ABW": 12, "AFG": 0, "AGO": 6, "AIA": 7, "ALA": 1, "ALB": 2, "AND": 5, "ARE": 233, "ARG": 10, "ARM": 11, "ASM": 4, "ATA": 8, "ATF": 79, "ATG": 9, "AUS": 13, "AUT": 14, "AZE": 15, "BDI": 36, "BEL": 21, "BEN": 23, "BES": 27, "BFA": 35, "BGD": 18, "BGR": 34, "BHR": 17, "BHS": 16, "BIH": 28, "BLM": 185, "BLR": 20, "BLZ": 22, "BMU": 24, "BOL": 26, "BRA": 31, "BRB": 19, "BRN": 33, "BTN": 25, "BVT": 30, "BWA": 29, "CAF": 42, "CAN": 40, "CCK": 47, "CHE": 215, "CHL": 44, "CHN": 45, "CIV": 54, "CMR": 39, "COD": 51, "COG": 50, "COK": 52, "COL": 48, "COM": 49, "CPV": 37, "CRI": 53, "CUB": 56, "CUW": 57, "CXR": 46, "CYM": 41, "CYP": 58, "CZE": 59, "DEU": 83, "DJI": 61, "DMA": 62, "DNK": 60, "DOM": 63, "DZA": 3, "ECU": 64, "EGY": 65, "ERI": 68, "ESH": 245, "ESP": 209, "EST": 69, "ETH": 71, "FIN": 75, "FJI": 74, "FLK": 72, "FRA": 76, "FRO": 73, "FSM": 144, "GAB": 80, "GBR": 234, "GEO": 82, "GGY": 92, "GHA": 84, "GIB": 85, "GIN": 93, "GLP": 89, "GMB": 81, "GNB": 94, "GNQ": 67, "GRC": 86, "GRD": 88, "GRL": 87, "GTM": 91, "GUF": 77, "GUM": 90, "GUY": 95, "HKG": 100, "HMD": 97, "HND": 99, "HRV": 55, "HTI": 96, "HUN": 101, "IDN": 104, "IMN": 108, "IND": 103, "IOT": 32, "IRL": 107, "IRN": 105, "IRQ": 106, "ISL": 102, "ISR": 109, "ITA": 110, "JAM": 111, "JEY": 113, "JOR": 114, "JPN": 112, "KAZ": 115, "KEN": 116, "KGZ": 121, "KHM": 38, "KIR": 117, "KNA": 187, "KOR": 119, "KWT": 120, "LAO": 122, "LBN": 124, "LBR": 126, "LBY": 127, "LCA": 188, "LIE": 128, "LKA": 210, "LSO": 125, "LTU": 129, "LUX": 130, "LVA": 123, "MAC": 131, "MAF": 189, "MAR": 150, "MCO": 146, "MDA": 145, "MDG": 132, "MDV": 135, "MEX": 143, "MHL": 138, "MKD": 164, "MLI": 136, "MLT": 137, "MMR": 152, "MNE": 148, "MNG": 147, "MNP": 165, "MOZ": 151, "MRT": 140, "MSR": 149, "MTQ": 139, "MUS": 141, "MWI": 133, "MYS": 134, "MYT": 142, "NAM": 153, "NCL": 157, "NER": 160, "NFK": 163, "NGA": 161, "NIC": 159, "NIU": 162, "NLD": 156, "NOR": 166, "NPL": 155, "NRU": 154, "NZL": 158, "OMN": 167, "PAK": 168, "PAN": 171, "PCN": 176, "PER": 174, "PHL": 175, "PLW": 169, "PNG": 172, "POL": 177, "PRI": 179, "PRK": 118, "PRT": 178, "PRY": 173, "PSE": 170, "PYF": 78, "QAT": 180, "REU": 181, "ROU": 182, "RUS": 183, "RWA": 184, "SAU": 195, "SDN": 211, "SEN": 196, "SGP": 200, "SGS": 207, "SHN": 186, "SJM": 213, "SLB": 204, "SLE": 199, "SLV": 66, "SMR": 193, "SOM": 205, "SPM": 190, "SRB": 197, "SSD": 208, "STP": 194, "SUR": 212, "SVK": 202, "SVN": 203, "SWE": 214, "SWZ": 70, "SXM": 201, "SYC": 198, "SYR": 216, "TCA": 229, "TCD": 43, "TGO": 222, "THA": 220, "TJK": 218, "TKL": 223, "TKM": 228, "TLS": 221, "TON": 224, "TTO": 225, "TUN": 226, "TUR": 227, "TUV": 230, "TWN": 217, "TZA": 219, "UGA": 231, "UKR": 232, "UMI": 236, "URY": 237, "USA": 235, "UZB": 238, "VAT": 98, "VCT": 191, "VEN": 240, "VGB": 242, "VIR": 243, "VNM": 241, "VUT": 239, "WLF": 244, "WSM": 192, "YEM": 246, "ZAF": 206, "ZMB": 247, "ZWE": 248}
	numeric = map[string]int{"004": 0, "008": 2, "010": 8, "012": 3, "016": 4, "020": 5, "024": 6, "028": 9, "031": 15, "032": 10, "036": 13, "040": 14, "044": 16, "048": 17, "050": 18, "051": 11, "052": 19, "056": 21, "060": 24, "064": 25, "068": 26, "070": 28, "072": 29, "074": 30, "076": 31, "084": 22, "086": 32, "090": 204, "092": 242, "096": 33, "100": 34, "104": 152, "108": 36, "112": 20, "116": 38, "120": 39, "124": 40, "132": 37, "136": 41, "140": 42, "144": 210, "148": 43, "152": 44, "156": 45, "158": 217, "162": 46, "166": 47, "170": 48, "174": 49, "175": 142, "178": 50, "180": 51, "184": 52, "188": 53, "191": 55, "192": 56, "196": 58, "203": 59, "204": 23, "208": 60, "212": 62, "214": 63, "218": 64, "222": 66, "226": 67, "231": 71, "232": 68, "233": 69, "234": 73, "238": 72, "239": 207, "242": 74, "246": 75, "248": 1, "250": 76, "254": 77, "258": 78, "260": 79, "262": 61, "266": 80, "268": 82, "270": 81, "275": 170, "276": 83, "288": 84, "292": 85, "296": 117, "300": 86, "304": 87, "308": 88, "312": 89, "316": 90, "320": 91, "324": 93, "328": 95, "332": 96, "334": 97, "336": 98, "340": 99, "344": 100, "348": 101, "352": 102, "356": 103, "360": 104, "364": 105, "368": 106, "372": 107, "376": 109, "380": 110, "384": 54, "388": 111, "392": 112, "398": 115, "400": 114, "404": 116, "408": 118, "410": 119, "414": 120, "417": 121, "418": 122, "422": 124, "426": 125, "428": 123, "430": 126, "434": 127, "438": 128, "440": 129, "442": 130, "446": 131, "450": 132, "454": 133, "458": 134, "462": 135, "466": 136, "470": 137, "474": 139, "478": 140, "480": 141, "484": 143, "492": 146, "496": 147, "498": 145, "499": 148, "500": 149, "504": 150, "508": 151, "512": 167, "516": 153, "520": 154, "524": 155, "528": 156, "531": 57, "533": 12, "534": 201, "535": 27, "540": 157, "548": 239, "554": 158, "558": 159, "562": 160, "566": 161, "570": 162, "574": 163, "578": 166, "580": 165, "581": 236, "583": 144, "584": 138, "585": 169, "586": 168, "591": 171, "598": 172, "600": 173, "604": 174, "608": 175, "612": 176, "616": 177, "620": 178, "624": 94, "626": 221, "630": 179, "634": 180, "638": 181, "642": 182, "643": 183, "646": 184, "652": 185, "654": 186, "659": 187, "660": 7, "662": 188, "663": 189, "666": 190, "670": 191, "674": 193, "678": 194, "682": 195, "686": 196, "688": 197, "690": 198, "694": 199, "702": 200, "703": 202, "704": 241, "705": 203, "706": 205, "710": 206, "716": 248, "724": 209, "728": 208, "729": 211, "732": 245, "740": 212, "744": 213, "748": 70, "752": 214, "756": 215, "760": 216, "762": 218, "764": 220, "768": 222, "772": 223, "776": 224, "780": 225, "784": 233, "788": 226, "792": 227, "795": 228, "796": 229, "798": 230, "800": 231, "804": 232, "807": 164, "818": 65, "826": 234, "831": 92, "832": 113, "833": 108, "834": 219, "840": 235, "850": 243, "854": 35, "858": 237, "860": 238, "862": 240, "876": 244, "882": 192, "887": 246, "894": 247}
)

const (
	ta2 = iota
	ta3
	tn
)

func indexName(t int, code string) (i int, ok bool) {
	switch t {
	case ta2:
		if len(code) != 2 {
			return -1, false
		}
		i, ok = alpha2[code]
		return
	case ta3:
		if len(code) != 3 {
			return -1, false
		}
		i, ok = alpha3[code]
		return
	case tn:
		if len(code) != 3 {
			return -1, false
		}
		i, ok = numeric[code]
		return
	}
	return -1, false
}

func typeIs(t int, code string) (ok bool) {
	_, ok = indexName(t, code)
	return
}

func typeName(t int, code string) string {
	if i, ok := indexName(t, code); ok {
		return names[i]
	}
	return ""
}

// IsAlpha2 ...
func IsAlpha2(code string) bool {
	return typeIs(ta2, code)
}

// IsAlpha3 ...
func IsAlpha3(code string) bool {
	return typeIs(ta3, code)
}

// IsNumeric ...
func IsNumeric(code string) bool {
	return typeIs(tn, code)
}

// Is ...
func Is(code string) (r bool) {
	if len(code) == 2 {
		r = typeIs(ta2, code)
	} else if len(code) == 3 {
		r = typeIs(ta3, code)
		if !r {
			r = typeIs(tn, code)
		}
	}
	return
}

// NameAlpha2 ...
func NameAlpha2(code string) string {
	return typeName(ta2, code)
}

// NameAlpha3 ...
func NameAlpha3(code string) string {
	return typeName(ta3, code)
}

// NameNumeric ...
func NameNumeric(code string) string {
	return typeName(tn, code)
}

// Name ...
func Name(code string) (r string) {
	if len(code) == 2 {
		r = typeName(ta2, code)
	} else if len(code) == 3 {
		r = typeName(ta3, code)
		if r == "" {
			r = typeName(tn, code)
		}
	}
	return
}
