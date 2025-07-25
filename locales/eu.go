package main

var localeStrings = map[string]string{
	"removeItemConfirm": "Ziur zaude betirako ezabatu nahi duzula sarrera?",
	"removeItem":        "Sarrera ezabatu",
	"editItem":          "Sarrera editatu",
	"invite":            "Pertsonak gonbidatu",

	"gridViewTitle": "Sarrera bista",
	"listViewTitle": "Sarrera zerrenda",

	"inviteViewTitle":     "Gonbidapena ",
	"inviteMessage":       "Gonbidatu zaituzte ",
	"invitePersonsJoiner": "pertsona hauekin",
	"inviteInstructions":  "Parte hartu nahi baduzu, gehitu zure izena hemen eta sarrera automatikoki eguneratuko da.",

	"labelTitle":          "Izenburua",
	"labelPersons":        "Parte-hartzaileak (komekin separatuta)",
	"labelLocation":       "Lekua",
	"labelDate":           "Data",
	"labelTime":           "Ordua (hautazkoa)",
	"labelEndDate":        "Amaiera-data (hautazkoa)",
	"labelEndTime":        "Amaiera-ordua (hautazkoa)",
	"labelAddDescription": "Deskribapena gehitu",
	"labelDescription":    "Deskribapena (hautazkoa)",
	"labelName":           "Izena",
	"labelCreate":         "Sortu",
	"labelJoin":           "Gehitu",
	"labelNew":            "Berria",
	"labelUpdate":         "Eguneratu",

	"dateToday":     "Gaur",
	"dateTomorrow":  "Bihar",
	"dateThisWeek":  "Aste honetan",
	"dateNextWeek":  "Hurrengo astean",
	"dateThisMonth": "Hilabete honetan",
	"dateNextMonth": "Hurrengo hilabetean",

	"dayMon": "Asl",
	"dayTue": "Asa",
	"dayWed": "Asz",
	"dayThu": "Ose",
	"dayFri": "Ost",
	"daySat": "Lar",
	"daySun": "Iga",

	"messageNoItems": "Ez daude sarrerarik",
	"wordItem":       "sarrera",
}

var monthNames = []string{
	"Urtarrila", "Otsaila", "Martxoa", "Apirila", "Maiatza", "Ekaina",
	"Uztaila", "Abuztua", "Iraila", "Urria", "Azaroa", "Abendua",
}

var dayNames = []string{
	"Astelehena", "Asteartea", "Asteazkena", "Osteguna", "Ostirala", "Larunbata", "Igandea",
}

var localeFormats = map[string]string{
	"daysLeft":  " (%d egun geratzen dira)",
	"daysPast":  " (duela %d egun)",
	"today":     " (gaur)",
	"untilDate": " — egunerarte ",
	"untilTime": " — ordurarte ",
	"at":        " orduan ",
}

func getLocaleString(key string) string {
	return localeStrings[key]
}
