package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type code struct {
	Code string
}

type allCode []code

var codes = allCode{
	{
		Code: "console.log(Deno.version) //example code",
	},
}

func GetCode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(codes)
}

func IndexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<h1 align="center">make a http post request, more info in <a href="https://github.com/ELPanaJose/api-deno-compiler">Github Repository</a></h1>`)
}
