package account

import (
	"gorm.io/gorm"
)

type AccountRepoInterface interface {
	Save(account *Actors) error
	FindByUsername(username string) (Actors, error)
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

// func (r accountRepo) UpdateByID(id int, actor *Actors) error {
// 	return r.db.Model(&Actors{}).Where("id = ?", id).Updates(actor).Error
// }

func (r accountRepo) FindByUsername(username string) (Actors, error) {
	var actor Actors

	err := r.db.Where("username = ?", string(username)).First(&actor).Error
	if err != nil {
		return actor, err
	}

	return actor, nil
}
