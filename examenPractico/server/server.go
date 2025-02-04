package server

import (
	"github.com/gin-gonic/gin"
	"time"
	"examenPractico/routes"
	"examenPractico/controllers"
)


func StartServer(port string) {
	router := gin.Default()
	routes.SetupRoutes(router)
	go router.Run(":" + port)
}

func StartPolling() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			
			go controllers.ShortPulling(nil)
			go controllers.LongPulling(nil)
		}
	}
}
