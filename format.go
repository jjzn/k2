package k2

import "time"
import "strings"

func formatTime(t time.Time, fmt string) string {
	s := t.Format(fmt)

	s = strings.Replace(s, "Monday", "Lunes", -1)
	s = strings.Replace(s, "Tuesday", "Martes", -1)
	s = strings.Replace(s, "Wednesday", "Miércoles", -1)
	s = strings.Replace(s, "Thursday", "Jueves", -1)
	s = strings.Replace(s, "Friday", "Viernes", -1)
	s = strings.Replace(s, "Saturday", "Sábado", -1)
	s = strings.Replace(s, "Sunday", "Domingo", -1)

	return s
}
