package main

import "net/http"

// model for courses

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}
type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

var courses []Course

// middlewares :

func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

func main() {

}

// controllers

// serve home route

func serverHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1> welcome to home <h1>"))
}
func getAllCourse(w http.ResponseWriter, r *http.Request) {

}
