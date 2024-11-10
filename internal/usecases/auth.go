package usecases

type UserRepository interface{}

type authUsecase struct {
	repo UserRepository
}

func NewAuthUsecase(repo UserRepository) *authUsecase {
	return &authUsecase{
		repo: repo,
	}
}
