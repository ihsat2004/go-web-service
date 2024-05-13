package controller

import (
	"database/sql"
	"encoding/json"
	"myapp/model"
	httpResp "myapp/util/httpresponse"
	"net/http"

	"strconv"

	"github.com/gorilla/mux"
)

func AddStudent(w http.ResponseWriter, r *http.Request) {
	if !VerifyCookie(w, r) {
		return
	}
	var stud model.Student

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&stud); err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, "invalid json body")
		return
	}

	defer r.Body.Close()

	saveErr := stud.Create()
	if saveErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, saveErr.Error())
		return
	}

	httpResp.RespondWithJSON(w, http.StatusCreated, map[string]string{"status": "student added"})
}

func GetStud(w http.ResponseWriter, r *http.Request) {
	if !VerifyCookie(w, r) {
		return
	}
	sid := mux.Vars(r)["sid"]
	stdId, idErr := getUserId(sid)
	if idErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, idErr.Error())
		return
	}

	s := model.Student{StdId: stdId}
	getErr := s.Read()
	if getErr != nil {
		switch getErr {
		case sql.ErrNoRows:
			httpResp.RespondWithError(w, http.StatusNotFound, "Student not found")
		default:
			httpResp.RespondWithError(w, http.StatusInternalServerError, getErr.Error())
		}
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, s)
}

func getUserId(userIdParam string) (int64, error) {

	userId, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if userErr != nil {
		return 0, userErr
	}
	return userId, nil
}

func UpdateStud(w http.ResponseWriter, r *http.Request) {
	if !VerifyCookie(w, r) {
		return
	}
	old_sid := mux.Vars(r)["sid"]
	old_stdId, idErr := getUserId(old_sid)
	if idErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, idErr.Error())
		return
	}

	var stud model.Student
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&stud); err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, "invalid json body")
		return
	}
	defer r.Body.Close()

	updateErr := stud.Update(old_stdId)
	if updateErr != nil {
		switch updateErr {
		case sql.ErrNoRows:
			httpResp.RespondWithError(w, http.StatusNotFound, "Student not found")
		default:
			httpResp.RespondWithError(w, http.StatusInternalServerError, updateErr.Error())
		}
	} else {
		httpResp.RespondWithJSON(w, http.StatusOK, stud)
	}

}

func DeleteStud(w http.ResponseWriter, r *http.Request) {
	if !VerifyCookie(w, r) {
		return
	}
	sid := mux.Vars(r)["sid"]
	stdId, idErr := getUserId(sid)
	if idErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, idErr.Error())
		return
	}
	s := model.Student{StdId: stdId}
	if err := s.Delete(); err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "deleted"})
}

func GetAllStuds(w http.ResponseWriter, r *http.Request) {
	if !VerifyCookie(w, r) {
		return
	}
	students, getErr := model.GetAllStudents()
	if getErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, getErr.Error())
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, students)
}
