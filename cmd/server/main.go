package main

import (
	controlador "github.com/NicoAvellaneda1/GoCapas/cmd/server/handler"
	"github.com/NicoAvellaneda1/GoCapas/internal/users"
	"github.com/gin-gonic/gin"
)

func main() {
	repo := users.NewRepository()
	service := users.NewService(repo)
	u := controlador.NewUser(service)

	router := gin.Default()
	//agrupo las rutas
	us := router.Group("/users")
	us.POST("/", u.Store())
	us.GET("/", u.GetAll())

	router.Run()
}
