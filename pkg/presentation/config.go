package presentation

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func StartServer(port int) {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	SetupRoutes(r)

	addr := fmt.Sprintf(":%d", port)

	if err := r.Run(addr); err != nil {
		log.Fatalf("an error occured while running the server: %v", err)
	}
}

func SetupRoutes(r *gin.Engine) {

}
