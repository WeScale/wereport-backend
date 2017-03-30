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
		"Connect",
		"GET",
		"/connect",
		Connect,
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
		"/clients/{id}",
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
		"/consultants/{id}",
		GetOneConsultant,
	},
	Route{
		"ConsultantCreate",
		"POST",
		"/consultants",
		ConsultantCreate,
	},
	Route{
		"FactureIndex",
		"GET",
		"/factures",
		GetFactures,
	},
	Route{
		"FactureShow",
		"GET",
		"/factures/{id}",
		GetOneFacture,
	},
	Route{
		"FactureCreate",
		"POST",
		"/factures",
		FactureCreate,
	},
	Route{
		"ReportShowAll",
		"GET",
		"/reports/{year}/{month}",
		GetReports,
	},
	Route{
		"ReportShowOne",
		"GET",
		"/reports/{year}/{month}/consultant/{id}",
		GetReportsOneConsultant,
	},
	Route{
		"ReportCreate",
		"POST",
		"/reports",
		ReportCreate,
	},
	Route{
		"ReportDayIndex",
		"GET",
		"/reportdays",
		GetReportDays,
	},
	Route{
		"ReportDayCreate",
		"POST",
		"/reportsdays",
		ReportDayCreate,
	},
	Route{
		"ContratIndex",
		"GET",
		"/contrats",
		GetContrats,
	},
	Route{
		"ContratShow",
		"GET",
		"/contrats/{id}",
		GetOneContrat,
	},
	Route{
		"ContratConsultants",
		"GET",
		"/contrats/consultant/{id}",
		GetContratsConsultant,
	},
	Route{
		"ContratCreate",
		"POST",
		"/contrats",
		ContratCreate,
	},
}
