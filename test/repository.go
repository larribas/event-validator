package test

import (
    "testing"
    "github.com/sp-lorenzo-arribas/event_validator/domain"
)

// GenericRepositoryTest encapsulates a series of tests that validate that a particular Repository implementation
// complies with the desired functionality for such interface. Prior to using it, the domain.Current
// environment is supposed to reflect the particular implementation under test
func GenericRepositoryTest(t *testing.T, setUp, tearDown func()) {

    do(setUp, testCreateAssignsNewVersion, t, tearDown)
    do(setUp, testGetNextVersion, t, tearDown)
    do(setUp, testGetNextVersionForEmptyType, t, tearDown)
    do(setUp, testInspect, t, tearDown)
    do(setUp, testInspectNonexistentValidator, t, tearDown)
}

func do(setUp func(), testCase func(*testing.T), t *testing.T, tearDown func()) {
    setUp()
    testCase(t)
    tearDown()
}

func testCreateAssignsNewVersion(t *testing.T) {
    repo := domain.Current.GetRepository()

    validator :=
    repo.Create()

}

func testGetNextVersion(t *testing.T) {
    repo := domain.Current.GetRepository()

}

func testGetNextVersionForEmptyType(t *testing.T) {
    repo := domain.Current.GetRepository()

}

func testInspect(t *testing.T) {
    repo := domain.Current.GetRepository()

}

func testInspectNonexistentValidator(t *testing.T) {
    repo := domain.Current.GetRepository()

}