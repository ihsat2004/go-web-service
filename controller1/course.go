package controller

import (
	"database/sql"
	"encoding/json"
	"myapp/model"
	httpResp "myapp/util/httpresponse"

	"net/http"

	"github.com/gorilla/mux"
)

func AddCourse(w http.ResponseWriter, r *http.Request) {
	var cour model.Course
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&cour); err != nil {
		// httpResp.ResponseWithError(w, http.StatusBadRequest, "invalid json body")
		httpResp.RespondWithError(w, http.StatusBadRequest, "Invalid json body")
		return
	}
	defer r.Body.Close()

	saveErr := cour.Create()
	if saveErr != nil {
		httpResp.RespondWithError(w, http.StatusInternalServerError, saveErr.Error())
		return
	}

	// httpResp.ResponseWithJSON(w, http.StatusCreated, map[string]string{"status": "course added"})
	httpResp.RespondWithJSON(w, http.StatusCreated, map[string]string{"status": "course added"})
}

func GetCourse(w http.ResponseWriter, r *http.Request) {
	CId := mux.Vars(r)["cid"]

	c := model.Course{CID: CId}
	getErr := c.Read()
	if getErr != nil {
		switch getErr {
		case sql.ErrNoRows:

			httpResp.RespondWithError(w, http.StatusNotFound, "Course not found")
		default:
			// httpResp.ResponseWithError(w, http.StatusInternalServerError, getErr.Error())
			httpResp.RespondWithError(w, http.StatusInternalServerError, getErr.Error())
		}
		return
	}

	// httpResp.ResponseWithJSON(w, http.StatusOK, c)
	httpResp.RespondWithJSON(w, http.StatusOK, c)
}

func UpdateCourse(w http.ResponseWriter, r *http.Request) {
	oldCID := mux.Vars(r)["cid"]

	var cour model.Course
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&cour); err != nil {
		// httpResp.ResponseWithError(w, http.StatusBadRequest, "invalid json body")
		httpResp.RespondWithError(w, http.StatusBadRequest, "invalid json body")
		return
	}
	defer r.Body.Close()

	updateErr := cour.Update(oldCID)
	if updateErr != nil {
		switch updateErr {
		case sql.ErrNoRows:
			// httpResp.ResponseWithError(w, http.StatusNotFound, "course not found")
			httpResp.RespondWithError(w, http.StatusNotFound, "course not found")
		default:
			// httpResp.ResponseWithError(w, http.StatusInternalServerError, updateErr.Error())
			httpResp.RespondWithError(w, http.StatusInternalServerError, updateErr.Error())
		}
		return
	}

	// httpResp.ResponseWithJSON(w, http.StatusOK, cour)
	httpResp.RespondWithJSON(w, http.StatusOK, cour)
}

func DeleteCourse(w http.ResponseWriter, r *http.Request) {
	CId := mux.Vars(r)["cid"]

	c := model.Course{CID: CId}
	if err := c.Delete(); err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "deleted"})
}

func GetAllCourse(w http.ResponseWriter, r *http.Request) {
	courses, getErr := model.GetAllCourses()
	if getErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, getErr.Error())
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, courses)
}
