package main

var localStrings = map[string]string{
	"removeItemConfirm": "Segur que vols eliminar permanentement aquesta entrada?",
	"removeItem": "Eliminar entrada",
	"editItem": "Editar entrada",
	"invite": "Convidar persones",

	"gridViewTitle": "Vista d'entrades",

	"inviteViewTitle": "Invitació a",
	"inviteMessage": "T'han convidat a",
	"invitePersonsJoiner": "amb",
	"inviteInstructions": "Si hi vols participar, afegeix aquí el teu nom i l'entrada s'actualitzarà automàticament.",

	"labelTitle": "Títol",
	"labelPersons": "Participants (separats per comes)",
	"labelDate": "Data",
	"labelTime": "Hora (opcional)",
	"labelEndDate": "Data final (opcional)",
	"labelEndTime": "Hora final (opcional)",
	"labelAddDescription": "Afegir descripció",
	"labelDescription": "Descripció (opcional)",
	"labelName": "Nom",
	"labelCreate": "Crear",
	"labelJoin": "Unir-se",
	"labelNew": "Nova",
	"labelUpdate": "Actualitzar",

	"dateToday": "Avui",
	"dateTomorrow": "Demà",
	"dateThisWeek": "Aquesta setmana",
	"dateNextWeek": "Setmana següent",
	"dateThisMonth": "Aquest mes",
	"dateNextMonth": "Mes següent",

	"messageNoItems": "No hi ha entrades",
	"wordItem": "entrada",
}

var monthNames = []string{
	"Gener", "Febrer", "Març", "Abril", "Maig", "Juny",
	"Juliol", "Agost", "Setembre", "Octubre", "Novembre", "Desembre",
}

var dayNames = []string{
	"Dilluns", "Dimarts", "Dimecres", "Dijous", "Divendres", "Dissabte", "Diumenge",
}

var localFormats = map[string]string{
	"daysLeft": " (queden %d dies)",
	"daysPast": " (fa %d dies)",
	"today": " (avui)",
	"untilDate": " — fins el dia ",
	"untilTime": " — fins les ",
	"at": " a les ",
}

func getLocalString(key string) string {
	return localStrings[key]
}
