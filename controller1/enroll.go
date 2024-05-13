package controller

import (
	"database/sql"
	"encoding/json"
	"myapp/model"
	httpResp "myapp/util/httpresponse"
	"myapp/util/httpresponse/date"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func Enroll(w http.ResponseWriter, r *http.Request) {
	var e model.Enroll
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&e); err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, "invalid json body")
		return
	}

	e.Date_Enrolled = date.GetDate()
	defer r.Body.Close()

	saveErr := e.EnrollStud()
	if saveErr != nil {
		if strings.Contains(saveErr.Error(), "duplicate key") {
			httpResp.RespondWithError(w, http.StatusForbidden, "Duplicate keys")
			return
		} else {
			httpResp.RespondWithError(w, http.StatusInternalServerError, saveErr.Error())
		}
	}
	// The if statement checks if the error message (saveErr.Error()) contains the substring "duplicate key". This is commonly used
	// in database operations to detect errors related to violating unique constraints or primary key constraints, which typically occur
	// when trying to insert duplicate records and if the error message string contains the string "duplicate key" then the error message Duplicate key is shown to user.

	// no error
	httpResp.RespondWithJSON(w, http.StatusCreated, map[string]string{"status": "enrolled"})
}

func GetEnroll(w http.ResponseWriter, r *http.Request) {
	// get url parameters
	sid := mux.Vars(r)["sid"]
	cid := mux.Vars(r)["cid"]
	// get string sid to int type
	stdid, _ := strconv.ParseInt(sid, 10, 64)

	e := model.Enroll{StdId: stdid, CourseID: cid}
	getErr := e.Get()
	if getErr != nil {
		switch getErr {
		case sql.ErrNoRows:
			httpResp.RespondWithError(w, http.StatusNotFound, "No such enrollments")
		default:
			httpResp.RespondWithError(w, http.StatusInternalServerError, getErr.Error())
		}
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, e)
}

func GetEnrolls(w http.ResponseWriter, r *http.Request) {
	enrolls, getErr := model.GetAllEnrolls()
	if getErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, getErr.Error())
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, enrolls)
}

func DeleteEnroll(w http.ResponseWriter, r *http.Request) {
	sid := mux.Vars(r)["sid"]
	cid := mux.Vars(r)["cid"]
	stdid, _ := strconv.ParseInt(sid, 10, 64) //The strconv.ParseInt returns two values, the parsed integer value of the provided string and an Error if it occurs during the parsing.

	e := model.Enroll{StdId: stdid, CourseID: cid}

	if err := e.Delete(); err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "deleted"})
}
