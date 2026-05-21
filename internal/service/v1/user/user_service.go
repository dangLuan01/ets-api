package v1service

import (
	"fmt"
	"log"
	"time"

	v1dto "github.com/dangLuan01/ets-api/internal/dto/v1"
	"github.com/dangLuan01/ets-api/internal/models"
	repository "github.com/dangLuan01/ets-api/internal/repository/user"
	repositoryAttempt "github.com/dangLuan01/ets-api/internal/repository/user_attempt"
	"github.com/dangLuan01/ets-api/internal/utils"
	"github.com/gin-gonic/gin"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	repo repository.UserRepository
	repoAttempt repositoryAttempt.UserAttemptRepository
}

func NewUserService(repo repository.UserRepository, repoAttempt repositoryAttempt.UserAttemptRepository) UserService {
	return &userService{
		repo: repo,
		repoAttempt: repoAttempt,
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
		return models.User{}, utils.NewError(string(utils.ErrCodeUnauthorized), "User not logged.")
	}

	user, err := us.repo.FindBYUUID(userLogged.UserUUID.String());
	if err != nil {
		return models.User{}, utils.NewError(string(utils.ErrCodeNotFound), "Not found user")
	}
	
	return user, nil
}

func (us *userService) CreateUser(user models.User) (models.User, error) {
	user.Email = utils.NormailizeString(user.Email)
	if user, existed, err := us.repo.FindExistByEmail(user.Email); err == nil && existed {
		
		return models.User{}, utils.NewError(
			string(utils.ErrCodeConflict), 
			fmt.Sprintf("Email: %v already existed.", user.Email),
		)
	}
	user.UUID = uuid.New()
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(*user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {

		return models.User{}, utils.WrapError(
			string(utils.ErrCodeInternal), 
			"Faile hash password", 
			err,
		)
	}
	passStr := string(hashPassword)
	user.PasswordHash = &passStr
	if err := us.repo.Create(user); err != nil {

		return models.User{}, utils.WrapError(
			string(utils.ErrCodeInternal), 
			"Faile create user", 
			err,
		)
	}
	
	return user, nil
}
func (us *userService) UpdateUser(ctx *gin.Context, params v1dto.UpdateUserInput) error {
	if params.ExamDate != nil {
		
		date, err := time.Parse("2006-01-02", *params.ExamDate)
		if err != nil {
			return utils.WrapError(
				string(utils.ErrCodeBadRequest),
				"exam_date",
				fmt.Errorf("Ngày không đúng định dạng YYYY-MM-DD"),
			)
		}

		now :=  time.Now()
		today := time.Date(
			now.Year(),
			now.Month(),
			now.Day(),
			0,0,0,0,
			now.Location(),
		)

		if date.Before(today) {
			return utils.WrapError(string(utils.ErrCodeBadRequest), "exam_date", fmt.Errorf("Ngày dự kiến không hợp lệ!"))				
		}
	}

	user, exists := utils.GetUserLogged(ctx)
	if !exists {
		return utils.NewError(string(utils.ErrCodeInternal), "Người dùng chưa đăng nhập.")
	}

	if err := us.repo.Update(
		user.UserUUID, 
		models.UserUpdate{
			UserName: params.Name,
			Target: params.Target,
			ExamDate: params.ExamDate,
		},
	); err != nil {
		log.Println(err)
		return utils.NewError(string(utils.ErrCodeInternal), "Cập nhật thông tin không thành công.")
	}

	return nil
}

func (us *userService) DeleteUser(uuid uuid.UUID) error {
	if err := us.repo.Delete(uuid); err != nil {
		return utils.WrapError(string(utils.ErrCodeInternal), "Faile delete user", err)
	}

	return nil
}

func (us *userService) UpdatePassword(ctx *gin.Context, params v1dto.UpdatePasswordInput) error {
	
	user, exists := utils.GetUserLogged(ctx)
	if !exists {
		return utils.NewError(string(utils.ErrCodeBadRequest), "Người dùng chưa đăng nhập.")
	}

	userExist, err := us.repo.FindBYUUID(user.UserUUID.String())

	if err != nil {	
		return utils.NewError(string(utils.ErrCodeInternal), "Lỗi hệ thống!")
	}

	if params.HasPassword {
		if err := bcrypt.CompareHashAndPassword([]byte(*userExist.PasswordHash), []byte(params.CurrentPassword)); err != nil {
			return utils.WrapError(string(utils.ErrCodeBadRequest), "current_password", fmt.Errorf("Mật khẩu không đúng!"))
		}
	}
	
	if params.NewPassword != params.ConfirmPassword {
		return utils.WrapError(string(utils.ErrCodeBadRequest), "confirm_password", fmt.Errorf("Mật khẩu xác nhận không khớp!"))
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(params.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return utils.NewError(string(utils.ErrCodeInternal), "Unable error hash password.")
	}

	if err := us.repo.UpdatePassword(user.UserUUID.String(), string(hashPassword)); err != nil {
		return utils.NewError(string(utils.ErrCodeInternal), "Cập nhật mật khẩu thất bại!")
	}

	return nil
}

func (us *userService) GetAttemptByUserUUID(ctx *gin.Context, params v1dto.GetAttemptByUserUuuidParams) ([]v1dto.UserAttemptDTO, int64, error) {
	user, _ := utils.GetUserLogged(ctx)
	return us.repoAttempt.FindAttemptByUserUUID(user.UserUUID.String(), params)
}