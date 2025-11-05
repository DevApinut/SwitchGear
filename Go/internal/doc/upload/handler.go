package doc

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type IHandler interface {
	Upload(ctx *gin.Context)
}
type Handler struct {
	Service IService
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{
		Service: NewService(db),
	}
}
