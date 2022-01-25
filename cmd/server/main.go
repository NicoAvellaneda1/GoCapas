package main

import (
	"os"

	"github.com/NicoAvellaneda1/GoCapas/cmd/server/docs"
	controlador "github.com/NicoAvellaneda1/GoCapas/cmd/server/handler"
	"github.com/NicoAvellaneda1/GoCapas/internal/users"
	"github.com/NicoAvellaneda1/GoCapas/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// @title MeLi Bootcamp API
// @version 1.0
// @description This API Handle MeLi products.
// @termsOfService http://hola.quetal

// @contact.name API Nico
// @contact.url http://www.nico.io/support
// @contact.email support@nico.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	//cargo la variable de entorno
	_ = godotenv.Load()

	db := store.New(store.FileType, "./users.json")
	repo := users.NewRepository(db)
	service := users.NewService(repo)
	u := controlador.NewUser(service)

	router := gin.Default()

	//swagger
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//agrupo las rutas
	us := router.Group("/users")
	us.POST("/", u.Store())
	us.GET("/", u.GetAll())
	us.PUT("/:id", u.Update())
	us.PATCH("/:id", u.UpdateName())
	us.DELETE("/:id", u.Delete())

	err := router.Run()
	if err != nil {
		panic(err)
	}
}
