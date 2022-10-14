package utils

import (
	"fmt"
	"time"
)

func FormatDate() string {
	t := time.Now()

	date := fmt.Sprintf("%d-%01d", int(t.Month()), t.Day())

	return date

}
func TimeNowFormatted() time.Time {
	now := time.Now()
	now.Format(time.RFC822)

	return now
}
