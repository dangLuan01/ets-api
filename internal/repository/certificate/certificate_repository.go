package repository

import (
	v1dto "github.com/dangLuan01/ets-api/internal/dto/v1"
	"github.com/dangLuan01/ets-api/internal/models"
	"github.com/dangLuan01/ets-api/internal/utils"
	"github.com/doug-martin/goqu/v9"
)

const (
	TABLE_CERTIFICATE 			= "certificates"
)


type SqlCertificateRepository struct {
	db *goqu.Database
}

func NewSqlCertificateRepository(DB *goqu.Database) CertificateRepository {
	return &SqlCertificateRepository{
		db: DB,
	}
}

func (cr *SqlCertificateRepository) GetAllCertificates() ([]models.Certificate, error) {
	var certificates []models.Certificate
	err := cr.db.From(TABLE_CERTIFICATE).ScanStructs(&certificates)
	if err != nil {
		return nil, err
	}
	
	return certificates, nil
}

func (cr *SqlCertificateRepository) CreateCertificate(params v1dto.CertificateParamsInput) error {
	_, err := cr.db.From(TABLE_CERTIFICATE).Insert().Rows(params).Executor().Exec()
	if err != nil {
		return err
	}

	return nil
}

func (cr *SqlCertificateRepository) FindCertificateById(id int) (certificate models.Certificate, err error) {

	found, err := cr.db.From(TABLE_CERTIFICATE).Where(goqu.C("id").Eq(id)).ScanStruct(&certificate)
	if err != nil {
		return models.Certificate{}, err
	}

	if !found {
		return models.Certificate{}, utils.NewError(string(utils.ErrCodeNotFound),"Not found certificate.")
	}

	return certificate, nil
}

func (cr *SqlCertificateRepository) UpdateCertificateById(id int, params goqu.Record) error {

	_, err := cr.db.From(TABLE_CERTIFICATE).
		Update().Set(params).
		Where(goqu.C("id").Eq(id)).
		Executor().Exec()

	return err
}