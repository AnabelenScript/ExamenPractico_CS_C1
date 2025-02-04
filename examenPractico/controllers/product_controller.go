package controllers

import (
	"sync"
	"time"
	"github.com/gin-gonic/gin"
	"examenPractico/models"
	"log"
)

var (
	productos            []models.Producto
	productosMu          sync.Mutex
	ultimaActualizacion  time.Time
)

func CrearProducto(c *gin.Context) {
	var nuevoProducto models.Producto
	if err := c.ShouldBindJSON(&nuevoProducto); err != nil {
		c.JSON(400, gin.H{"error": "Error al decodificar el cuerpo de la solicitud"})
		return
	}

	productosMu.Lock()
	productos = append(productos, nuevoProducto)
	ultimaActualizacion = time.Now()
	productosMu.Unlock()

	c.JSON(201, gin.H{"message": "Producto creado", "producto": nuevoProducto})
}

func ObtenerProductos(c *gin.Context) {
	productosMu.Lock()
	defer productosMu.Unlock()

	c.JSON(200, gin.H{"productos": productos})
}

func ObtenerProductoPorID(c *gin.Context) {
	id := c.Param("id")

	productosMu.Lock()
	defer productosMu.Unlock()

	for _, producto := range productos {
		if producto.ID == id {
			c.JSON(200, gin.H{"producto": producto})
			return
		}
	}

	c.JSON(404, gin.H{"error": "Producto no encontrado"})
}

func ActualizarProducto(c *gin.Context) {
	id := c.Param("id")

	var productoActualizado models.Producto
	if err := c.ShouldBindJSON(&productoActualizado); err != nil {
		c.JSON(400, gin.H{"error": "Error al decodificar el cuerpo de la solicitud"})
		return
	}

	productosMu.Lock()
	defer productosMu.Unlock()

	for i, producto := range productos {
		if producto.ID == id {
			productos[i] = productoActualizado
			ultimaActualizacion = time.Now()
			c.JSON(200, gin.H{"message": "Producto actualizado", "producto": productoActualizado})
			return
		}
	}

	c.JSON(404, gin.H{"error": "Producto no encontrado"})
}

func EliminarProducto(c *gin.Context) {
	id := c.Param("id")

	productosMu.Lock()
	defer productosMu.Unlock()

	for i, producto := range productos {
		if producto.ID == id {
			productos = append(productos[:i], productos[i+1:]...)
			ultimaActualizacion = time.Now()
			c.JSON(200, gin.H{"message": "Producto eliminado"})
			return
		}
	}

	c.JSON(404, gin.H{"error": "Producto no encontrado"})
}

func ShortPulling(c *gin.Context) {
	productosMu.Lock()
	defer productosMu.Unlock()
	log.Printf("Productos consultados: %+v\nÚltima actualización: %s\n", productos, ultimaActualizacion)

	c.JSON(200, gin.H{
		"productos":          productos,
		"ultima_actualizacion": ultimaActualizacion,
	})
}

func LongPulling(c *gin.Context) {
	productosMu.Lock()
	ultimaModificacion := ultimaActualizacion
	productosMu.Unlock()

	timeout := time.After(30 * time.Second)
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			productosMu.Lock()
			if !ultimaActualizacion.Equal(ultimaModificacion) {
				productosMu.Unlock()
				log.Printf("Cambio detectado en los productos: %+v\nÚltima actualización: %s\n", productos, ultimaActualizacion)
				c.JSON(200, gin.H{
					"productos":          productos,
					"ultima_actualizacion": ultimaActualizacion,
				})
				return
			}
			productosMu.Unlock()
		case <-timeout:
			c.JSON(200, gin.H{"message": "No hay cambios"})
			return
		}
	}
}
