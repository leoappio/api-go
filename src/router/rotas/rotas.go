package rotas

import (
	"api/src/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

type Rota struct {
	URI        string
	Metodo     string
	Funcao     func(http.ResponseWriter, *http.Request)
	RequerAuth bool
}

func Configurar(r *mux.Router) *mux.Router {
	rotas := rotasUsuarios
	rotas = append(rotas, rotaLogin)
	rotas = append(rotas, rotasPublicacoes...)

	for _, rota := range rotas {
		if rota.RequerAuth {
			r.HandleFunc(rota.URI, middleware.Logger(middleware.Autenticar(rota.Funcao))).Methods(rota.Metodo)
		} else {
			r.HandleFunc(rota.URI, middleware.Logger(rota.Funcao)).Methods(rota.Metodo)
		}
	}

	return r
}
