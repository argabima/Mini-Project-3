package account

import (
	"gorm.io/gorm"
)

type AccountRepoInterface interface {
	Save(account *Actors) error
	FindByUsername(username string) (Actors, error)
	FindAll() ([]Actors, error)
	Update(actor *Actors) error
	// Delete(actor *Actors) error
	FindByID(id any) (Actors, error)
}

type accountRepo struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) AccountRepoInterface {
	return accountRepo{db}
}

func (r accountRepo) Save(account *Actors) error {
	return r.db.Create(account).Error
}

func (r accountRepo) FindAll() ([]Actors, error) {
	var actors []Actors

	err := r.db.Find(&actors).Error
	if err != nil {
		return actors, err
	}

	return actors, nil
}

func (r accountRepo) FindByUsername(username string) (Actors, error) {
	var actor Actors

	err := r.db.Where("username = ?", string(username)).First(&actor).Error
	if err != nil {
		return actor, err
	}

	return actor, nil
}


func (r accountRepo) Update(actor *Actors) error {
	return r.db.Save(actor).Error
}


// func (r accountRepo) Delete(actor *Actors) error {
// 	return r.db.Delete(actor).Error
// }

func (r accountRepo) FindByID(id any) (Actors, error) {
	var actor Actors
	err := r.db.First(&actor, id).Error

	return actor, err
}


