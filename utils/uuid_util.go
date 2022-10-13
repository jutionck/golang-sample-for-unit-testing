package utils

import "github.com/google/uuid"

func UuidGenerate() string {
	return uuid.New().String()
}
