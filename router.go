package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	//handler := cors.Default().Handler(router)

	return router
}

// func NewRouter() *vestigo.Router {
// 	router := vestigo.NewRouter()
// 	// you can enable trace by setting this to true
// 	vestigo.AllowTrace = true

// 	// Setting up router global  CORS policy
// 	// These policy guidelines are overriddable at a per resource level shown below
// 	router.SetGlobalCors(&vestigo.CorsAccessControl{
// 		AllowOrigin:      []string{"*", "test.com"},
// 		AllowCredentials: true,
// 		ExposeHeaders:    []string{"*"},
// 		MaxAge:           3600 * time.Second,
// 		AllowHeaders:     []string{"authorization", "content-type"},
// 	})

// 	for _, route := range routes {
// 		var handler http.Handler

// 		handler = route.HandlerFunc
// 		handler = Logger(handler, route.Name)

// 		if strings.EqualFold("POST", route.Method) {
// 			router.Post(route.Pattern, route.HandlerFunc)
// 		} else if strings.EqualFold("GET", route.Method) {
// 			router.Get(route.Pattern, route.HandlerFunc)
// 		} else if strings.EqualFold("PUT", route.Method) {
// 			router.Put(route.Pattern, route.HandlerFunc)
// 		} else if strings.EqualFold("DELETE", route.Method) {
// 			router.Delete(route.Pattern, route.HandlerFunc)
// 		}

// 	}

// 	return router
// }
