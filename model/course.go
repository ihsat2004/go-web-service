package model

import postgres "myapp/datastore"

type Course struct {
	CID        string `json:"cid"`
	CourseName string `json:"coursename"`
}

const (
	queryInsertCourse = "INSERT INTO course(cid,coursename) VALUES($1, $2)"
	queryGetCour      = "SELECT cid, coursename From course WHERE cid=$1;"
	queryUpdateCour   = "UPDATE course SET cid=$1, coursename=$2 WHERE cid=$3 RETURNING cid;"
	queryDeleteCour   = "DELETE FROM course WHERE cid=$1 RETURNING cid;"
)

func (c *Course) Create() error {
	_, err := postgres.Db.Exec(queryInsertCourse, c.CID, c.CourseName)
	return err
}

func (c *Course) Read() error {
	return postgres.Db.QueryRow(queryGetCour, c.CID).Scan(&c.CID, &c.CourseName)
}

func (c *Course) Update(oldCID string) error {
	_, err := postgres.Db.Exec(queryUpdateCour, c.CID, c.CourseName, oldCID)
	return err
}

func (c *Course) Delete() error {
	if err := postgres.Db.QueryRow(queryDeleteCour, c.CID).Scan(&c.CID); err != nil {
		return err
	}
	return nil
}

func GetAllCourses() ([]Course, error) {
	rows, getErr := postgres.Db.Query("SELECT * FROM course;")
	if getErr != nil {
		return nil, getErr
	}
	courses := []Course{}

	for rows.Next() {
		var c Course
		dbErr := rows.Scan(&c.CID, &c.CourseName)
		if dbErr != nil {
			return nil, dbErr
		}

		courses = append(courses, c)
	}
	rows.Close()
	return courses, nil
}
