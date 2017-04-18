package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {
	response, err := http.Get("http://api.theysaidso.com/qod.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Before printing response")
	fmt.Println(string(body))
}
