package student_schedule

type Faculty struct {
	Id     string  `json:"-"`
	Code   string  `json:"code,omitempty"`
	Groups []Group `json:"groups,omitempty"`
}
