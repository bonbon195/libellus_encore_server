package student_schedule

type Group struct {
	Name string `json:"name,omitempty"`
	Days []Day  `json:"days,omitempty"`
}
