package product

import (
	"log"

	"github.com/Sadham-Hussian/go-learnings/gofiber-grpc/server/database"
)

// CreateProductQuery function to add product in the database
func CreateProductQuery(product *database.ProductItem) error {
	db, err := database.SQLConnection()
	if err != nil {
		log.Println(err)
		return err
	}

	result := db.Create(&product)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// ReadProductQuery function to read a product in the database
func ReadProductQuery(id int32) (database.ProductItem, error) {
	db, err := database.SQLConnection()
	if err != nil {
		log.Println(err)
		return database.ProductItem{}, err
	}

	var product database.ProductItem
	result := db.First(&product, id)
	if result.Error != nil {
		return database.ProductItem{}, result.Error
	}

	return product, nil
}

// UpdateProductQuery function to update product in the database
func UpdateProductQuery(product *database.ProductItem) error {
	db, err := database.SQLConnection()
	if err != nil {
		log.Println(err)
		return err
	}

	result := db.Model(&product).Updates(database.ProductItem{Name: product.Name, Type: product.Type, Prize: product.Prize})
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// DeleteProductQuery function to delete product in the database
func DeleteProductQuery(id int32) error {
	db, err := database.SQLConnection()
	if err != nil {
		log.Println(err)
		return err
	}

	result := db.Delete(&database.ProductItem{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// ReadProductsQuery function to read all products in the database
func ReadProductsQuery() ([]database.ProductItem, error) {
	db, err := database.SQLConnection()
	if err != nil {
		log.Println(err)
		return []database.ProductItem{}, err
	}

	var products []database.ProductItem
	result := db.Find(&products)
	if result.Error != nil {
		return []database.ProductItem{}, result.Error
	}

	return products, nil
}
