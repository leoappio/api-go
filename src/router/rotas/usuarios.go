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
	{
		URI:        "/usuarios/{usuarioId}/seguir",
		Metodo:     http.MethodPost,
		Funcao:     controllers.SeguirUsuario,
		RequerAuth: true,
	},
	{
		URI:        "/usuarios/{usuarioId}/seguidores",
		Metodo:     http.MethodGet,
		Funcao:     controllers.BuscarSeguidores,
		RequerAuth: true,
	},
	{
		URI:        "/usuarios/{usuarioId}/seguindo",
		Metodo:     http.MethodGet,
		Funcao:     controllers.BuscarSeguindo,
		RequerAuth: true,
	},
	{
		URI:        "/usuarios/{usuarioId}/atualizar-senha",
		Metodo:     http.MethodPost,
		Funcao:     controllers.AtualizarSenha,
		RequerAuth: true,
	},
}
