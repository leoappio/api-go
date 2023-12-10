package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasUsuarios = []Rota{
	{
		URI:        "/usuarios",
		Metodo:     http.MethodPost,
		Funcao:     controllers.CriarUsuario,
		RequerAuth: false,
	},
	{
		URI:        "/usuarios",
		Metodo:     http.MethodGet,
		Funcao:     controllers.BuscarUsuarios,
		RequerAuth: true,
	},
	{
		URI:        "/usuarios/{usuarioId}",
		Metodo:     http.MethodGet,
		Funcao:     controllers.BuscarUsuario,
		RequerAuth: true,
	},
	{
		URI:        "/usuarios/{usuariosId}",
		Metodo:     http.MethodPut,
		Funcao:     controllers.AtualizarUsuario,
		RequerAuth: true,
	},
	{
		URI:        "/usuarios/{usuariosId}",
		Metodo:     http.MethodDelete,
		Funcao:     controllers.DeletarUsuario,
		RequerAuth: true,
	},
}
