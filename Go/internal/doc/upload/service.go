package doc

import (
	"gorm.io/gorm"
)

type IService interface {
	// Upload(file model.Upload)
}

type Service struct {
	Repository IRepository
}

func NewService(db *gorm.DB) IService {
	return &Service{
		Repository: NewRepository(db),
	}
}
