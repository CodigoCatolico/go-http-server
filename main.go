package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

const port = "8080"

func main() {

	mux := http.NewServeMux()

	// Usando um logger criado podemos passar flags para customizar o comportamento do logger.
	// Como exemplo está sendo utilizando as flags de micro segundos, prefixo de mensagem e
	// arquivo de origem.
	logger := log.New(os.Stdout, "HTTP: ", log.Lmicroseconds|log.Lmsgprefix|log.Lshortfile)

	// Adiconando o logMiddleware criamos um encapsulamento do handler, agora a func "path"
	// está envolvida pelo middleware. O controle da requisicao é passado na seguinte ordem:
	//
	// - logMiddleware
	// - handler
	// - logMiddleware
	//
	mux.Handle("/path/", logMiddleware(logger, http.HandlerFunc(path)))
	mux.Handle("/body", logMiddleware(logger, http.HandlerFunc(body)))
	mux.Handle("/query", logMiddleware(logger, http.HandlerFunc(query)))
	mux.Handle("/header", logMiddleware(logger, http.HandlerFunc(header)))

	fmt.Println("server running at port:", port)
	http.ListenAndServe(":"+port, mux)
}

func path(rw http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	response := struct {
		PathParam string `json:"path_param"`
	}{
		PathParam: strings.TrimPrefix(path, "/path/"),
	}
	rw.Header().Add("content-type", "application/json")
	json.NewEncoder(rw).Encode(&response)
}

func body(rw http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := struct {
		Body string `json:"body"`
	}{
		Body: string(body),
	}
	rw.Header().Add("content-type", "application/json")
	json.NewEncoder(rw).Encode(&response)
}

func query(rw http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	rw.Header().Add("content-type", "application/json")
	json.NewEncoder(rw).Encode(&query)
}

func header(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("content-type", "application/json")
	json.NewEncoder(rw).Encode(&r.Header)
}
