package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	msg := []byte("JF DEV CONTACT")
	resp, err := http.Post("http://0.0.0.0:3030", "text/plain", bytes.NewBuffer(msg))
	if err != nil {
		log.Println("Erro", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Erro", err)
	}

	fmt.Println(string(body))

}
