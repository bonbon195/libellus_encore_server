package firebasesdk

import (
	"context"
	"log"

	"encore.app/model/consultations"
	"encore.app/model/student_schedule"
	"encore.app/model/teacher_schedule"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/db"
	"google.golang.org/api/option"
)

var secrets struct {
	FirebaseCredentials string
	FirebaseURL         string
}

var studentSchedulePath = "student_schedule"
var teacherSchedulePath = "teacher_schedule"
var teacherConsultationsPath = "teacher_consultations"
var ctx = context.Background()

func initFirebaseClient() (*db.Client, error) {

	conf := &firebase.Config{
		DatabaseURL: secrets.FirebaseURL,
	}
	opt := option.WithCredentialsJSON([]byte(secrets.FirebaseCredentials))
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {

		return nil, err
	}

	client, err := app.Database(ctx)
	if err != nil {
		return nil, err
	}
	return client, nil
}
func SendData(studentSchedule *[]student_schedule.Faculty, teacherSchedule *[]teacher_schedule.Teacher, teacherConsultations *map[string][]consultations.ConsultTeacher, consultations *map[string]map[string][]consultations.ConsultDay, schedule *map[string]map[string]map[string]student_schedule.Day, updateDate *string) error {
	client, err := initFirebaseClient()
	if err != nil {
		return err
	}

	studentScheduleRef := client.NewRef(studentSchedulePath)
	if err := studentScheduleRef.Set(ctx, studentSchedule); err != nil {
		return err
	}

	teacherScheduleRef := client.NewRef(teacherSchedulePath)
	if err := teacherScheduleRef.Set(ctx, teacherSchedule); err != nil {
		return err
	}

	teacherConsultRef := client.NewRef(teacherConsultationsPath)
	if err := teacherConsultRef.Set(ctx, teacherConsultations); err != nil {
		return err
	}

	updateDateRef := client.NewRef("update_date")
	if err := updateDateRef.Set(ctx, updateDate); err != nil {
		return err
	}

	// backwards compatibility
	consultRef := client.NewRef("consultations")
	if err := consultRef.Set(ctx, consultations); err != nil {
		return err
	}
	specialitiesRef := client.NewRef("specialities")
	if err := specialitiesRef.Set(ctx, schedule); err != nil {
		return err
	}

	log.Println("Done!")
	return nil
}

func GetStudentSchedule(schedule *[]student_schedule.Faculty) error {
	client, err := initFirebaseClient()
	if err != nil {
		return err
	}
	ref := client.NewRef(studentSchedulePath)
	log.Println(ref.Path)
	if err := ref.Get(ctx, schedule); err != nil {
		return err
	}
	return nil
}

func GetTeacherSchedule(schedule *[]teacher_schedule.Teacher) error {
	client, err := initFirebaseClient()
	if err != nil {
		log.Println(err)
		return err
	}
	ref := client.NewRef("teacher_schedule")
	if err := ref.Get(ctx, schedule); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func GetTeacherConsultations(consultations *map[string][]consultations.ConsultTeacher) error {
	client, err := initFirebaseClient()
	if err != nil {
		return err
	}
	ref := client.NewRef(teacherConsultationsPath)
	if err := ref.Get(ctx, consultations); err != nil {
		return err
	}
	return nil
}

func GetUpdateDate(date *string) error {
	client, err := initFirebaseClient()
	if err != nil {
		return err
	}
	ref := client.NewRef("update_date")
	if err := ref.Get(ctx, date); err != nil {
		return err
	}
	return nil
}
