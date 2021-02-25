package main

import (
	"net/http"

	"github.com/igmarin/webservice/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
