package handlers

import (
	"gs-1/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOrder(service services.OrderService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newOrder services.NewOrder

		if err := c.BindJSON(&newOrder); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "post body not in correct format"})
			return
		}

		if newOrder.Quantity > 99999 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "quantity can't be greater than 99999"}) // No thanks, I run this on hardware that's serving 6 other sites 🤣
			return
		}

		summary, err := service.Create(newOrder)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error", "status": "this error has been flagged and escalated with our engineers"}) // LOL!
			return
		}

		c.IndentedJSON(http.StatusCreated, summary)
		return
	}
}
