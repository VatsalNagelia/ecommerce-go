package main

import (
	"log"
	"os"

	"github.com/VatsalNagelia/ecommerce-yt/controllers"
	"github.com/VatsalNagelia/ecommerce-yt/database"
	"github.com/VatsalNagelia/ecommerce-yt/middleware"
	"github.com/VatsalNagelia/ecommerce-yt/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))

	router := gin.New()

	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.GET("/addToCart", app.AddToCart())
	router.GET("/removeItem", app.RemoveItem())
	router.GET("/listCart", controllers.GetItemFromCart())
	router.POST("/addAddress", controllers.AddAddress())
	router.PUT("/editHomeAddress", controllers.EditHomeAddress())
	router.PUT("/editWorkAddress", controllers.EditWorkAddress())
	router.GET("/deleteAddresses", controllers.DeleteAddress())
	router.GET("/cartCheckout", app.BuyFromCart())
	router.GET("/instantBuy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))

}
