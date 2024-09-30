package main

// load required packages
import (
	"fmt"
	"log"
	"os"

	"github.com/alifhaider/carehub-backend/controller"
	"github.com/alifhaider/carehub-backend/database"
	"github.com/alifhaider/carehub-backend/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// load environment file
	loadEnv()
	// load database configuration and connection
	loadDatabase()

		// Set Gin mode to release
		gin.SetMode(gin.ReleaseMode)

	// Create a new router
	router := gin.New()

	// Configure trusted proxies (example: trusting only localhost)
	router.SetTrustedProxies([]string{os.Getenv("DB_HOST")}) // Adjust as necessary
	// start the server
	serveApplication()

}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Println(".env file loaded successfully")
}

func loadDatabase() {
    database.InitDb()
    database.Db.AutoMigrate(&models.Role{})
    database.Db.AutoMigrate(&models.User{})
    seedData()
}

func seedData() {
    var roles = []models.Role{{Name: "doctor", Description: "Doctor role"}, {Name: "user", Description: "User role"}}
    var user = []models.User{{Username: "alif", Email: "alif.dev", Password: "222222", RoleID: 1}}
    database.Db.Save(&roles)
    database.Db.Save(&user)
}

func serveApplication() {
    router := gin.Default()
    authRoutes := router.Group("/auth/user")
        // registration route
    authRoutes.POST("/register", controller.Register)
        // login route
    authRoutes.POST("/login", controller.Login)

    router.Run(":8000")
    fmt.Println("Server running on port 8000")
}