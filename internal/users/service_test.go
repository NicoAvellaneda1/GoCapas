package users

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

//al storeSpy lo cree en repositore_test!!

func TestUpadate(t *testing.T) {
	//cargo un usuario en la bd de prueba
	us1 := Usuarios{1, "Beafore", "Perez", "jose@perez.com", 45, 1.65, true, "2021-12-14"}
	us := []Usuarios{us1}

	usJson, _ := json.Marshal(us)
	spy := &storeSpy{Data: usJson}
	//paso la bd al repo
	repo := repository{db: spy}
	//paso el repo a el service
	service := service{repository: &repo}

	//pruebo que funcione el metodo Get
	usuario, err := service.Get(1)
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, usuario.Nombre, "Beafore")
	assert.True(t, spy.BanderaRead)
	assert.Nil(t, err)

	//Ahora pruebo el Update en el service
	usuarioUpdate, err := service.Update(1, "After", "Perez", "jose@perez.com", 45, 1.65, true, "2021-12-14")
	if err != nil {
		t.Fail()
	}
	esperado := Usuarios{1, "After", "Perez", "jose@perez.com", 45, 1.65, true, "2021-12-14"}

	assert.Equal(t, esperado, usuarioUpdate)
	assert.Nil(t, err)
}

func TestDelete(t *testing.T) {
	us1 := Usuarios{1, "Juan", "Perez", "jose@perez.com", 45, 1.65, true, "2021-12-14"}
	us := []Usuarios{us1}

	usJson, _ := json.Marshal(us)

	repo := repository{db: &storeMock{Data: usJson}}
	service := service{repository: &repo}
	err := service.Delete(1)
	us, err1 := repo.GetAll()

	assert.Equal(t, 0, len(us))
	assert.Nil(t, err)
	assert.Nil(t, err1)
}

func TestDeleteError(t *testing.T) {
	expectedError := errors.New("Usuario 1 no encontrado")

	repo := repository{db: &storeMock{Err: expectedError}}
	service := service{repository: &repo}
	err := service.Delete(1)

	assert.NotNil(t, err)
	assert.Equal(t, expectedError, err)
}
