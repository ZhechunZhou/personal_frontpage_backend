package common

import "github.com/google/uuid"

func GetUUID() string {
	return uuid.New().String()
}

func GenerateFileName(extension string) string {
	return GetUUID() + "." + extension
}
