package get_teacher_consultations

import (
	"encore.app/model/consultations"
	"encore.app/send_json"
	"log"
	"net/http"
	"slices"

	"encore.app/firebasesdk"
)

//encore:api public raw method=GET path=/getTeacherConsultations
func getTeacherConsultations(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	var teacherConsultations = make(map[string][]consultations.ConsultTeacher)
	err := firebasesdk.GetTeacherConsultations(&teacherConsultations)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var week = query.Get("week")
	var teacher = query.Get("teacher")
	if teachers, weekExists := teacherConsultations[week]; week != "" && weekExists {
		if teacherIndex := getTeacherIndex(teacher, teachers); teacher != "" && teacherIndex > 0 {
			send_json.SendJson(&w, teachers[teacherIndex])
			return
		} else if teacher == "" {
			send_json.SendJson(&w, teachers)
			return
		}
	} else if week == "" && teacher == "" {
		send_json.SendJson(&w, teacherConsultations)
		return
	}
	w.WriteHeader(http.StatusNotFound)

}

func getTeacherIndex(teacher string, teachers []consultations.ConsultTeacher) int {
	return slices.IndexFunc(teachers, func(t consultations.ConsultTeacher) bool {
		log.Println(t.Name)
		return teacher == t.Name
	})
}
