package handlers

import (
	"errors"
	"gs-1/repositories"
	"gs-1/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetByProductId locates the product by it's Id
func GetByProductId(service services.ProductService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		product, err := service.Get(id)
		if err != nil {
			if errors.Is(err, repositories.ErrProductNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
				return
			}

			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}

		c.JSON(http.StatusOK, product)
	}
}
