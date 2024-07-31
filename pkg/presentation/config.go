package presentation

import (
	"fmt"
	"log"

	"github.com/daprieto1/tracing/pkg/infrastructure/database"
	"github.com/daprieto1/tracing/pkg/presentation/handlers"
	"github.com/daprieto1/tracing/pkg/usecase"
	"github.com/gin-gonic/gin"
)

func StartServer(port int) {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.UseRawPath = true
	r.UnescapePathValues = false

	SetupRoutes(r)

	addr := fmt.Sprintf(":%d", port)

	if err := r.Run(addr); err != nil {
		log.Fatalf("an error occured while running the server: %v", err)
	}
}

func SetupRoutes(r *gin.Engine) {
	db, err := database.NewPGInstance()
	if err != nil {
		log.Fatalf("an error occured connecting to the database: %v", err)
	}

	store := database.NewPostgresDataStore(db.DB)
	productUseCase := usecase.NewUseCaseImplementation(store)
	handler := handlers.NewHandlersImplementation(*productUseCase)

	v1 := r.Group("/api/v1")
	{
		v1.POST("/product", handler.CreateProduct)
	}
}
