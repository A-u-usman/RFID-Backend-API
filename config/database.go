package config

import (
	models "github.com/A-u-usman/RFID-Backend-API.git/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ****************************************
// MySQ: db configuration
// *************************************
// func SetupDatabaseConnection() *gorm.DB {
// 	err := godotenv.Load()
// 	if err != nil {
// 		panic("Failed to load env file")

// 	}

// 	//*********************************
// 	//localhost MySQL db configuration
// 	//*********************************
// 	dbDatabase := os.Getenv("DB_DATABASE")
// 	dbUser := os.Getenv("DB_USER")
// 	dbPassword := os.Getenv("DB_PASSWORD")
// 	dbHost := os.Getenv("DB_HOST")
// 	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbDatabase)

// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic("Failed to setup database connection")
// 	}
// 	db.AutoMigrate(&models.User{}, &models.UserActivityLog{})
// 	return db
// }

// func CloseDatabase(db *gorm.DB) {
// 	dbMySQL, err := db.DB()
// 	if err != nil {
// 		panic("Failed")
// 	}
// 	dbMySQL.Close()
// }

//***********End*****************

//***************************
//Sqlite database connection
//***************************

// func SetupDatabaseConnection() *gorm.DB {
// 	// err := godotenv.Load()
// 	// if err != nil {
// 	// 	panic("Failed to load env file")

// 	// }

// 	//*********************************
// 	// online sqlite3 db configuration
// 	//*********************************

// 	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
// 	if err != nil {
// 		panic("Failed to setup database connection")

// 	}
// 	db.AutoMigrate(&models.Teacher{})
// 	return db
// }

// func CloseDatabase(db *gorm.DB) {
// 	dbSQlite3, err := db.DB()
// 	if err != nil {
// 		panic("Failed")
// 	}
// 	dbSQlite3.Close()
// }

//***********End*******************

// ****************************
// Postgre database connection
// ****************************
func SetupDatabaseConnection() *gorm.DB {
	// err := godotenv.Load()
	// if err != nil {
	// 	panic("Failed to load env file")

	// }

	//*********************************
	//Postgres db configuration
	//*********************************
	// dbDatabase := os.Getenv("DB_DATABASE")
	// dbUser := os.Getenv("DB_USER")
	// dbPassword := os.Getenv("DB_PASSWORD")
	// dbHost := os.Getenv("DB_HOST")
	//dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai", dbHost, dbUser, dbPassword,  dbDatabase)
	dsn := "postgres://zmrnufln:NJuQE21hxQPS3OMEsOlQMdOrhYLAx1ff@rosie.db.elephantsql.com/zmrnufln"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to setup database connection")
	}
	db.AutoMigrate(&models.User{}, &models.UserActivityLog{})
	return db
}

func CloseDatabase(db *gorm.DB) {
	dbPostgresSQL, err := db.DB()
	if err != nil {
		panic("Failed")
	}
	dbPostgresSQL.Close()
}

//***********End*****************
