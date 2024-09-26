package student_schedule

type Lesson struct {
	Name      string `json:"name,omitempty"`
	Teacher   string `json:"teacher,omitempty"`
	Classroom string `json:"classroom,omitempty"`
	Time      string `json:"time,omitempty"`
	Subgroup  int    `json:"subgroup,omitempty"`
	Height    int    `json:"height,omitempty"`
}
