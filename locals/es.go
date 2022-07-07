package main

var localStrings = map[string]string{
	"removeItemConfirm": "Seguro que quieres eliminar permanentemente esta entrada?",
	"removeItem": "Eliminar entrada",
	"editItem": "Editar entrada",
	"invite": "Invitar personas",

	"gridViewTitle": "Vista de entradas",
	"listViewTitle": "Lista de entradas",

	"inviteViewTitle": "Invitación a",
	"inviteMessage": "Te han invitado a",
	"invitePersonsJoiner": "con",
	"inviteInstructions": "Si quieres participar, añade aquí tu nombre y la entrada se actualizará automáticamente.",

	"labelTitle": "Título",
	"labelPersons": "Participantes (separados por comas)",
	"labelDate": "Fecha",
	"labelTime": "Hora (opcional)",
	"labelEndDate": "Fecha final (opcional)",
	"labelEndTime": "Hora final (opcional)",
	"labelAddDescription": "Añadir descripción",
	"labelDescription": "Descripción (opcional)",
	"labelName": "Nombre",
	"labelCreate": "Crear",
	"labelJoin": "Unirse",
	"labelNew": "Nueva",
	"labelUpdate": "Actualizar",

	"dateToday": "Hoy",
	"dateTomorrow": "Mañana",
	"dateThisWeek": "Esta semana",
	"dateNextWeek": "Semana siguiente",
	"dateThisMonth": "Este mes",
	"dateNextMonth": "Mes siguiente",

	"dayMon": "Lun",
	"dayTue": "Mar",
	"dayWed": "Mié",
	"dayThu": "Jue",
	"dayFri": "Vie",
	"daySat": "Sáb",
	"daySun": "Dom",

	"messageNoItems": "No hay entradas",
	"wordItem": "entrada",
}

var monthNames = []string{
	"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio",
	"Julio", "Agosto", "Septiembre", "Octubre", "Noviembre", "Diciembre",
}

var dayNames = []string{
	"Lunes", "Martes", "Miércoles", "Jueves", "Viernes", "Sábado", "Domingo",
}

var localFormats = map[string]string{
	"daysLeft": " (quedan %d días)",
	"daysPast": " (hace %d días)",
	"today": " (hoy)",
	"untilDate": " — hasta el día ",
	"untilTime": " — hasta las ",
	"at": " a las ",
}

func getLocalString(key string) string {
	return localStrings[key]
}
