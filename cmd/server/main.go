package main

import (
	controlador "github.com/NicoAvellaneda1/GoCapas/cmd/server/handler"
	"github.com/NicoAvellaneda1/GoCapas/internal/users"
	"github.com/NicoAvellaneda1/GoCapas/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	//cargo la variable de entorno
	_ = godotenv.Load()

	db := store.New(store.FileType, "./users.json")
	repo := users.NewRepository(db)
	service := users.NewService(repo)
	u := controlador.NewUser(service)

	router := gin.Default()
	//agrupo las rutas
	us := router.Group("/users")
	us.POST("/", u.Store())
	us.GET("/", u.GetAll())
	us.PUT("/:id", u.Update())
	us.PATCH("/:id", u.UpdateName())
	us.DELETE("/:id", u.Delete())

	router.Run()
}
