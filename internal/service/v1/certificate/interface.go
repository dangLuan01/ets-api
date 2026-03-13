package v1service

import (
	v1dto "github.com/dangLuan01/ets-api/internal/dto/v1"
	"github.com/dangLuan01/ets-api/internal/models"
)

type CertificateService interface {
	GetAllCertificates(params v1dto.GetAllCertificateParams) ([]models.Certificate, int64, error)
	CreateCertificate(params v1dto.CertificateParamsInput) error
	EditCertificate(id int) (models.Certificate, error)
	UpdateCertificate(params v1dto.CertificateParamsUpdate) error
}