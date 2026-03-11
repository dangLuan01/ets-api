package repository

import (
	"fmt"
	"slices"
	"time"

	"github.com/dangLuan01/ets-api/internal/models"
	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
)

type SqlUserRepository struct {
	users []models.User
	db *goqu.Database
}

func NewSqlUserRepository(DB *goqu.Database) UserRepository {
	return &SqlUserRepository{
		users : make([]models.User, 0),
		db: DB,
	}
}

func (ur *SqlUserRepository) FindAll() ([]models.User, error){
	
	ds := ur.db.From(goqu.T("users")).
	Select(
		goqu.I("uuid"),
		goqu.I("username"),
		goqu.I("email"),
		goqu.I("level"),
		goqu.I("status"),
	)
	var users []models.User
	if err := ds.ScanStructs(&users); err != nil {
		return nil, fmt.Errorf("faile get all user:%v", err)
	}

	return users, nil
}

func (ur *SqlUserRepository) FindBYUUID(uuid string) (models.User, error) {
	ds := ur.db.From(goqu.T("users")).
	Where(	
		goqu.C("uuid").Eq(uuid),
	).
	Select(
		goqu.I("uuid"),
		goqu.I("username"),
		goqu.I("email"),
		goqu.I("level"),
		goqu.I("is_member"),
		goqu.I("upload_count"),
		goqu.I("expried_date"),
		goqu.I("status"),
	)
	var user models.User

	found, err := ds.ScanStruct(&user)
	if err != nil || !found {
		return  models.User{}, err
	}

	return user, err
}

func (ur *SqlUserRepository) Create(user models.User) error {
	insertUser := ur.db.Insert("users").Rows(user).Executor()
	if _, err := insertUser.Exec(); err != nil {
       return fmt.Errorf("faile insert rows user")
	}

	return nil
}

func (ur *SqlUserRepository) Update(uuid uuid.UUID, user models.User) error {
	for i, u := range ur.users{
		if u.UUID == uuid {
			ur.users[i] = user
			return nil
		}
	}
	return fmt.Errorf("user not found")
}
func (ur *SqlUserRepository) Delete(uuid uuid.UUID) error {

	for i, u := range ur.users{
		if u.UUID == uuid {
			ur.users = slices.Delete(ur.users, i, i + 1)
			return nil
		}
	}

	return fmt.Errorf("user not found")
}
func (ur *SqlUserRepository) FindByEmail(email string) (models.User, error) {
	
	ds := ur.db.From(goqu.T("users")).Where(
		goqu.C("email").Eq(email),
	).Limit(1)
	
    var user models.User
    found, err := ds.ScanStruct(&user)
	if err != nil {
		return models.User{}, err
	}
	
	if found {
		return user, nil
	}

	return models.User{}, err
}

func (ur *SqlUserRepository) UpdateMember(uuid string, is_member int, expriedDate time.Time) error {
	
	_, err := ur.db.Update(goqu.T("users")).
		Set(
			goqu.Record{
				"is_member": is_member,
				"expried_date": expriedDate,
				"upload_count": 9999,
			},
		).
		Where(
			goqu.C("uuid").Eq(uuid),
		).Executor().Exec()

	if err != nil {
		return err
	}
	
	return nil
}

func (ur *SqlUserRepository) UpdatePassword(uuid, password string) error {
	_, err := ur.db.Update(goqu.T("users")).
		Set(goqu.Record{"password": password}).
		Where(goqu.C("uuid").Eq(uuid)).
		Executor().Exec()
	if err != nil {
		return err	
	}

	return nil
}

func (ur *SqlUserRepository) UpdateCountUpload(uuid string) error {

	user , err := ur.FindBYUUID(uuid)
	if err != nil {
		return err
	}

	uploadCountAmount := user.UploadCount - 1
	if uploadCountAmount < 0 {
		uploadCountAmount = 0
	}

	_, err = ur.db.Update(goqu.T("users")).
		Set(goqu.Record{
			"upload_count": uploadCountAmount,
		}).
		Where(goqu.C("uuid").Eq(uuid)).Executor().Exec()
	
	if err != nil {
		return err
	}
	
	return nil
}
