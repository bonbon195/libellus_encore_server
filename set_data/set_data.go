package set_data

import (
	"encoding/json"
	"encore.app/firebasesdk"
	"encore.app/model/consultations"
	"encore.app/model/student_schedule"
	"encore.app/model/teacher_schedule"
	"io"
	"log"
	"net/http"
)

var secrets struct {
	AdminToken string
}

//encore:api public raw method=POST path=/setData
func SetData(w http.ResponseWriter, r *http.Request) {
	if r.Header["Token"] == nil || r.Header["Token"][0] != secrets.AdminToken {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}(r.Body)
	b, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var body struct {
		StudentScheduleData      *[]student_schedule.Faculty                            `json:"student_schedule"`
		TeacherScheduleData      *[]teacher_schedule.Teacher                            `json:"teacher_schedule"`
		TeacherConsultationsData *map[string][]consultations.ConsultTeacher             `json:"teacher_consultations"`
		ConsultationsData        *map[string]map[string][]consultations.ConsultDay      `json:"consultations"`
		ScheduleData             *map[string]map[string]map[string]student_schedule.Day `json:"schedule"`
		UpdateDate               *string                                                `json:"update_date"`
	}

	if err = json.Unmarshal(b, &body); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err = firebasesdk.SendData(body.StudentScheduleData, body.TeacherScheduleData, body.TeacherConsultationsData, body.ConsultationsData, body.ScheduleData, body.UpdateDate); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
}
