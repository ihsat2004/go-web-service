package controller

import (
	"encoding/json"
	"myapp/model"
	httpResp "myapp/util/httpresponse"
	"net/http"
	"time"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	var admin model.Admin
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&admin); err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, "invalid json body")
		return
	}
	defer r.Body.Close()
	saveErr := admin.Create()
	if saveErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, saveErr.Error())
		return
	}
	// no error
	httpResp.RespondWithJSON(w, http.StatusCreated, map[string]string{"status": "admin added"})
}

func Login(w http.ResponseWriter, r *http.Request) {
	var admin model.Admin
	err := json.NewDecoder(r.Body).Decode(&admin)
	if err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, "invalid json body")
		return
	}
	defer r.Body.Close()
	getErr := admin.Get()
	if getErr != nil {
		httpResp.RespondWithError(w, http.StatusUnauthorized, getErr.Error())
		return
	}

	//set cookie
	cookie := http.Cookie{
		Name:    "my-cookie",
		Value:   "my-value",
		Expires: time.Now().Add(30 * time.Minute),
		Secure:  true,
	}
	http.SetCookie(w, &cookie)

	httpResp.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "success"})
}

func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    "my-cookie",
		Expires: time.Now(),
	})
	httpResp.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "cookie deleted"})
}

func VerifyCookie(w http.ResponseWriter, r *http.Request) bool {
	// Retrieve the "my-cookie" cookie from the request
	cookie, err := r.Cookie("my-cookie")
	if err != nil {
		if err == http.ErrNoCookie {
			// No cookie found, redirect to login page or return an error
			httpResp.RespondWithError(w, http.StatusSeeOther, "cookie not found")
			return false
		}
		// Some other error occurred
		httpResp.RespondWithError(w, http.StatusInternalServerError,
			"internal server error")
		return false
	}
	// Verify the cookie value
	if cookie.Value != "my-value" {
		// Invalid cookie value, redirect to login page or return an error
		httpResp.RespondWithError(w, http.StatusUnauthorized, "cookie does not match")
		return false
	}
	return true
}
