package users

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

var usuarios []Usuarios
var lastID int

type Repository interface {
	GetAll() ([]Usuarios, error)
	Store(id int, nombre string, apellido string, email string, edad int, altura float64, activo bool, fechaCreacion string) ([]Usuarios, error)
	LastID() (int, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]Usuarios, error) {
	return usuarios, nil
}

func (r *repository) Store(id int, nombre string, apellido string, email string, edad int, altura float64, activo bool, fechaCreacion string) ([]Usuarios, error) {
	us := Usuarios{id, nombre, apellido, email, edad, altura, activo, fechaCreacion}
	usuarios = append(usuarios, us)
	return usuarios, nil
}

func (r *repository) LastID() (int, error) {
	return lastID, nil
}
