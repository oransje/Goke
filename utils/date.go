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
