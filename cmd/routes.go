package main

import (
	"github.com/Asylniet/golang_lab3/controllers"
	"github.com/gin-gonic/gin"
)

func setUpRoutes(r *gin.Engine) {
	//Store
	r.POST("/store", controllers.AddItem) // name, price
	r.GET("/store", controllers.GetItems)
	r.GET("/store/:id", controllers.GetItem)
	r.GET("/store/search", controllers.SearchItems) // filter, sort
	r.DELETE("/store/:id", controllers.RemoveItem)
	r.PUT("/store/:id", controllers.UpdateItem) // name, price

	// Users
	r.POST("/user", controllers.AddUser) // name, password
	r.GET("/user", controllers.GetUsers)
	r.DELETE("/user/:id", controllers.RemoveUser)
	r.GET("/user/login", controllers.LoginUser)
	r.GET("/user/:id", controllers.GetUserById)
	r.PUT("/user/:id", controllers.UpdateUser)         // name, password
	r.POST("/user/rate/:id", controllers.RateItem)     // rating
	r.POST("/user/saveItem/:id", controllers.SaveItem) // item(id)
	r.DELETE("/user/removeItem/:id", controllers.RemoveSavedItem)
	r.GET("/user/getSavedItems/:id", controllers.GetSavedItems)
}
