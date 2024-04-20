package student_schedule

type Day struct {
	Name    string   `json:"name"`
	Date    string   `json:"date"`
	Lessons []Lesson `json:"lessons"`
}
