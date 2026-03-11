package v1service

import (
	v1dto "github.com/dangLuan01/ets-api/internal/dto/v1"
	"github.com/dangLuan01/ets-api/internal/models"
)

type CertificateService interface {
	GetAllCertificates() ([]models.Certificate, error)
	CreateCertificate(params v1dto.CertificateParamsInput) error
	EditCertificate(id int) (models.Certificate, error)
	UpdateCertificate(params v1dto.CertificateParamsUpdate) error
}