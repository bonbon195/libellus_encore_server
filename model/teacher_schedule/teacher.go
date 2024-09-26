package teacher_schedule

type Teacher struct {
	Name string       `json:"name,omitempty"`
	Week []TeacherDay `json:"week,omitempty"`
}
