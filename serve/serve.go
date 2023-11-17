package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	log.Println("Serve rodando......")
	http.HandleFunc("/", Serve)

	http.ListenAndServe("0.0.0.0:3030", nil)
	fmt.Println("Serve")

}

func Serve(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Println(w, "Erro ao ler o corpo da solicitação", http.StatusBadRequest)
		return
	}

	//Aqui o servidor vai esta respodendo nosso cliente
	msg_client := fmt.Sprintf(string(body))

	if msg_client == "JF DEV CONTACT" {
		fmt.Fprintf(w, "Ola cliente sua mensagem chegou")
	} else {
		fmt.Fprintf(w, "Ola cliente. Sua mensagem chegou mas nao foi oque esperamos")
	}

}
