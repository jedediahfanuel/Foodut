package repository

import (
	dbController "github.com/Foodut/backend/database"
	model "github.com/Foodut/backend/modules/product/domain/model"
	"gorm.io/gorm"
)

func GetProductsAssociation(products []model.Product) {
	// For each product
	for i := 0; i < len(products); i++ {
		// with picture
		GetOneProductAssociation(&products[i])
	}
}

func GetOneProductAssociation(product *model.Product) {
	// Check connection
	con := dbController.GetConnection()

	// Picture Association
	con.Model(&product).Association("Picture").Find(&product.Picture)
}

func ReadAllProducts(productId []string) []model.Product {
	// Check connection
	con := dbController.GetConnection()

	// Get all products from database
	var products []model.Product

	// Extra query by Id
	if productId != nil {
		con.Find(&products, productId[0])
	} else {
		con.Find(&products)
	}

	return products
}

func ReadProductsByName(productName []string) []model.Product {
	// Check connection
	con := dbController.GetConnection()

	// Get product from database, filter by product_name
	var products []model.Product
	if productName != nil {
		con.Where("product_name = ?", productName[0]).Find(&products)
	}

	return products
}

func ReadProductsByNameAlike(productName []string) []model.Product {
	// Check connection
	con := dbController.GetConnection()

	// Get product from database, filter by name LIKE
	var products []model.Product
	if productName != nil {
		like := "%" + productName[0] + "%"
		con.Where("product_name LIKE ?", like).Find(&products)
	}

	return products
}

func ReadProductsByCategoryId(categoryId int) []model.Product {
	// Check connection
	con := dbController.GetConnection()

	// Get product from database, filter by name LIKE
	var products []model.Product
	if categoryId > 0 {
		con.Where("category_id = ?", categoryId).Find(&products)
	}

	return products
}

func ReadProductsByCategoryName(categoryName []string) []model.Product {
	// Get product from database, filter by name LIKE
	var products []model.Product
	if categoryName != nil {
		category := ReadCategoryByName(categoryName)
		products = ReadProductsByCategoryId(category.ID)
	}

	return products
}

func DeleteProductById(productId string) *gorm.DB {
	// Check connection
	con := dbController.GetConnection()

	// Delete Product on database by its id
	var product model.Product

	result := con.Delete(&product, productId)

	return result
}

func CreateProduct(product model.Product) *gorm.DB {
	// Check connection
	con := dbController.GetConnection()

	// Insert object to database
	result := con.Create(&product)

	return result
}

func ReadProductById(productId int) model.Product {
	// Check connection
	con := dbController.GetConnection()

	var product model.Product

	con.Find(&product, productId)

	return product
}

func ReadProductStock(productId int) int {
	// Check connection
	con := dbController.GetConnection()

	var product model.Product
	con.Select("product_stock").Find(&product, productId)

	return product.ProductStock
}

func UpdateProductStock(productId int, newStock int) *gorm.DB {
	// Check connection
	con := dbController.GetConnection()

	// Update product stock using raw SQL
	result := con.Exec(
		"UPDATE `products` "+
			"SET `product_stock`= ? "+
			"WHERE id = ?",
		newStock,
		productId)

	return result
}

func UpdateProduct(product model.Product) *gorm.DB {
	// Check connection
	con := dbController.GetConnection()

	// Save updated prdouct data
	result := con.Save(&product)

	return result
}
