package domain

// The current execution environment, from which all components reclaim particular implementations for domain
// interfaces. By default, the current environment is set to ground (nil) implementations of all the domain interfaces.
// This initialization is done for the sake of testability
var Current = &Environment{
	NewRepository:    func() Repository { return &GroundRepository{} },
	NewFormatChecker: func() FormatChecker { return &GroundFormatChecker{} },
}

// Environment holds the particular parameters and interface implementations for an Event Validator instance.
// Its purpose is to centralize the definition of the application's topology, and to enable different configurations
// to be switched at runtime (for testing purposes) or be loaded from a file (for the sake of portability)
type Environment struct {
	repositoryInstance Repository
	checkerInstance    FormatChecker

	NewRepository    func() Repository
	NewFormatChecker func() FormatChecker
}

// GetRepository returns a (lazily instantiated) singleton repository instance
func (env *Environment) GetRepository() Repository {
	if env.repositoryInstance == nil {
		env.repositoryInstance = env.NewRepository()
	}

	return env.repositoryInstance
}

// GetFormatChecker returns a (lazily instantiated) singleton format checker instance
func (env *Environment) GetFormatChecker() FormatChecker {
	if env.checkerInstance == nil {
		env.checkerInstance = env.NewFormatChecker()
	}

	return env.checkerInstance
}
