package student_schedule

type Day struct {
	Name    string   `json:"name,omitempty"`
	Date    string   `json:"date,omitempty"`
	Lessons []Lesson `json:"lessons,omitempty"`
}
