package main

import (
	"net/http"
)

func what(w http.ResponseWriter, r *http.Request) {
	htmlCode := `
<html>
	<head>
		<title>What?</title>
	</head>
	<body>
		<h1>What H1?</h1>
	</body>
</html>`
	w.Write([]byte(htmlCode))
}

func main() {
	http.HandleFunc("/what", what)
	http.ListenAndServe(":8080", nil)
}
