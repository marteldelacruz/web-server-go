package admin

import (
	"fmt"
	"net/http"
	"strconv"

	Util "../util"
)

type Data struct {
	Subject map[string]map[string]float64
	Student map[string]map[string]float64
}

type Args struct {
	Name    string
	Subject string
	Grade   float64
}

type Server struct {
	Maps Data
}

// Inits the maps
func (server *Server) Init() {
	(*server).Maps.Student = make(map[string]map[string]float64)
	(*server).Maps.Subject = make(map[string]map[string]float64)

	fmt.Println("Successful creation of maps!")
}

// Checks if the student already has a grade for a subject,
// if so, returns true
func (server *Server) studentHasGrade(name string, subject string) bool {
	for s, _ := range (*server).Maps.Student[name] {
		if s == subject {
			return true
		}
	}

	return false
}

// Adds a new stutent to the map thing
func (server *Server) Add(args Args) {
	// create new student if not exist
	if (*server).Maps.Student[args.Name] == nil {
		(*server).Maps.Student[args.Name] = make(map[string]float64)
	}

	// check if grade is already saved
	if (*server).studentHasGrade(args.Name, args.Subject) {
		fmt.Println("Student already has a grade for that subject")
		return
	}
	(*server).Maps.Student[args.Name][args.Subject] = args.Grade

	// create new subject
	if (*server).Maps.Subject[args.Subject] == nil {
		(*server).Maps.Subject[args.Subject] = make(map[string]float64)
	}
	(*server).Maps.Subject[args.Subject][args.Name] = args.Grade

	fmt.Println("New student added!")
}

// Returns the student average grade of all its subjects
func (server *Server) StudentAverage(args Args) float64 {
	var result float64 = 0
	fmt.Println("Client ask for student average")

	for _, g := range (*server).Maps.Student[args.Name] {
		result += g
	}

	return result / float64(len((*server).Maps.Student[args.Name]))
}

// Returns the subject average grade of all its students
func (server *Server) SubjectAverage(args Args) float64 {
	var result float64 = 0
	fmt.Println("Client ask for subject average")

	for _, g := range (*server).Maps.Subject[args.Subject] {
		result += g
	}

	return result / float64(len((*server).Maps.Subject[args.Subject]))
}

// Returns the average of all students and subjects
func (server *Server) GeneralAverage(args Args) float64 {
	var total float64
	var result float64 = 0
	fmt.Println("Client ask for general average")

	for _, subject := range (*server).Maps.Student {
		for _, g := range subject {
			result += g
			total += 1
		}
	}

	return result / total
}

// Default home page where the client can save the
// students data
func (server *Server) HomePage(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	// sending home.html
	fmt.Fprintf(
		res, Util.LoadHtml("home.html"),
	)
}

//
func (server *Server) GetGradesTable() string {
	data := ""
	var total int64 = 1
	fmt.Println("Client ask for all data")

	for name, s := range (*server).Maps.Student {
		for subject, g := range s {
			data += "<tr><th scope='row'>" + fmt.Sprint(total) + "</th>"
			data += "<td>" + name + "</td>"
			data += "<td>" + subject + "</td>"
			data += "<td>" + fmt.Sprintf("%.2f", g) + "</td></tr>"
			total += 1
		}
	}

	return data
}

// This page will show all the data stored by the client
// on a single table
func (server *Server) GradesPage(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	// sending grades.html
	gradesTable := server.GetGradesTable()

	fmt.Fprintf(
		res, Util.LoadHtml("grades.html"), gradesTable,
	)
}

// this function will handle the requested methods from the
// home form in order to save new data
func (server *Server) Save(res http.ResponseWriter, req *http.Request) {
	fmt.Println("SAVE method => " + req.Method)

	switch req.Method {
	case "POST":
		if err := req.ParseForm(); err != nil {
			fmt.Fprintf(res, "ParseForm() error %v", err)
			return
		}
		(*server).SaveToMaps(req)

		fmt.Fprintf(
			res, Util.LoadHtml("saved.html"),
		)
		break
	default:
		fmt.Println("Cannot handle requested method :(")
		break
	}
}

// saves the req to the maps using the parsing the received method
// to the args struct
func (server *Server) SaveToMaps(req *http.Request) {
	// getting grade
	g, _ := strconv.ParseFloat(req.FormValue("grade"), 64)

	// new args struct
	args := Args{
		Name:    req.FormValue("name"),
		Grade:   g,
		Subject: req.FormValue("subject"),
	}

	// add to the list
	server.Add(args)
}
