package routes

import (
	"net/http"
)

type code struct {
	Code string
}

func GetCode(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://github.com/paij0se/api-deno-compiler#this-a-simple-api-that-execute-your-deno-code-and-send-you-the-output-has-not-limit-per-request", http.StatusMovedPermanently)

}

func IndexRoute(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://github.com/paij0se/api-deno-compiler#this-a-simple-api-that-execute-your-deno-code-and-send-you-the-output-has-not-limit-per-request", http.StatusMovedPermanently)
}
