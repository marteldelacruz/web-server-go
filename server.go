package main

import (
	"fmt"
	"net/http"

	Util "./util"
)

// Default home page where the client can save the
// students data
func HomePage(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	// sending home.html
	fmt.Fprintf(
		res, Util.LoadHtml("home.html"),
	)
}

// This page will show all the data stored by the client
// on a single table
func GradesPage(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	// sending grades.html
	fmt.Fprintf(
		res, Util.LoadHtml("grades.html"),
	)
}

// this function will handle the requested methods from the
// home form in order to save new data
func Save(res http.ResponseWriter, req *http.Request) {
	fmt.Println("SAVE method =>" + req.Method)
	switch req.Method {
	case "POST":
		if err := req.ParseForm(); err != nil {
			fmt.Fprintf(res, "ParseForm() error %v", err)
			return
		}
		fmt.Println(req.PostForm)
		break
	default:
		fmt.Println("Cannot handle requested method :(")
		break
	}
}

func main() {
	// handle styles
	http.Handle("/", http.FileServer(http.Dir("./style")))

	// home page
	http.HandleFunc("/home", HomePage)
	http.HandleFunc("/grades", GradesPage)
	http.HandleFunc("/save", Save)

	// debug purposes :D
	fmt.Println("----------------------------------")
	fmt.Println("STARTING SERVER...")
	http.ListenAndServe(":9000", nil)
	Util.ScanString()
}
