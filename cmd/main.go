package main

import (
	"fmt"
	"time"
	"xsis/assignment/internal/app/controller"
	"xsis/assignment/internal/app/model"
	"xsis/assignment/internal/app/repository"
	"xsis/assignment/internal/app/service"
	"xsis/assignment/internal/pkg/config"
	"xsis/assignment/internal/pkg/db"
	"xsis/assignment/internal/pkg/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	cfg    config.Config
	DBConn *gorm.DB
)

func init() {

	configLoad, err := config.LoadConfig(".")
	if err != nil {
		log.Panic("cannot load app config")
	}
	cfg = configLoad

	db, err := db.ConnectDB(cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)
	if err != nil {
		log.Panic("db not established")
	}
	DBConn = db

	// check existing constraint if no longer necessary
	// if db.Migrator().HasConstraint(&repo.Products{}, "products_name_key") {
	// 	db.Migrator().DropConstraint(&repo.Products{}, "products_name_key")
	// }
	db.AutoMigrate(
		&model.Movies{},
	)

	// Setup logrus
	logLevel, err := log.ParseLevel("debug")
	if err != nil {
		logLevel = log.InfoLevel
	}

	log.SetLevel(logLevel)                 // appyly log level
	log.SetFormatter(&log.JSONFormatter{}) // define format using json

}

const (
	limit  = 10
	offset = 0
	asc    = 0
)

func main() {

	r := gin.Default()

	// implement middleware
	r.Use(
		middleware.LoggingMiddleware(),
		middleware.RecoveryMiddleware(),
		cors.New(cors.Config{
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{"OPTIONS", "GET", "POST", "PATCH", "DELETE"},
			AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "authorization"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			AllowOriginFunc: func(origin string) bool {
				return origin == "http://localhost"
			},
			MaxAge: 12 * time.Hour,
		}),
	)

	imageService := service.NewUploadService(cfg.CloudinaryName, cfg.CloudinaryAPIKey, cfg.CloudinarySecretKey, cfg.CloudinaryDir)

	movieRepo := repository.NewMoviesRepo(DBConn)
	movieService := service.NewMoviesService(movieRepo, imageService)
	movieController := controller.NewMoviesController(movieService)

	r.POST("/movies", movieController.Create)
	r.GET("/movies", middleware.PaginationMiddleware(offset, limit, asc), movieController.GetMovies)
	r.GET("/movies/:id", movieController.GetMoviesByID)
	r.PATCH("/movies/:id", movieController.UpdateMoviesByID)
	r.DELETE("/movies/:id", movieController.DeleteMoviesByID)

	appPort := fmt.Sprintf(":%s", cfg.ServerPort)

	r.Run(appPort)
}
