package controllers

import (
	"log"
	"net/http"

	"github.com/A-u-usman/RFID-Backend-API.git/dto"
	models "github.com/A-u-usman/RFID-Backend-API.git/entities"
	helper "github.com/A-u-usman/RFID-Backend-API.git/helpers"
	"github.com/A-u-usman/RFID-Backend-API.git/services"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	// Login(ctx *gin.Context)
	Register(ctx *gin.Context)
	EnableDisableUser(ctx *gin.Context)
	SaveUser(ctx *gin.Context)
	ShowUserForm(ctx *gin.Context)
	AllUsers(ctx *gin.Context)
	ShowUsers(ctx *gin.Context)
	ShowUsersActivity(ctx *gin.Context)
	// FindByID(context *gin.Context)
	FindUserByID(ctx *gin.Context)
	// InsertCategory(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	DelUser(ctx *gin.Context)
	// BlockAccount(ctx *gin.Context)
	// ForgotPassword(ctx *gin.Context)
	// ChangePassword(ctx *gin.Context)
	// ResetPassword(ctx *gin.Context)
	// SendOTP(ctx *gin.Context)
	// VerifyOTP(ctx *gin.Context)
	// InviteUser(ctx *gin.Context)
}

type userController struct {
	userService services.UserService
	// jwtService  services.JWTService
	mailService services.MailService
}

// mailService services.MailService to be addede letter
func UserControllerImp(userService services.UserService, mailService services.MailService) UserController {
	return &userController{
		userService: userService,
		// jwtService:  jwtService,
		mailService: mailService,
	}
}

// view functions ***************************
func (c *userController) ShowUsers(ctx *gin.Context) {
	users := c.userService.AllUser()
	data := gin.H{
		"title": "Users page",
		"users": users,
	}
	ctx.HTML(http.StatusOK, "index.html", data)
}

func (c *userController) ShowUsersActivity(ctx *gin.Context) {
	users := c.userService.AllUserActivity()
	data := gin.H{
		"title": "Users page",
		"users": users,
	}
	ctx.HTML(http.StatusOK, "activity.html", data)
}

func (c *userController) ShowUserForm(ctx *gin.Context) {

	data := gin.H{}
	ctx.HTML(http.StatusOK, "newuser.html", data)
}

func (c *userController) SaveUser(ctx *gin.Context) {
	var registerDTO dto.RegisterDTO
	log.Println(registerDTO)
	errDTO := ctx.ShouldBind(&registerDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		data := gin.H{
			"message": response,
		}
		ctx.HTML(http.StatusBadRequest, "newuser.html", data)
		return
	}

	if !c.userService.IsDuplicateRfid(registerDTO.Rfid) {
		response := helper.BuildErrorResponse("Failed to process request", "RFID already Exist", helper.EmptyObj{})
		data := gin.H{
			"message": response,
		}
		ctx.HTML(http.StatusBadRequest, "newuser.html", data)
		return
	} else {
		c.userService.CreateUser(registerDTO)
		// token := c.jwtService.GenerateToken(strconv.FormatUint(createdUser.ID, 10))
		// createdUser.Token = token
		data := gin.H{
			"message": "User Successfully registered",
		}
		ctx.HTML(http.StatusOK, "newuser.html", data)
		return
	}
}

