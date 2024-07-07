package controllers

import (
	"net/http"
	"project2/database"
	"project2/models"

	"github.com/gin-gonic/gin"
)

func GetFruits(c *gin.Context) {
	rows, err := database.DB.Query("SELECT id, name, price, amount FROM fruits")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var fruits []models.Fruit
	for rows.Next() {
		var fruit models.Fruit
		err := rows.Scan(&fruit.ID, &fruit.Name, &fruit.Price, &fruit.Amount)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		fruits = append(fruits, fruit)
	}

	if err = rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, fruits)
}
func InsertFruit(c *gin.Context) {
	var newFruit models.Fruit

	if err := c.ShouldBindJSON(&newFruit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err})
		return
	}
	sqlStatement := `INSERT INTO fruits (name, price, amount) VALUES ($1, $2, $3)`
	_, err := database.DB.Exec(sqlStatement, newFruit.Name, newFruit.Price, newFruit.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "fruit inserted succesfuly"})

}
