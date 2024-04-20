package student_schedule

type Lesson struct {
	Name      string `json:"name"`
	Teacher   string `json:"teacher"`
	Classroom string `json:"classroom"`
	Time      string `json:"time"`
	Subgroup  int    `json:"subgroup"`
	Height    int    `json:"height"`
}
