package doc

import "gorm.io/gorm"

type IRepository interface {
	Create()
}

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) IRepository {
	return &Repository{
		DB: db,
	}
}

func (repo Repository) Create() {

}
