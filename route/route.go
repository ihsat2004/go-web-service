package routes

import (
	"log"
	"myapp/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func InitilizeRoutes() {
	var port = 8080
	router := mux.NewRouter()
	router.HandleFunc("/student", controller.AddStudent).Methods("POST")
	router.HandleFunc("/student/{sid}", controller.GetStud).Methods("GET")
	router.HandleFunc("/student/{sid}", controller.UpdateStud).Methods("PUT")
	router.HandleFunc("/student/{sid}", controller.DeleteStud).Methods("DELETE")
	router.HandleFunc("/students", controller.GetAllStuds)

	//course
	router.HandleFunc("/course", controller.AddCourse).Methods("POST")
	router.HandleFunc("/course/{cid}", controller.GetCourse).Methods("GET")
	router.HandleFunc("/course/{cid}", controller.UpdateCourse).Methods("PUT")
	router.HandleFunc("/course/{cid}", controller.DeleteCourse).Methods("DELETE")
	router.HandleFunc("/courses", controller.GetAllCourse)

	//enroll
	router.HandleFunc("/enroll", controller.Enroll).Methods("POST")
	router.HandleFunc("/enroll/{sid}/{cid}", controller.GetEnroll).Methods("GET")
	router.HandleFunc("/enrolls", controller.GetEnrolls).Methods("GET")
	router.HandleFunc("/enroll/{sid}/{cid}", controller.DeleteEnroll).Methods("DELETE")

	//router
	router.HandleFunc("/signup", controller.Signup).Methods("POST")

	//login
	router.HandleFunc("/login", controller.Login).Methods("POST")

	//logouta
	router.HandleFunc("/logout", controller.Logout)

	//serve static files
	fhandler := http.FileServer(http.Dir("./view"))
	router.PathPrefix("/").Handler(fhandler)

	log.Println("Application running on port", port)
	log.Fatal(http.ListenAndServe(":8080", router))
}
