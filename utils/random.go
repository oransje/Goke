package utils

import (
	"math/rand"
	"time"
)

const randomize = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789"

func RandomMigrationName() string {
	rand.Seed(time.Now().UnixNano())
	migration_name := make([]byte, 5)

	for i := range migration_name {
		migration_name[i] = randomize[rand.Intn(len(randomize))]
	}

	return string(migration_name)
}
