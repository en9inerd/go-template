package config

import (
	"github.com/en9inerd/go-pkgs/flagpair"
	// "strconv" // Uncomment when using getEnvInt
)

// Config holds application configuration
type Config struct {
	Port string
	// Add your application-specific config fields here
	// Example:
	// DatabaseURL string
	// APIKey      string
	// Timeout     time.Duration

	// Runtime
	Verbose bool
}

// ParseConfig parses command-line flags and environment variables
func ParseConfig(args []string, getenv func(string) string) (*Config, error) {
	getEnv := func(key, fallback string) string {
		if v := getenv(key); v != "" {
			return v
		}
		return fallback
	}

	// getEnvInt parses an integer from an environment variable with a fallback.
	// Uncomment when adding integer config fields.
	// getEnvInt := func(key string, fallback int) int {
	// 	if v := getenv(key); v != "" {
	// 		if i, err := strconv.Atoi(v); err == nil {
	// 			return i
	// 		}
	// 	}
	// 	return fallback
	// }

	r := flagpair.New("app")

	port := r.String("port", "p", getEnv("APP_PORT", "8000"), "Port to listen on")
	// Add your application-specific flags here (long, short, default, usage)
	// Example:
	// databaseURL := r.String("database-url", "", getEnv("DATABASE_URL", ""), "Database connection URL")
	// apiKey := r.String("api-key", "", getEnv("API_KEY", ""), "API key")

	// Runtime
	verbose := r.Bool("verbose", "v", false, "Enable verbose logging")

	if err := r.Parse(args[1:]); err != nil {
		return nil, err
	}

	return &Config{
		Port:    *port,
		Verbose: *verbose,
		// Add your application-specific config assignments here
		// Example:
		// DatabaseURL: *databaseURL,
		// APIKey:      *apiKey,
	}, nil
}
