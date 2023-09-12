package utils

import (
	"fmt"
	"time"
)

func FormatDate() string {
	t := time.Now()

	date := fmt.Sprintf("%d_%01d_%02d", int(t.Year()), t.Month(), t.Day())

	return date

}
