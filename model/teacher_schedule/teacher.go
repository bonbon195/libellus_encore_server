package teacher_schedule

type Teacher struct {
	Name string       `json:"name"`
	Week []TeacherDay `json:"week"`
}
