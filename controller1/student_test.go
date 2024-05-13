package controller

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddStudent(t *testing.T) {
	url := "http://localhost:8080/student"
	var jsonStr = []byte(`{"stdid":1002, "fname": "Sangay", "lname": "Lhamo",
	"email":"sl@gmail.com"}`)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
	expResp := `{"status":"student added"}`
	assert.JSONEq(t, expResp, string(body))
}
func TestGetStudent(t *testing.T) {
	c := http.Client{}
	r, _ := c.Get("http://localhost:8080/student/1002")
	body, _ := io.ReadAll(r.Body)
	assert.Equal(t, http.StatusOK, r.StatusCode)
	expResp := `{"stdid":1002, "fname":"Sangay", "lname":"Lhamo",
	"email":"sl@gmail.com"}`
	assert.JSONEq(t, expResp, string(body))
}
func TestDeleteStudent(t *testing.T) {
	url := "http://localhost:8080/student/1002"
	req, _ := http.NewRequest("DELETE", url, nil)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	expResp := `{"status":"deleted"}`
	assert.JSONEq(t, expResp, string(body))
}
func TestStudentNotFound(t *testing.T) {
	assert := assert.New(t)
	c := http.Client{}
	r, _ := c.Get("http://localhost:8080/student/1002")
	body, _ := io.ReadAll(r.Body)
	assert.Equal(http.StatusNotFound, r.StatusCode)
	expResp := `{"error":"Student not found"}`
	assert.JSONEq(expResp, string(body))
}
