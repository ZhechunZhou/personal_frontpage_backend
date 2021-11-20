package main

import (
	"os"
)

//Config ...
type Config struct {
	DBHost   string
	DBName   string
	User     string
	Password string

	CloakClientRealm    string
	CloakUrl            string
	CloakUser           string
	CloakPassword       string
	CloakAdminCliId     string
	ClockAdminCliSecret string

	AWSAccessKeyID     string
	AWSSecretAccessKey string
	AWSRegion          string
	AWSBucket          string
}

//Value ...
func Value() *Config {
	return &Config{
		//database
		DBHost:   getEnv("DB_HOST", ""),
		DBName:   getEnv("DB_NAME", ""),
		User:     getEnv("DB_USER", ""),
		Password: getEnv("DB_PASSWORD", ""),

		//S3 bucket
		AWSAccessKeyID:     getEnv("AWS_ACCESS_KEY_ID", ""),
		AWSSecretAccessKey: getEnv("SECRET_ACCESS_KEY", ""),
		AWSRegion:          getEnv("AWS_REGION", ""),
		AWSBucket:          getEnv("AWS_BUCKET", ""),
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
