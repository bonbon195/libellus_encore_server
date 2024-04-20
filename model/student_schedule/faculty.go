package student_schedule

type Faculty struct {
	Id     string  `json:"-"`
	Code   string  `json:"code"`
	Groups []Group `json:"groups"`
}
