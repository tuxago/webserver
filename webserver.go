package main

import (
	"net/http"
)

func simple(w http.ResponseWriter, r *http.Request) {
	htmlCode := `
<html>
	<head>
		<title>Simple Title</title>
	</head>
	<body>
		<h1>Simple</h1>
	</body>
</html>`
	w.Write([]byte(htmlCode))
}

func main() {
	http.HandleFunc("/simple", simple)
	http.ListenAndServe(":8080", nil)
}
