package helpers

import "time"

func FormatTanggal(t time.Time) string {
	return t.Format("02-01-2006")
}
