package main

import (
	"fmt"
	"net/http"

	Util "./util"
)

func root(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	fmt.Fprintf(
		res,
		`<DOCTYPE html>
		<html>
			<head>
				<title>Ejemplo Hola</title>
			</head>
			<body>
				Hola Mundo!
			</body>
		</html>`,
	)
}

func main() {
	http.HandleFunc("/", root)
	fmt.Println("----------------------------------")
	fmt.Println("STARTING SERVER...")
	http.ListenAndServe(":9000", nil)
	Util.ScanString()
}
