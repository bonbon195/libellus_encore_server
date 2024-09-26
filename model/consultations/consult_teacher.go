package consultations

type ConsultTeacher struct {
	Name string       `json:"name,omitempty"`
	Week []ConsultDay `json:"week,omitempty"`
}
