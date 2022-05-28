package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"

	noansi "github.com/ELPanaJose/api-deno-compiler/src/routes/others"
	"github.com/labstack/echo/v4"
	re "github.com/paij0se/heroku-echo-ip-dashboard/src/controllers"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func PostCode(e echo.Context) error {
	re.Requester(e.Scheme() + "://" + e.Request().Host)

	var inputCode code
	reqBody, err := ioutil.ReadAll(e.Request().Body)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err)
	}

	json.Unmarshal([]byte(reqBody), &inputCode)

	input := inputCode.Code
	// check if the request is empty
	if input == "" {
		return e.JSON(http.StatusBadRequest, "The request is empty")
	} else {

		fmt.Println("input:", input)

		// create the program
		program := RandStringRunes(10) + ".ts"
		f, err := os.Create(program)
		if err != nil {
			fmt.Println("some error creating the archive", err)
		}

		defer f.Close()

		// write the program
		_, err2 := f.WriteString(input)
		if err2 != nil {
			fmt.Println("error writing the archive", err2)
		}
		// execute deno and kill the process
		fmt.Println("archive edited")
		var stdout, stderr bytes.Buffer
		cmd := exec.Command("sh", "-c", `
	./deno run --allow-net --no-check `+program+`&`+` sleep 1;kill $! 2>&1`)
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr
		peo := cmd.Run()
		if peo != nil {
			fmt.Println(err)
		}
		// capture the stderr and stdout
		executedOut := stdout.String() + stderr.String()
		noAnsii := noansi.NoAnsi(executedOut)
		coolOut := strings.ReplaceAll(noAnsii, "sh: 2: kill: No such process", "")
		// fmt.Println(coolOut)
		// delete the archive
		err = os.Remove(program)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("archive deleted")
		return e.JSON(http.StatusOK, coolOut)
	}
}
