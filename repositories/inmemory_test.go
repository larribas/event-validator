package repositories

import (
	"github.com/sp-lorenzo-arribas/event_validator/domain"
	"github.com/sp-lorenzo-arribas/event_validator/test"
	"testing"
)

func TestInMemoryRepository(t *testing.T) {
	domain.Current.NewRepository = NewInMemoryRepository
	test.GenericRepositoryTest(t, func() {}, func() {})
}
