package domain

type Patient struct {
	Rut       string
	FirstName string
	LastName  string
}

func NewPatient(rut, firstName, lastName string) (*Patient, error) {
	return &Patient{
		Rut:       rut,
		FirstName: firstName,
		LastName:  lastName,
	}, nil
}
