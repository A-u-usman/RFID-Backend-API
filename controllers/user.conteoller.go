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
	AllUsers(ctx *gin.Context)
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
	// mailService services.MailService
}

// mailService services.MailService to be addede letter
func UserControllerImp(userService services.UserService) UserController {
	return &userController{
		userService: userService,
		// jwtService:  jwtService,
		// mailService: mailService,
	}
}

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
}

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
		response := getstruct{Message: "true"}
		ctx.JSON(http.StatusCreated, response)
	} else {
		// createdUser := c.userService.CreateUser(registerDTO)
		// // token := c.jwtService.GenerateToken(strconv.FormatUint(createdUser.ID, 10))
		// // createdUser.Token = token
		// response := helper.BuildResponse(true, "OK!", createdUser)
		// ctx.JSON(http.StatusCreated, response)

		// response := helper.BuildErrorResponse("false", "Invalid", helper.EmptyObj{})
		response := getstruct{Message: "false"}
		ctx.JSON(http.StatusConflict, response)

	}
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
}

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
