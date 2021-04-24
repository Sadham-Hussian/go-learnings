package services

import (
	"log"

	"github.com/Sadham-Hussian/go-learnings/crud/app/models"
	"github.com/Sadham-Hussian/go-learnings/crud/platform/database"
)

// CreateProduct service to create a product
func CreateProduct(product *models.Product) (interface{}, int) {
	response := make(map[string]interface{})
	db, err := database.SQLConnection()
	if err != nil {
		log.Println(err)
		response["error"] = "internal server error."
		response["status"] = "Fail"
		return response, 500
	}

	result := db.Create(&product)
	if result.Error != nil {
		response["error"] = "Invalid data"
		response["status"] = "Fail"
		return response, 400
	}

	response["status"] = "Ok"
	return response, 201
}

// GetProductByID service to get a product by id
func GetProductByID(id int) (interface{}, int) {
	response := make(map[string]interface{})
	db, err := database.SQLConnection()
	if err != nil {
		log.Println(err)
		response["error"] = "internal server error"
		response["status"] = "Fail"
		return response, 500
	}

	var product models.Product
	result := db.Find(&product, id)
	if result.Error != nil || result.RowsAffected == 0 {
		response["error"] = "Invalid id"
		response["status"] = "Fail"
		return response, 204
	}

	response["status"] = "Ok"
	response["product"] = product
	return response, 200
}

// GetAllProducts function to read all products
func GetAllProducts() (interface{}, int) {
	response := make(map[string]interface{})
	db, err := database.SQLConnection()
	if err != nil {
		log.Println(err)
		response["error"] = "internal server error"
		response["status"] = "Fail"
		return response, 500
	}

	var products []models.Product
	result := db.Find(&products)
	if result.Error != nil {
		response["error"] = "internal server error"
		response["status"] = "Fail"
		return response, 500
	}

	response["products"] = products
	response["status"] = "Ok"
	return response, 200
}

// DeleteProductByID service to delete a product by id
func DeleteProductByID(id int) (interface{}, int) {
	response := make(map[string]interface{})
	db, err := database.SQLConnection()
	if err != nil {
		log.Println(err)
		response["error"] = "internal server error"
		response["status"] = "Fail"
		return response, 500
	}

	result := db.Delete(&models.Product{}, id)
	if result.Error != nil {
		response["error"] = "Not able to delete"
		response["status"] = "Fail"
		return response, 500
	}

	response["status"] = "Ok"
	return response, 200
}
