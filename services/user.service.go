package services

import (
	"fmt"
	"log"

	"github.com/A-u-usman/RFID-Backend-API.git/dto"
	models "github.com/A-u-usman/RFID-Backend-API.git/entities"
	repository "github.com/A-u-usman/RFID-Backend-API.git/repositories"
	"github.com/mashingan/smapping"
)

type UserService interface {
	// InsertUser(user models.User) models.User
	AllUserActivity() []models.UserActivityLog
	CreateUser(user dto.RegisterDTO) models.User
	RecordActivity(user models.UserActivityLog)
	IsDuplicateRfid(rfid string) bool
	FindUserByRfid(rfid string) interface{}
	UpdateUser(c dto.UpdateDTO) models.User
	UpdateAccessStatus(c models.User) models.User
	DeleteUser(c dto.DeleteDTO)
	AllUser() []models.User
	FindUserByID(userID string) models.User
	IsAllowedToEditUser(userID string) bool
	// // BlockAccount(user models.Admin) models.Admin
	// // // UpdateUser(user models.User) models.User
	// VerifyCredential(email string, password string) interface{}

	// FindByID(id uint64) models.Admin
	// // // ProfileUser(userID string) models.Teacher
	// // // VerifyOldPassword(id uint64, password string) interface{}
	// ResetPwd(user models.Admin, newPassword string) models.Admin
	// VerifyEmail(email string) bool
	// StoreOtp(userOtp models.AdminOTP) models.AdminOTP
	// VerifyOTP(email string, otp string) interface{}
	// FindByEmail(email string) models.Admin
}

type userService struct {
	//userRepository repository.UserRepository
	userRepository repository.UserRepository
}

func UserServiceImp(userRep repository.UserRepository) UserService {
	return &userService{
		userRepository: userRep,
	}
}

//******view service************

func (service *userService) AllUserActivity() []models.UserActivityLog {
	return service.userRepository.AllUsersActivity()
}

// *******endview*************
func (service *userService) CreateUser(user dto.RegisterDTO) models.User {
	userToCreate := models.User{}
	err := smapping.FillStruct(&userToCreate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res := service.userRepository.InsertUser(userToCreate)
	return res
}

func (service *userService) RecordActivity(user models.UserActivityLog) {

	// userToCreate := models.UserActivityLog{}
	// userToCreate.Name = user.Name
	// userToCreate.Email = user.Email
	// userToCreate.AccessStatus = user.AccessStatus
	// userToCreate.Status = user.Status
	// userToCreate.Rfid = user.Rfid
	// err := smapping.FillStruct(&userToCreate, smapping.MapFields(&user))
	// if err != nil {
	// 	log.Fatalf("Failed map %v", err)
	// }
	// userToCreate.UserID = user.ID
	service.userRepository.RecordActivity(user)
}

func (service *userService) IsDuplicateRfid(rfid string) bool {
	res := service.userRepository.IsDuplicateRfid(rfid)
	log.Println(res.Error)
	return !(res.Error == nil)

}

func (service *userService) FindUserByRfid(rfid string) interface{} {
	res := service.userRepository.FindUserByRfid(rfid)
	return res
}

func (service *userService) AllUser() []models.User {
	return service.userRepository.AllUsers()
}

func (service *userService) DeleteUser(c dto.DeleteDTO) {
	user := models.User{}
	err := smapping.FillStruct(&user, smapping.MapFields(&c))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	service.userRepository.DeleteUser(user)
}

func (service *userService) FindUserByID(userID string) models.User {
	return service.userRepository.FindUserByID(userID)
}

func (service *userService) IsAllowedToEditUser(userID string) bool {
	// user := service.userRepository.FindUserByID(userID)
	// id := fmt.Sprintf("%v", user.Rfid)
	// return userID == id
	user := service.userRepository.FindUserByID(userID)
	id := fmt.Sprintf("%v", user.ID)
	return userID == id
}

func (service *userService) UpdateUser(c dto.UpdateDTO) models.User {
	user := models.User{}
	err := smapping.FillStruct(&user, smapping.MapFields(&c))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	res := service.userRepository.UpdateUser(user)
	return res
}

func (service *userService) UpdateAccessStatus(c models.User) models.User {
	user := models.User{}
	err := smapping.FillStruct(&user, smapping.MapFields(&c))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	res := service.userRepository.UpdateUser(user)
	return res
}
