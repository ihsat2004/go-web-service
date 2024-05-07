package model

import postgres "myapp/datastore"

type Student struct {
	StdId     int64  `json:"stdid"`
	FirstName string `json:"fname"`
	LastName  string `json:"lname"`
	Email     string `json:"email"`
}

const queryInsertUser = "INSERT INTO student(stdid, firstname, lastname, email) VALUES($1, $2, $3, $4);"
const queryGetUser = "SELECT stdid, firstname, lastname, email FROM student WHERE stdid=$1;"
const queryUpdate = "UPDATE student SET stdid=$1, firstname=$2, lastname=$3, email=$4 WHERE stdid=$5 RETURNING stdid"
const queryDeleteUser = "DELETE FROM student WHERE stdid=$1 RETURNING stdid;"

func (s *Student) Create() error {
	_, err := postgres.Db.Exec(queryInsertUser, s.StdId, s.FirstName, s.LastName, s.Email)
	return err
}

func (s *Student) Read() error {
	return postgres.Db.QueryRow(queryGetUser, s.StdId).Scan(&s.StdId, &s.FirstName, &s.LastName, &s.Email)
}

func (s *Student) Update(oldID int64) error {
	err := postgres.Db.QueryRow(queryUpdate, s.StdId, s.FirstName, s.LastName, s.Email, oldID).Scan(&s.StdId)
	return err
}

func (s *Student) Delete() error {
	if err := postgres.Db.QueryRow(queryDeleteUser, s.StdId).Scan(&s.StdId); err != nil {
		return err
	}
	return nil
}

func GetAllStudents() ([]Student, error) {
	rows, getErr := postgres.Db.Query("SELECT * from student;")
	if getErr != nil {
		return nil, getErr
	}
	students := []Student{}

	for rows.Next() {
		var s Student
		dbErr := rows.Scan(&s.StdId, &s.FirstName, &s.LastName, &s.Email)
		if dbErr != nil {
			return nil, dbErr
		}

		students = append(students, s)
	}
	rows.Close()
	return students, nil
}
