package consultations

type ConsultTeacher struct {
	Name string       `json:"name"`
	Week []ConsultDay `json:"week"`
}
