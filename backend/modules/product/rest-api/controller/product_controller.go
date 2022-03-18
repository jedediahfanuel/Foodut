package controller

import (
	"encoding/json"
	"net/http"

	model "github.com/Foodut/backend/modules/product/domain/model"
	srvc "github.com/Foodut/backend/modules/product/domain/service"
	rspn "github.com/Foodut/backend/responses"
)

/**
  Regular "GET" without any extra queries
  Retrieve all registered products from database.
*/
func GetAllProducts(writer http.ResponseWriter, req *http.Request) {

	// Check id query
	productId := req.URL.Query()["id"]

	// Get list of product object
	var products []model.Product = srvc.SearchById(productId)

	// Set response
	var response rspn.Response
	if len(products) > 0 {
		response.Response_200(products)
	} else {
		response.Response_204()
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(response)
}

/**
  Extra "query" using "product name".
  Retrieve all matched product by product name.
  If it is none exact name, then use 'LIKE' sql.
*/
func GetProductByName(writer http.ResponseWriter, req *http.Request) {

	// Check product_name query
	nameToFind := req.URL.Query()["product_name"]

	// Get products using query
	var products []model.Product = srvc.SearchByName(nameToFind)

	// Set response
	var response rspn.Response
	if len(products) > 0 {
		response.Response_200(products)
	} else {
		response.Response_204()
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(response)
}

func DeleteProductById(writer http.ResponseWriter, req *http.Request) {
	// Check id query
	productId := req.URL.Query()["id"]

	deleteErr := srvc.DeleteById(productId)

	var response rspn.Response
	if deleteErr == nil {
		response.Response_200("data has been deleted")
	} else {
		response.Response_400(deleteErr)
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(response)
}
