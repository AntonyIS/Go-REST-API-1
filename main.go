package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type techFirm struct {
	Location string `json "location"`
	Name     string `json "name"`
	CEO      string `json "ceo"`
	ID       int    `json "id"`
}

// Define companies slice[dynamic array]
var companies = []techFirm{
	{ID: 1, Location: "Menlo Park, USA", Name: "Facebook", CEO: "Mark Zuckerberg"},
	{ID: 2, Location: "Palo Alto, USA", Name: "Tesla", CEO: "Elon Musk"},
	{ID: 3, Location: "Seattle , USA", Name: "Amazon", CEO: "Andy Jassy"},
	{ID: 4, Location: "Redmond USA", Name: "MicroSoft", CEO: "Satya Nadella"},
	{ID: 5, Location: "Mountain View, USA", Name: "Google", CEO: "Sundra Pichai"},
	{ID: 6, Location: "Cupertino", Name: "Apple", CEO: "Tim Cook"},
}

func main() {
	router := gin.Default()

	router.GET("/api/v1/companies", GetCompanies)
	router.GET("/api/v1/companies/:id", GetCompany)
	router.POST("/api/v1/companies", PostCompany)
	router.PUT("/api/v1/companies/:id", EditCompany)
	router.DELETE("/api/v1/companies/:id", DeleteCompany)
	// Run server
	router.Run("localhost:8080")
}

func GetCompanies(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, companies)
}

func GetCompany(c *gin.Context) {
	requestID := c.Param("id")
	id, err := strconv.Atoi(requestID)
	if err != nil {
		log.Fatalf("ERROR =====> %v", err)
	}
	for _, company := range companies {
		if company.ID == id {
			c.IndentedJSON(http.StatusOK, company)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"statusCode": 404, "message": "Company not found"})
}

func PostCompany(c *gin.Context) {
	var newCompany techFirm
	if err := c.BindJSON(&newCompany); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
	}
	companies = append(companies, newCompany)
	c.IndentedJSON(http.StatusCreated, newCompany)
}

func EditCompany(c *gin.Context) {
	var newCompany techFirm
	if err := c.BindJSON(&newCompany); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
	}
	for index, company := range companies {
		if company.ID == newCompany.ID {
			company.Location = newCompany.Location
			company.Name = newCompany.Name
			company.CEO = newCompany.CEO
			// Insert into company slice
			companies = append(companies[:index], companies[index+1:]...)
			companies = append(companies, newCompany)
			c.IndentedJSON(http.StatusCreated, newCompany)
			return
		}

	}
}

func DeleteCompany(c *gin.Context) {
	requestID := c.Param("id")
	id, err := strconv.Atoi(requestID)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Internal server error"})
	}

	for index, company := range companies {
		if company.ID == id {
			companies = append(companies[:index], companies[index+1:]...)
		}
	}
	c.IndentedJSON(http.StatusOK, companies)
}
