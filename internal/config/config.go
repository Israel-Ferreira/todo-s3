package config

import (
	"fmt"
	"log"
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

var ConfigVars *Config

func LoadEnvVars() {
	if err := dotenv.Load(); err != nil {
		log.Println(err.Error())
	}

	portVar := os.Getenv("PORT")

	if portVar == "" {
		portVar = fmt.Sprintf(":%d", 9000)
	} else {
		portVar = fmt.Sprintf(":%s", portVar)
	}

	awsConfig := &AwsConfig{}

	awsConfig.AwsBucketName = os.Getenv("AWS_BUCKET_NAME")
	awsConfig.AwsAccessKey = os.Getenv("AWS_ACCESS_KEY")
	awsConfig.AwsSecretKey = os.Getenv("AWS_SECRET_KEY")

	dbConfig := DbConfig{}

	dbConfig.DbName = os.Getenv("DB_NAME")
	dbConfig.DbHost = os.Getenv("DB_HOST")
	dbConfig.DbPort = os.Getenv("DB_PORT")
	dbConfig.DbUsername = os.Getenv("DB_USER")
	dbConfig.DbPass = os.Getenv("DB_PASS")

	ConfigVars = &Config{
		Port:      portVar,
		AwsConfig: *awsConfig,
		DbConfig:  dbConfig,
	}

}

func GetDbConfig() DbConfig {
	return ConfigVars.DbConfig
}
