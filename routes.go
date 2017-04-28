package main

import (
	"net/http"

	"github.com/WeScale/wereport-backend/handlers"
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
		Handlers.Index,
	},
	Route{
		"Connect",
		"GET",
		"/connect",
		Handlers.Connect,
	},
	Route{
		"ClientsIndex",
		"GET",
		"/clients",
		Handlers.GetClients,
	},
	Route{
		"ClientShow",
		"GET",
		"/clients/{id}",
		Handlers.GetOneClient,
	},
	Route{
		"ClientCreate",
		"POST",
		"/clients",
		Handlers.ClientCreate,
	}, Route{
		"ConsultantIndex",
		"GET",
		"/consultants",
		Handlers.GetConsultants,
	},
	Route{
		"ConsultantShow",
		"GET",
		"/consultants/{id}",
		Handlers.GetOneConsultant,
	},
	Route{
		"ConsultantCreate",
		"POST",
		"/consultants",
		Handlers.ConsultantCreate,
	},
	Route{
		"ContratShow",
		"GET",
		"/contrats/{id}",
		Handlers.GetOneContrat,
	},
	Route{
		"ContratConsultants",
		"GET",
		"/contrats/consultant/{id}",
		Handlers.GetContratsConsultant,
	},
	Route{
		"ContratCreate",
		"POST",
		"/contrats",
		Handlers.ContratCreate,
	},
	Route{
		"ContratIndex",
		"GET",
		"/contrats",
		Handlers.GetContrats,
	},
	Route{
		"FactureShow",
		"GET",
		"/factures/{year}/{month}/",
		Handlers.GetOneMonthFactures,
	},
	Route{
		"ReportShowAll",
		"GET",
		"/reports/{year}/{month}",
		Handlers.GetReports,
	},
	Route{
		"ReportShowOne",
		"GET",
		"/reports/{year}/{month}/consultant/{id}",
		Handlers.GetReportsOneConsultant,
	},
	Route{
		"ReportCreate",
		"POST",
		"/reports",
		Handlers.ReportCreate,
	},
	Route{
		"ReportDayIndex",
		"GET",
		"/reportdays",
		Handlers.GetReportDays,
	},
	Route{
		"ReportDayCreate",
		"POST",
		"/reportsdays",
		Handlers.ReportDayCreate,
	},
}
