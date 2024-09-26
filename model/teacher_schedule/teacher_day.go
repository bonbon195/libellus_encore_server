package teacher_schedule

type TeacherDay struct {
	Date    string          `json:"date,omitempty"`
	Lessons []TeacherLesson `json:"lessons,omitempty"`
}
