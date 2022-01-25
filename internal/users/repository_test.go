package users

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type storeMock struct {
	Data []byte
	Err  error
}

func (s *storeMock) Read(data interface{}) error {
	_ = json.Unmarshal(s.Data, &data)

	return nil
}

func (s *storeMock) Write(data interface{}) error {
	s.Data, _ = json.Marshal(data)
	return nil
}

func TestGetAll(t *testing.T) {
	us1 := Usuarios{1, "Juan", "Perez", "jose@perez.com", 45, 1.65, true, "2021-12-14"}
	us2 := Usuarios{2, "Miguelo", "Perez", "jose@perez.com", 45, 1.65, true, "2021-12-14"}
	us := []Usuarios{us1, us2}

	usJson, _ := json.Marshal(us)

	repo := repository{db: &storeMock{Data: usJson}}
	us, err := repo.GetAll()

	assert.Equal(t, 2, len(us))
	assert.Equal(t, err, nil)
}

//------------------------------------------------------------------

//creo un spy que lo voy a poder utlizar en el Service_test porque estan dentro del mismo paquete
type storeSpy struct {
	Data        []byte
	BanderaRead bool
}

func (s *storeSpy) Read(data interface{}) error {
	s.BanderaRead = true
	err := json.Unmarshal(s.Data, &data)
	if err != nil {
		return nil
	}

	return nil
}

func (s *storeSpy) Write(data interface{}) error {
	s.Data, _ = json.Marshal(data)

	return nil
}

func TestUpadateName(t *testing.T) {
	//cargo un usuario en la bd de prueba
	us1 := Usuarios{1, "Beafore", "Perez", "jose@perez.com", 45, 1.65, true, "2021-12-14"}
	us := []Usuarios{us1}

	usJson, _ := json.Marshal(us)
	spy := &storeSpy{Data: usJson}
	repo := repository{db: spy}

	//pruebo que funcione el metodo Get
	usuario, err := repo.Get(1)
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, usuario.Nombre, "Beafore")
	assert.True(t, spy.BanderaRead)
	assert.Nil(t, err)

	//Ahora pruebo el UpdateName
	usuarioUpdateName, err := repo.UpdateName(1, "After")
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, usuarioUpdateName.Nombre, "After")
	assert.Nil(t, err)
}
