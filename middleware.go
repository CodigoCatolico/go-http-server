package main

import (
	"log"
	"net/http"
	"time"
)

// Um middleware é uma camada de processamento que a mensagem
// vai passar antes de chegar as camadas interiores.
// https://www.alexedwards.net/blog/making-and-using-middleware
func logMiddleware(logger *log.Logger, next http.Handler) http.Handler {
	middleware := func(w http.ResponseWriter, r *http.Request) {

		// Executamos a lógica que precisamos antes do processamento da requisicao.
		// Nesse caso logamos que a mensagem chegou, seu metodo e path.
		logger.Printf("message started   : %s %s",
			r.Method,
			r.URL.Path,
		)
		// Capturamos o timestamp de inicio do processamento.
		t := time.Now()

		// Passamos a mensagem para o proximo handler.
		next.ServeHTTP(w, r)

		// Depois do tratamento do handler interno, temos acesso novamente a requisicao.
		// Assim podemos logar o tempo decorrido desde que passamos a mensagem para o
		// handler interno.
		logger.Printf("message completed : %s %s : %s",
			r.Method,
			r.URL.Path,
			time.Since(t).String(),
		)
	}

	// Por fim retornamos nosso middleware usando a funcao HandleFunc que recebe uma
	// simples funcao com a assinatura (http.ResponseWriter, *http.Request) e retorna
	// um http.Handler.
	return http.HandlerFunc(middleware)
}
