package get_teacher_schedule

import (
	"encore.app/firebasesdk"
	"encore.app/model/teacher_schedule"
	"encore.app/send_json"
	"net/http"
	"slices"
)

//encore:api public raw method=GET path=/getTeacherSchedule
func getTeacherSchedule(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	var schedule []teacher_schedule.Teacher
	err := firebasesdk.GetTeacherSchedule(&schedule)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	teacher := query.Get("teacher")

	if index := getTeacherIndex(teacher, schedule); teacher != "" && index > 0 {
		send_json.SendJson(&w, schedule[index])
		return
	} else if teacher == "" {
		send_json.SendJson(&w, schedule)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

func getTeacherIndex(teacher string, schedule []teacher_schedule.Teacher) int {
	return slices.IndexFunc(schedule, func(t teacher_schedule.Teacher) bool {
		return teacher == t.Name
	})
}
