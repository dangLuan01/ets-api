package repository

import (
	"fmt"

	"github.com/dangLuan01/ets-api/internal/models"
	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
)

const	TABLE_USER = "users"

type SqlUserRepository struct {
	db *goqu.Database
}

func NewSqlUserRepository(DB *goqu.Database) UserRepository {
	return &SqlUserRepository{
		db: DB,
	}
}

func (ur *SqlUserRepository) FindAll() ([]models.User, error){
	
	ds := ur.db.From(goqu.T(TABLE_USER)).
	Select(
		goqu.I("uuid"),
		goqu.I("username"),
		goqu.I("email"),
		goqu.I("role"),
		goqu.I("status"),
	)
	var users []models.User
	if err := ds.ScanStructs(&users); err != nil {
		return nil, fmt.Errorf("faile get all user:%v", err)
	}

	return users, nil
}

func (ur *SqlUserRepository) FindBYUUID(uuid string) (models.User, error) {
	var user models.User

	ds := ur.db.From(goqu.T(TABLE_USER)).
	Where(
		goqu.C("uuid").Eq(uuid),
	).
	Select(
		goqu.C("uuid"),
		goqu.C("username"),
		goqu.C("email"),
		goqu.C("password_hash"),
		goqu.C("avatar"),
		goqu.C("target"),
		goqu.C("role"),
		goqu.C("exam_date"),
	)

	found, err := ds.ScanStruct(&user)
	if err != nil || !found {
		return  models.User{}, err
	}

	return user, err
}

func (ur *SqlUserRepository) Create(user models.User) error {
	insertUser := ur.db.Insert(TABLE_USER).Rows(user).Executor()
	if _, err := insertUser.Exec(); err != nil {
       return fmt.Errorf("faile insert rows user")
	}

	return nil
}

func (ur *SqlUserRepository) Update(uuid uuid.UUID, user models.UserUpdate) error {
	ds := ur.db.Update(TABLE_USER).Set(user).
		Where(goqu.C("uuid").Eq(uuid))

	if _, err := ds.Executor().Exec(); err != nil {
		return err
	}

	return nil
}

func (ur *SqlUserRepository) Delete(uuid uuid.UUID) error {
	panic("")
}

func (ur *SqlUserRepository) FindExistByEmail(email string) (models.User, bool, error) {
	
	ds := ur.db.From(goqu.T(TABLE_USER)).Where(
		goqu.C("email").Eq(email),
	).Limit(1)
	
    var user models.User
    found, err := ds.ScanStruct(&user)
	if err != nil {
		return models.User{}, false, err
	}
	
	if found {
		return user, true, nil
	}

	return models.User{}, false, err
}

func (ur *SqlUserRepository) UpdatePassword(uuid, password string) error {
	_, err := ur.db.Update(goqu.T(TABLE_USER)).
		Set(goqu.Record{"password_hash": password}).
		Where(goqu.C("uuid").Eq(uuid)).
		Executor().Exec()
	if err != nil {
		return err	
	}

	return nil
}