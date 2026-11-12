package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice string  `json:"price"`
	Author      *Author `json:"author" `
}
type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

var courses []Course

func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}
func main() {

}

//controller

func serveHome(w http.ResponseWriter , r *http.Request){
	w.Write([]byte("<h1>Welcome to serve home Api </h1>"))
}

func getAllCourses(w http.ResponseWriter , r *http.Request){
	fmt.Println("All Courses </h1>")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}