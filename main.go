package main // import "github.com/sp-lorenzo-arribas/event_validator"

import (
	"github.com/sp-lorenzo-arribas/event_validator/api"
	"github.com/sp-lorenzo-arribas/event_validator/domain"
	"github.com/sp-lorenzo-arribas/event_validator/formats"
	"github.com/sp-lorenzo-arribas/event_validator/repositories"
	"os"
)

// SetCurrentEnvironment defines the topology of the current domain. That is, the particular implementations for each of the defined interfaces
func SetCurrentEnvironment() {
	domain.Current = &domain.Environment{
		NewRepository: func() domain.Repository {
			return repositories.NewRedisRepository(EnvOrDefault("EV_REDIS_HOST", "localhost"), EnvOrDefault("EV_REDIS_PORT", "6379"))
		},
		NewFormatChecker: func() domain.FormatChecker {
			return &formats.JSONSchemaFormatChecker{}
		},
	}
}

// EnvOrDefault obtains a value from the OS environment, or a default value if such a key was not specified or is empty
func EnvOrDefault(key, _default string) string {
	val := os.Getenv(key)
	if val == "" {
		val = _default
	}

	return val
}

func main() {
	SetCurrentEnvironment()
	api.New(EnvOrDefault("EV_API_ADDRESS", "localhost"), EnvOrDefault("EV_API_PORT", "8190")).Expose()
}
