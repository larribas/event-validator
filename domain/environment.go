package domain

var Current = &Environment{
    NewRepository: func() Repository { return &GroundRepository{} },
    NewFormatChecker: func() FormatChecker { return &GroundFormatChecker{} },
}

type Environment struct {
    repositoryInstance Repository
    checkerInstance FormatChecker

    NewRepository func() Repository
    NewFormatChecker func() FormatChecker
}

func (env *Environment) GetRepository() Repository {
    if env.repositoryInstance == nil {
        env.repositoryInstance = env.NewRepository()
    }

    return env.repositoryInstance
}

func (env *Environment) GetFormatChecker() FormatChecker {
    if env.checkerInstance == nil {
        env.checkerInstance = env.NewFormatChecker()
    }

    return env.checkerInstance
}