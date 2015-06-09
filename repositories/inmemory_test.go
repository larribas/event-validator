package repositories

import (
    "github.com/sp-lorenzo-arribas/event_validator/domain"
    "testing"
    "github.com/sp-lorenzo-arribas/event_validator/test"
)

func TestMemStreamRepository(t *testing.T) {
    domain.Current.NewRepository = NewInMemoryValidatorRepository
    test.GenericRepositoryTest(t, setUp, tearDown)
}

func setUp() {}

func tearDown() {}