package v1service

import (
	v1dto "github.com/dangLuan01/ets-api/internal/dto/v1"
	"github.com/dangLuan01/ets-api/internal/models"
	repository "github.com/dangLuan01/ets-api/internal/repository/certificate"
	"github.com/doug-martin/goqu/v9"
)

type certificateService struct {
	repo repository.CertificateRepository
}

func NewCertificateService(repo repository.CertificateRepository) CertificateService {
	return &certificateService{
		repo: repo,
	}
}

func (cs *certificateService) GetAllCertificates() ([]models.Certificate, error) {
	return cs.repo.GetAllCertificates()
}

func (cs *certificateService) CreateCertificate(params v1dto.CertificateParamsInput) error {
	return cs.repo.CreateCertificate(params)
}

func (cs *certificateService) EditCertificate(id int) (models.Certificate, error) {
	return cs.repo.FindCertificateById(id)
}

func (cs *certificateService) UpdateCertificate(params v1dto.CertificateParamsUpdate) error {
	updateData := goqu.Record{}
	if params.Code != nil {
		updateData["code"] = params.Code
	}
	if params.Name != nil {
		updateData["name"] = params.Name
	}
	if params.Description != nil {
		updateData["description"] = params.Description
	}
	if params.Status != nil {
		updateData["status"] = params.Status
	}

	return cs.repo.UpdateCertificateById(params.Id, updateData)
}