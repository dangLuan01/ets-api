package v1service

import (
	"fmt"

	v1dto "github.com/dangLuan01/ets-api/internal/dto/v1"
	"github.com/dangLuan01/ets-api/internal/models"
	repository "github.com/dangLuan01/ets-api/internal/repository/user"
	"github.com/dangLuan01/ets-api/internal/utils"
	"github.com/gin-gonic/gin"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (us *userService) GetAllUser()  ([]models.User, error) {
	users, err := us.repo.FindAll()
	if err != nil {
		
		return nil, utils.WrapError(
			string(utils.ErrCodeInternal), 
			"Faile fetch users.", 
			err,
		)
	}

	return users, nil
}

func (us *userService) GetUserByUUID(ctx *gin.Context) (models.User, error) {
	
	userLogged, exists := utils.GetUserLogged(ctx)
	if !exists {
		return models.User{}, utils.NewError(string(utils.ErrCodeInternal), "Failed get user logged.")
	}

	user, err := us.repo.FindBYUUID(userLogged.UserUUID.String());
	if err != nil {

		return models.User{}, utils.NewError(string(utils.ErrCodeNotFound), "No user")
	}
	
	return user, nil
}

func (us *userService) CreateUser(user models.User) (models.User, error) {
	user.Email = utils.NormailizeString(user.Email)
	if user, existed, err := us.repo.FindByEmail(user.Email); err == nil && existed {
		
		return models.User{}, utils.NewError(
			string(utils.ErrCodeConflict), 
			fmt.Sprintf("Email: %v already existed.", user.Email),
		)
	}
	user.UUID = uuid.New()
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {

		return models.User{}, utils.WrapError(
			string(utils.ErrCodeInternal), 
			"Faile hash password", 
			err,
		)
	}
	user.PasswordHash = string(hashPassword)
	if err := us.repo.Create(user); err != nil {

		return models.User{}, utils.WrapError(
			string(utils.ErrCodeInternal), 
			"Faile create user", 
			err,
		)
	}
	
	return user, nil
}
func (us *userService) UpdateUser(uuid uuid.UUID, user models.User) (models.User, error) {
	user.Email = utils.NormailizeString(user.Email)
	if u, _, err := us.repo.FindByEmail(user.Email); err != nil && u.UUID != uuid{
		
		return models.User{}, utils.NewError(
			string(utils.ErrCodeConflict), 
			fmt.Sprintf("Email: %v already existed.", u.Email),
		)
	}
	currencyUser, err := us.repo.FindBYUUID(uuid.String())
	if err != nil {
		return models.User{}, utils.NewError(string(utils.ErrCodeNotFound), "user not found")
	}
	
	currencyUser.UserName = user.UserName
	currencyUser.Email = user.Email

	if user.PasswordHash != "" {
		hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
		if err != nil {
			return models.User{}, utils.WrapError(string(utils.ErrCodeInternal), "Faile hash pass", err)
		}
		currencyUser.PasswordHash = string(hashPassword)
		
	}
	if user.Role != 0 {
		currencyUser.Role = user.Role
	}
	
	if user.Status != 0 {
		currencyUser.Status = user.Status	
	}
	
	if err := us.repo.Update(uuid, currencyUser); err != nil {
		return models.User{}, utils.WrapError(string(utils.ErrCodeInternal), "Faile update user", err)
	}
	return currencyUser, nil
}

func (us *userService) DeleteUser(uuid uuid.UUID) error {
	if err := us.repo.Delete(uuid); err != nil {
		return utils.WrapError(string(utils.ErrCodeInternal), "Faile delete user", err)
	}

	return nil
}

func (us *userService) ChangePassword(ctx *gin.Context, params v1dto.ChangerPasswordParams) error {
	
	user, exists := utils.GetUserLogged(ctx)
	if !exists {
		return utils.NewError(string(utils.ErrCodeInternal), "Failed get user login.")
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		return utils.NewError(string(utils.ErrCodeInternal), "Unable error hash password.")
	}

	if err := us.repo.UpdatePassword(user.UserUUID.String(), string(hashPassword)); err != nil {
		return utils.WrapError(string(utils.ErrCodeInternal), "Failed update password", err)
	}

	return nil
}