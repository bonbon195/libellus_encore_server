package consultations

type ConsultDay struct {
	Name      string `json:"name,omitempty"`
	Date      string `json:"date,omitempty"`
	Time      string `json:"time,omitempty"`
	Classroom string `json:"classroom,omitempty"`
}
