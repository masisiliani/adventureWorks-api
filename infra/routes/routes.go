package routes

import (
	"adventureWorks-api/controller"
	"net/http"
)

// Route - Objeto que agrupa as informações das rotas.
// Name: Nome da Rota
// Method: Verbo HTTP
// Pattern:Path da Rota
// HandlerFunc: Método da aplicação (controller) que corresponde ao endpoint relacionado.
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes : agrupa todas as rotas disponíveis
type Routes []Route

// RoutesList var config the routes that api provides
var RoutesList = Routes{
	Route{
		"GetCustomerByID",
		"GET",
		"/customer/{id}",
		controller.GetCustomerByID,
	},
}
