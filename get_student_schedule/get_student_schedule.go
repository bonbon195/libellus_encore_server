package get_student_schedule

import (
	"encore.app/firebasesdk"
	"encore.app/model/student_schedule"
	"encore.app/send_json"
	"net/http"
	"slices"
)

//encore:api public raw method=GET path=/getStudentSchedule
func getStudentSchedule(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	var schedule []student_schedule.Faculty

	err := firebasesdk.GetStudentSchedule(&schedule)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var faculty = query.Get("faculty")
	var group = query.Get("group")
	if faculty != "" && group == "" {
		if facultyIndex := getFacultyIndex(faculty, schedule); 0 < facultyIndex {
			send_json.SendJson(&w, schedule[facultyIndex])
			return
		}
	} else if faculty != "" && group != "" {
		if facultyIndex := getFacultyIndex(faculty, schedule); 0 < facultyIndex {
			if groupIndex := getGroupIndex(group, schedule[facultyIndex].Groups); 0 < groupIndex {
				send_json.SendJson(&w, schedule[facultyIndex].Groups[groupIndex])
				return
			}
		}
	} else if faculty == "" && group == "" {
		send_json.SendJson(&w, schedule)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

func getFacultyIndex(faculty string, schedule []student_schedule.Faculty) int {
	return slices.IndexFunc(schedule, func(f student_schedule.Faculty) bool {
		return f.Code == faculty
	})
}

func getGroupIndex(group string, groups []student_schedule.Group) int {
	return slices.IndexFunc(groups, func(g student_schedule.Group) bool {
		return g.Name == group
	})
}
