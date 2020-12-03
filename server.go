package main

import (
	"fmt"
	"net/http"

	Admin "./admin"
	Util "./util"
)

func main() {
	// new server
	server := Admin.Server{}
	server.Init()

	// handle styles
	http.Handle("/", http.FileServer(http.Dir("./style")))

	// available pages
	http.HandleFunc("/home", server.HomePage)
	http.HandleFunc("/grades", server.GradesPage)
	http.HandleFunc("/average", server.AveragePage)
	http.HandleFunc("/general", server.GeneralPage)

	// forms
	http.HandleFunc("/save", server.Save)
	http.HandleFunc("/search", server.Search)

	// debug purposes :D
	fmt.Println("----------------------------------")
	fmt.Println("STARTING SERVER...")

	// start server
	http.ListenAndServe(":9000", nil)
	Util.ScanString()
}