func (c *userController) EnableDisableUser(ctx *gin.Context) {
	rfid := ctx.Param("rfid")
	if rfid == "" {
		res := getstruct{Message: "false"}
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	getRfidUser := c.userService.FindUserByID(rfid)
	userStatus := getRfidUser.Status
	if userStatus == true {
		getRfidUser.Status = false
		getRfidUser.AccessStatus = false
		//change access to false (outside)
		c.userService.UpdateAccessStatus(getRfidUser)
		//get all list and return response
		users := c.userService.AllUser()
		data := gin.H{
			"title": "Users page",
			"users": users,
		}
		ctx.HTML(http.StatusOK, "index.html", data)
		return
	}
	getRfidUser.Status = true
	c.userService.UpdateAccessStatus(getRfidUser)
	users := c.userService.AllUser()
	data := gin.H{
		"title": "Users page",
		"users": users,
	}
	ctx.HTML(http.StatusOK, "index.html", data)
	return
}

//end view fucn *******************************

// API functions*********************************
func (c *userController) Register(ctx *gin.Context) {
	var registerDTO dto.RegisterDTO
	log.Println(registerDTO)
	errDTO := ctx.ShouldBind(&registerDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if !c.userService.IsDuplicateRfid(registerDTO.Rfid) {
		response := helper.BuildErrorResponse("Failed to process request", "RFID already Exist", helper.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
		return
	} else {
		createdUser := c.userService.CreateUser(registerDTO)
		// token := c.jwtService.GenerateToken(strconv.FormatUint(createdUser.ID, 10))
		// createdUser.Token = token
		response := helper.BuildResponse(true, "OK!", createdUser)
		ctx.JSON(http.StatusCreated, response)
		return
	}
}

func (c *userController) AllUsers(ctx *gin.Context) {
	var users []models.User = c.userService.AllUser()
	res := helper.BuildResponse(true, "OK", users)
	ctx.JSON(http.StatusOK, res)
}

// delete fcontroller
func (c *userController) DelUser(ctx *gin.Context) {
	// var category models.Category
	// id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	// if err != nil {
	// 	response := helper.BuildErrorResponse("Failed to get id", "No param id was found", helper.EmptyObj{})
	// 	ctx.JSON(http.StatusBadRequest, response)
	// }
	// category.ID = id

	var userDeleteDTO dto.DeleteDTO
	errDTO := ctx.ShouldBind(&userDeleteDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	// authHeader := ctx.GetHeader("Authorization")
	// tokenstring := GetTokenString(authHeader)
	// token, errToken := c.jwtService.ValidateToken(tokenstring)
	// if errToken != nil {
	// 	panic(errToken.Error())
	// }
	// claims := token.Claims.(jwt.MapClaims)
	// adminID := fmt.Sprintf("%v", claims["user_id"])

	//*************** to be used letter
	// if c.userService.IsAllowedToEditUser(userDeleteDTO.ID) {
	// 	c.userService.DeleteUser(userDeleteDTO)
	// 	res := helper.BuildResponse(true, "User Deleted", helper.EmptyObj{})
	// 	ctx.JSON(http.StatusOK, res)
	// } else {
	// 	response := helper.BuildErrorResponse("You dont have permission to delete this user", "you are not the admin", helper.EmptyObj{})
	// 	ctx.JSON(http.StatusForbidden, response)
	// }

	c.userService.DeleteUser(userDeleteDTO)
	res := helper.BuildResponse(true, "User Deleted", helper.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}

type getstruct struct {
	Message string `json:"message"`
	Access  string `json:"acces"`
}

// new validate rfid
func (c *userController) FindUserByID(ctx *gin.Context) {
	rfid := ctx.Param("rfid")
	if rfid == "" {
		res := getstruct{Message: "false"}
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	// var getUserDTO dto.GetDTO
	// log.Println(getUserDTO)
	// errDTO := ctx.ShouldBind(&getUserDTO)
	// if errDTO != nil {
	// 	res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
	// 	ctx.JSON(http.StatusBadRequest, res)
	// 	return
	// }
	if !c.userService.IsDuplicateRfid(rfid) {

		getRfidUser := c.userService.FindUserByRfid(rfid)
		if v, ok := getRfidUser.(models.User); ok {
			accesStatus := v.AccessStatus
			if v.Status == false {
				var intruder models.UserActivityLog
				intruder.Rfid = rfid
				intruder.Name = "Blocked User"
				c.userService.RecordActivity(models.User(intruder))
				mgs := c.mailService.GenerateIntruderNotificationMessage(intruder.Rfid)
				mailSubjet := "Blocked User Access Attempt"
				c.mailService.SendMail(mgs, mailSubjet, "muhammadabdullahi190@gmail.com")
				response := getstruct{Message: "false",
					Access: "Blocked User intruder",
				}
				ctx.JSON(http.StatusConflict, response)
				return

			} else {
				if accesStatus == true {
					v.AccessStatus = false
					//change access to false (outside)
					c.userService.UpdateAccessStatus(v)
					//save activity
					c.userService.RecordActivity(v)
					//return response
					response := getstruct{Message: "true",
						Access: "close",
					}
					ctx.JSON(http.StatusFound, response)
					return
				}
				v.AccessStatus = true
				c.userService.UpdateAccessStatus(v)
				//save activity***************
				c.userService.RecordActivity(v)
				//return response*************
				response := getstruct{Message: "true",
					Access: "open",
				}
				ctx.JSON(http.StatusFound, response)
				return
			}
		}
		var intruder models.UserActivityLog
		intruder.Rfid = rfid
		intruder.Name = "Intruder"
		c.userService.RecordActivity(models.User(intruder))
		mgs := c.mailService.GenerateIntruderNotificationMessage(intruder.Rfid)
		mailSubjet := "Intruder Access Attempt"
		c.mailService.SendMail(mgs, mailSubjet, "muhammadabdullahi190@gmail.com")
		response := getstruct{Message: "false",
			Access: "intruder",
		}
		ctx.JSON(http.StatusConflict, response)
		return
	} else {
		var intruder models.UserActivityLog
		intruder.Rfid = rfid
		intruder.Name = "Intruder"
		mgs := c.mailService.GenerateIntruderNotificationMessage(intruder.Rfid)
		mailSubjet := "Intruder Access Attempt"
		c.mailService.SendMail(mgs, mailSubjet, "muhammadabdullahi190@gmail.com")
		c.userService.RecordActivity(models.User(intruder))
		response := getstruct{Message: "false",
			Access: "intruder",
		}
		ctx.JSON(http.StatusConflict, response)

	}
}

//old validate uid
// func (c *userController) FindUserByID(ctx *gin.Context) {
// 	rfid := ctx.Param("rfid")
// 	if rfid == "" {
// 		res := getstruct{Message: "false"}
// 		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
// 		return
// 	}

// 	// var getUserDTO dto.GetDTO
// 	// log.Println(getUserDTO)
// 	// errDTO := ctx.ShouldBind(&getUserDTO)
// 	// if errDTO != nil {
// 	// 	res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
// 	// 	ctx.JSON(http.StatusBadRequest, res)
// 	// 	return
// 	// }
// 	if !c.userService.IsDuplicateRfid(rfid) {

// 		response := getstruct{Message: "true"}
// 		ctx.JSON(http.StatusFound, response)
// 	} else {
// 		// createdUser := c.userService.CreateUser(registerDTO)
// 		// // token := c.jwtService.GenerateToken(strconv.FormatUint(createdUser.ID, 10))
// 		// // createdUser.Token = token
// 		// response := helper.BuildResponse(true, "OK!", createdUser)
// 		// ctx.JSON(http.StatusCreated, response)

// 		// response := helper.BuildErrorResponse("false", "Invalid", helper.EmptyObj{})
// 		response := getstruct{Message: "false"}
// 		ctx.JSON(http.StatusConflict, response)

// 	}
// authHeader := ctx.GetHeader("Authorization")
// tokenstring := GetTokenString(authHeader)
// token, errToken := c.jwtService.ValidateToken(tokenstring)
// if errToken != nil {
// 	panic(errToken.Error())
// }
// claims := token.Claims.(jwt.MapClaims)
// adminId := fmt.Sprintf("%v", claims["user_id"]) //strconv.ParseUint(fmt.Sprintf("%v", claims["user_id"]), 10, 64)
// ******* to be used letter
// if c.userService.IsAllowedToEditUser(getUserDTO.Rfid) {
// 	// result := sc.subCategoryService.UpdateSubCategory(subCategoryUpdatedDTO)
// 	var user models.User = c.userService.FindUserByID(getUserDTO.Rfid)
// 	if (user == models.User{}) {
// 		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
// 		ctx.JSON(http.StatusNotFound, res)
// 		return
// 	} else {
// 		res := helper.BuildResponse(true, "OK", user)
// 		ctx.JSON(http.StatusOK, res)
// 		return
// 	}

// 	// response := helper.BuildResponse(true, "OK", result)
// 	// ctx.JSON(http.StatusOK, response)
// }

// to be used letter
// var user models.User = c.userService.FindUserByID(getUserDTO.Rfid)
// 	if (user == models.User{}) {
// 		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
// 		ctx.JSON(http.StatusNotFound, res)
// 		return
// 	} else {
// 		res := helper.BuildResponse(true, "OK", user)
// 		ctx.JSON(http.StatusOK, res)
// 		return
// 	}

// response := helper.BuildErrorResponse("failed to get user", "False", helper.EmptyObj{})
// ctx.JSON(http.StatusForbidden, response)

// id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
// if err != nil {
// 	res := helper.BuildErrorResponse("No param id w found", err.Error(), helper.EmptyObj{})
// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
// 	return
// }
// var category models.Category = c.categoryService.FindCategoryByID(id)
// if (category == models.Category{}) {
// 	res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
// 	ctx.JSON(http.StatusNotFound, res)
// } else {
// 	res := helper.BuildResponse(true, "OK", category)
// 	ctx.JSON(http.StatusOK, res)
// }
// }

func (c *userController) UpdateUser(ctx *gin.Context) {
	var userUpdatedDTO dto.UpdateDTO
	errDTO := ctx.ShouldBind(&userUpdatedDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	// authHeader := ctx.GetHeader("Authorization")
	// tokenstring := GetTokenString(authHeader)
	// token, errToken := c.jwtService.ValidateToken(tokenstring)
	// if errToken != nil {
	// 	panic(errToken.Error())
	// }
	// claims := token.Claims.(jwt.MapClaims)
	// adminID := fmt.Sprintf("%v", claims["user_id"])
	if c.userService.IsAllowedToEditUser(userUpdatedDTO.ID) {
		// id, errID := strconv.ParseUint(adminID, 10, 64)
		// if errID == nil {
		// 	categoryUpdatedDTO.AdminID = id
		// }
		result := c.userService.UpdateUser(userUpdatedDTO)
		response := helper.BuildResponse(true, "OK", result)
		ctx.JSON(http.StatusOK, response)
	} else {
		response := helper.BuildErrorResponse("Failed to process the request", "You dont have permission to update the user", helper.EmptyObj{})
		ctx.JSON(http.StatusForbidden, response)
	}
}
