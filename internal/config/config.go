package config

import "os"

// Config содержит конфигурационные параметры приложения
type Config struct {
	StandardToken string
	AdvToken      string
	StatToken     string
	LogLevel      string
	Port          string
}

// New создает и возвращает новый экземпляр Config, загружая параметры из переменных окружения
func New() *Config {
	return &Config{
		StandardToken: os.Getenv("WB_STANDARD_TOKEN"),
		AdvToken:      os.Getenv("WB_ADV_TOKEN"),
		StatToken:     os.Getenv("WB_STAT_TOKEN"),
		LogLevel:      getEnvOrDefault("LOG_LEVEL", "info"),
		Port:          getEnvOrDefault("PORT", "8080"),
	}
}

// getEnvOrDefault возвращает значение переменной окружения или значение по умолчанию
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
