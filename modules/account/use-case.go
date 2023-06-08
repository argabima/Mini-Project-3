package account

//	func (u accountUsecase) UpdateByID(id int, actor *Actors) error {
//		return u.repo.UpdateByID(id, actor)
//	}
type AccountUsecaseInterface interface {
	Create(account *Actors) error
	getByUsername(username string) (Actors, error)
}

type accountUsecase struct {
	repo AccountRepoInterface
}

func NewAccountUsecase(repo AccountRepoInterface) AccountUsecaseInterface {
	return accountUsecase{
		repo: repo,
	}
}

func (u accountUsecase) Create(account *Actors) error {
	return u.repo.Save(account)
}

func (u accountUsecase) getByUsername(username string) (Actors, error) {
	return u.repo.FindByUsername(username)
}
