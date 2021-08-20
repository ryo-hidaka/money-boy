// url_mappings.go

package app

import "/github.com/ryo.hidaka/money-boy/api/app"

func mapUrls() {
	router.GET("/products/:product_id", products.GetProduct)
	router.POST("/products", products.CreateProduct)
}
