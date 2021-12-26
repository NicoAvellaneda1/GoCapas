package controlador

import (
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
		if token != "1234" {
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
		if token != "1234" {
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
