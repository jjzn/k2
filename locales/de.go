package main

var localeStrings = map[string]string{
	"removeItemConfirm": "Bist du sicher, dass du diesen Eintrag löschen möchtest?",
	"removeItem": "Eintrag löschen",
	"editItem": "Eintrag bearbeiten",
	"invite": "Personen einladen",

	"gridViewTitle": "Kalenderansicht",
	"listViewTitle": "Listenansicht",

	"inviteViewTitle": "Einladung zu",
	"inviteMessage": "Du wurdest zu",
	"invitePersonsJoiner": "eingeladen, mit",
	"inviteInstructions": "Wenn du teilnehmen möchtest, füge deinen Namen hinzu, und der Eintrag wird automatisch geändert.",

	"labelTitle": "Titel",
	"labelPersons": "Teilnehmende (durch Kommata getrennt)",
	"labelDate": "Datum",
	"labelTime": "Uhrzeit (freiwillig)",
	"labelEndDate": "Enddatum (freiwillig)",
	"labelEndTime": "Enduhrzeit (freiwillig)",
	"labelAddDescription": "Beschreibung hinzufügen",
	"labelDescription": "Beschreibung (freiwillig)",
	"labelName": "Name",
	"labelCreate": "Erstellen",
	"labelJoin": "Teilnehmen",
	"labelNew": "Neu",
	"labelUpdate": "Ändern",

	"dateToday": "Heute",
	"dateTomorrow": "Morgen",
	"dateThisWeek": "Diese Woche",
	"dateNextWeek": "Nächste Woche",
	"dateThisMonth": "Dieser Monat",
	"dateNextMonth": "Nächster Monat",

	"dayMon": "Mo",
	"dayTue": "Di",
	"dayWed": "Mw",
	"dayThu": "Do",
	"dayFri": "Fr",
	"daySat": "Sa",
	"daySun": "So",

	"messageNoItems": "Keine Einträge",
	"wordItem": "Eintrag",
}

var monthNames = []string{
	"Januar", "Februar", "März", "April", "May", "Juni",
	"Juli", "August", "September", "October", "November", "Dezember",
}

var dayNames = []string{
	"Montag", "Dienstag", "Mitwoch", "Donnerstag", "Freitag", "Samstag", "Sonntag",
}

var localeFormats = map[string]string{
	"daysLeft": " (noch %d Tage)",
	"daysPast": " (vor %d Tagen)",
	"today": " (heute)",
	"untilDate": " — bis zum ",
	"untilTime": " — bis ",
	"at": " um ",
}

func getLocaleString(key string) string {
	return localeStrings[key]
}
