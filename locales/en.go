package main

var localeStrings = map[string]string{
	"removeItemConfirm": "Are you sure you want to remove this item forever?",
	"removeItem": "Remove item",
	"editItem": "Edit item",
	"invite": "Invite people",

	"gridViewTitle": "Grid view",
	"listViewTitle": "List view",

	"inviteViewTitle": "Invitation for",
	"inviteMessage": "You have been invited to",
	"invitePersonsJoiner": "with",
	"inviteInstructions": "If you want to participate, add your name here and the item will update automatically.",

	"labelTitle": "Title",
	"labelPersons": "Persons (comma separated)",
	"labelDate": "Date",
	"labelTime": "Time (optional)",
	"labelEndDate": "End date (optional)",
	"labelEndTime": "End time (optional)",
	"labelAddDescription": "Add description",
	"labelDescription": "Description (optional)",
	"labelName": "Name",
	"labelCreate": "Create",
	"labelJoin": "Join",
	"labelNew": "New",
	"labelUpdate": "Update",

	"dateToday": "Today",
	"dateTomorrow": "Tomorrow",
	"dateThisWeek": "This week",
	"dateNextWeek": "Next week",
	"dateThisMonth": "This month",
	"dateNextMonth": "Next month",

	"dayMon": "Mon",
	"dayTue": "Tue",
	"dayWed": "Wed",
	"dayThu": "Thu",
	"dayFri": "Fri",
	"daySat": "Sat",
	"daySun": "Sun",

	"messageNoItems": "No items",
	"wordItem": "item",
}

var monthNames = []string{
	"January", "February", "March", "April", "May", "June",
	"Julio", "August", "September", "October", "November", "December",
}

var dayNames = []string{
	"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday",
}

var localeFormats = map[string]string{
	"daysLeft": " (%d days left)",
	"daysPast": " (%d days ago)",
	"today": " (today)",
	"untilDate": " — until the ",
	"untilTime": " — until ",
	"at": " at ",
}

func getLocaleString(key string) string {
	return localeStrings[key]
}
