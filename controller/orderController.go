package controller

import (
	"assignment-2/databases"
	"assignment-2/entity"
	"assignment-2/helpers"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

var (
	appJSON = "application/json"
)

func CreateOrder(c *gin.Context) {
	db := databases.GetDB()
	contentType := helpers.GetContentType(c)
	Order := entity.Order{}

	if contentType == appJSON {
		c.ShouldBindJSON(&Order)
	} else {
		c.ShouldBind(&Order)
	}

	err := db.Debug().Create(&Order).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Order)
}

func GetOrders(c *gin.Context) {
	db := databases.GetDB()
	Orders := []entity.Order{}

	err := db.Preload("Items").Find(&Orders).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, Orders)

}

func DeleteOrder(c *gin.Context) {
	db := databases.GetDB()
	orderID, err := strconv.Atoi(c.Param("OrderID"))
	orders := entity.Order{}
	items := entity.Item{}

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    "500",
			"message": "Invalid param orderId",
		})
		return
	}

	err = db.First(&orders, orderID).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Data Not Found",
		})
		return
	}

	err = db.Clauses(clause.Returning{}).Where("order_id = ?", orderID).Delete(&items).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	} else {
		err := db.Delete(&orders).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error":   "error deleting",
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "success delete order",
		})
	}

}

func UpdateOrder(c *gin.Context) {
	db := databases.GetDB()
	contentType := helpers.GetContentType(c)
	OrderID, err := strconv.Atoi(c.Param("OrderID"))
	UpdateOrder := entity.Order{}

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    "500",
			"message": "Invalid param orderId",
		})
		return
	}

	if contentType == appJSON {
		c.ShouldBindJSON(&UpdateOrder)
	} else {
		c.ShouldBind(&UpdateOrder)
	}

	for i := range UpdateOrder.Items {
		err = db.Model(&UpdateOrder.Items[i]).Where("item_id=?", UpdateOrder.Items[i].ItemID).Updates(&UpdateOrder.Items[i]).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error":   "Error updating item",
				"message": err.Error(),
			})
			return
		}
	}

	err = db.Model(&UpdateOrder).Where("order_id=?", OrderID).Omit("Items").Updates(&UpdateOrder).Error

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   "Error updating order",
			"message": err.Error(),
		})
		return
	}

	err = db.Preload("Items").Where("order_id=?", OrderID).Find(&UpdateOrder).Error

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   "Error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfuly update data",
		"data":    UpdateOrder,
	})
}
