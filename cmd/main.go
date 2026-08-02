package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "hospital-middleware/docs"
	"hospital-middleware/pkg/database"

	"hospital-middleware/internal/service"
	httpHandler "hospital-middleware/internal/delivery/http"
)

const (
	host     = "postgres"        // or the Docker service name if running in another container
	port     = "5432"             // default PostgreSQL port
	user     = "postgres"         // as defined in docker-compose.yml
	password = "postgrespassword" // as defined in docker-compose.yml
	dbname   = "agnos_db"         // as defined in docker-compose.yml
)

// @title           Hospital Middleware API
// @version         1.0
// @description     Middleware system for searching patients across hospitals.
// @host            localhost:8080
// @BasePath        /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer " followed by a space and your JWT token.
func main() {
	db, err := database.InitDB(host, port, user, password, dbname)
	if err != nil {
		log.Fatalf("Could not initialize database: %v", err)
	}

	_ = db

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/ping", ping)

	staffService := service.NewStaffService(db)
	staffHandler := httpHandler.NewStaffHandler(staffService)

	patientService := service.NewPatientService(db)
	patientHandler := httpHandler.NewPatientHandler(patientService)

	router.POST("/staff/create", staffHandler.CreateStaff)
	router.POST("/staff/login", staffHandler.Login)

	protected := router.Group("/")
	protected.Use(httpHandler.AuthMiddleware())
	{
		protected.POST("/patient/search", patientHandler.SearchPatients)
	}

	router.GET("/patient/search/:id", httpHandler.MockHospitalASearchHandler)

	router.Run(":8080")
}

// ping godoc
// @Summary Health check
// @Description Check API status
// @Tags system
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /ping [get]
func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}