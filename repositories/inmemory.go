package repositories

import (
	"github.com/sp-lorenzo-arribas/event_validator/domain"
)

// InMemoryRepository is a lightweight implementation of a repository that stores validators in the current
// application's heap. Thus, they are not persisted between sessions, which makes it the perfect implementation
// to test the application's functionality.
type InMemoryRepository struct {
	validators map[string][]*domain.Validator
}

// NewInMemoryRepository instantiates a new InMemoryRepository and exposes it as a domain.Repository interface
func NewInMemoryRepository() domain.Repository {
	return &InMemoryRepository{
		validators: make(map[string][]*domain.Validator),
	}
}

func (r *InMemoryRepository) Create(validator *domain.Validator) (version int) {
	if _, ok := r.validators[validator.Type]; !ok {
		r.validators[validator.Type] = make([]*domain.Validator, 0)
	}

	validator.Version = len(r.validators[validator.Type])
	r.validators[validator.Type] = append(r.validators[validator.Type], validator)

	return validator.Version
}

func (r *InMemoryRepository) GetNextVersion(_type string) int {
	versions, _ := r.validators[_type]
	return len(versions)
}

func (r *InMemoryRepository) Inspect(_type string, version int) (*domain.Validator, error) {
	if version < 0 || version >= len(r.validators[_type]) {
		return nil, domain.ErrValidatorDoesNotExist{_type, version}
	}

	return r.validators[_type][version], nil
}
