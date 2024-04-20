package student_schedule

type Group struct {
	Name string `json:"name"`
	Days []Day  `json:"days"`
}
