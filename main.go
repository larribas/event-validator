package main // import "github.com/sp-lorenzo-arribas/event_validator"

import (
    "github.com/sp-lorenzo-arribas/event_validator/domain"
    "github.com/sp-lorenzo-arribas/event_validator/formats"
    "github.com/sp-lorenzo-arribas/event_validator/repositories"
    "os"
    "github.com/sp-lorenzo-arribas/event_validator/api"
)

// SetCurrentEnvironment defines the topology of the current domain. That is, the particular implementations for each of the defined interfaces
func SetCurrentEnvironment() {
    domain.Current = &domain.Environment{
        NewRepository: func() domain.Repository {
            return repositories.NewRedisRepository(os.Getenv("EV_REDIS_HOST"), os.Getenv("EV_REDIS_PORT"))
        },
        NewFormatChecker: func() domain.FormatChecker {
            return &formats.JSONSchemaFormatChecker{}
        },
    }
}

func main() {
    SetCurrentEnvironment()
    api.New(os.Getenv("EV_API_ADDRESS"), os.Getenv("EV_API_PORT")).Expose()
}




