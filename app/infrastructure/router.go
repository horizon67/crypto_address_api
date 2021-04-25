package infrastructure

import (
	gin "github.com/gin-gonic/gin"

	"github.com/horizon67/crypto_address_api/app/interfaces/controllers"
)

var Router *gin.Engine

func init() {
	router := gin.Default()
	router.Use(gin.Logger())

	cryptoAddressController := controllers.NewCryptoAddressController(NewSqlHandler())

	router.GET("/crypto_addresses", func(c *gin.Context) { cryptoAddressController.Index(c) })
	router.GET("/crypto_addresses/:id", func(c *gin.Context) { cryptoAddressController.Show(c) })
	router.POST("/crypto_addresses", func(c *gin.Context) { cryptoAddressController.Create(c) })
	router.PUT("/crypto_addresses/:id", func(c *gin.Context) { cryptoAddressController.Save(c) })
	router.DELETE("/crypto_addresses/:id", func(c *gin.Context) { cryptoAddressController.Delete(c) })

	Router = router
}
