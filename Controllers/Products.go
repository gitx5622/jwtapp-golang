package Controllers

import (
	"github.com/gin-gonic/gin"
	"jwtapp/Models"
	"net/http"
)

// Get all the Products
func GetProducts(c *gin.Context) {
	// Get all the products
	var product []Models.Product
	err := Models.GetAllProducts(&product)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"message": "Unable to get products",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":200,
			"message": "success",
			"products": product,
		})
	}
}


// Create a product
func CreateProduct(c *gin.Context) {
	var product Models.Product
	_ = c.BindJSON(&product)
	err := Models.CreateAProduct(&product)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"message": "Unable to create products",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":200,
			"message": "Product created",
			"products": product,
		})
	}

}


// Get a Single product
func GetProduct(c *gin.Context) {
	id := c.Params.ByName("id")
	var product Models.Product
	err := Models.GetAProduct(&product, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":200,
			"message": "success",
			"products": product,
		})
	}
}


func UpdateProduct(c *gin.Context) {
	var product Models.Product
	id := c.Params.ByName("id")
	err := Models.GetAProduct(&product, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"message": "Unable to get product to be updated",
		})
	}
	_ = c.BindJSON(&product)
	err = Models.UpdateAProduct(&product, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"message": "Unable to update product",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":200,
			"message": "Product updated",
			"products": product,
		})
	}
}



