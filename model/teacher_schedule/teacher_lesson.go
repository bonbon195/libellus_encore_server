package teacher_schedule

type TeacherLesson struct {
	Name      string `json:"name,omitempty"`
	Group     string `json:"group,omitempty"`
	Classroom string `json:"classroom,omitempty"`
	Time      string `json:"time,omitempty"`
	Subgroup  int    `json:"subgroup,omitempty"`
	Height    int    `json:"height,omitempty"`
}
