package repository

import (
	v1dto "github.com/dangLuan01/ets-api/internal/dto/v1"
	"github.com/dangLuan01/ets-api/internal/models"
	"github.com/doug-martin/goqu/v9"
)

type CertificateRepository interface {
	GetAllCertificates() ([]models.Certificate, error)
	CreateCertificate(params v1dto.CertificateParamsInput) error
	FindCertificateById(id int) (models.Certificate, error)
	UpdateCertificateById(id int, params goqu.Record) error
}