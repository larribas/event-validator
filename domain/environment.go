package domain

var Current Environment

type Environment struct {
    repositoryInstance Repository

    NewRepository func() Repository
}

func (env *Environment) GetRepository() Repository {
    if env.repositoryInstance == nil {
        env.repositoryInstance = env.NewRepository()
    }

    return env.repositoryInstance
}