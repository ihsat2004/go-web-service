package model

import postgres "myapp/datastore"

// The type Enroll defines a structure with fields for student ID, course ID, and enrollment date.
// @property {int64} stdId - The stdId property in the Enroll struct represents the student ID of
// the student enrolling in a course. It is of type int64 and is tagged with json:"stdid" for JSON
// marshaling and unmarshaling.
// @property {string} courseID - The courseID property in the Enroll struct represents the unique
// identifier for the course that a student is enrolling in. It is of type string and is tagged with
// json:"cid" for JSON marshaling and unmarshaling purposes.
// @property {string} Date_Enrolled - The Date_Enrolled property in the Enroll struct represents
// the date when a student enrolled in a course. It is a string type field and is tagged with
// json:"date" for JSON marshaling and unmarshaling purposes.
type Enroll struct {
	StdId         int64  `json:"stdid"`
	CourseID      string `json:"cid"`
	Date_Enrolled string `json:"date"`
}

const (
	queryEnrollStd = "INSERT INTO enroll(std_id, course_id, date_enrolled) VALUES($1, $2, $3) RETURNING std_id;"
)

func (e *Enroll) EnrollStud() error {
	row := postgres.Db.QueryRow(queryEnrollStd, e.StdId, e.CourseID, e.Date_Enrolled)
	err := row.Scan(&e.StdId)
	return err
}

const queryGetEnroll = "SELECT std_id, course_id, date_enrolled FROM enroll WHERE std_id=$1 and course_id=$2;"

func (e *Enroll) Get() error {
	return postgres.Db.QueryRow(queryGetEnroll, e.StdId, e.CourseID).Scan(&e.StdId, &e.CourseID, &e.Date_Enrolled)
}

func GetAllEnrolls() ([]Enroll, error) {
	rows, getErr := postgres.Db.Query("SELECT std_id, course_id, date_enrolled from enroll;")
	if getErr != nil {
		return nil, getErr
	}
	// create a slice of type Course
	enrolls := []Enroll{}

	for rows.Next() {
		var e Enroll
		dbErr := rows.Scan(&e.StdId, &e.CourseID, &e.Date_Enrolled)
		if dbErr != nil {
			return nil, dbErr
		}
		enrolls = append(enrolls, e)
	}
	rows.Close()
	return enrolls, nil
}

const queryDeleteEnroll = "DELETE FROM enroll WHERE std_id=$1 and course_id=$2 RETURNING std_id;"

func (e *Enroll) Delete() error {
	row := postgres.Db.QueryRow(queryDeleteEnroll, e.StdId, e.CourseID)
	err := row.Scan(&e.StdId)
	return err
}
