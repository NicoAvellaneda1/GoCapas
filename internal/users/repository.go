package users

import (
	"fmt"

	"github.com/NicoAvellaneda1/GoCapas/pkg/store"
)

type Usuarios struct {
	Id            int     `json:"id"`
	Nombre        string  `json:"nombre"`
	Apellido      string  `json:"apellido"`
	Email         string  `json:"email"`
	Edad          int     `json:"edad"`
	Altura        float64 `json:"altura"`
	Activo        bool    `json:"activo"`
	FechaCreacion string  `json:"fechaCreacion"`
}

//var usuarios []Usuarios
//var lastID int

type Repository interface {
	GetAll() ([]Usuarios, error)
	Store(id int, nombre string, apellido string, email string, edad int, altura float64, activo bool, fechaCreacion string) ([]Usuarios, error)
	LastID() (int, error)
	Update(id int, nombre string, apellido string, email string, edad int, altura float64, activo bool, fechaCreacion string) (Usuarios, error)
	UpdateName(id int, nombre string) (Usuarios, error)
	Delete(id int) error
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll() ([]Usuarios, error) {
	var usuarios []Usuarios
	r.db.Read(&usuarios)
	return usuarios, nil
}

func (r *repository) Store(id int, nombre string, apellido string, email string, edad int, altura float64, activo bool, fechaCreacion string) ([]Usuarios, error) {
	var usuarios []Usuarios
	r.db.Read(&usuarios)
	us := Usuarios{id, nombre, apellido, email, edad, altura, activo, fechaCreacion}
	usuarios = append(usuarios, us)
	if err := r.db.Write(usuarios); err != nil {
		return []Usuarios{}, err
	}
	return usuarios, nil
}

func (r *repository) LastID() (int, error) {
	var usuarios []Usuarios
	if err := r.db.Read(&usuarios); err != nil {
		return 0, err
	}
	if len(usuarios) == 0 {
		return 0, nil
	}

	return usuarios[len(usuarios)-1].Id, nil
}

func (r *repository) Update(id int, nombre string, apellido string, email string, edad int, altura float64, activo bool, fechaCreacion string) (Usuarios, error) {
	var usuarios []Usuarios
	if err := r.db.Read(&usuarios); err != nil {
		return Usuarios{}, err
	}
	us := Usuarios{Nombre: nombre, Apellido: apellido, Email: email, Edad: edad, Altura: altura, Activo: activo, FechaCreacion: fechaCreacion}
	update := false
	for i := range usuarios {
		if usuarios[i].Id == id {
			us.Id = id
			usuarios[i] = us
			update = true
		}
	}
	if !update {
		return Usuarios{}, fmt.Errorf("Usuario %d no encontrado", id)
	}
	if err := r.db.Write(usuarios); err != nil {
		return Usuarios{}, err
	}
	return us, nil
}

func (r *repository) UpdateName(id int, name string) (Usuarios, error) {
	var usuarios []Usuarios
	r.db.Read(&usuarios)
	var us Usuarios
	update := false
	for i := range usuarios {
		if usuarios[i].Id == id {
			usuarios[i].Nombre = name
			update = true
			us = usuarios[i]
		}
	}
	if !update {
		return Usuarios{}, fmt.Errorf("Usuario %d no encontrado", id)
	}

	if err := r.db.Write(usuarios); err != nil {
		return Usuarios{}, err
	}

	return us, nil
}

func (r *repository) Delete(id int) error {
	var usuarios []Usuarios
	//leo
	if err := r.db.Read(&usuarios); err != nil {
		return err
	}

	deleted := false
	var index int
	for i := range usuarios {
		if usuarios[i].Id == id {
			index = i
			deleted = true
		}
	}
	if !deleted {
		return fmt.Errorf("Usuario %d no encontrado", id)
	}
	usuarios = append(usuarios[:index], usuarios[index+1:]...)

	//escribo
	if err := r.db.Write(usuarios); err != nil {
		return err
	}
	return nil
}
