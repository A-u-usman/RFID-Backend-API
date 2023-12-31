package repository

// models "github.com/QT-Solution-Services/Inventory-API.git/entities"
import (
	"log"

	models "github.com/A-u-usman/RFID-Backend-API.git/entities"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository interface {
	AllUsersActivity() []models.UserActivityLog
	InsertUser(user models.User) models.User
	RecordActivity(user models.UserActivityLog) models.UserActivityLog
	IsDuplicateRfid(rfid string) (tx *gorm.DB)
	FindUserByRfid(rfid string) interface{}
	UpdateUser(user models.User) models.User
	DeleteUser(user models.User)
	AllUsers() []models.User
	FindUserByID(userID string) models.User
	// BlockAccount(user models.Admin) models.Admin
	// // UpdateUser(user models.User) models.User
	// VerifyCredential(email string, password string) interface{}

	// FindByID(id uint64) models.User
	// FindByEmail(email string) models.Admin
	// VerifyEmail(email string) bool
	// // ProfileUser(userID string) models.Teacher
	// // VerifyOldPassword(id uint64, password string) interface{}
	// ResetPwd(user models.User, newPassword string) models.Admin
	// VerifyEmail(email string) bool
	// StoreOtp(userOtp models.AdminOTP) models.AdminOTP
	// VerifyOTP(email string, otp string) interface{}
	// FindByEmail(email string) models.Admin
}

type userConnection struct {
	connection *gorm.DB
}

// // NewUserRepository create a new instance of UserRepository
func UserRepositoryImp(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}

// ******view repo
func (db *userConnection) AllUsersActivity() []models.UserActivityLog {
	var users []models.UserActivityLog
	db.connection.Find(&users)
	return users
}

// *******end view
func (db *userConnection) InsertUser(user models.User) models.User {
	// user.Rfid = hashAndSalt([]byte(user.Rfid))
	user.AccessStatus = false
	db.connection.Save(&user)
	return user
}

func (db *userConnection) RecordActivity(user models.UserActivityLog) models.UserActivityLog {
	// user.Rfid = hashAndSalt([]byte(user.Rfid))
	db.connection.Save(&user)
	return user
}

func (db *userConnection) IsDuplicateRfid(rfid string) (tx *gorm.DB) {
	var user models.User
	// hashed_rfid := hashAndSalt([]byte(rfid))
	// db.connection.Where("rfid = ?", hashed_rfid).Take(&user)
	return db.connection.Where("rfid = ?", rfid).Take(&user)
}

func (db *userConnection) AllUsers() []models.User {
	var users []models.User
	db.connection.Find(&users)
	return users
}

func (db *userConnection) DeleteUser(user models.User) {
	db.connection.Delete(&user)
}

func (db *userConnection) FindUserByID(userID string) models.User {
	var user models.User
	db.connection.Find(&user, userID)
	return user
}

func (db *userConnection) FindUserByRfid(rfid string) interface{} {
	var user models.User
	// db.connection.Find(&user, userID)
	db.connection.Where("rfid = ? and status=true", rfid).Take(&user)
	return user
}

func (db *userConnection) UpdateUser(user models.User) models.User {
	db.connection.Save(&user)
	// db.connection.Preload("Admin").Find(&c)
	return user
}

// // func (db *userConnection) BlockAccount(user models.Admin) models.Admin {

// // 	user.AccountStatus = false
// // 	db.connection.Save(&user)
// // 	return user
// // }

// func (db *adminConnection) VerifyCredential(email string, password string) interface{} {
// 	var user models.Admin
// 	res := db.connection.Where("email = ?", email).Take(&user)
// 	if res.Error == nil {
// 		return user
// 	}
// 	return nil
// }

// func (db *adminConnection) FindByID(id uint64) models.Admin {
// 	var user models.Admin
// 	db.connection.Where("id = ?", id).Take(&user)
// 	return user
// }

// func (db *adminConnection) StoreOtp(userOtp models.AdminOTP) models.AdminOTP {
// 	db.connection.Save(&userOtp)
// 	return userOtp
// }

// func (db *adminConnection) ResetPwd(user models.Admin, newPassword string) models.Admin {
// 	user.Password = hashAndSalt([]byte(newPassword))
// 	db.connection.Save(&user)
// 	return user
// }

// func (db *adminConnection) VerifyEmail(email string) bool {
// 	var user models.Admin
// 	res := db.connection.Where("email = ?", email).Take(&user)
// 	if res.Error == nil {
// 		return true
// 	}
// 	return false
// }

// func (db *adminConnection) FindByEmail(email string) models.Admin {
// 	var user models.Admin
// 	db.connection.Where("email = ?", email).Take(&user)
// 	return user
// }

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)

	if err != nil {
		log.Panicln(err)
		panic("Failed to hashed a paaword")
	}

	return string(hash)
}

// // func (db *userConnection) UpdateUser(user models.Teacher) models.Teacher {
// // 	if user.Password != "" {
// // 		user.Password = hashAndSalt([]byte(user.Password))
// // 	} else {
// // 		var tempUser models.User
// // 		db.connection.Find(&tempUser, user.ID)
// // 		user.Password = tempUser.Password
// // 	}
// // 	db.connection.Save(&user)
// // 	return user
// // }

// func (db *adminConnection) VerifyOTP(email string, otp string) interface{} {
// 	var user models.AdminOTP
// 	res := db.connection.Where("user_email = ? AND otp = ?", email, otp).Take(&user)
// 	if res.Error == nil {
// 		return user
// 	}
// 	return nil
// }

// // //	func (db *userConnection) IsDuplicateEmail(email string) (tx *gorm.DB) {
// // //		var user models.User
// // //		return db.connection.Where("email = ?", email).Take(&user)
// // //	}
// // func (db *userConnection) ProfileUser(userID string) models.User {
// // 	var user models.User
// // 	db.connection.Find(&user, userID)
// // 	return user
// // }

// // func (db *userConnection) VerifyOldPassword(id uint64, password string) interface{} {
// // 	var user models.User
// // 	res := db.connection.Where("id = ?", id).Take(&user)
// // 	if res.Error == nil {
// // 		return user
// // 	}
// // 	return nil
// // }
