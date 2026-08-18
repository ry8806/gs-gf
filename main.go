package main

import (
	"gs-1/handlers"
	"gs-1/repositories"
	"gs-1/services"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

func main() {

	// I'm aware that Go can serve the files directly (which would be embedded in the single executable), however, I'll let nginx do it

	// Could have used uber:FX - but that's possibly outside the scope of this
	var container = dig.New()
	container.Provide(func() *repositories.ProductRepository {
		return &repositories.ProductRepository{}
	})
	container.Provide(func() *repositories.OrderRepository {
		return &repositories.OrderRepository{}
	})

	container.Provide(services.NewProductService)
	container.Provide(services.NewOrderService)

	var router = gin.Default()

	if err := container.Invoke(func(service services.ProductService) {
		router.GET("/products/:id", handlers.GetByProductId(service))
	}); err != nil {
		panic(err)
	}

	if err := container.Invoke(func(service services.OrderService) {
		router.POST("/orders", handlers.CreateOrder(service))
	}); err != nil {
		panic(err)
	}

	router.Run(":9999") // Bind to all interfaces
}
