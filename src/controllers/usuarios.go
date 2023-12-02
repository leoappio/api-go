package controllers

import (
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := io.ReadAll(r.Body)

	if erro != nil {
		log.Fatal(erro)
	}

	var usuario modelos.Usuario

	if erro := json.Unmarshal(corpoRequest, &usuario); erro != nil {
		log.Fatal(erro)
	}

	db, erro := banco.Conectar()
	if erro != nil {
		log.Fatal(erro)
	}

	repository := repositorios.NovoRepositorioDeUsuarios(db)
	repository.Criar(usuario)
}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {

}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {

}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {

}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {

}
