package repository

import (
	"time"

	"github.com/dangLuan01/ets-api/internal/models"
	"github.com/google/uuid"
)

type UserRepository interface {
	FindAll() ([]models.User, error)
	FindBYUUID(uuid string) (models.User, error)
	Create(user models.User) error
	Update(uuid uuid.UUID, user models.User) error
	Delete(uuid uuid.UUID) error
	FindByEmail(email string) (models.User, error)
	UpdateMember(uuid string, is_member int, expriedDate time.Time) error
	UpdatePassword(uuid string, password string) error
	UpdateCountUpload(uuid string) error
}