package backend

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv  string
	AppPort string
	AppName string
	JWTKey  string

	DbHost     string
	DbPort     string
	DbName     string
	DbUser     string
	DbPassword string
}

func LoadConfig() Config {
	loadEnvFile()

	return Config{
		AppEnv:  os.Getenv("APP_ENV"),
		AppPort: os.Getenv("APP_PORT"),
		AppName: os.Getenv("APP_NAME"),
		JWTKey:  os.Getenv("JWT_KEY"),

		DbHost:     os.Getenv("DB_HOST"),
		DbPort:     os.Getenv("DB_PORT"),
		DbName:     os.Getenv("DB_NAME"),
		DbUser:     os.Getenv("DB_USER"),
		DbPassword: os.Getenv("DB_PASSWORD"),
	}
}

func loadEnvFile() {
	dir, err := os.Getwd()
	if err != nil {
		log.Println("Warning: cannot get working directory, using system environment variables")
		return
	}

	for {
		envPath := filepath.Join(dir, ".env")
		if _, statErr := os.Stat(envPath); statErr == nil {
			if loadErr := godotenv.Load(envPath); loadErr != nil {
				log.Printf("Warning: failed to load %s: %v", envPath, loadErr)
			}
			return
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			log.Println("Warning: .env file not found, using system environment variables")
			return
		}
		dir = parent
	}
}
