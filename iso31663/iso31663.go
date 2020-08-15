package iso31663

var (
	names = []string{"British Antarctic Territory", "Burma", "Byelorussian Soviet Socialist Republic", "Canton and Enderbury Islands", "Czechoslovakia", "East Germany", "East Timor", "French Southern and Antarctic Lands", "French Territory of the Afars and the Issas", "Gilbert and Ellice Islands", "Johnston Atoll", "Metropolitan France", "Midway Atoll", "Netherlands Antilles", "New Hebrides", "North Vietnam", "Panama Canal Zone", "Queen Maud Land", "Republic of Dahomey", "Republic of Upper Volta", "Rhodesia", "Saudi-Iraqi neutral zone", "Serbia and Montenegro", "Sikkim", "South Yemen", "Soviet Union", "Trust Territory of the Pacific Islands", "United States Miscellaneous Pacific Islands", "Wake Island", "Yugoslavia", "Zaire"}
	codes = map[string]int{"AIDJ": 8, "ANHH": 13, "BQAQ": 0, "BUMM": 1, "BYAA": 2, "CSHH": 4, "CSXX": 22, "CTKI": 3, "DDDE": 5, "DYBJ": 18, "FQHH": 7, "FXFR": 11, "GEHH": 9, "HVBF": 19, "JTUM": 10, "MIUM": 12, "NHVU": 14, "NQAQ": 17, "NTHH": 21, "PCHH": 26, "PUUM": 27, "PZPA": 16, "RHZW": 20, "SKIN": 23, "SUHH": 25, "TPTL": 6, "VDVN": 15, "WKUM": 28, "YDYE": 24, "YUCS": 29, "ZRCD": 30}
)

// Is ...
func Is(code string) (ok bool) {
	_, ok = codes[code]
	return
}

// Name ...
func Name(code string) string {
	if len(code) != 4 {
		return ""
	}
	if i, ok := codes[code]; ok {
		return names[i]
	}
	return ""
}
