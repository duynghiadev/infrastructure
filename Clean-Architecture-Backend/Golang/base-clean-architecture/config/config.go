package config

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

// Configuration contains static info required to run the apps
// It contains DB info
type Configuration struct {
	Port                  string `env:"PORT" envDefault:"8081"`
	HashSalt              string `env:"HASH_SALT,required"`
	SigningKey            string `env:"SIGNING_KEY,required"`
	TokenTTL              int64  `env:"TOKEN_TTL,required"`
	JwtSecret             string `env:"JWT_SECRET,required"`
	DatabaseConnectionURL string `env:"CONNECTION_URL,required"`
}

// NewConfig will read the config data from given .env file
func NewConfig(files ...string) *Configuration {
	envFile := getEnvFilePath(files...)
	if err := loadEnvFile(envFile); err != nil {
		log.Printf("Warning: %v\n", err)
	}

	cfg := &Configuration{}
	if err := env.Parse(cfg); err != nil {
		log.Printf("Error when parse environment: %+v\n", err)
	}

	return cfg
}

// getEnvFilePath returns the path to the .env file
func getEnvFilePath(files ...string) []string {
	if len(files) > 0 {
		return files
	}

	_, b, _, _ := runtime.Caller(0)
	projectRoot := filepath.Join(filepath.Dir(b), "..")
	return []string{filepath.Join(projectRoot, ".env")}
}

// loadEnvFile loads the environment variables from the .env file
func loadEnvFile(files []string) error {
	err := godotenv.Load(files...)
	if err != nil {
		return fmt.Errorf("Can't read file .env at path: %v", files)
	}
	return nil
}
