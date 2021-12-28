package controlador

import (
	"fmt"
	"os"
	"strconv"

	"github.com/NicoAvellaneda1/GoCapas/internal/users"
	"github.com/gin-gonic/gin"
)

type request struct {
	Nombre        string  `json:"nombre"`
	Apellido      string  `json:"apellido"`
	Email         string  `json:"email"`
	Edad          int     `json:"edad"`
	Altura        float64 `json:"altura"`
	Activo        bool    `json:"activo"`
	FechaCreacion string  `json:"fechaCreacion"`
}

type User struct {
	service users.Service
}

func NewUser(u users.Service) *User {
	return &User{
		service: u,
	}
}

func (u *User) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(401, gin.H{
				"error": "token invalido",
			})
			return
		}

		usuarios, err := u.service.GetAll()
		if err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, usuarios)
	}
}

func (u *User) Store() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(404, gin.H{
				"error": "token ivalido",
			})
			return
		}

		var req request
		if err := c.Bind(&req); err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		usuarios, err := u.service.Store(req.Nombre, req.Apellido, req.Email, req.Edad, req.Altura, req.Activo, req.FechaCreacion)
		if err != nil {
			c.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, usuarios)
	}
}

func (u *User) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(401, gin.H{
				"error": "token invalido",
			})
			return
		}

		//convierto a int el parametro recibido como string
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(400, gin.H{
				"error": "ID invalido",
			})
			return
		}

		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		//realizo las validaciones
		if req.Nombre == "" {
			c.JSON(400, gin.H{"error": "el nombre es requerido"})
			return
		}
		if req.Apellido == "" {
			c.JSON(400, gin.H{"error": "el apellido es requerido"})
			return
		}
		if req.Email == "" {
			c.JSON(400, gin.H{"error": "el email es requerido"})
			return
		}
		if req.Edad == 0 {
			c.JSON(400, gin.H{"error": "la edad es requerida"})
			return
		}
		if req.Altura == 0 {
			c.JSON(400, gin.H{"error": "la altura es requerida"})
			return
		}
		if req.FechaCreacion == "" {
			c.JSON(400, gin.H{"error": "el fechaCreacion es requerido"})
			return
		}

		us, err := u.service.Update(int(id), req.Nombre, req.Apellido, req.Email, req.Edad, req.Altura, req.Activo, req.FechaCreacion)
		if err != nil {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, us)
	}
}

func (u *User) UpdateName() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(401, gin.H{"error": "token invalido"})
			return
		}

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(401, gin.H{"error": "ID invalido"})
			return
		}

		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(404, gin.H{"error": err.Error()})
		}

		if req.Nombre == "" {
			c.JSON(400, gin.H{"error": "el nombre es requerido"})
		}

		us, err := u.service.UpdateName(int(id), req.Nombre)
		if err != nil {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, us)
	}
}

func (u *User) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(401, gin.H{"error": "token invalido"})
			return
		}

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(401, gin.H{"error": "ID invalido"})
			return
		}

		err = u.service.Delete(int(id))
		if err != nil {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{
			"data": fmt.Sprintf("El usuario %d ha sido eliminado", id),
		})
	}
}
