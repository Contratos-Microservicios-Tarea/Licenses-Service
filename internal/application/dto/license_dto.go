package dto

type LicenseDTO struct {
	Folio     string
	PatientID string
	DoctorID  string
	Diagnosis string
	StartDate string
	Days      uint8
	Status    string
}
