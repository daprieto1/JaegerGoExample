package presentation

import (
	"fmt"
	"log"

	"github.com/Salaton/tracing/pkg/infrastructure/database"
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

	_ = database.NewPostgresDataStore(db.DB)
}
