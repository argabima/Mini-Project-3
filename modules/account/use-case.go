package account

//	func (u accountUsecase) UpdateByID(id int, actor *Actors) error {
//		return u.repo.UpdateByID(id, actor)
//	}
type AccountUsecaseInterface interface {
	Create(account *Actors) error
	getByUsername(username string) (Actors, error)
	// FindAll() ([]Actors, error)
	// Update(actor *Actors) error
	// Delete(actor *Actors) error
	// FindByID(id any) (Actors, error)
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

/*
func (u accountUsecase) FindAll() ([]Actors, error) {
	return u.repo.FindAll()
}

func (u accountUsecase) Update(actor *Actors) error {
	return u.repo.Update(actor)
}

func (u accountUsecase) Delete(actor *Actors) error {
	return u.repo.Delete(actor)
}

func (u accountUsecase) FindByID(id any) (Actors, error) {
	return u.repo.FindByID(id)
}
*/