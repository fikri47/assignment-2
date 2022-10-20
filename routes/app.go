package routes

import (
	"assignment-2/controller"
	"assignment-2/databases"
	"fmt"

	"github.com/gin-gonic/gin"
)

const PORT = ":8080"

func StartApp() {
	route := gin.Default()
	databases.StartDB()
	orderRoute := route.Group("/orders")
	{
		orderRoute.POST("/", controller.CreateOrder)
		orderRoute.GET("/", controller.GetOrders)
		orderRoute.PUT("/:OrderID", controller.UpdateOrder)
		orderRoute.DELETE("/:OrderID", controller.DeleteOrder)
	}

	fmt.Println("Server running on port ==>", PORT)
	route.Run(PORT)
}
