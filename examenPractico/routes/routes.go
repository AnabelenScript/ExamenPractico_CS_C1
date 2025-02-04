package routes

import (
	"github.com/gin-gonic/gin"
	"examenPractico/controllers"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/productos/create", controllers.CrearProducto)
	router.GET("/productos/getProductos", controllers.ObtenerProductos)
	router.GET("/productos/:id", controllers.ObtenerProductoPorID)
	router.PUT("/productos/:id", controllers.ActualizarProducto)
	router.DELETE("/productos/:id", controllers.EliminarProducto)
	router.GET("/productos/shortpulling", controllers.ShortPulling)
	router.GET("/productos/longpulling", controllers.LongPulling)
}
