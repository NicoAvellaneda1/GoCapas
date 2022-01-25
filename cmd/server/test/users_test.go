package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	controlador "github.com/NicoAvellaneda1/GoCapas/cmd/server/handler"
	"github.com/stretchr/testify/assert"

	"github.com/NicoAvellaneda1/GoCapas/internal/users"
	"github.com/NicoAvellaneda1/GoCapas/pkg/store"
	"github.com/gin-gonic/gin"
)

func createServer() *gin.Engine {
	_ = os.Setenv("TOKEN", "1234")
	db := store.New(store.FileType, "../users.json")
	repo := users.NewRepository(db)
	service := users.NewService(repo)
	u := controlador.NewUser(service)
	r := gin.Default()

	us := r.Group("/users")
	us.PUT("/:id", u.Update())
	us.DELETE("/:id", u.Delete())

	return r
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "1234")

	return req, httptest.NewRecorder()
}

func TestUpadate(t *testing.T) {
	//crea el seriver y define las rutas
	r := createServer()
	//crea el request y la reponse de Update para obtener el resultado
	req, rr := createRequestTest(http.MethodPut, "/users/1", `{
		"id": 1,
		"nombre": "Juan",
		"apellido": "Gomez",
		"email": "jose@perez.com",
		"edad": 45,
		"altura": 1.65,
		"activo": true,
		"fechaCreacion": "2021-12-14"
	   }`)

	//indica al servidor quepuede atender la solicitud
	r.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)
}

func TestDelete(t *testing.T) {
	r := createServer()
	req, rr := createRequestTest(http.MethodDelete, "/users/4", "")

	r.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)
}
