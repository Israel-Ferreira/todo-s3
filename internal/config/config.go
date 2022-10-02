package config

import (
	"fmt"
	"os"

	dotenv "github.com/joho/godotenv"
)

type AwsConfig struct {
	AwsAccessKey  string
	AwsSecretKey  string
	AwsBucketName string
}

type DbConfig struct {
	DbHost     string
	DbName     string
	DbPort     string
	DbUsername string
	DbPass     string
}

type Config struct {
	Port string
	AwsConfig
	DbConfig
}

var ConfigVars Config

func LoadEnvVars() {
	dotenv.Load()

	portVar := os.Getenv("PORT")

	if portVar == "" {
		ConfigVars.Port = fmt.Sprintf(":%d", 9000)
	} else {
		ConfigVars.Port = fmt.Sprintf(":%s", portVar)
	}

	ConfigVars.AwsBucketName = os.Getenv("AWS_BUCKET_NAME")
	ConfigVars.AwsAccessKey = os.Getenv("AWS_ACCESS_KEY")
	ConfigVars.AwsSecretKey = os.Getenv("AWS_SECRET_KEY")

	ConfigVars.DbName = os.Getenv("DB_NAME")
	ConfigVars.DbHost = os.Getenv("DB_HOST")
	ConfigVars.DbUsername = os.Getenv("DB_USER")
	ConfigVars.DbPass = os.Getenv("DB_PASS")

}
