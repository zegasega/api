package controllers

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"project2/database"
	"project2/models"

	"github.com/gin-gonic/gin"
)

func GetPeople(c *gin.Context) {
	rows, err := database.DB.Query("SELECT id, name, age FROM people")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var people []models.Person
	for rows.Next() {
		var person models.Person
		err := rows.Scan(&person.ID, &person.Name, &person.Age)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		people = append(people, person)
	}
	c.JSON(http.StatusOK, people)
}

func GetPersonByID(c *gin.Context) {
	// URL parametresinden ID'yi alır ve integer'a çevirir.
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Veritabanı sorgusu, belirli bir ID'ye sahip kişiyi seçer.
	row := database.DB.QueryRow("SELECT id, name, age FROM people WHERE id = ?", intID)

	var person models.Person
	err = row.Scan(&person.ID, &person.Name, &person.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			// Eğer sonuç yoksa, 404 HTTP durum kodu ile hata döndürülür.
			c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
		} else {
			// Diğer hatalar için 500 HTTP durum kodu ile hata döndürülür.
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	// Kişi JSON formatında döndürülür.
	c.JSON(http.StatusOK, person)
}

func DeleteUserbyID(c *gin.Context) {
	id := c.Param("id")

	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
	}

	row := database.DB.QueryRow("SELECT * FROM people WHERE id = ?", intID)

	var person models.Person
	err = row.Scan(&person.ID, &person.Name, &person.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			// Eğer sonuç yoksa, 404 HTTP durum kodu ile hata döndürür
			c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
		} else {
			// Diğer hatalar için 500 HTTP durum kodu ile hata döndürür
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	// Kişi JSON formatında döndürülür
	c.JSON(http.StatusOK, person)
}

func InsertPerson(c *gin.Context) {
	var newPerson models.Person
	if err := c.ShouldBindJSON(&newPerson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sqlStatement := `INSERT INTO people (name, age) VALUES ($1, $2)`
	_, err := database.DB.Exec(sqlStatement, newPerson.Name, newPerson.Age)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Person inserted successfully!"})
}

func InsertSamplePerson(name string, age int) {
	sqlStatement := `INSERT INTO people (name, age) VALUES ($1, $2)`
	_, err := database.DB.Exec(sqlStatement, name, age)
	if err != nil {
		log.Fatalf("Unable to insert data: %v\n", err)
	}
	log.Println("Data inserted successfully!")
}
