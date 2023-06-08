package customers

type Usecase struct {
	repo RepositoryInterface
}

type UsecaseInterface interface {
	Create(customers *Customers) error
	Read() ([]Customers, error)
	ReadByID(id any) (Customers, error)
	Save(customers *Customers) error
	FindAll() ([]Customers, error)
	// FindById(id int) (Customers, error)	
	UpdateByID(customers *Customers) error
	DeleteByID(customers *Customers) error
	getByEmail(email string) (Customers, error)
}

func NewUsecase(repo RepositoryInterface) UsecaseInterface {
	return Usecase{
		repo: repo,
	}
}

func (u Usecase) Create(customers *Customers) error {
	return u.repo.Save(customers)
}

func (u Usecase) Read() ([]Customers, error) {
	return u.repo.FindAll()
}

func (u Usecase) ReadByID(id any) (Customers, error) {
	return u.repo.FindById(id)
}

func (u Usecase) Save(customers *Customers) error {
	return u.repo.Save(customers)
}

func (u Usecase) FindAll() ([]Customers, error) {
	return u.repo.FindAll()
}

func (u Usecase) UpdateByID(customers *Customers) error {
	return u.repo.UpdateByID(customers)
}

func (u Usecase) DeleteByID(customers *Customers) error {
	return u.repo.DeleteByID(customers)
}

func (u Usecase) getByEmail(email string) (Customers, error) {
	return u.repo.FindByEmail(email)
}
