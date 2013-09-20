package relative

import (
	"fmt"
	"time"
)

const (
	Day = 24 * time.Hour
	Week = 7 * Day
	Month = 30 * Day
	Year = 365 * Day
)

func RelativeDate(t time.Time) string {
	switch d := time.Since(t); {
	case d < time.Second:
		return "Just now"
	case d < time.Minute:
		return fmt.Sprintf("%d seconds ago", (d / time.Second))
	case d < time.Hour:
		return fmt.Sprintf("%d minutes ago", (d / time.Minute))
	case d < Day:
		return fmt.Sprintf("%d hours ago", (d / time.Hour))
	case d < Week:
		return fmt.Sprintf("%d days ago", (d / Day))
	case d < Month:
		return fmt.Sprintf("%d weeks ago", (d / Week))
	case d < Year:
		return fmt.Sprintf("%d months ago", (d / Month))
	default:
		return t.Format("Jan 2, 2006")
	}
}
