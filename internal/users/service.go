package users

type Service interface {
	GetAll() ([]Usuarios, error)
	Store(nombre string, apellido string, email string, edad int, altura float64, activo bool, fechaCreacion string) ([]Usuarios, error)
	//ObtenerID() int
	Update(id int, nombre string, apellido string, email string, edad int, altura float64, activo bool, fechaCreacion string) (Usuarios, error)
	UpdateName(id int, nombre string) (Usuarios, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]Usuarios, error) {
	us, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return us, nil
}

func (s *service) Store(nombre string, apellido string, email string, edad int, altura float64, activo bool, fechaCreacion string) ([]Usuarios, error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return nil, err
	}

	lastID++

	// nid := s.ObtenerID()

	us, err := s.repository.Store(lastID, nombre, apellido, email, edad, altura, activo, fechaCreacion)
	if err != nil {
		return nil, err
	}

	return us, nil
}

// func (s *service) ObtenerID() int {
// 	u, err := s.repository.GetAll()
// 	if err != nil {
// 		return 0
// 	}

// 	if u != nil {
// 		u := u[len(u)-1]
// 		nuevoID := u.Id
// 		nuevoID++
// 		return nuevoID
// 	} else {
// 		return 1
// 	}
// }

func (s *service) Update(id int, nombre string, apellido string, email string, edad int, altura float64, activo bool, fechaCreacion string) (Usuarios, error) {
	return s.repository.Update(id, nombre, apellido, email, edad, altura, activo, fechaCreacion)
}

func (s *service) UpdateName(id int, nombre string) (Usuarios, error) {
	return s.repository.UpdateName(id, nombre)
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}
