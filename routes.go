package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"ClientsIndex",
		"GET",
		"/clients",
		GetClients,
	},
	Route{
		"ClientShow",
		"GET",
		"/clients/{clientid}",
		GetOneClient,
	},
	Route{
		"ClientCreate",
		"POST",
		"/clients",
		ClientCreate,
	},
	Route{
		"ConsultantIndex",
		"GET",
		"/consultants",
		GetConsultants,
	},
	Route{
		"ConsultantShow",
		"GET",
		"/consultants/{consultantid}",
		GetOneConsultant,
	},
	Route{
		"ConsultantCreate",
		"POST",
		"/consultants",
		ConsultantCreate,
	},
}
