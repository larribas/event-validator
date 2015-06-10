package test

import (
	"bytes"
	"github.com/sp-lorenzo-arribas/event_validator/domain"
	"testing"
)

// GenericRepositoryTest encapsulates a series of tests that validate that a particular Repository implementation
// complies with the desired functionality for such interface. Prior to using it, the domain.Current
// environment is supposed to reflect the particular implementation under test
func GenericRepositoryTest(t *testing.T, setUp, tearDown func()) {

	do(setUp, testCreateAndInspect, t, tearDown)
	do(setUp, testGetNextVersion, t, tearDown)
	do(setUp, testInspectNonexistentValidator, t, tearDown)
}

func do(setUp func(), testCase func(*testing.T), t *testing.T, tearDown func()) {
	setUp()
	testCase(t)
	tearDown()
}

func testCreateAndInspect(t *testing.T) {
	repo := domain.Current.NewRepository()

	testCases := []struct {
		Type            string
		Rules           []byte
		ExpectedVersion int
	}{
		{"someEventType", []byte("first version rules"), 0},
		{"someEventType", []byte("second version rules"), 1},
	}

	for _, testCase := range testCases {
		validator := mustInstantiateValidator(t, testCase.Type, testCase.Rules)

		version := repo.Create(validator)
		if version != testCase.ExpectedVersion {
			t.Errorf("Expected Repository::Create to return version %d. Instead it returned %d", testCase.ExpectedVersion, version)
		}

		if validator.Version != testCase.ExpectedVersion {
			t.Errorf("Expected Repository::Create to update the supplied validator's version to %d. Instead it contains the value %d", testCase.ExpectedVersion, validator.Version)
		}

		// Test that we can retrieve the stored validator afterwards
		retrievedValidator, err := repo.Inspect(testCase.Type, version)
		if err != nil {
			t.Errorf("Expected Repository::Inspect not to return an error for a validator we just created. Instead it returned %s", err)
		}

		if bytes.Compare(retrievedValidator.Rules, testCase.Rules) != 0 {
			t.Error("Expected Repository::Inspect to return a validator with the same rules than the one we just created")
		}
	}
}

func testGetNextVersion(t *testing.T) {
	repo := domain.Current.NewRepository()

	repo.Create(mustInstantiateValidator(t, "someEventType", []byte("first version rules")))

	testCases := []struct {
		Type            string
		ExpectedVersion int
	}{
		{"someEventType", 1},
		{"nonexistentType", 0},
	}

	for _, testCase := range testCases {
		nextVersion := repo.GetNextVersion(testCase.Type)
		if nextVersion != testCase.ExpectedVersion {
			t.Errorf("Expected Repository::GetNextVersion to return %d for type '%s'. Instead, it returned %d", testCase.ExpectedVersion, testCase.Type, nextVersion)
		}
	}
}

func testInspectNonexistentValidator(t *testing.T) {
	repo := domain.Current.NewRepository()

	_, err := repo.Inspect("nonexistentType", 0)
	if err == nil {
		t.Errorf("Expected Repository::Inspect to return an error for a nonexistent validator")
	}
}

func mustInstantiateValidator(t *testing.T, _type string, rules []byte) *domain.Validator {
	validator, err := domain.NewValidator(_type, rules)
	if err != nil {
		t.Error("Expected NewValidator not to return an error (check which Current.GetValidator instance is being executed!)")
	}

	return validator
}
