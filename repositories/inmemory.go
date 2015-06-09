package repositories

import (
    "github.com/sp-lorenzo-arribas/event_validator/domain"
)

type InMemoryRepository struct {
    validators map[string][]*domain.Validator
}

func NewInMemoryValidatorRepository() domain.Repository {
    return InMemoryRepository{
        validators: make(map[string][]*domain.Validator),
    }
}

func (r *InMemoryRepository) Create(validator *domain.Validator) (version int) {
    versions, ok := r.validators[validator.Type]
    if !ok {
        r.validators[validator.Type] = make([]*domain.Validator, 0)
    }

    validator.Version = len(r.validators[validator.Type])
    versions = append(versions, validator)

    return validator.Version
}

func (r *InMemoryRepository) GetNextVersion(_type string) int {
    versions, _ := r.validators[_type]
    return len(versions)
}

func (r *InMemoryRepository) Inspect(_type string, version int) (*domain.Validator, error) {

    // TODO Try removing this bit and seeing if it works
    if _, ok := r.validators[_type]; !ok {
        return nil, domain.ErrValidatorDoesNotExist{_type, version}
    }

    if version < 0 || version >= len(r.validators[_type]) {
        return nil, domain.ErrValidatorDoesNotExist{_type, version}
    }

    return r.validators[_type][version], nil
}


