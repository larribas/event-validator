package repositories

import (
    "github.com/sp-lorenzo-arribas/event_validator/domain"
    "testing"
    "github.com/sp-lorenzo-arribas/event_validator/test"
)

func TestInMemoryRepository(t *testing.T) {
    domain.Current.NewRepository = NewInMemoryRepository
    test.GenericRepositoryTest(t, func(){}, func(){})
}