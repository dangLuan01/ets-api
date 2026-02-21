package v1service

import (
	"fmt"

	v1dto "github.com/dangLuan01/ets-api/internal/dto/v1"
	"github.com/dangLuan01/ets-api/internal/models"
	"github.com/dangLuan01/ets-api/internal/repository"
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
	
	userLogged, err := utils.GetUserLogged(ctx)
	if err != nil {
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
	if user, err := us.repo.FindByEmail(user.Email); err != nil {
		
		return models.User{}, utils.NewError(
			string(utils.ErrCodeConflict), 
			fmt.Sprintf("Email: %v already existed.", user.Email),
		)
	}
	user.UUID = uuid.New()
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {

		return models.User{}, utils.WrapError(
			string(utils.ErrCodeInternal), 
			"Faile hash password", 
			err,
		)
	}
	user.Password = string(hashPassword)
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
	if u, err := us.repo.FindByEmail(user.Email); err != nil && u.UUID != uuid{
		
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

	if user.Password != "" {
		hashPassword, err :=bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return models.User{}, utils.WrapError(string(utils.ErrCodeInternal), "Faile hash pass", err)
		}
		currencyUser.Password = string(hashPassword)
		
	}
	if user.Level != 0 {
		currencyUser.Level = user.Level	
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

func (us *userService) CheckStatus(ctx *gin.Context, uuid string) error {
	
	user, err := us.repo.FindBYUUID(uuid)
	if err != nil {
		return utils.NewError(string(utils.ErrCodeInternal), "Error fetch user.")
	}
	
	if user.UploadCount == 0 {
		return utils.NewError(string(utils.ErrCodeNotFound), "Bạn đã hết lượt upload. Vui lòng đăng ký thành viên để upload không giới hạn.")
	}

	return nil
}

func (us *userService) ChangePassword(ctx *gin.Context, params v1dto.ChangerPasswordParams) error {
	
	user, err := utils.GetUserLogged(ctx)
	if err != nil {
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

func (us *userService) UpdateCountUpload(uuid string) error {

	if err := us.repo.UpdateCountUpload(uuid); err != nil {
		return utils.WrapError(string(utils.ErrCodeInternal), "Falied update count upload", err)
	}

	return  nil
}