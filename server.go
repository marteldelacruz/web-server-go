package main

import (
	"fmt"
	"net/http"

	Util "./util"
)

func HomePage(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	fmt.Fprintf(
		res, Util.LoadHtml("index.html"),
	)
}

func main() {
	// handle styles
	http.Handle("/", http.FileServer(http.Dir("./style")))

	// home page
	http.HandleFunc("/home", HomePage)

	// debug purposes :D
	fmt.Println("----------------------------------")
	fmt.Println("STARTING SERVER...")
	http.ListenAndServe(":9000", nil)
	Util.ScanString()
}
