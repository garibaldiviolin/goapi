package config

import (
	"github.com/joho/godotenv"
	"os"
)
 
type Config struct {
    DB *DBConfig
}
 
type DBConfig struct {
    Dialect  string
    Username string
    Password string
    Name     string
    Charset  string
}
 
func GetConfig() *Config {
    godotenv.Load()

    return &Config{
        DB: &DBConfig{
            Dialect:  "postgres",
            Username: os.Getenv("DATABASE_USERNAME"),
            Password: os.Getenv("DATABASE_PASSWORD"),
            Name:     os.Getenv("DATABASE_NAME"),
            Charset:  "utf8",
        },
    }
}
