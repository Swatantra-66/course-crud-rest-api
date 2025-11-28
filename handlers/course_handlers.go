package handlers

import (
	"buildCourseAPI/models"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var Courses []models.Course //fake DB

// seeding some fake DB
func SeedCourses() {
	Courses = append(Courses, models.Course{
		CourseId:    "2",
		CourseName:  "Go Backend Developer Course",
		CoursePrice: 799,
		Author:      &models.Author{FullName: "Swatantra Yadav", Website: "go.dev"}})
	Courses = append(Courses, models.Course{
		CourseId:    "4",
		CourseName:  "ReactJS",
		CoursePrice: 599,
		Author:      &models.Author{FullName: "Swatantra Yadav", Website: "MenScriptCourses.com"}})
	Courses = append(Courses, models.Course{
		CourseId:    "6",
		CourseName:  "Solana Blockchain Full Course",
		CoursePrice: 999,
		Author:      &models.Author{FullName: "Swatantra Yadav", Website: "MenScriptCourses.com"}})
	Courses = append(Courses, models.Course{
		CourseId:    "10",
		CourseName:  "Rust Course",
		CoursePrice: 1499,
		Author:      &models.Author{FullName: "Swatantra Yadav", Website: "MenScriptCourses.com"}})
}

// handlers

func GetAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Courses)
}

func GetOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for _, course := range Courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode("No Course found with given id")
}

func CreateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	var course models.Course
	if err := json.NewDecoder(r.Body).Decode(&course); err != nil {
		json.NewEncoder(w).Encode("Invalid JSON")
		return
	}
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	course.CourseId = strconv.Itoa(rand.Intn(1000))
	Courses = append(Courses, course) //save in "DB"

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(course)
}

func UpdateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["id"]

	for index, c := range Courses {
		if c.CourseId == id {
			Courses = append(Courses[:index], Courses[index+1:]...)

			var updated models.Course
			if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode("Invalid JSON payload")
				return
			}

			if updated.IsEmpty() {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode("No data inside JSON")
				return
			}

			updated.CourseId = id
			Courses = append(Courses, updated)

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(updated)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode("No course found with given id")
}

func DeleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r) //fetching the id to be deleted
	id := params["id"]

	for index, course := range Courses {
		if course.CourseId == id {
			Courses = append(Courses[:index], Courses[index+1:]...)
			json.NewEncoder(w).Encode("Course with given id is deleted")
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode("No course found with given id")
}
