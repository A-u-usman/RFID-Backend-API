package main

import (
	"github.com/A-u-usman/RFID-Backend-API.git/config"
	"github.com/A-u-usman/RFID-Backend-API.git/controllers"
	repository "github.com/A-u-usman/RFID-Backend-API.git/repositories"
	"github.com/A-u-usman/RFID-Backend-API.git/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.SetupDatabaseConnection()
	userRepository repository.UserRepository = repository.UserRepositoryImp(db)
	// 	categoryRepository    repository.CategoryRepository     = repository.CategoryRepositoryImp(db)
	// 	subCategoryRepository repository.SubCategoryRepository  = repository.SubCategoryRepositoryImp(db)
	// 	producRepository      repository.ProductRepository      = repository.ProductImp(db)
	// 	jwtService            services.JWTService               = services.JWTServiceImp()
	mailService services.MailService = services.MailServiceImp()
	userService services.UserService = services.UserServiceImp(userRepository)
	// 	categoryService       services.CategoryService          = services.CategoryServiceImp(categoryRepository)
	// 	subCategoryService    services.SubCategoryService       = services.SubCategoryServiceImp(subCategoryRepository, categoryRepository)
	// 	productService        services.ProductService           = services.ProductServiceImp(producRepository, subCategoryRepository, categoryRepository)
	userController controllers.UserController = controllers.UserControllerImp(userService, mailService) // mailService
	// 	categoryController    controllers.CategoryController    = controllers.CategoryControllerImp(categoryService, jwtService)
	// 	subCategoryController controllers.SubCategoryController = controllers.SubCategoryControllerImp(categoryService, subCategoryService, jwtService)
	// 	productController     controllers.ProductController     = controllers.ProductControllerImp(categoryService, subCategoryService, productService, jwtService)
)

func main() {

	r := gin.Default()

	r.Static("/templates", "./templates/")
	r.LoadHTMLGlob("templates/*.html")
	userViews := r.Group("/view")
	{
		userViews.GET("/users", userController.ShowUsers)
		userViews.GET("/access-activity", userController.ShowUsersActivity)
		userViews.GET("/add-user", userController.ShowUserForm)
		userViews.POST("/add-user", userController.SaveUser)
		userViews.GET("enable-disable/:rfid", userController.EnableDisableUser)
	}
	userRoutes := r.Group("api/rfid")
	{
		// authRoutes.POST("/login", authController.Login)
		userRoutes.POST("/add", userController.Register)
		userRoutes.GET("/", userController.AllUsers)
		// userRoutes.POST("/create", userController.InsertCategory)
		userRoutes.GET("/get/:rfid", userController.FindUserByID)
		userRoutes.PUT("/update", userController.UpdateUser)
		userRoutes.DELETE("/del", userController.DelUser)
		// 	authRoutes.POST("/forgotpassword", authController.ForgotPassword)
		// 	//authRoutes.POST("/send-otp", authController.SendOTP)
		// 	//authRoutes.POST("/verify-otp", authController.VerifyOTP)
		// 	//authRoutes.POST("/inviteuser", authController.InviteUser)
	}

	// authRoutes2 := r.Group("api/auth", middleware.Authorize(jwtService))
	// {
	// 	authRoutes2.POST("/resetpassword", authController.ResetPassword)
	// 	authRoutes2.POST("/changepassword", authController.ChangePassword)
	// 	authRoutes.POST("/send-otp", authController.SendOTP)
	// 	authRoutes.POST("/verify-otp", authController.VerifyOTP)
	// 	// authRoutes2.POST("/blockaccount", authController.BlockAccount
	// 	// 	authRoutes2.POST("/inviteuser", authController.InviteUser)
	// }

	// categoryRoutes := r.Group("api/categories", middleware.Authorize(jwtService))
	// {
	// 	categoryRoutes.GET("/", categoryController.AllCategory)
	// 	categoryRoutes.POST("/create", categoryController.InsertCategory)
	// 	categoryRoutes.POST("/get", categoryController.FindCategoryByID)
	// 	categoryRoutes.PUT("/update", categoryController.UpdateCategory)
	// 	categoryRoutes.DELETE("/del", categoryController.DelCategory)
	// }

	// subCategoryRoutes := r.Group("api/categories/subcategories", middleware.Authorize(jwtService))
	// {
	// 	subCategoryRoutes.GET("/", subCategoryController.AllSubCategory)
	// 	subCategoryRoutes.POST("/create", subCategoryController.InsertSubCategory)
	// 	subCategoryRoutes.POST("/get", subCategoryController.FindSubCategoryByID)
	// 	subCategoryRoutes.PUT("/update", subCategoryController.UpdateSubCategory)
	// 	subCategoryRoutes.DELETE("/del", subCategoryController.DelSubCategory)
	// }
	// //product url
	// productRoutes := r.Group("api/categories/subcategories/product", middleware.Authorize(jwtService))
	// {
	// 	productRoutes.GET("/", productController.AllProduct)
	// 	productRoutes.POST("/create", productController.InsertProduct)
	// 	productRoutes.POST("/get", productController.FindProductByID)
	// 	productRoutes.PUT("/update", productController.UpdateProduct)
	// 	productRoutes.DELETE("/del", productController.DelProduct)
	// }

	// swggerRoute := r.Group("api")
	// {
	// 	swggerRoute.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// }
	// port := os.Getenv("PORT")

	// if port == "" {
	// 	port = "8080"
	// }
	// r.Run("0.0.0.0:" + port)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	// //online host = https://school-lite.fly.dev/
}
